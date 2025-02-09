package rest

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sudoblockio/icon-go-api/config"
	"github.com/sudoblockio/icon-go-api/service"
	"go.uber.org/zap"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var CirculatingSupply float64
var TotalSupply float64

func GetCirculatingSupply() (float64, error) {
	totalSupply, err := service.IconNodeServiceGetTotalSupply()
	TotalSupply = totalSupply
	if err != nil {
		return 0, err
	}

	burnBalance, err := service.IconNodeServiceGetBalance("hx1000000000000000000000000000000000000000")
	if err != nil {
		return 0, err
	}
	circulatingSupply := totalSupply - burnBalance
	return circulatingSupply, err
}

var LastUpdatedTimeCirculatingSupply time.Time

func UpdateCirculatingSupply() {
	timeDiff := time.Now().Sub(LastUpdatedTimeCirculatingSupply)
	if timeDiff > config.Config.StatsCirculatingSupplyUpdateTime {
		circulatingSupply, err := GetCirculatingSupply()
		if err != nil {
			zap.S().Info("Error getting circulating-supply: ", err)
		}
		CirculatingSupply = circulatingSupply
		LastUpdatedTimeCirculatingSupply = time.Now()
	}
}

var MarketCap float64

func GetMarketCap() (float64, error) {
	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/icon")
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0.0, err
	}

	response := make(map[string]interface{})
	err = json.Unmarshal(body, &response)
	if err != nil {
		// Just return cached value if there is an error
		return MarketCap, err
	}
	usdPrice, ok := response["market_data"].(map[string]interface{})["current_price"].(map[string]interface{})["usd"].(float64)
	if !ok {
		return MarketCap, errors.New(fmt.Sprintf("Error parsing coingecko response: %T", response))
	}
	UpdateCirculatingSupply()
	MarketCap = CirculatingSupply * usdPrice
	return MarketCap, err
}

var LastUpdatedTimeMarketCap time.Time

func UpdateMarketCap() {
	timeDiff := time.Now().Sub(LastUpdatedTimeMarketCap)
	if timeDiff > config.Config.StatsMarketCapUpdateTime {
		marketCap, err := GetMarketCap()
		if err != nil {
			zap.S().Info("Error getting market-cap: ", err)
		}
		MarketCap = marketCap
		LastUpdatedTimeMarketCap = time.Now()
	}
}
