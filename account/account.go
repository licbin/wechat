package account

import (
	"crypto/tls"
	"net/http"
)

//Account store wechat account info.
type Account struct {
	ID        string `json:"id"`               //主体账号ID,选填
	Domain    string `json:"domain,omitempty"` //站点信息
	AppID     string `json:"appid"`            //APPID 对应公众号APPID,或者小程序APPID
	AppSecret string `json:"appsecret"`        //APPID 对应的 APPSECRET ，主要用来获取token信息，如果需要生成token 则必须传入。
	AppToken  string `json:"apptoken"`         // APPTOKEN 在公众号或小程序的微信端后台管理里配置的信息
	// AESKeys   AESKeys     `json:"asekeys"`

	MerchantID     string          `json:"mchid,omitempty"`   //商户ID,用于微信支付。如果需要使用微信的支付信息则必填。
	SignKey        string          `json:"signkey,omitempty"` //微信支付需要用到的签名。如果需要使用微信的支付信息则必填。
	Cert           tls.Certificate `json:"-"`
	HTTPCertClient *http.Client    `json:"-"`
	HTTPClient     *http.Client    `json:"-"`
	certPEMBlock   []byte          `json:"-"`
	keyPEMBlock    []byte          `json:"-"`

	Token *TokenStore `json:"-"`
}
