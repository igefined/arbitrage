package response

type Response struct {
	Ch         string      `json:"ch"`
	Code       int         `json:"ts"`
	Status     string      `json:"status"`
	Result     interface{} `json:"tick"`
	ErrMessage string      `json:"err-msg"`
	ErrCode    string      `json:"err-code"`
}
