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

// UpdateUserRemarkRequest - 设置用户备注
type UpdateUserRemarkRequest struct {
	OpenID string `json:"openid,omitempty"`
	Remark string `json:"remark,omitempty"`
}

// UpdateUserRemarkResponse - 设置用户备注
type UpdateUserRemarkResponse struct {
	utils.Error
}

// GetUserInfoResponse - user info
type GetUserInfoResponse struct {
	utils.Error
	Subscribe      int    `json:"subscribe,omitempty"`       //用户是否订阅该公众号标识，值为0时，代表此用户没有关注该公众号，拉取不到其余信息。
	OpenID         string `json:"openid,omitempty"`          //用户的标识，对当前公众号唯一
	Nickname       string `json:"nickname,omitempty"`        //用户的昵称
	Sex            int    `json:"sex,omitempty"`             //用户的性别，值为1时是男性，值为2时是女性，值为0时是未知
	Language       string `json:"language,omitempty"`        //用户所在城市
	City           string `json:"city,omitempty"`            //用户所在国家
	Province       string `json:"province,omitempty"`        //用户所在省份
	Country        string `json:"country,omitempty"`         //用户的语言，简体中文为zh_CN
	HeadImgURL     string `json:"headimgurl,omitempty"`      //用户头像，最后一个数值代表正方形头像大小（有0、46、64、96、132数值可选，0代表640*640正方形头像），用户没有头像时该项为空。若用户更换头像，原有头像URL将失效。
	SubscribeTime  int64  `json:"subscribe_time,omitempty"`  //用户关注时间，为时间戳。如果用户曾多次关注，则取最后关注时间
	UnionID        string `json:"unionid,omitempty"`         //只有在用户将公众号绑定到微信开放平台帐号后，才会出现该字段。
	Remark         string `json:"remark,omitempty"`          //公众号运营者对粉丝的备注，公众号运营者可在微信公众平台用户管理界面对粉丝添加备注
	GroupID        int    `json:"groupid,omitempty"`         //用户所在的分组ID（兼容旧的用户分组接口）
	TagIDList      []int  `json:"tagid_list,omitempty"`      //	用户被打上的标签ID列表
	SubscribeScene string `json:"subscribe_scene,omitempty"` //返回用户关注的渠道来源，ADD_SCENE_SEARCH 公众号搜索，ADD_SCENE_ACCOUNT_MIGRATION 公众号迁移，ADD_SCENE_PROFILE_CARD 名片分享，ADD_SCENE_QR_CODE 扫描二维码，ADD_SCENE_PROFILE_LINK 图文页内名称点击，ADD_SCENE_PROFILE_ITEM 图文页右上角菜单，ADD_SCENE_PAID 支付后关注，ADD_SCENE_WECHAT_ADVERTISEMENT 微信广告，ADD_SCENE_OTHERS 其他
	QRScene        int    `json:"qr_scene,omitempty"`        //二维码扫码场景（开发者自定义）
	QRSceneStr     string `json:"qr_scene_str,omitempty"`    //二维码扫码场景描述（开发者自定义）
}
