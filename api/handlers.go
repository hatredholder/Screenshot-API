package api

import (
	"fmt"
	"net/http"

	"github.com/hatredholder/Screenshot-API/common/helpers"
	"github.com/hatredholder/Screenshot-API/common/types"
	"github.com/hatredholder/Screenshot-API/service"
)

// StorageHandler returns FileServer handler
// with custom filesystem and /tmp as it's directory
func StorageHandler() http.Handler {
	return http.StripPrefix("/storage/", http.FileServer(
		helpers.NoListFileSystem{Fsys: http.Dir("/tmp")},
	))
}

// ScreenshotWebsiteHandler is a handler function
// with the main service of the API
func ScreenshotWebsiteHandler(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Query().Get("url")

	fileName, err := service.ScreenshotWebsite(url)
	if err != nil {
		errResp := types.ErrorResponse{
			Error: err.Error(),
		}
		helpers.WriteJSON(w, http.StatusUnprocessableEntity, &errResp)
		return
	}

	screenshotResp := types.ScreenshotResponse{
		ScreenshotURL: fmt.Sprint(r.Host + "/storage/" + fileName),
	}

	helpers.WriteJSON(w, http.StatusOK, &screenshotResp)
}
