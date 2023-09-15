package server

import (
	contextpkg "context"
	"time"

	"github.com/tliron/glsp"
)

//
// Server
//

type Server struct {
	Handler     glsp.Handler
	LogBaseName string
	Debug       bool

	Log          glsp.Log
	Context      contextpkg.Context
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewServer(handler glsp.Handler, Log glsp.Log, debug bool) *Server {
	return &Server{
		Handler:      handler,
		Debug:        debug,
		Log:          Log,
		Context:      contextpkg.Background(),
		ReadTimeout:  75 * time.Second,
		WriteTimeout: 60 * time.Second,
	}
}
