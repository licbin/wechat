package officialaccount

import (
	"crypto/tls"
	"net"
	"net/http"
	"time"

	"github.com/licbin/wechat/account"
)

// OfficialAccount - 公众号账号
type OfficialAccount struct {
	*account.Account
	tokenStore *account.TokenStore
	client     http.Client
}

// NewOfficialAccount - 创建公众号请求内容
func NewOfficialAccount(account *account.Account, tokenStore *account.TokenStore) *OfficialAccount {
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
	return &OfficialAccount{
		account,
		tokenStore,
		client,
	}
}
