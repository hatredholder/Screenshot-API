package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/hatredholder/Screenshot-API/api"
)

func TestStorageHandlerStatusCode(t *testing.T) {
	req, err := http.NewRequest("GET", "/storage/", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.Handler(api.StorageHandler())

	handler.ServeHTTP(rr, req)

	want := http.StatusNotFound
	got := rr.Code

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestScreenshotWebsiteHandlerBodySuccess(t *testing.T) {
	req, err := http.NewRequest("GET", "/storage/?url=google.com", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.ScreenshotWebsiteHandler)

	handler.ServeHTTP(rr, req)

	if !strings.Contains(rr.Body.String(), "screenshotUrl") {
		t.Errorf(
			`expected: %s to contain: screenshotUrl`,
			rr.Body.String(),
		)
	}
}

func TestScreenshotWebsiteHandlerBodyFailure(t *testing.T) {
	req, err := http.NewRequest("GET", "/storage/?url=google.com.fake", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(api.ScreenshotWebsiteHandler)

	handler.ServeHTTP(rr, req)

	if !strings.Contains(rr.Body.String(), "error") {
		t.Errorf(
			`expected: %s to contain: error`,
			rr.Body.String(),
		)
	}
}
