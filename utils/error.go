package utils

// Error - 通用错误
type Error struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	ErrDesc string `json:"errdesc"`
}

// GetErrDesc - 获取错误描述
func GetErrDesc(code int) string {
	return ErrCodeToDesc[code]
}
