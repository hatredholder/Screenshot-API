// most of the code is taken from:
//
// https://github.com/chromedp/examples
package service

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"

	"github.com/chromedp/chromedp"
)

// ScreenshotWebsite takes the url, captures the screen
// and returns the file name
func ScreenshotWebsite(url string) (string, error) {
	// validate the url
	url = validateUrl(url)

	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var buf []byte

	hash := generateRandomHash()
	fileName := fmt.Sprintf("%s.png", hash)

	// capture entire browser viewport, returning png
	if err := chromedp.Run(ctx, fullScreenshot(url, 100, &buf)); err != nil {
		return "", err
	}

	if err := os.WriteFile(fmt.Sprintf("/tmp/%s", fileName), buf, 0o644); err != nil {
		return "", err
	}

	log.SetPrefix("[INFO] ")
	log.Printf("wrote %s", fileName)
	return fileName, nil
}

// fullScreenshot takes a screenshot of the entire browser viewport
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}

// validateUrl takes the url and returns a valid one
func validateUrl(url string) string {
	var validUrl = regexp.MustCompile("^(http|https)://")

	if !validUrl.MatchString(url) {
		url = "http://" + url
	}

	return url
}

// generateRandomHash returns a random 6 character string
func generateRandomHash() string {
	var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	s := make([]rune, 6)
	for i := range s {
		s[i] = chars[rand.Intn(len(chars))]
	}

	return string(s)
}
