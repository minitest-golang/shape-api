package api

import (
	"fmt"
	v1 "minitest/api/v1"
	"minitest/utils"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

// Use this function to enable the CORS.
// Only use this for debugger.
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "*")
	(*w).Header().Set("Access-Control-Allow-Headers", "*")
}

func CORS(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// A http server using Gorilla Mux
type GMuxServer struct {
	router *mux.Router
}

// Return a Gorilla Mux Server
func NewGMuxServer() *GMuxServer {
	return &GMuxServer{
		router: mux.NewRouter(),
	}
}

// A wrapper for pre-process client requests before actual serving
func (m *GMuxServer) midleware(r *http.Request) {
	// ToDo: Pre-handle any incoming request
}

// A wrapper of HTTP handler route.
// This function return a handler wrapper in which the middleware function and/or
// the CORS function are called. After that the real user handler shall  be called
// to handle user's request.
func (m *GMuxServer) route(handler HandleFunc) HandleFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		m.midleware(r)
		if utils.EnableCORS {
			enableCors(&rw)
		}
		if r.Method == "OPTIONS" {
			if utils.EnableCORS {
				CORS(rw, r)
			}
			return
		}
		handler(rw, r)
	}
}

// Call this function to start a HTTP server.
// This function must be called after Route API.
func (m *GMuxServer) Start(address string) error {
	utils.InfoLog("Start server at: %s.\n", address)
	return http.ListenAndServe(address, m.router)
}

// Use this API to route HTTP requests
func (m *GMuxServer) Route(path, method string, handler HandleFunc) {
	if utils.EnableCORS {
		m.router.HandleFunc(path, m.route(handler)).Methods(method, "OPTIONS")
	} else {
		m.router.HandleFunc(path, m.route(handler)).Methods(method)
	}
}

// Just for swagger web
// Call this function if we want to expose our API via swagger Web.
func (m *GMuxServer) SwaggerRoute(path string) {
	m.router.PathPrefix(path).Handler(httpSwagger.WrapHandler)
}

func CreateRestApis() {
	server := NewGMuxServer()

	// http://localhost:8081/app/v1/swagger/index.html
	server.SwaggerRoute("/app/v1/swagger")

	// User APIs
	server.Route("/app/v1/user/signup", "POST", v1.UserSignUpApiHandler)
	server.Route("/app/v1/user/login", "POST", v1.UserLoginApiHandler)

	// Shape APIs
	server.Route("/app/v1/shape/create", "POST", v1.CreateShapeApiHandler)
	server.Route("/app/v1/shape", "GET", v1.GetAllShapeApiHandler)
	server.Route("/app/v1/shape/{shape_id}", "GET", v1.GetShapeApiHandler)
	server.Route("/app/v1/shape/calculate", "POST", v1.CalculateShapeApiHandler)
	server.Route("/app/v1/shape/{shape_id}", "PUT", v1.UpdateShapeApiHandler)
	server.Route("/app/v1/shape/{shape_id}", "DELETE", v1.DeleteShapeApiHandler)

	// Start server now
	server.Start(fmt.Sprintf(":%s", os.Getenv("SERVER_PORT")))
}
