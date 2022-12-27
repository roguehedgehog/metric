package importdata

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"testing"
	"time"
)

type LongHistoryTestCase struct {
	json     []byte
	expected []LongHistory
}

func TestLongHistoryImport(t *testing.T) {
	tests := map[string]LongHistoryTestCase{
		"single":   singleEntry(),
		"2 tracks": MultipleEntries(),
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			// When
			im := LongHistoryImporter{}
			c := len(test.expected)
			entries := make(chan LongHistory, c)
			actual := make([]LongHistory, 0, c)

			err := im.extract(context.Background(), entries, test.json)

			close(entries)
			if err != nil {
				t.Error(err)
			}

			// Then
			for entry := range entries {
				actual = append(actual, entry)
			}

			diff := cmp.Diff(
				test.expected,
				actual,
				cmpopts.SortSlices(func(a, b LongHistory) bool { return a.Uri < b.Uri }),
			)
			if diff != "" {
				t.Error(diff)
			}
		})
	}
}

func singleEntry() LongHistoryTestCase {
	return LongHistoryTestCase{
		json: []byte(`[{
			"ts":"2019-12-11T13:08:50Z",
			"username":"fayjxwssg8o18dm1iyy0jl9v1",
			"platform":"Android OS 9 API 28 (samsung, SM-G950F)",
			"ms_played":518293,
			"conn_country":"GB",
			"ip_addr_decrypted":"89.197.12.205",
			"user_agent_decrypted":"unknown",
			"master_metadata_track_name":"Lillo",
			"master_metadata_album_artist_name":"Pryda",
			"master_metadata_album_album_name":"Lillo",
			"spotify_track_uri":"spotify:track:0wIblt271U2oH60h7w8ECA",
			"episode_name":null,
			"episode_show_name":null,
			"spotify_episode_uri":null,
			"reason_start":"trackdone",
			"reason_end":"trackdone",
			"shuffle":false,
			"skipped":null,
			"offline":false,
			"offline_timestamp":1576069210606,
			"incognito_mode":false
		}]`),
		expected: []LongHistory{
			{
				Uri:       "spotify:track:0wIblt271U2oH60h7w8ECA",
				Title:     "Lillo",
				Artist:    "Pryda",
				Album:     "Lillo",
				Timestamp: time.Date(2019, 12, 11, 13, 8, 50, 0, time.UTC),
				MsPlayed:  518293,
				Skipped:   false,
				Shuffle:   false,
			},
		},
	}
}

func MultipleEntries() LongHistoryTestCase {
	return LongHistoryTestCase{
		json: []byte(`[{
    "ts": "2015-01-12T21:27:17Z",
    "username": "!tahir",
    "platform": "Android OS 5.0.1 API 21 (LGE, Nexus 4)",
    "ms_played": 43374,
    "conn_country": "GB",
    "ip_addr_decrypted": "89.168.84.178",
    "user_agent_decrypted": "unknown",
    "master_metadata_track_name": "Something Good Can Work (The Twelves Tabloid dub)",
    "master_metadata_album_artist_name": "Two Door Cinema Club",
    "master_metadata_album_album_name": "Kitsuné Tabloid by The Twelves",
    "spotify_track_uri": "spotify:track:4wbBdRY3h4hU2LW937MOfM",
    "episode_name": null,
    "episode_show_name": null,
    "spotify_episode_uri": null,
    "reason_start": "appload",
    "reason_end": "appload",
    "shuffle": false,
    "skipped": true,
    "offline": false,
    "offline_timestamp": 0,
    "incognito_mode": false
  },
  {
    "ts": "2015-01-05T16:19:52Z",
    "username": "!tahir",
    "platform": "Android OS 5.0.1 API 21 (LGE, Nexus 4)",
    "ms_played": 274797,
    "conn_country": "GB",
    "ip_addr_decrypted": "89.168.84.178",
    "user_agent_decrypted": "unknown",
    "master_metadata_track_name": "Embrace",
    "master_metadata_album_artist_name": "Goldroom",
    "master_metadata_album_album_name": "Embrace",
    "spotify_track_uri": "spotify:track:6G1xBCUqxa18lRgyCPmpW4",
    "episode_name": null,
    "episode_show_name": null,
    "spotify_episode_uri": null,
    "reason_start": "trackdone",
    "reason_end": "trackdone",
    "shuffle": false,
    "skipped": false,
    "offline": false,
    "offline_timestamp": 0,
    "incognito_mode": false
  }
]`),
		expected: []LongHistory{
			{
				Uri:       "spotify:track:6G1xBCUqxa18lRgyCPmpW4",
				Title:     "Embrace",
				Artist:    "Goldroom",
				Album:     "Embrace",
				Timestamp: time.Date(2015, 1, 5, 16, 19, 52, 0, time.UTC),
				MsPlayed:  274797,
				Skipped:   false,
				Shuffle:   false,
			},
			{
				Uri:       "spotify:track:4wbBdRY3h4hU2LW937MOfM",
				Title:     "Something Good Can Work (The Twelves Tabloid dub)",
				Artist:    "Two Door Cinema Club",
				Album:     "Kitsuné Tabloid by The Twelves",
				Timestamp: time.Date(2015, 1, 12, 21, 27, 17, 0, time.UTC),
				MsPlayed:  43374,
				Skipped:   true,
				Shuffle:   false,
			},
		},
	}
}
