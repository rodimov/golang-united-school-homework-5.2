package cache

import "time"

type Cache struct {
	data      map[string]string
	deadlines map[string]time.Time
}

func NewCache() Cache {
	return Cache{make(map[string]string),
		map[string]time.Time{}}
}

func (c *Cache) Get(key string) (string, bool) {
	if val, ok := c.data[key]; ok {
		deadline, isHasDeadline := c.deadlines[key]
		if isHasDeadline && time.Now().After(deadline) {
			delete(c.data, key)
			delete(c.deadlines, key)
		} else {
			return val, true
		}
	}

	return "", false
}

func (c *Cache) Put(key, value string) {
	_, isHasDeadline := c.deadlines[key]

	if isHasDeadline {
		delete(c.deadlines, key)
	}

	c.data[key] = value
}

func (c *Cache) Keys() []string {
	var keys []string

	for key, _ := range c.data {
		deadline, isHasDeadline := c.deadlines[key]
		if !isHasDeadline || time.Now().Before(deadline) {
			keys = append(keys, key)
		}
	}

	return keys
}

func (c *Cache) PutTill(key, value string, deadline time.Time) {
	c.data[key] = value
	c.deadlines[key] = deadline
}
