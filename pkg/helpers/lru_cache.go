//package lru_cache
// run with: go run lru_cache.go
// run with: go test -v lru_cache.go
// race detector: CGO_ENABLED=1 go run -race --cover lru_cache.go
// go doc lru_cache.go
// CGO_ENABLED=1 go test --race --cover lru_cache.go

package main

import (
	"errors"
	"log"
)

type Cache struct {
	cache map[int64]int64
	queue []int64
	size  int
}

type LRU interface {
	Get(idx int64) (int64, error)
	Set(idx int64, value int64) error
	Evict() error
}

// New function creates a new cache with a size of 3
// Cache created like this contains a map with double linked list
// which allows to evict the least recently used (stale) index or key
func New(size int) (*Cache, error) {
	return &Cache{
		cache: make(map[int64]int64),
		queue: make([]int64, 0),
		size:  size,
	}, nil
}

func (c *Cache) Get(idx int64) (int64, error) {
	if val, ok := c.cache[idx]; ok {
		c.queue = append(c.queue, idx) // add new idx to the end of the queue
		return val, nil
	}
	return 0.00, errors.New("no new idx added to the end of the queue")
}

func (c *Cache) Set(idx int64, value int64) error {

	c.queue = append(c.queue, idx) // append the new idx to the end of the queue
	c.cache[idx] = value           // add the new idx to the cache
	log.Printf("queue: %v, cwhache: %v", c.queue, c.cache)
	c.Evict()
	return nil
}

func (c *Cache) Evict() error {
	if len(c.queue) == c.size && len(c.cache) == c.size { // if the queue and cache are full
		oldestIdx := c.queue[len(c.queue)-1] // naturally logically index of a candidate for eviction
		if _, ok := c.cache[oldestIdx]; ok { // delete the oldest idx from the cache
			// _ = delete(c.cache, oldestIdx)
			c.queue = c.queue[len(c.cache)-1:]
		}
		return errors.New("no oldest idx found")
	}

	return nil
}

func main() {
	cache, err := New(3)
	if err != nil {
		log.Fatal(err)
	}
	cache.Set(1, 1)
	cache.Set(2, 2)
	cache.Set(3, 3)
	cache.Get(1)
	cache.Get(2)
	cache.Get(3)
}
