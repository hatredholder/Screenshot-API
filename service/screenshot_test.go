package service_test

import (
	"testing"

	"github.com/hatredholder/Screenshot-API/service"
)

func TestScreenshotRealWebsite(T *testing.T) {
	_, err := service.ScreenshotWebsite("google.com")
	if err != nil {
		T.Error(err)
	}
}

func TestScreenshotFakeWebsite(T *testing.T) {
	_, err := service.ScreenshotWebsite("anActuallyFakeWebsite.com")

	got := err.Error()
	want := "page load error net::ERR_NAME_NOT_RESOLVED"

	if got != want {
		T.Error(err)
	}
}
