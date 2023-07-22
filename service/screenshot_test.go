package service_test

import (
	"testing"

	"github.com/hatredholder/Screenshot-API/service"
)

func TestScreenshotWebsiteSuccess(t *testing.T) {
	_, err := service.ScreenshotWebsite("google.com")
	if err != nil {
		t.Error(err)
	}
}

func TestScreenshotWebsiteFailure(t *testing.T) {
	_, err := service.ScreenshotWebsite("google.com.fake")
	if err == nil {
		t.Fatal("no error was returned for fake website")
	}

	got := err.Error()
	want := "page load error net::ERR_NAME_NOT_RESOLVED"

	if want != got {
		t.Errorf("expected: %v, got: %v", want, got)
	}
}
