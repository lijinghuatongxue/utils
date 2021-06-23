package test

import (
	"github.com/lijinghuatongxue/utils"
	"testing"
)

func TestGetIPV4IpICMP(t *testing.T) {
	DomainName := "google.com"
	//DomainName := "8.8.8.8"
	utils.GetIpv4IpICMP(DomainName)
}
