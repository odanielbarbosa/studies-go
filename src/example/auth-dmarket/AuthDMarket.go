package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"strconv"
	"time"
)

type Offer struct {
	GameId string `json:"gameId"`
	Title  string `json:"title"`
	Image  string ` json:"image"`
	Extra  struct {
		CategoryPath string `json:"categoryPath"`
	}
}

type Keys struct {
	Private string `json:"private"`
	Public  string `json:"public"`
}

type MarketResponse struct {
	Objects []Offer `json:"objects"`
}

func GetPrivateKey(s string) (*[64]byte, error) {
	b, err := hex.DecodeString(s)
	if err != nil {
		return nil, err
	}
	var privateKey [64]byte
	copy(privateKey[:], b[:64])

	return &privateKey, nil
}

func Sign(pk, msg string) (signature string, err error) {
	b, err := GetPrivateKey(pk)
	return sign(b, []byte(msg)), nil
}

func sign(pk *[64]byte, msg []byte) string {
	return hex.EncodeToString(ed25519.Sign((*pk)[:], msg))
}

func getSignature(keys Keys, method string, path string, body string) string {
	timestamp := strconv.Itoa(int(time.Now().UTC().Unix()))
	unsigned := method + path + body + timestamp
	signature, _ := Sign(keys.Private, unsigned)
	return signature
}
