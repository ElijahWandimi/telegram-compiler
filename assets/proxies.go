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
			{"191.102.73.108",	999},
			{"38.127.179.19",	55994},
			{"203.154.232.25",	4153},
		}
)

func (p *Proxy) FullAddress() string {
	return fmt.Sprintf("http://%s:%d", p.Server, p.Port)
}