package tests

import (
	"fmt"
	"testing"
)

type Data struct {
	id  string
	val int
}

// immutable map
var im map[string]*Data = map[string]*Data{
	"123": {"ajklhdslah", 1},
	"345": {"ajqshdslah", 1},
	"567": {"ajhdslahit", 1},
	"890": {"ajklhdshit", 1},
	"012": {"ajklhdslah", 1},
	"001": {"ajklhdslah", 1},
}

func getMapKey(appID int, partnerID int, tagID string, adType string) string {
	return fmt.Sprintf("%d#$#%d#$#%s#$#%s", appID, partnerID, tagID, adType)
}

func TestGetMapKey(t *testing.T) {
	tstr := getMapKey(1, 1, "Home_Screen", "banner")
	t.Logf("%s\n", tstr)
}
