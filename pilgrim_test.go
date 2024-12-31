package pilgrim

import (
	"net/http"
	"testing"
)

func TestNewRouter(t *testing.T) {
	r := NewRouter()
	if r.ServeMux == nil {
		t.Errorf("ServeMux is nil")
	}
	if len(r.set) != 0 {
		t.Errorf("set is not empty")
	}
}

func TestUse(t *testing.T) {
	r := NewRouter()
	md := func(next http.Handler) http.Handler {
		return next
	}
	r.Use(md)
	if len(r.set) != 1 {
		t.Errorf("set is not updated")
	}
}

func TestRoutesGroup(t *testing.T) {
	r := NewRouter()
	var called bool
	r.RoutesGroup(func(r *Router) {
		called = true
	})
	if !called {
		t.Errorf("function is not called")
	}
}
