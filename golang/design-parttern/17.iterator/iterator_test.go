package iterator

import (
	"fmt"
	"testing"
)

func TestMyIterator_Current(t *testing.T) {
	agreegate := &MyAggregate{container: []int{9, 8, 7, 6, 5, 4, 3, 2}}
	iterator := agreegate.Iterator()
	for iterator.HasNext() {
		current := iterator.Current().(int)
		fmt.Println(current)
	}
}
