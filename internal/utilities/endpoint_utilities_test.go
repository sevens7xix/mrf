package utilities

import (
	"bytes"
	"testing"

	"github.com/sevens7xix/mrf/internal/model"
)

func TestProcessTitleTracks(t *testing.T) {
	AssertEqual(t,
		ProcessTitleTracks([]model.MockProcessedItem{}),
		"=======================TRACKLIST=======================\n")
}

func TestPrintTracks(t *testing.T) {
	buffer := bytes.Buffer{}

	PrintTracks(&buffer, "a lot of tracks")

	got := buffer.String()
	wanted := "a lot of tracks"

	AssertEqual(t, got, wanted)
}
