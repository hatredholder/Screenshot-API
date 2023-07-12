package helpers

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
)

// GetURLPath returns value after the first "/" symbol
func GetURLPath(r *http.Request) string {
	p := strings.Split(r.URL.Path[1:], "/")

	// get rid of empty strings
	result := []string{}
	for _, i := range p {
		if i != "" {
			result = append(result, string(i))
		}
	}

	if len(result) > 1 {
		return strings.ToLower(result[1]) // converting to lowercase
	}

	return ""
}

// WriteJSON writes JSON to ResponseWriter with encoding of v
func WriteJSON(w http.ResponseWriter, s int, v any) {
	w.WriteHeader(s)
	json.NewEncoder(w).Encode(v)
}

// hide directory listings in http.FileServer using a custom filesystem

type NoListFileSystem struct {
	Fsys http.FileSystem
}

// Open hides directories by returning os.ErrNotExist if file is a directory
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
