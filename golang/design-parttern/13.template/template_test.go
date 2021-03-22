package template

import "testing"

func TestTaskTemplate_Do(t *testing.T) {
	mytask := NewMyTaskTemplate()
	mytask.Do()
}
