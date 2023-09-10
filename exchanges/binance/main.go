package binance

import (
	"github.com/jomasuleymen/p2p_mining/exchange"
)

var exchangeName string = "Binance"

var instance *exchange.Exchange

func GetInstance() *exchange.Exchange {
	if instance == nil {
		instance = &exchange.Exchange{
			Name:          exchangeName,
			Spot:          &spotDataFetcher{},
			P2PAdsFetcher: &adsFetcher{},
		}
	}

	return instance
}
