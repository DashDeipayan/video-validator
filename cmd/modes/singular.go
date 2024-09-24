package modes

import (
	"log/slog"
	"net/url"

	"github.com/video-validator/internal/hls"
	"github.com/video-validator/internal/models"
)

func StartSingular() {
	streamUrl := "https://stream-akamai.castr.com/5b9352dbda7b8c769937e459/live_2361c920455111ea85db6911fe397b9e/index.fmp4.m3u8"

	hls.ProcessHLS(streamUrl)

	masterManifest := hls.ProcessMasterManifest(streamUrl)

	mediaURL := masterManifest.MasterPlaylist.Variants[0].URI

	parsedURL, err := url.Parse(mediaURL)
	var mediaManifest models.PlaylistResult
	if err != nil {
		mediaManifest = hls.ProcessMediaManifest(mediaURL)
		slog.Info("error", slog.Any("err", mediaManifest))
		return
	} else {
		mediaURL = parsedURL.Scheme + "://" + parsedURL.Host + parsedURL.Path
		mediaManifest = hls.ProcessMediaManifest(mediaURL)
	}

	slog.Info("error", slog.Any("err", parsedURL))

	slog.Info("error", slog.Any("err", masterManifest))
	slog.Info("error", slog.Any("err", mediaManifest))
}
