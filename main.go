package main

import (
	"fmt"
	"time"

	service "vang247.com/timers/services"
)

func HandleSaveChartDatas() {
	urlRequest := "https://ws_service.vang247.com/api/v1/save-chart-data"
	resp, err := service.GetRequest(urlRequest, "")
	if err != nil {
		fmt.Printf("err with get from goldenfund: %v\n", err)
		return
	}
	fmt.Printf("resp.StatusCode HandleSaveChartDatas: %v\n", resp.StatusCode)
}

func HandleUpdateSJCBankPrices() {
	urlRequest := "https://ws_other_price.vang247.com/updates"
	resp, err := service.GetRequest(urlRequest, "")
	if err != nil {
		fmt.Printf("err with get from goldenfund: %v\n", err)
		return
	}
	fmt.Printf("resp.StatusCode HandleUpdateSJCBankPrices: %v\n", resp.StatusCode)
}

// func init() {
// 	HandleSaveChartDatas()
// 	HandleUpdateSJCBankPrices()
// }

func main() {
	ticker5Minutes := time.NewTicker(5 * time.Minute)
	// Ticker cho mỗi 60 phút
	ticker30Minutes := time.NewTicker(30 * time.Minute)

	// Kênh để kết thúc chương trình
	done := make(chan bool)

	for {
		select {
		case <-ticker5Minutes.C:
			HandleUpdateSJCBankPrices()
		case <-ticker30Minutes.C:
			HandleSaveChartDatas()
		case <-done:
			return
		}
	}
}
