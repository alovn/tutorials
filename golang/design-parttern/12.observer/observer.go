package observer

import "fmt"

type Observer interface {
	OnChange(Subject)
	Subscribe(*Subject)
}

type Subject struct {
	observers []Observer
	state     int
}

func (sub *Subject) Publish(state int) {
	sub.state = state
	sub.notifyAllObservers()
}
func (sub *Subject) Attach(o Observer) {
	sub.observers = append(sub.observers, o)
}
func (sub Subject) notifyAllObservers() {
	fmt.Println("notify all")
	for _, observer := range sub.observers {
		observer.OnChange(sub)
	}
}

type AObserver struct {
}

func (o AObserver) OnChange(sub Subject) {
	fmt.Println("AObserver receive:", sub.state)
}
func (o AObserver) Subscribe(s *Subject) {
	s.Attach(o)
}

type BObserver struct {
}

func (o BObserver) OnChange(sub Subject) {
	fmt.Println("BObserver receive:", sub.state)
}

func (o BObserver) Subscribe(s *Subject) {
	s.Attach(o)
}
