package helpers_test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/hatredholder/Screenshot-API/common/helpers"
	"github.com/hatredholder/Screenshot-API/common/types"
)

func TestWriteJSONStatusCode(t *testing.T) {
	rr := httptest.NewRecorder()
	v := types.ScreenshotResponse{
		ScreenshotURL: "test",
	}

	helpers.WriteJSON(rr, http.StatusOK, v)

	want := http.StatusOK
	got := rr.Code

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestWriteJSONBody(t *testing.T) {
	rr := httptest.NewRecorder()
	v := types.ScreenshotResponse{
		ScreenshotURL: "test",
	}

	helpers.WriteJSON(rr, http.StatusOK, v)

	// WriteJSON adds a newline character
	want := `{"screenshotUrl":"test"}` + "\n"
	got := rr.Body.String()

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestValidateUrlNoHttpPrefix(t *testing.T) {
	want := "http://google.com"
	got := helpers.ValidateUrl("google.com")

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestValidateUrlWithHttpsPrefix(t *testing.T) {
	want := "https://google.com"
	got := helpers.ValidateUrl("https://google.com")

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestGenerateRandomHashLen(t *testing.T) {
	hash := helpers.GenerateRandomHash()

	want := 6
	got := len(hash)

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestNoListFileSystemOpenDir(t *testing.T) {
	noListFs := helpers.NoListFileSystem{Fsys: http.Dir("/tmp")}
	_, err := noListFs.Open(".")

	want := os.ErrNotExist
	got := err

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}

func TestNoListFileSystemOpenFakeFile(t *testing.T) {
	noListFs := helpers.NoListFileSystem{Fsys: http.Dir("/tmp")}
	_, err := noListFs.Open("fakefile")

	if err == nil {
		t.Error("no error was returned for fake file")
	}
}
