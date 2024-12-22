package utils

import (
	"math/rand"
	"strings"
	"time"
)

var (
	chars      string   = "absdefghklmnoprst"
	maxLen     int      = 10
	lenStr              = len(chars)
	currencies []string = []string{EUR, USD}
)

const (
	EUR = "EUR"
	USD = "USD"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func MakeRandomName() string {
	var sb strings.Builder
	for i := 0; i < maxLen; i++ {
		sb.WriteByte(chars[rand.Intn(lenStr)])
	}
	return sb.String()
}

func MakeRandomInt64() int64 {
	return rand.Int63n(10)
}

func MakeRandonCurrency() string {
	return currencies[rand.Intn(2)]
}
