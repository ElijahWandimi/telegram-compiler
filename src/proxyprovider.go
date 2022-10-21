package src


import (
	"github.com/oyamo/telegram-compiler/assets"
	"math/rand"
)

func NextProxy() assets.Proxy {
	return assets.PROXY_SERVERS[rand.Intn(len(assets.PROXY_SERVERS))]
}

