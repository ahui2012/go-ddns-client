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

var DNSProviders []DNSProvider

func Init(domains []config.DomainConfig) {
	DNSProviders = []DNSProvider{}
	for _, domain := range domains {
		if domain.Provider == "tencent" {
			var provider = &tencent.TencentCloud{Config: domain}
			provider.Init()
			DNSProviders = append(DNSProviders, provider)
		} else {
			panic("not support dns service provider: " + domain.Provider)
		}
	}
}

func Update(ip string) {
	for _, provider := range DNSProviders {
		if err := provider.UpdateRecord(ip); err != nil {
			log.Println(err)
		}
	}
}
