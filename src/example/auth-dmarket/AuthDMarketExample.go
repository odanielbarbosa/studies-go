package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func getRootUrl() string {
	return "https://api.dmarket.com"
}

func getFirstOfferFromMarket() (offer Offer) {
	resp, _ := http.Get(getRootUrl() + "/exchange/v1/market/items?gameId=a8db&limit=1&currency=USD")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	var marketResponse MarketResponse
	err := json.Unmarshal(bodyBytes, &marketResponse)
	if err != nil {
		return Offer{}
	}
	return marketResponse.Objects[0]
}

func buildTargetBodyFromOffer(offer Offer) string {
	return `{
			"targets": [{
				"amount": 1,
				"gameId": "` + offer.GameId + `",
				"price": {"amount": "2", "currency": "USD"},
				"attributes": {
					"gameId": "` + offer.GameId + `",
					"categoryPath": "` + offer.Extra.CategoryPath + `",
					"title": "` + offer.Title + `",
					"name": "` + offer.Title + `",
					"image": "` + offer.Image + `",
					"ownerGets": {"amount": "1", "currency": "USD"}
				}
			}]}`
}

func mainn() {
	offer := getFirstOfferFromMarket()
	keys := Keys{
		Private: "d0bd3bd93cc26453c37ec7c068ab94124b2a87d332cccb9e754ff851f1e82c5f16f687d46aa6f3add3223655df32b70e39e6e612461c9e694770c99af65c343a",
		Public:  "16f687d46aa6f3add3223655df32b70e39e6e612461c9e694770c99af65c343a",
	}
	body := buildTargetBodyFromOffer(offer)
	method := "POST"
	path := "/exchange/v1/target/create"
	timestamp := strconv.Itoa(int(time.Now().UTC().Unix()))
	unsigned := method + path + body + timestamp
	signature, _ := Sign(keys.Private, unsigned)
	client := &http.Client{}
	req, _ := http.NewRequest(method, getRootUrl()+path, ioutil.NopCloser(strings.NewReader(body)))
	req.Header.Set("X-Sign-Date", timestamp)
	req.Header.Set("X-Request-Sign", "dmar ed25519 "+signature)
	req.Header.Set("X-Api-Key", keys.Public)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	res, _ := client.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	_, err := io.Copy(os.Stdout, res.Body)
	if err != nil {
		return
	}
}
