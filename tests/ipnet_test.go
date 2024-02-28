package tests

import (
	"net"
	"testing"
)

func TestIpCIDR(t *testing.T) {
	ipSingle := "15.10.0.1/32"
	ip := net.ParseIP("15.10.0.1")
	i, n, err := net.ParseCIDR(ipSingle)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Ip: %s, IpNet: %s \n", i.String(), n.String())
	if !n.Contains(ip) {
		t.Fatal("Expected Match")
	}
}
