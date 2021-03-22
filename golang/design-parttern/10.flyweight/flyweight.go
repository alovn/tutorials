package flyweight

import (
	"sync"
)

type Person struct {
	Name string
}
type Cache struct {
	cache sync.Map
}

func (c *Cache) Get(name string) *Person {
	if p, ok := c.cache.Load(name); ok {
		return p.(*Person)
	}
	p := &Person{Name: name}
	c.cache.Store(name, p)
	return p
}
