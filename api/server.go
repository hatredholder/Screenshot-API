package api

import (
	"fmt"
	"net/http"
)

type Server struct {
	listenAddr string
}

func NewServer(listenAddr string) *Server {
	return &Server{
		listenAddr: listenAddr,
	}
}

func (s *Server) Start() error {
	http.Handle("/storage/", StorageHandler())
	http.HandleFunc("/screenshot/", ScreenshotWebsiteHandler)

	fmt.Printf("[INFO] Server started on port: \"%s\"\n", s.listenAddr)

	return http.ListenAndServe(s.listenAddr, nil)
}
