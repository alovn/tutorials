package iterator

//容器接口
type IAggregate interface {
	Iterator() Iterator
}

//迭代器接口
type Iterator interface {
	HasNext() bool
	Current() interface{}
}

//一个容器
type MyAggregate struct {
	container []int //容器中装载int类型
}

//容器的迭代器
func (a *MyAggregate) Iterator() Iterator {
	return &MyIterator{aggregate: a}
}

//实现一个具体的迭代器
type MyIterator struct {
	cursor    int          // 当前的游标
	aggregate *MyAggregate // 对应的容器
}

// 判斷是否迭代到最后，如果沒有，則返回true
func (i *MyIterator) HasNext() bool {
	return i.cursor < len(i.aggregate.container)
}

// 获取当前迭代元素（榕容器中取出当前游标对应的元素）
func (i *MyIterator) Current() interface{} {
	current := i.aggregate.container[i.cursor]
	i.next() //游标指向下一个
	return current
}

// 将游标指向下一元素
func (i *MyIterator) next() {
	if i.cursor < len(i.aggregate.container) {
		i.cursor++
	}
}
