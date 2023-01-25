package binance

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestListHistoricalTwapOrderService_Do(t *testing.T) {
	c := NewClient("2n2Lq3Iu1tFWPk4WPOXIVlm8WSJud9O0xD0NVPRDiwX0hwdfGgU9PDLa0O1oyuIY",
		"31wsK0Ung09SSAgse73BPY3rpa56ZUM2anlXxgyzs65k5vQFd1b4Uc8Xdwyh4mWE")

	_, err := c.NewCreateTwapOrderSerivce().Symbol("BTCUSDT").Quantity(0.45).
		Side(SideTypeBuy).Duration(10 * 60).Do(context.TODO())

	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(10 * time.Second)
	list, err := c.NewOpenTwapOrderService().Do(context.TODO())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(fmt.Sprintf("%+v", list))

	res1, err := c.NewCancelTwapOrderService().AlgoId(fmt.Sprintf("%d", list.Orders[0].AlgoID)).Do(context.TODO())

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fmt.Sprintf("%+v\n", res1))
}
