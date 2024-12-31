package pilgrim

import (
	"fmt"
	"net/http"
	"time"
)

func BasicLogger() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			startTime := time.Now()
			next.ServeHTTP(w, r)
			elapsedTime := time.Since(startTime)
			fmt.Printf("\033[1;30mINFO: \033[1;31m%s %s\033[1;30m in \033[1;36m%s\033[0m\n",
				r.Method, r.URL.Path, elapsedTime.String(),
			)
		})
	}
}
