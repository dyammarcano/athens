package external

import (
	"net/http/httptest"
	"testing"

	"github.com/gomods/athens/internal/storage/compliance"
	"github.com/gomods/athens/internal/storage/mem"
)

func TestExternal(t *testing.T) {
	strg, err := mem.NewStorage()
	if err != nil {
		t.Fatal(err)
	}
	handler := NewServer(strg)
	srv := httptest.NewServer(handler)
	defer srv.Close()
	externalStrg := NewClient(srv.URL, nil)
	clear := strg.(interface{ Clear() error }).Clear
	compliance.RunTests(t, externalStrg, clear)
}
