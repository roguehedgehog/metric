package entity

import "github.com/roguehedgehog/metric/domain/importdata"

type Track struct {
	Uri    string
	Title  string
	Artist string
	Album  string
}

func TrackFrom(lh importdata.LongHistory) Track {
	return Track{
		Uri:    lh.Uri,
		Title:  lh.Title,
		Artist: lh.Artist,
		Album:  lh.Album,
	}
}
