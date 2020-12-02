package account

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

//TokenRequester - interface to reflash wechat token
type TokenRequester interface {
	RefreshToken() (string, time.Time)
}

//TokenStore store token locally
type TokenStore struct {
	Token          string         `json:"token,omitempty"`
	ExpireOn       time.Time      `json:"expire_on,omitempty"`
	Requester      TokenRequester `json:"requester,omitempty"`
	RequestCounter int32          `json:"request_counter,omitempty"`
	NextTokenChan  chan string    `json:"next_token_chan,omitempty"`
	mutex          *sync.RWMutex  `json:"-"`
	Client         *http.Client
}

// NewTokenStore -
func NewTokenStore(requester TokenRequester) *TokenStore {
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
	// if requester == nil {
	// 	requester = requester
	// }
	return &TokenStore{
		Client:        &client,
		Requester:     requester,
		NextTokenChan: make(chan string, 0),
		mutex:         &sync.RWMutex{},
	}
}

//Get token from store
func (t *TokenStore) Get() (token string) {
	t.mutex.RLock()
	if t.ExpireOn.After(time.Now()) && t.Token != "" {
		token = t.Token
		t.mutex.RUnlock()
		return
	}
	t.mutex.RUnlock()
	return t.refreshToken()
}

//FreshNewtoken 刷新新的token
func (t *TokenStore) FreshNewtoken() {
	t.refreshToken()
}

//GetNewToken 刷新新的token
func (t *TokenStore) GetNewToken() string {
	return t.refreshToken()
}

//refresh token from server
func (t *TokenStore) refreshToken() (token string) {
	if t.Requester == nil {
		//如果Requester未设， 永远返回空token
		return ""
	}
	t.Token, t.ExpireOn = t.Requester.RefreshToken()
	return t.Token
}

// Do - do
func (t *TokenStore) Do(method, reqURL string, body, result interface{}) error {
	var req *http.Request
	var err error
	switch method {
	case http.MethodGet:
		req, err = http.NewRequest(http.MethodGet, reqURL, nil)
		if err != nil {
			return err
		}
	case http.MethodPost:
		if body == nil {
			return errors.New("body is nil")
		}
		data, err := json.Marshal(body)
		if err != nil {
			return err
		}

		req, err = http.NewRequest(http.MethodPost, reqURL, bytes.NewReader(data))
		if err != nil {
			return err
		}
	default:
		return errors.New("unsupport http method")

	}

	resp, err := t.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("http status code:%d, message:%s", resp.StatusCode, resp.Status)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(data, result); err != nil {
		return err
	}
	return nil
}
