package user

import "github.com/licbin/wechat/utils"

// GetUserListResponse - 用户关注者
type GetUserListResponse struct {
	utils.Error
	Total      int    `json:"total"`       //关注该公众账号的总用户数
	Count      int    `json:"count"`       //拉取的OPENID个数，最大值为10000
	NextOpenID string `json:"next_openid"` //拉取列表的最后一个用户的OPENID
	Data       struct {
		OpenID []string `json:"openid"`
	} `json:"data"` //列表数据，OPENID的列表
}

// CreateTagRequest - 创建tag请求
type CreateTagRequest struct {
	Tag Tag `json:"tag,omitempty"`
}

// CreateTagResponse - 创建 tag response
type CreateTagResponse struct {
	utils.Error
	Tag Tag `json:"tag,omitempty"`
}

// UpdateTagRequest - 更新tag请求
type UpdateTagRequest struct {
	Tag Tag `json:"tag,omitempty"`
}

// UpdateTagResponse - 更新标签返回结果
type UpdateTagResponse struct {
	*utils.Error
}

// DeleteTagRequest - 删除tag请求
type DeleteTagRequest struct {
	Tag Tag `json:"tag,omitempty"`
}

// DeleteTagResponse - 删除标签返回结果
type DeleteTagResponse struct {
	*utils.Error
}

// Tag - 标签
type Tag struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// GetTagsResponse - 获取标签response
type GetTagsResponse struct {
	utils.Error
	Tags []Tag `json:"tags,omitempty"`
}

// GetTagUserRequest - 获取标签下的粉丝的请求
type GetTagUserRequest struct {
	TagID      int    `json:"tagid,omitempty"`
	NextOpenID string `json:"next_openid,omitempty"`
}

// BatchTaggingRequest - 批量为用户打标签
type BatchTaggingRequest struct {
	TagID      int      `json:"tagid,omitempty"`
	OpenIDList []string `json:"openid_list,omitempty"`
}

// BatchTaggingResponse - 批量为用户标签
type BatchTaggingResponse struct {
	utils.Error
}

// BatchUnTaggingRequest - batch un tagging
type BatchUnTaggingRequest struct {
	TagID      int      `json:"tagid,omitempty"`
	OpenIDList []string `json:"openid_list,omitempty"`
}

// BatchUnTaggingResponse - batch un tagging
type BatchUnTaggingResponse struct {
	utils.Error
}

// GetUserTagsResponse - 获取用户TAGS
type GetUserTagsResponse struct {
	utils.Error
	TagIDList []int `json:"tagid_list,omitempty"`
}
