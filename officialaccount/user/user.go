package user

import (
	"fmt"
	"net/http"

	"github.com/licbin/wechat/utils"
)

const (
	getUserURL       = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s"
	updateUserRemark = "https://api.weixin.qq.com/cgi-bin/user/info/updateremark?access_token=%s"
	getUserInfo      = "https://api.weixin.qq.com/cgi-bin/user/info?access_token=%s&openid=%s&lang=zh_CN"
)

func (srv *defaultService) GetUserInfo(openid string) (*GetUserInfoResponse, error) {
	reqURL := fmt.Sprintf(getUserInfo, srv.Get(), openid)
	result := new(GetUserInfoResponse)
	err := srv.Do(http.MethodGet, reqURL, nil, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// UpdateUserRemark - implement Service
func (srv *defaultService) UpdateUserRemark(openid, remark string) (*UpdateUserRemarkResponse, error) {
	reqURL := fmt.Sprintf(updateUserRemark, srv.Get())

	req := new(UpdateUserRemarkRequest)
	req.OpenID = openid
	req.Remark = remark

	result := new(UpdateUserRemarkResponse)
	err := srv.Do(http.MethodPost, reqURL, req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// GetUserList - get user，implement Service
func (srv *defaultService) GetUserList(nextOpenID string) (*GetUserListResponse, error) {
	reqURL := fmt.Sprintf(getUserURL, srv.Get(), nextOpenID)
	result := new(GetUserListResponse)
	err := srv.Do(http.MethodGet, reqURL, nil, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}
