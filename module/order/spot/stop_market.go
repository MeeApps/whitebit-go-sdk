package spot

import whitebit "github.com/whitebit-exchange/go-sdk"

const stopMarketOrderEndpointUrl = "/api/v4/order/stop_market"

type StopMarketOrder struct {
	MarketOrder
	Activated           int64  `json:"activated"`
	ActivationPrice     string `json:"activation_price"`
	ActivationCondition string `json:"activationCondition"`
}

type StopMarketOrderParams struct {
	Market                  string `json:"market"`
	Amount                  string `json:"amount"`
	Side                    string `json:"side"`
	ActivationPrice         string `json:"activation_price"`
	ClientOrderId           string `json:"clientOrderId,omitempty"`
	SelfTradePreventionMode string `json:"stp,omitempty"` // no, cancel_both, cancel_new, cancel_old. Default: no
}

type stopMarketEndpoint struct {
	whitebit.AuthParams
	StopMarketOrderParams
}

func newStopMarketEndpoint(params StopMarketOrderParams) *stopMarketEndpoint {
	return &stopMarketEndpoint{
		AuthParams:            whitebit.NewAuthParams(stopMarketOrderEndpointUrl),
		StopMarketOrderParams: params,
	}
}
