package account

import (
	"testing"
)

func Test_defaultTokenStore_RefreshToken(t *testing.T) {
	ts := NewDefaultTokenRequester("wx950b68e93b5c0f8f", "9708e08d26f615abc025f9df0a03ac4f")
	gotToken, _ := ts.RefreshToken()
	if gotToken == "" {
		t.Errorf("defaultTokenStore.RefreshToken() gotToken = %v", gotToken)
		return
	}
	t.Log("token value:", gotToken)
}
