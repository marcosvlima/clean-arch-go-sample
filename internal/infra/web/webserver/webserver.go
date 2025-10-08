package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router chi.Router
	// Handlers[path][method] = handler. method can be HTTP verb (GET/POST/...) or "ANY" for all methods
	Handlers      map[string]map[string]http.HandlerFunc
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make(map[string]map[string]http.HandlerFunc),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc) {
	// Register handler for any method (backwards compatible)
	if s.Handlers[path] == nil {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[path]["ANY"] = handler
}

// AddHandlerMethod registers a handler for a specific HTTP method (e.g. "GET", "POST")
func (s *WebServer) AddHandlerMethod(path string, method string, handler http.HandlerFunc) {
	if s.Handlers[path] == nil {
		s.Handlers[path] = make(map[string]http.HandlerFunc)
	}
	s.Handlers[path][method] = handler
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for path, methods := range s.Handlers {
		for method, handler := range methods {
			if method == "ANY" {
				s.Router.Handle(path, handler)
			} else {
				s.Router.Method(method, path, handler)
			}
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
