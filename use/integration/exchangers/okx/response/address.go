package response

type DepositAddresses []DepositAddress

type DepositAddress struct {
	Chain    string `json:"chain"`
	CtAddr   string `json:"ctAddr"`
	Ccy      string `json:"ccy"`
	To       string `json:"to"`
	Addr     string `json:"addr"`
	Selected bool   `json:"selected"`
}
