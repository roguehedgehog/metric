package importdata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/roguehedgehog/metric/domain/repo"
	"time"
)

type LongHistory struct {
	Uri    string `json:"spotify_track_uri"`
	Title  string `json:"master_metadata_track_name,omitempty"`
	Artist string `json:"master_metadata_album_artist_name,omitempty"`
	Album  string `json:"master_metadata_album_album_name,omitempty"`

	Timestamp time.Time `json:"ts,omitempty"`
	MsPlayed  int       `json:"ms_played,omitempty"`
	Shuffle   bool      `json:"shuffle,omitempty"`
	Skipped   bool      `json:"skipped,omitempty"`
}

type LongHistoryImporter struct {
	playRepo *repo.Play
}

func (importer *LongHistoryImporter) Process(file ImportFile) error {
	if file.what != ImportTypeLongHistory {
		return fmt.Errorf("cannot import file of type %s as Long History", file.what)
	}

	// Create workers
	// Read file
	// Extract
	// Wait

	return nil
}

func (importer *LongHistoryImporter) extract(ctx context.Context, entries chan<- LongHistory, entry []byte) error {
	dec := json.NewDecoder(bytes.NewReader(entry))
	t, err := dec.Token()
	if err != nil {
		return fmt.Errorf("token %s is not allowed, causes %v", t, err)
	}

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if !dec.More() {
				return nil
			}

			var entry LongHistory
			err = dec.Decode(&entry)
			if err != nil {
				return fmt.Errorf("could not decode track, %v", err)
			}

			entries <- entry
		}
	}
}
