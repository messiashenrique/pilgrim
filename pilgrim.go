package pilgrim

import (
	"net/http"
	"slices"
)

type (
	Middleware func(next http.Handler) http.Handler
	Router     struct {
		*http.ServeMux
		set []Middleware
	}
)

func NewRouter(md ...Middleware) *Router {
	return &Router{ServeMux: &http.ServeMux{}, set: md}
}

func (r *Router) Use(md ...Middleware) {
	r.set = append(r.set, md...)
}

func (r *Router) RoutesGroup(fn func(r *Router)) {
	fn(&Router{ServeMux: r.ServeMux, set: slices.Clone(r.set)})
}

func (r *Router) HandleFunc(path string, fn http.HandlerFunc) {
	r.Handle(path, r.resolve(fn))
}

func (r *Router) resolve(fn http.HandlerFunc) (out http.Handler) {
	mdx := make([]Middleware, 0)
	mdx = append(slices.Clone(r.set), mdx...)
	out = http.Handler(fn)

	slices.Reverse(mdx)

	for _, m := range mdx {
		out = m(out)
	}
	return
}
