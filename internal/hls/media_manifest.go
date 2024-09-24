package hls

import (
	"log/slog"

	"github.com/grafov/m3u8"
	"github.com/video-validator/internal/models"
)

func ProcessMediaManifest(url string) models.PlaylistResult {
	playlist, listType, err := ProcessManifest(url)
	if err != nil {
		slog.Error("Failed to process the master manifest:", slog.Any("error", err))
		return models.PlaylistResult{Err: err}
	}
	if listType != m3u8.MEDIA {
		slog.Error("wrong playlist returned:", slog.Any("error", err))
		return models.PlaylistResult{Err: err}
	}

	result := GetList(listType, playlist)
	if result.Err != nil {
		slog.Error("Failed to process the master manifest:", slog.Any("error", result.Err))
		return models.PlaylistResult{Err: result.Err}
	}

	slog.Info("results", slog.Any("res", result))
	return result
}
