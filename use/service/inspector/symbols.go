package inspector

import (
	"context"
)

var symbols = []string{
	"BTCUSDT", "ETHUSDT", "LTCUSDT", "ATOMUSDT", "DOTUSDT", "OPUSDT", "DOGEUSDT", "APTUSDT",
	"IMXUSDT", "SOLUSDT", "ARBUSDT", "TRONUSDT", "ADAUSDT", "TONUSDT", "HFTUSDT", "KAVAUSDT",
	"SNXUSDT", "NEOUSDT", "XRPUSDT", "EOSUSDT", "TRXUSDT", "VENUSDT", "VETUSDT", "LINKUSDT",
	"WAVEUSDT", "BATUSDT", "XMRUSDT", "CELRUSDT", "DASHUSDT", "ALGOUSDT", "BANDUSDT", "CHRUSDT",
}

func (s *service) Symbols(ctx context.Context) ([]string, error) {
	return symbols, nil
}
