package user

import (
	"testing"

	"github.com/licbin/pkg/xstring"
	"github.com/licbin/wechat/account"
)

const (
	testAppID  = "wx950b68e93b5c0f8f"
	testSecret = "9708e08d26f615abc025f9df0a03ac4f"
	testOpenID = "oGFjE6H05ZiaEtvZ_VPjdeSsI--w"
	testTagID  = 0
)

func Test_tag(t *testing.T) {
	tr := account.NewDefaultTokenRequester(testAppID, testSecret)
	ts := account.NewTokenStore(tr)
	token := ts.Get()
	if token == "" {
		t.Error("token is empty")
		return
	}

	srv := NewService(ts)
	tagName := "测试" + xstring.RandomNumber(2)

	// 测试创建tag
	resp, err := srv.CreateTag(tagName)
	if err != nil {
		t.Errorf("CreateTag Error:%v", err)
		return
	}
	if resp.ErrCode != 0 {
		t.Errorf("CreateTag Errcode:%d, ErrMsg:%s, ErrDesc:%s", resp.ErrCode, resp.ErrMsg, resp.ErrDesc)
		return
	}

	// 获取tag
	gt, err := srv.GetTags()
	if err != nil {
		t.Errorf("GetTags Error:%v", err)
		return
	}
	if gt.ErrCode != 0 {
		t.Errorf("GetTags Errcode:%d, ErrMsg:%s, ErrDesc:%s", gt.ErrCode, gt.ErrMsg, gt.ErrDesc)
		return
	}
	if len(gt.Tags) == 0 {
		t.Error("GetTags: tags length should greater 0, but is 0,")
		return
	}

	// 更新Tag
	for _, v := range gt.Tags {
		t.Logf("Range tags name:%s, id:%d", v.Name, v.ID)
		if v.ID != 0 && v.ID != 1 && v.ID != 2 {
			ut, err := srv.UpdateTag(v.ID, v.Name+"修改")
			if err != nil {
				t.Errorf("update Tag error:%v", err)
				return
			}
			if ut.ErrCode != 0 {
				t.Errorf("UpdateTag Errcode:%d, ErrMsg:%s, ErrDesc:%s", ut.ErrCode, ut.ErrMsg, ut.ErrDesc)
				return
			}
		}
		// srv.UpdateTag(v.ID)
	}

	// 测试删除TAG
	for _, v := range gt.Tags {
		if v.ID != 0 && v.ID != 1 && v.ID != 2 {
			ut, err := srv.DeleteTag(v.ID)
			if err != nil {
				t.Errorf("DeleteTag  error:%v", err)
				return
			}
			if ut.ErrCode != 0 {
				t.Errorf("DeleteTag Errcode:%d, ErrMsg:%s, ErrDesc:%s", ut.ErrCode, ut.ErrMsg, ut.ErrDesc)
				return
			}
		}
	}

}

func Test_GetTagUserList(t *testing.T) {
	tr := account.NewDefaultTokenRequester(testAppID, testSecret)
	ts := account.NewTokenStore(tr)
	token := ts.Get()
	if token == "" {
		t.Error("token is empty")
		return
	}

	srv := NewService(ts)
	resp, err := srv.GetTagUserList(0, "")
	if err != nil {
		t.Error("GetTagUserList is empty")
		return
	}

	t.Logf("GetTagUserList: %v", resp.Data.OpenID)
}

func Test_BatchTaggingAndUnTagging(t *testing.T) {
	tr := account.NewDefaultTokenRequester(testAppID, testSecret)
	ts := account.NewTokenStore(tr)
	token := ts.Get()
	if token == "" {
		t.Error("token is empty")
		return
	}
	srv := NewService(ts)

	tagName := "测试" + xstring.RandomNumber(2)
	// 测试创建tag
	resp1, err := srv.CreateTag(tagName)
	if err != nil {
		t.Errorf("CreateTag Error:%v", err)
		return
	}

	resp, err := srv.BatchTagging(resp1.Tag.ID, []string{testOpenID})
	if err != nil {
		t.Error("BatchTagging ERR:", err)
		return
	}
	if resp.ErrCode != 0 {
		t.Errorf("BatchTagging errcode:%d,errmessgae:%s,errdesc:%s",
			resp.ErrCode, resp.ErrMsg, resp.ErrDesc)
		return
	}

	resp2, err := srv.BatchUnTagging(resp1.Tag.ID, []string{testOpenID})
	if err != nil {
		t.Error("BatchTagging ERR:", err)
		return
	}
	if resp2.ErrCode != 0 {
		t.Errorf("BatchTagging errcode:%d,errmessgae:%s,errdesc:%s",
			resp.ErrCode, resp.ErrMsg, resp.ErrDesc)
		return
	}
}
