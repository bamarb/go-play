package netx

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/mccutchen/go-httpbin/v2/httpbin"
)

func TestSlowServer(t *testing.T) {
	htb := httpbin.New()
	testServer := httptest.NewServer(htb)
	defer testServer.Close()
	t.Logf("starting http get with delay %d, %s\n", 10, time.Now())
	_, err := http.Get(testServer.URL + "/delay/10")
	t.Logf("received response at %s , Err:%s\n", time.Now(), err)
}
