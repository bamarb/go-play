package randx

import (
	"crypto/rand"
	"errors"
	"math/big"
	mrand "math/rand"
	"time"
)

const (
	Alphabet     = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	Numerals     = "1234567890"
	Alphanumeric = Alphabet + Numerals
	Ascii        = Alphanumeric + "~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"
)

var ErrMinMax = errors.New("min must be strictly less than  max")

func IntBetween(min, max int) int {
	if min > max {
		panic(ErrMinMax.Error())
	}
	if min == max {
		min = 0
	}
	ret, err := rand.Int(rand.Reader, big.NewInt(int64(max-min)))
	if err != nil {
		panic(err)
	}
	return min + int(ret.Int64())
}

func String(n int, charset string) string {
	randStr := make([]byte, n)
	charLen := big.NewInt(int64(len(charset)))
	for i := 0; i < n; i++ {
		r, err := rand.Int(rand.Reader, charLen)
		if err != nil {
			panic(err)
		}
		randStr[i] = charset[int(r.Int64())]
	}
	return string(randStr)
}

func StringBetween(min, max int, charset string) string {
	strlen := IntBetween(min, max)
	return String(strlen, charset)
}

// CoinToss returns a bool interpret as u will
func CoinToss() bool {
	res := [2]bool{true, false}
	return res[mrand.Intn(2)]
}

// SleepBetween sleep between min max millis
func SleepBetween(min, max int) {
	if min == max {
		time.Sleep(time.Duration(min) * time.Millisecond)
	}
	if min < max {
		r := IntBetween(min, max)
		time.Sleep(time.Duration(r) * time.Millisecond)
	}
}

// Shuffle shuffles a slice in-place
func Shuffle[S ~[]E, E any](x S) {
	mrand.Shuffle(len(x), func(i, j int) {
		x[i], x[j] = x[j], x[i]
	})
}
