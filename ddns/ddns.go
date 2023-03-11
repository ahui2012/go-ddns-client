package ddns

import (
	"log"

	"github.com/ahui2012/go-ddns-client/config"
	"github.com/ahui2012/go-ddns-client/ddns/tencent"
)

type DNSProvider interface {
	Init()
	UpdateRecord(ip string) error
}

var dnsProviders []DNSProvider

func Init(domains []config.DomainConfig) {
	dnsProviders = []DNSProvider{}
	for _, domain := range domains {
		if domain.Provider == "tencent" {
			var provider = &tencent.TencentCloud{Config: domain}
			provider.Init()
			dnsProviders = append(dnsProviders, provider)
		} else {
			panic("not support dns service provider: " + domain.Provider)
		}
	}
}

func Update(ip string) {
	for _, provider := range dnsProviders {
		if err := provider.UpdateRecord(ip); err != nil {
			log.Println(err)
		}
	}
}
