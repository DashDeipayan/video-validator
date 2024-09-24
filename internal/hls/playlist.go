package hls

import (
	"errors"
	"fmt"

	"github.com/grafov/m3u8"
	"github.com/video-validator/internal/models"
)

func GetList(listType m3u8.ListType, playlist m3u8.Playlist) models.PlaylistResult {
	switch listType {
	case m3u8.MASTER:
		masterPlaylist := playlist.(*m3u8.MasterPlaylist)
		fmt.Println("Master playlist:")
		for _, variant := range masterPlaylist.Variants {
			fmt.Printf("Variant URI: %s, Bandwidth: %d\n", variant.URI, variant.Bandwidth)
		}
		return models.PlaylistResult{MasterPlaylist: masterPlaylist}
	case m3u8.MEDIA:
		mediaPlaylist := playlist.(*m3u8.MediaPlaylist)
		fmt.Println("This is a media playlist, not a master playlist.")
		return models.PlaylistResult{MediaPlaylist: mediaPlaylist}
	default:
		fmt.Println("Unknown playlist type.")
		return models.PlaylistResult{Err: errors.New("unknown playlist type")}
	}
}
