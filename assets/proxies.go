package assets

import (
	"fmt"
)

type Proxy struct {
	Server string
	Port uint32
}

var (
	PROXY_SERVERS = [...] Proxy {
			{"45.77.56.114",	30205},
			{"82.196.11.105",	1080},
			{"51.254.69.243",	3128},
		}
)

func (p *Proxy) FullAddress() string {
	return fmt.Sprintf("http://%s:%d", p.Server, p.Port)
}