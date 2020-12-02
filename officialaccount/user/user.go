package user

import (
	"fmt"
	"net/http"

	"github.com/licbin/wechat/utils"
)

const (
	getUserURL = "https://api.weixin.qq.com/cgi-bin/user/get?access_token=%s&next_openid=%s"
)

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
