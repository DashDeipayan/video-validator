package hls

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/grafov/m3u8"
	"github.com/video-validator/internal/helper"
)

func ProcessManifest(url string) (m3u8.Playlist, m3u8.ListType, error) {
	resp, err := helper.HttpClient.Get(url)
	if err != nil {
		fmt.Println("Failed to fetch the M3U8 URL:", err)
		return nil, 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Failed to fetch the M3U8 URL: HTTP status", resp.StatusCode)
		return nil, 0, fmt.Errorf("failed to fetch the M3U8 URL: HTTP status %d", resp.StatusCode)
	}

	playlist, listType, err := m3u8.DecodeFrom(resp.Body, true)
	if err != nil {
		fmt.Println("Failed to parse M3U8 content:", err)
		return nil, 0, err
	}
	return playlist, listType, nil
}

func ProcessHLS(hlsUrl string) {
	parsedUrl, err := url.Parse(hlsUrl)

	if err != nil {
		fmt.Println("Failed to parse the HLS URL:", err)
	}
	fmt.Println(parsedUrl)
}
