package spot

import (
	whitebit "github.com/whitebit-exchange/go-sdk"
)

const limitEndpointUrl = "/api/v4/order/new"

type LimitOrder struct {
	MarketOrder
	Price string `json:"price"`
}

type LimitOrderParams struct {
	Market                  string `json:"market"`
	Amount                  string `json:"amount"`
	Side                    string `json:"side"`
	Price                   string `json:"price"`
	PostOnly                bool   `json:"postOnly"`
	IOC                     bool   `json:"ioc"`
	ClientOrderId           string `json:"clientOrderId,omitempty"`
	SelfTradePreventionMode string `json:"stp,omitempty"` // no, cancel_both, cancel_new, cancel_old. Default: no
}

type limitEndpoint struct {
	whitebit.AuthParams
	LimitOrderParams
}

func newLimitEndpoint(params LimitOrderParams) *limitEndpoint {
	return &limitEndpoint{
		AuthParams:       whitebit.NewAuthParams(limitEndpointUrl),
		LimitOrderParams: params,
	}
}
