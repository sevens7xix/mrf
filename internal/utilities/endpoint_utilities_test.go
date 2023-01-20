package utilities

import (
	"testing"

	"github.com/sevens7xix/mrf/internal/model"
)

func TestProcessTitleTracks(t *testing.T) {
	AssertEqual(t,
		ProcessTitleTracks([]model.MockProcessedItem{}),
		"=======================TRACKLIST=======================\n")
}
