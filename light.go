package lightserver

import (
	"fmt"
	"net/http"

	"github.com/GeoinovaDev/lower/exec"
)

type LightServer struct {
	onConfig func(w http.ResponseWriter, r *http.Request) string
}

type lightServerRoute struct {
	route    string
	onBefore func(w http.ResponseWriter, r *http.Request)
	onAfter  func(w http.ResponseWriter, r *http.Request)
	onConfig func(w http.ResponseWriter, r *http.Request)
}

func New() *LightServer {
	return &LightServer{nil}
}

func (s *LightServer) CreateRoute(route string) lightServerRoute {
	return lightServerRoute{route, nil, nil, nil}
}

func (s *lightServerRoute) OnGet(route string, handler func(QueryString) string) *lightServerRoute {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		qs := QueryString{r.URL.Query()}

		defer r.Body.Close()

		exec.Try(func() {
			s.invoke(w, r, s.onBefore)
			s.invoke(w, r, s.onConfig)

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-requested-with")

			text := handler(qs)
			s.invoke(w, r, s.onAfter)

			fmt.Fprint(w, text)
		})
	})

	return s
}

func (s *lightServerRoute) OnPost(route string, handler func(QueryString) string) *lightServerRoute {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		qs := QueryString{r.URL.Query()}

		defer r.Body.Close()

		exec.Try(func() {
			s.invoke(w, r, s.onBefore)
			s.invoke(w, r, s.onConfig)

			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, x-requested-with")

			text := handler(qs)
			s.invoke(w, r, s.onAfter)

			fmt.Fprint(w, text)
		})
	})

	return s
}

func (s *lightServerRoute) invoke(w http.ResponseWriter, r *http.Request, fn func(w http.ResponseWriter, r *http.Request)) {
	if fn == nil {
		return
	}

	fn(w, r)
}
