package response

type CoinInfoResp struct {
	Rows []coinInfo `json:"rows"`
}

type coinInfo struct {
	Name         string  `json:"name"`
	Coin         string  `json:"coin"`
	RemainAmount string  `json:"remainAmount"`
	Chains       []chain `json:"chains"`
}

type chain struct {
	ChainType             string `json:"chainType"`
	Confirmation          string `json:"confirmation"`
	WithdrawFee           string `json:"withdrawFee"`
	DepositMin            string `json:"depositMin"`
	WithdrawMin           string `json:"withdrawMin"`
	Chain                 string `json:"chain"`
	ChainDeposit          string `json:"chainDeposit"`
	ChainWithdraw         string `json:"chainWithdraw"`
	MinAccuracy           string `json:"minAccuracy"`
	WithdrawPercentageFee string `json:"withdrawPercentageFee"`
}
