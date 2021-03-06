package user

import "github.com/licbin/wechat/account"

// Service - 服务
type Service interface {
	// GetUserList - 获取粉丝信息,
	// 公众号可通过本接口来获取帐号的关注者列表，关注者列表由一串OpenID（加密后的微信号，每个用户对每个公众号的OpenID是唯一的）组成。
	// 一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求。
	// 当公众号关注者数量超过10000时，可通过填写next_openid的值，从而多次拉取列表的方式来满足需求。
	// 关注者列表已返回完时，返回next_openid为空
	GetUserList(nextOpenID string) (*GetUserListResponse, error)
	// CreateTag - 创建标签,
	// 一个公众号最多可以创建100个标签。
	CreateTag(tageName string) (*CreateTagResponse, error)

	// GetTags - 获取公众号已创建的标签
	GetTags() (*GetTagsResponse, error)

	// UpdateTag - 编辑标签,不能更新id是0/1/2的标签
	UpdateTag(tagid int, tageName string) (*UpdateTagResponse, error)

	// DeleteTag - 删除标签
	DeleteTag(tagid int) (*DeleteTagResponse, error)

	// GetTagUserList -  获取标签下粉丝列表，返回数据格式和GetUserList一样。
	// 一次拉取调用最多拉取10000个关注者的OpenID，可以通过多次拉取的方式来满足需求。
	// 当公众号关注者数量超过10000时，可通过填写next_openid的值，从而多次拉取列表的方式来满足需求。
	// 关注者列表已返回完时，返回next_openid为空
	GetTagUserList(tagid int, nextOpenID string) (*GetUserListResponse, error)

	// BatchTagging - 批量为用户打标签,需要注意的是一个用户最多打上20个标签，另外每次传入的openid列表个数不能超过50个
	BatchTagging(tagID int, openidList []string) (*BatchTaggingResponse, error)
	// BatchUnTagging - 批量为用户取消标签,需要注意的是一个用户最多打上20个标签，另外每次传入的openid列表个数不能超过50个
	BatchUnTagging(tagID int, openidList []string) (*BatchUnTaggingResponse, error)

	// GetUserTags - 获取用户身上的标签列表
	GetUserTags(openid string) (*GetUserTagsResponse, error)

	// UpdateUserRemark - 设置用户备注名
	UpdateUserRemark(openid, remark string) (*UpdateUserRemarkResponse, error)

	// GetUserInfo - 获取用户基本信息（包括UnionID机制）
	GetUserInfo(openid string) (*GetUserInfoResponse, error)
}

type defaultService struct {
	*account.TokenStore
}

// NewService - service
func NewService(tokenSotre *account.TokenStore) Service {
	return &defaultService{
		tokenSotre,
	}
}
