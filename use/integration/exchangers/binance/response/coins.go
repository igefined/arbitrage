package response

type CoinInformation struct {
	Coin        string    `json:"coin"`
	Name        string    `json:"name"`
	NetworkList []Network `json:"networkList"`
}

type Network struct {
	Coin           string `json:"coin"`
	Name           string `json:"name"`
	DepositEnable  bool   `json:"depositEnable"`
	WithdrawEnable bool   `json:"withdrawEnable"`
}
