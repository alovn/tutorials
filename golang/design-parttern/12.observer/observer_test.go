package observer

import "testing"

func TestBObserver_Update(t *testing.T) {
	subject := &Subject{}
	(&AObserver{}).Subscribe(subject)
	(&BObserver{}).Subscribe(subject)

	subject.Publish(1)
}
