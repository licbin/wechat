package user

import (
	"fmt"
	"net/http"

	"github.com/licbin/wechat/utils"
)

const (
	createTagURL      = "https://api.weixin.qq.com/cgi-bin/tags/create?access_token=%s"
	getTagsURL        = "https://api.weixin.qq.com/cgi-bin/tags/get?access_token=%s"
	updateTagURL      = "https://api.weixin.qq.com/cgi-bin/tags/update?access_token=%s"
	deleteTagURL      = "https://api.weixin.qq.com/cgi-bin/tags/delete?access_token=%s"
	getTagUserURL     = "https://api.weixin.qq.com/cgi-bin/user/tag/get?access_token=%s"
	batchTaggingURL   = "https://api.weixin.qq.com/cgi-bin/tags/members/batchtagging?access_token=%s"
	batchUnTaggingURL = "https://api.weixin.qq.com/cgi-bin/tags/members/batchuntagging?access_token=%s"
)

// CreateTag - create tag
func (srv *defaultService) CreateTag(name string) (*CreateTagResponse, error) {
	reqURL := fmt.Sprintf(createTagURL, srv.Get())
	req := CreateTagRequest{
		Tag: Tag{
			Name: name,
		},
	}
	result := new(CreateTagResponse)
	err := srv.Do(http.MethodPost, reqURL, &req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// UpdateTag - create tag
func (srv *defaultService) UpdateTag(id int, name string) (*UpdateTagResponse, error) {
	reqURL := fmt.Sprintf(updateTagURL, srv.Get())
	req := UpdateTagRequest{
		Tag: Tag{
			ID:   id,
			Name: name,
		},
	}
	result := new(UpdateTagResponse)
	err := srv.Do(http.MethodPost, reqURL, &req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// DeleteTag - create tag
func (srv *defaultService) DeleteTag(id int) (*DeleteTagResponse, error) {
	reqURL := fmt.Sprintf(deleteTagURL, srv.Get())
	req := DeleteTagRequest{
		Tag: Tag{
			ID: id,
		},
	}
	result := new(DeleteTagResponse)
	err := srv.Do(http.MethodPost, reqURL, &req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// GetTags - get tags - implement Service Interface
func (srv *defaultService) GetTags() (*GetTagsResponse, error) {
	reqURL := fmt.Sprintf(getTagsURL, srv.Get())

	result := new(GetTagsResponse)
	err := srv.Do(http.MethodGet, reqURL, nil, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// GetTagUserList - get user，implement Service
func (srv *defaultService) GetTagUserList(tagid int, nextOpenID string) (*GetUserListResponse, error) {
	reqURL := fmt.Sprintf(getTagUserURL, srv.Get())

	req := new(GetTagUserRequest)
	req.TagID = tagid
	req.NextOpenID = nextOpenID

	result := new(GetUserListResponse)
	err := srv.Do(http.MethodPost, reqURL, req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// BatchTagging - implement Service
func (srv *defaultService) BatchTagging(tagID int, openidList []string) (*BatchTaggingResponse, error) {
	reqURL := fmt.Sprintf(batchTaggingURL, srv.Get())

	req := new(BatchTaggingRequest)
	req.TagID = tagID
	req.OpenIDList = openidList

	result := new(BatchTaggingResponse)
	err := srv.Do(http.MethodPost, reqURL, req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}

// BatchUnTagging - implement Service
func (srv *defaultService) BatchUnTagging(tagID int, openidList []string) (*BatchUnTaggingResponse, error) {
	reqURL := fmt.Sprintf(batchUnTaggingURL, srv.Get())

	req := new(BatchUnTaggingRequest)
	req.TagID = tagID
	req.OpenIDList = openidList

	result := new(BatchUnTaggingResponse)
	err := srv.Do(http.MethodPost, reqURL, req, result)
	if err != nil {
		return nil, err
	}
	result.ErrDesc = utils.GetErrDesc(result.ErrCode)
	return result, nil
}
