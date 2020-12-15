package user

import (
	"fmt"
	"net/http"

	"github.com/licbin/wechat/utils"
)

const (
	//getBlackListURL - 获取公众号的黑名单列表
	getBlackListURL = "https://api.weixin.qq.com/cgi-bin/tags/members/getblacklist?access_token=%s"
	//batchBlackListURL - 拉黑用户
	batchBlackListURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchblacklist?access_token=%s"

	//batchUnBlackListURL - 取消拉黑用户
	batchUnBlackListURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchunblacklist?access_token=%s"
)

// GetBlackList - 获取公众号的黑名单列表
func (srv *defaultService) GetBlackList(opneid string) (*GetBlackUserListResponse, error) {
	reqURL := fmt.Sprintf(getBlackListURL, srv.Get())

	req := new(GetBlackUserListRequest)
	req.BeginOpenID = opneid

	result := new(GetBlackUserListResponse)
	err := srv.Do(http.MethodPost, reqURL, req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// BatchBlackList - 拉黑用户
func (srv *defaultService) BatchBlackList(openidList []string) (*BatchBlackListResponse, error) {
	reqURL := fmt.Sprintf(batchBlackListURL, srv.Get())

	req := new(BatchBlackListRequest)
	req.OpenIDList = openidList

	result := new(BatchBlackListResponse)
	err := srv.Do(http.MethodPost, reqURL, req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// BatchUnBlackList - 取消拉黑用户
func (srv *defaultService) BatchUnBlackList(openidList []string) (*BatchUnBlackListResponse, error) {
	reqURL := fmt.Sprintf(batchUnBlackListURL, srv.Get())

	req := new(BatchUnBlackListRequest)
	req.OpenIDList = openidList

	result := new(BatchUnBlackListResponse)
	err := srv.Do(http.MethodPost, reqURL, req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}
