package prototype

type Job struct {
	Company string
}
type Person struct {
	Name string
	Age  int
	Job  *Job
}

//浅拷贝
func (p *Person) Clone() *Person {
	ret := new(Person)
	*ret = *p
	newJob := *(p).Job
	ret.Job = &newJob
	return ret
}
