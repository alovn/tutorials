package chain

import (
	"testing"
)

func TestRequestChain_Handle(t *testing.T) {
	log := &LogHandler{}
	script := ScriptFilterHandler{NextHandler: log}
	ad := AdFilterHandler{NextHandler: script}

	content := "正常内容，广告内容 script"
	ad.Handle(content)
}
