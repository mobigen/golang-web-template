package controllers

// HTTPResponse 공통으로 사용할 응답 메시지
type HTTPResponse struct {
	Result  string      `yaml:"result" json:"result"`
	Code    int         `yaml:"code" json:"code"`
	Message string      `yaml:"message" json:"message"`
	Data    interface{} `yaml:"data" json:"data"`
}

// For Result
const (
	STRSuccess string = "success"
	STRError   string = "error"
)

// For Code ( messages 디렉토리에 json파일의 code 값과 동일해야 한다.
const (
	HTTPErrCode1000 int = 1000 + iota
	HTTPErrCode1001
)

// for Message
var (
	HTTPErrMsg map[int]string
)
