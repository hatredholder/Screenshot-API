package api

import (
	"net/http"
)

func NewServeMux() *http.ServeMux {
	sm := http.NewServeMux()

	sm.Handle("/storage/", StorageHandler())
	sm.HandleFunc("/screenshot/", ScreenshotWebsiteHandler)

	return sm
}

func NewServer(addr string) *http.Server {
	sm := NewServeMux()

	return &http.Server{
		Addr:    addr,
		Handler: sm,
	}
}
