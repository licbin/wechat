package user

import (
	"testing"

	"github.com/licbin/wechat/account"
)

func Test_defaultService_GetUserInfo(t *testing.T) {
	tr := account.NewDefaultTokenRequester(testAppID, testSecret)
	ts := account.NewTokenStore(tr)
	token := ts.Get()
	if token == "" {
		t.Error("token is empty")
		return
	}

	srv := NewService(ts)

	resp, err := srv.GetUserInfo(testOpenID)
	if err != nil {
		t.Error("GetUserInfo err:", err)
		return
	}
	t.Logf("response: %#v", resp)
}
