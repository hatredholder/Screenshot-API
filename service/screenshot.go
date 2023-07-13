package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/hatredholder/Screenshot-API/common/helpers"
)

// ScreenshotWebsite takes the url, captures the screen
// and returns the file name
func ScreenshotWebsite(url string) (string, error) {
	// validate the url
	url = helpers.ValidateUrl(url)

	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	var buf []byte

	hash := helpers.GenerateRandomHash()
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
