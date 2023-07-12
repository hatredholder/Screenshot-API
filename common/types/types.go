package types

type ScreenshotResponse struct {
	ScreenshotURL string `json:"screenshotUrl"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
