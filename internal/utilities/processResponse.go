package utilities

import (
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/sevens7xix/mrf/internal/model"
)

func ProcessTitleTracks[T model.ProcessedItem](items []T) string {
	var builder strings.Builder
	var wg sync.WaitGroup

	// Here we write the header of our processed Response
	builder.WriteString(strings.Repeat("=", 23))
	builder.WriteString("TRACKLIST")
	builder.WriteString(strings.Repeat("=", 23))
	builder.WriteString("\n")

	for _, item := range items {
		wg.Add(1)

		go func(item T) {
			defer wg.Done()
			builder.WriteString(fmt.Sprintf("- %s.\n", item.GetName()))
		}(item)
	}

	wg.Wait()

	return builder.String()
}

func PrintTracks(w io.Writer, tracklist string) {
	_, err := w.Write([]byte(tracklist))

	if err != nil {
		log.Fatalf("Error printing the tracklist '%s'", err)
		return
	}
}
