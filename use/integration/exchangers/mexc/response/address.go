package response

type DepositAddress struct {
	Coin    string `json:"coin"`
	Network string `json:"network"`
	Address string `json:"address"`
	Memo    string `json:"memo,omitempty"`
}

type WithdrawAddresses struct {
	Data         []DepositAddress `json:"data"`
	TotalRecords uint             `json:"totalRecords"`
	Page         uint8            `json:"page"`
	TotalPageNum uint             `json:"totalPageNum"`
}
