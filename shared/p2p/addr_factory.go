package p2p

import (
	"github.com/libp2p/go-libp2p/config"
	ma "github.com/multiformats/go-multiaddr"
)

// relayAddrsOnly returns an AddrFactory which will only return Multiaddr via
// specified relay string.
func relayAddrsOnly(relay string) config.AddrsFactory {
	return func(addrs []ma.Multiaddr) []ma.Multiaddr {
		if relay == "" {
			return addrs
		}

		var relayAddrs []ma.Multiaddr

		for _, a := range addrs {
			if a.String() == "/p2p-circuit" {
				continue
			}
			relayAddr, err := ma.NewMultiaddr(relay + "/p2p-circuit" + a.String())
			if err != nil {
				log.Errorf("Failed to create multiaddress for relay node: %v", err)
			} else {
				relayAddrs = append(relayAddrs, relayAddr)
			}
		}

		if len(relayAddrs) == 0 {
			log.Warn("Addresses via relay node are zero. Using non-relay addresses")
			return addrs
		}
		return relayAddrs
	}
}
