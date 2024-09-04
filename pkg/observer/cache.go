package observer

import (
	"fmt"

	"github.com/cilium/tetragon/pkg/api/dataapi"
	lru "github.com/hashicorp/golang-lru/v2"
)

type Cache struct {
	cache *lru.Cache[dataapi.DataEventId, *data]
	size  int
}

type data []byte

func NewCache(
	dataCacheSize int,
) (*Cache, error) {
	lruCache, err := lru.NewWithEvict(
		dataCacheSize,
		func(_ dataapi.DataEventId, _ *data) {
			dataCacheEvictions.Inc()
		},
	)
	if err != nil {
		return nil, err
	}
	cache := &Cache{
		cache: lruCache,
		size:  dataCacheSize,
	}
	// FIXME: inspect garbage collector
	// pm.cacheGarbageCollector()
	return cache, nil
}

func (c *Cache) get(dataEventId dataapi.DataEventId) (*data, error) {
	data, ok := c.cache.Get(dataEventId)
	if !ok {
		dataCacheMisses.WithLabelValues("get").Inc()
		// FIXME
		return nil, fmt.Errorf("invalid entry for data event ID: %v", dataEventId)
	}
	return data, nil
}

func (c *Cache) add(id dataapi.DataEventId, msgData *data) bool {
	evicted := c.cache.Add(id, msgData)
	if !evicted {
		dataCacheTotal.Inc()
	}
	return evicted
}

func (c *Cache) remove(desc dataapi.DataEventDesc) bool {
	present := c.cache.Remove(desc.Id)
	if present {
		dataCacheTotal.Dec()
	} else {
		dataCacheMisses.WithLabelValues("remove").Inc()
	}
	return present
}
