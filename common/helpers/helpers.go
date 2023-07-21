package helpers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"os"
	"regexp"
)

// WriteJSON writes header statuscode of sc,
// encodes JSON to ResponseWriter with encoding of v,
// followed by a newline character
func WriteJSON(rw http.ResponseWriter, sc int, v any) {
	rw.WriteHeader(sc)
	json.NewEncoder(rw).Encode(v)
}

// ValidateUrl takes the url and returns a valid one
func ValidateUrl(url string) string {
	var validUrl = regexp.MustCompile("^(http|https)://")

	if !validUrl.MatchString(url) {
		url = "http://" + url
	}

	return url
}

// GenerateRandomHash returns a random 6 character string
func GenerateRandomHash() string {
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, 6)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}

	return string(s)
}

// hide directory listings in http.FileServer using a custom filesystem

type NoListFileSystem struct {
	Fsys http.FileSystem
}

// Open hides directory listings by returning os.ErrNotExist if file is a directory
func (fs NoListFileSystem) Open(name string) (http.File, error) {
	file, err := fs.Fsys.Open(name)
	if err != nil {
		return nil, err
	}

	// check if file is a directory
	info, err := file.Stat()
	if err != nil {
		return nil, err
	}
	if info.IsDir() {
		return nil, os.ErrNotExist
	}

	return file, nil
}
