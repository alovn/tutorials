package flyweight

import (
	"fmt"
	"testing"
)

func TestCache_Get(t *testing.T) {
	cache := &Cache{}
	p1 := cache.Get("zhangshan")
	p2 := cache.Get("zhangshan")
	p3 := cache.Get("zhangshan")
	fmt.Printf("%p, %+v\n", p1, p1)
	fmt.Printf("%p, %+v\n", p2, p2)
	fmt.Printf("%p, %+v\n", p3, p3)
}
