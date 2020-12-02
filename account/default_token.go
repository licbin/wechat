package account

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

type defaultTokenStore struct {
	AppID  string
	Secret string
	client http.Client
}

// NewDefaultTokenRequester - 生成默认的token store
func NewDefaultTokenRequester(appid, secret string) TokenRequester {
	atr := &http.Transport{
		Dial: (&net.Dialer{
			Timeout:   10 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		Proxy: nil,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		ResponseHeaderTimeout: 6 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
		MaxIdleConns:          50,
		IdleConnTimeout:       30 * time.Second,
	}
	client := http.Client{Transport: atr}
	return &defaultTokenStore{
		AppID:  appid,
		Secret: secret,
		client: client,
	}
}

//RefreshToken - get access token from Wechat, implements TokenRequester
func (ts *defaultTokenStore) RefreshToken() (token string, expiresOn time.Time) {
	apiServers := []string{
		"https://api.weixin.qq.com",
		"https://sh.api.weixin.qq.com",
		"https://sz.api.weixin.qq.com",
		"https://hk.api.weixin.qq.com",
	}
	const apiURLTemplate = "%s/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
	serverIdx := 0
retry:
	// log.Println("[INFO] - Refreshing token from", apiServers[serverIdx])
	req, err := http.NewRequest("GET", fmt.Sprintf(apiURLTemplate, apiServers[serverIdx], ts.AppID, ts.Secret), nil)
	resp, err := ts.client.Do(req)
	if err != nil {
		log.Println("[ERR] - [Get Wechat Token] error:", err)
		return
	}
	if resp.StatusCode != 200 {
		log.Println("[ERR] - [Get Wechat Token - StatusCode <> 200] error:", err)
		return
	}
	JSONData, err := ParseJSONBody(resp.Body)
	if err != nil {
		log.Println("[ERR] - [Get Wechat Token] Parsing body error:", err)
		return
	}
	m := JSONData.(map[string]interface{})
	if errcode, ok := m["errcode"].(float64); ok {
		switch {
		case errcode == -1:
			serverIdx++
			if serverIdx > len(apiServers)-1 {
				return
			}
			time.Sleep(time.Second)
			goto retry
		case errcode == 40013:
			log.Println("[ERR] - [Get Wechat Token] Invalid APP ID:", ts.AppID)
			return
		case errcode != 0:
			log.Println("[ERR] - [ERR][Get Wechat Token]:", m["errcode"], m["errmsg"])
			return
		}
	}

	token, _ = m["access_token"].(string)
	expiresIn, _ := m["expires_in"].(float64)
	expiresOn = time.Now().Add(time.Second * time.Duration(expiresIn-50))
	return
}

//ParseJSONBody Parse Json object from body
func ParseJSONBody(body io.ReadCloser) (JSONData interface{}, err error) {
	b, e := ioutil.ReadAll(body)
	defer body.Close()
	if e != nil {
		if err == nil {
			err = e
		}
		return
	}
	//先去掉golang包不能处理的字符,只处理32控制字符
	// for i, ch := range b {
	// 	switch {
	// 	case ch == '\r':
	// 	case ch == '\n':
	// 	case ch == '\t':
	// 	case ch < ' ':
	// 		b[i] = ' '
	// 	}
	// }

	json.Unmarshal(b, &JSONData)
	// m := JSONData.(map[string]interface{})
	// for k, v := range m {
	// 	switch vv := v.(type) {
	// 	case string:
	// 		fmt.Println(k, "is string", vv)
	// 	case int:
	// 		fmt.Println(k, "is int", vv)
	// 	case float64:
	// 		fmt.Println(k, "is float64", vv)
	// 	case []interface{}:
	// 		fmt.Println(k, "is an array:")
	// 		for i, u := range vv {
	// 			fmt.Println(i, u)
	// 		}
	// 	default:
	// 		fmt.Println(k, "is of a type I don't know how to handle")
	// 	}
	// }
	return
}
