package utilities

import (
	"fmt"
	"strings"

	"github.com/sevens7xix/mrf/internal/model"
)

func ProcessTitleTracks[T model.ProcessedItem](items []T) string {
	var builder strings.Builder
	builder.WriteString(strings.Repeat("=", 23))
	builder.WriteString("TRACKLIST")
	builder.WriteString(strings.Repeat("=", 23))
	builder.WriteString("\n")
	for _, item := range items {
		builder.WriteString(fmt.Sprintf("- %s.\n", item.GetName()))
	}
	return builder.String()
}
