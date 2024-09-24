package models

import "github.com/grafov/m3u8"

type PlaylistResult struct {
	MasterPlaylist *m3u8.MasterPlaylist
	MediaPlaylist  *m3u8.MediaPlaylist
	Err            error
}
