package utils

//CdataString xml cdata 转义标签
type CdataString struct {
	Value string `xml:",cdata"`
}

//String  CdataString string
func (c CdataString) String() string {
	return c.Value
}
