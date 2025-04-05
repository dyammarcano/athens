package mem

import (
	"testing"

	"github.com/dyammarcano/athens/internal/index"
	"github.com/dyammarcano/athens/internal/index/compliance"
)

func TestMem(t *testing.T) {
	indexer := &indexer{}
	compliance.RunTests(t, indexer, indexer.clear)
}

func (i *indexer) clear() error {
	i.mu.Lock()
	i.lines = []*index.Line{}
	i.mu.Unlock()
	return nil
}
