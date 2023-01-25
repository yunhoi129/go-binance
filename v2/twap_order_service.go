package binance

import (
	"context"
	"fmt"
	"net/http"

	"github.com/adshao/go-binance/v2/futures"
)

type CreateTwapOrderService struct {
	c            *Client
	symbol       string
	side         SideType
	positionSide *futures.PositionSideType
	quantity     float64
	duration     int64
	clientAlgoId *string
	reduceOnly   *bool
	limitPrice   *float64
}

func (s *CreateTwapOrderService) Symbol(symbol string) *CreateTwapOrderService {
	s.symbol = symbol
	return s
}

func (s *CreateTwapOrderService) Side(side SideType) *CreateTwapOrderService {
	s.side = side
	return s
}

func (s *CreateTwapOrderService) PositionSide(side futures.PositionSideType) *CreateTwapOrderService {
	s.positionSide = &side
	return s
}

func (s *CreateTwapOrderService) Quantity(quantity float64) *CreateTwapOrderService {
	s.quantity = quantity
	return s
}

// Duration in second
func (s *CreateTwapOrderService) Duration(duration int64) *CreateTwapOrderService {
	s.duration = duration
	return s
}

func (s *CreateTwapOrderService) ClientAlgoId(algoId string) *CreateTwapOrderService {
	s.clientAlgoId = &algoId
	return s
}

func (s *CreateTwapOrderService) ReduceOnly(reduceOnly bool) *CreateTwapOrderService {
	s.reduceOnly = &reduceOnly
	return s
}

func (s *CreateTwapOrderService) LimitPrice(limitPrice float64) *CreateTwapOrderService {
	s.limitPrice = &limitPrice
	return s
}

func (s *CreateTwapOrderService) createOrder(ctx context.Context, endpoint string, opts ...RequestOption) (data []byte, header *http.Header, err error) {

	r := &request{
		method:   http.MethodPost,
		endpoint: endpoint,
		secType:  secTypeSigned,
	}
	m := params{
		"symbol":   s.symbol,
		"side":     s.side,
		"quantity": s.quantity,
		"duration": s.duration,
	}
	if s.positionSide != nil {
		m["positionSide"] = *s.positionSide
	}

	if s.clientAlgoId != nil {
		m["clientAlgoId"] = *s.clientAlgoId
	}

	if s.reduceOnly != nil {
		m["reduceOnly"] = *s.reduceOnly
	}

	if s.limitPrice != nil {
		m["limitPrice"] = *s.limitPrice
	}

	r.setFormParams(m)
	data, header, err = s.c.callAPIWithHeader(ctx, r, opts...)
	if err != nil {
		return []byte{}, &http.Header{}, err
	}
	return data, header, nil
}

func (s *CreateTwapOrderService) Do(ctx context.Context, opt ...RequestOption) (res *CreateTwapOrderResponse, err error) {
	data, _, err := s.createOrder(ctx, "/sapi/v1/algo/futures/newOrderTwap", opt...)
	if err != nil {
		return nil, err
	}

	res = new(CreateTwapOrderResponse)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}

	return
}

type CreateTwapOrderResponse struct {
	ClientAlgoID string `json:"clientAlgoId"`
	Success      bool   `json:"success"`
	Code         int    `json:"code"`
	Msg          string `json:"msg"`
}

type ListHistoricalTwapOrderService struct {
	c         *Client
	symbol    string
	side      *SideType
	startTime *int64
	endTime   *int64
	page      *int
	pageSize  *int
}

func (s *ListHistoricalTwapOrderService) Symbol(symbol string) *ListHistoricalTwapOrderService {
	s.symbol = symbol
	return s
}

func (s *ListHistoricalTwapOrderService) Side(side *SideType) *ListHistoricalTwapOrderService {
	s.side = side
	return s
}

func (s *ListHistoricalTwapOrderService) StartTime(time int64) *ListHistoricalTwapOrderService {
	s.startTime = &time
	return s
}

func (s *ListHistoricalTwapOrderService) EndTime(time int64) *ListHistoricalTwapOrderService {
	s.endTime = &time
	return s
}

// Page Default 1
func (s *ListHistoricalTwapOrderService) Page(page int) *ListHistoricalTwapOrderService {
	s.page = &page
	return s
}

