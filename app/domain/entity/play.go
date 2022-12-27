package entity

import (
	"github.com/roguehedgehog/metric/domain/importdata"
	"time"
)

type Play struct {
	Timestamp time.Time
	MsPlayed  int
	Shuffle   bool
	Skipped   bool
	Track     Track
}

func PlayFrom(lh importdata.LongHistory) Play {
	return Play{
		Timestamp: lh.Timestamp,
		MsPlayed:  lh.MsPlayed,
		Shuffle:   lh.Shuffle,
		Skipped:   lh.Skipped,
		Track:     TrackFrom(lh),
	}
}
