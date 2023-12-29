package hashmap

import (
	"errors"
)

type Server[K comparable, V any] struct {
	table *LinkedHashMap[K, V] 
}

//consistent hashing algortihm using a 'ring', preferred over 'distrbuteMap'
func (lhm *LinkedHashMap[K, V]) ConsistentHash(servers []Server[K, V], ring uint64) error {
	if len(servers) == 0 {
		return errors.New("no servers available for distribution")
	}
	serverDistance := ring / uint64(len(servers))

	var determineServer func(uint64, uint64, uint64) int
	determineServer = func(hash , ring, dist uint64) int {
		initialServer := 1
		
		for i:=int(dist) ; i<int(ring); i+=int(dist) {
			if hash < uint64(i) {
				return initialServer
			}
			initialServer++
		}

		return 0
	}

	for _, bucket := range lhm.buckets {
		for _, kv := range bucket {
      hash := Hasher(bytify(kv.key)) % ring //[0, 360)
			focusedServer := servers[determineServer(hash, serverDistance, ring)]
			focusedServer.table.Insert(kv.key, kv.value)
			continue
		}
	}
	return nil
}