// PageSize Min:1 Max:100 Default: 100
func (s *ListHistoricalTwapOrderService) PageSize(size int) *ListHistoricalTwapOrderService {
	s.pageSize = &size
	return s
}

func (s *ListHistoricalTwapOrderService) Do(ctx context.Context, opts ...RequestOption) (res *ListHistoricalTwapOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/historicalOrders",
		secType:  secTypeSigned,
	}
	r.setParam("symbol", s.symbol)

	if s.side != nil {
		r.setParam("side", *s.side)
	}

	if s.startTime != nil {
		r.setParam("startTime", *s.startTime)
	}

	if s.endTime != nil {
		r.setParam("endTime", *s.endTime)
	}

	if s.page != nil {
		r.setParam("page", *s.page)
	}

	if s.pageSize != nil {
		r.setParam("pageSize", *s.pageSize)
	}

	data, header, err := s.c.callAPIWithHeader(ctx, r, opts...)
	fmt.Println(header.Get("X-Sapi-Used-Ip-Weight-1m"))
	if err != nil {
		return nil, err
	}
	res = new(ListHistoricalTwapOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

type ListHistoricalTwapOrderResponse struct {
	Total  int `json:"total"`
	Orders []struct {
		AlgoID       int    `json:"algoId"`
		Symbol       string `json:"symbol"`
		Side         string `json:"side"`
		PositionSide string `json:"positionSide"`
		TotalQty     string `json:"totalQty"`
		ExecutedQty  string `json:"executedQty"`
		ExecutedAmt  string `json:"executedAmt"`
		AvgPrice     string `json:"avgPrice"`
		ClientAlgoID string `json:"clientAlgoId"`
		BookTime     int64  `json:"bookTime"`
		EndTime      int64  `json:"endTime"`
		AlgoStatus   string `json:"algoStatus"`
		AlgoType     string `json:"algoType"`
		Urgency      string `json:"urgency"`
	} `json:"orders"`
}

type CancelTwapOrderService struct {
	c      *Client
	algoId string
}

func (s *CancelTwapOrderService) AlgoId(algoId string) *CancelTwapOrderService {
	s.algoId = algoId
	return s
}

func (s *CancelTwapOrderService) Do(ctx context.Context, opts ...RequestOption) (res *CancelTwapOrderResponse, err error) {
	r := &request{
		method:   http.MethodDelete,
		endpoint: "/sapi/v1/algo/futures/order",
		secType:  secTypeSigned,
	}
	r.setParam("algoId", s.algoId)

	data, _, err := s.c.callAPIWithHeader(ctx, r, opts...)

	if err != nil {
		return nil, err
	}

	res = new(CancelTwapOrderResponse)
	err = json.Unmarshal(data, res)
	if err != nil {
		return nil, err
	}

	return
}

type CancelTwapOrderResponse struct {
	AlgoID  int    `json:"algoId"`
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Msg     string `json:"msg"`
}

type OpenTwapOrderService struct {
	c *Client
}

func (s *OpenTwapOrderService) Do(ctx context.Context, opt ...RequestOption) (res *OpenTwapOrderResponse, err error) {
	r := &request{
		method:   http.MethodGet,
		endpoint: "/sapi/v1/algo/futures/openOrders",
		secType:  secTypeSigned,
	}
	data, err := s.c.callAPI(ctx, r, opt...)
	if err != nil {
		return nil, err
	}
	res = new(OpenTwapOrderResponse)
	err = json.Unmarshal(data, res)

	if err != nil {
		return nil, err
	}
	return
}

type OpenTwapOrderResponse struct {
	Total  int `json:"total"`
	Orders []struct {
		AlgoID       int    `json:"algoId"`
		Symbol       string `json:"symbol"`
		Side         string `json:"side"`
		PositionSide string `json:"positionSide"`
		TotalQty     string `json:"totalQty"`
		ExecutedQty  string `json:"executedQty"`
		ExecutedAmt  string `json:"executedAmt"`
		AvgPrice     string `json:"avgPrice"`
		ClientAlgoID string `json:"clientAlgoId"`
		BookTime     int64  `json:"bookTime"`
		EndTime      int64  `json:"endTime"`
		AlgoStatus   string `json:"algoStatus"`
		AlgoType     string `json:"algoType"`
		Urgency      string `json:"urgency"`
	} `json:"orders"`
}
