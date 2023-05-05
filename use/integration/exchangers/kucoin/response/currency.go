package response

type CurrencyDetail struct {
	Currency        string  `json:"currency"`
	Name            string  `json:"name"`
	FullName        string  `json:"fullName"`
	Precision       int     `json:"precision"`
	Confirms        string  `json:"confirms,omitempty"`
	ContractAddress string  `json:"ContractAddress"`
	IsMarginEnabled bool    `json:"isMarginEnabled"`
	IsDebitEnabled  bool    `json:"isDebitEnabled"`
	Chains          []Chain `json:"chains"`
}

type Chain struct {
	Name              string `json:"chainName"`
	Chain             string `json:"chain"`
	WithdrawMinSize   string `json:"withdrawMinSize"`
	WithdrawMinFee    string `json:"withdrawMinFee"`
	IsWithdrawEnabled bool   `json:"isWithdrawEnabled"`
	IsDepositEnabled  bool   `json:"isDepositEnabled"`
	Confirms          int    `json:"confirms"`
	ContractAddress   string `json:"contractAddress"`
}
