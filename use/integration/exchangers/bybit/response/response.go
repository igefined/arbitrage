package response

type Response struct {
	Code          int         `json:"retCode"`
	Message       string      `json:"retMsg"`
	Result        interface{} `json:"result"`
	ExtensionInfo interface{} `json:"retExtInfo"`
	Time          int64       `json:"time"`
}
