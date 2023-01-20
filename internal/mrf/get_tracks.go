package mrf

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/sevens7xix/mrf/internal/utilities"
)

func GetTracks(bestFlag bool, printFlag bool, args []string) {
	service, err := GetService(args, bestFlag)

	if err != nil {
		log.Fatalf("Something happened geting the cli Service '%s'", err)
		return
	}

	tracklist, err := service.GetProcessedItems()

	if err != nil {
		log.Fatalf("Something happened parsing the items result '%s'", err)
		return
	}

	if printFlag {
		f, err := os.Create(fmt.Sprintf("%s.txt", strings.Join(args, "_")))

		if err != nil {
			if err != nil {
				log.Fatalf("Something wrong happened creating the file '%s'", err)
				return
			}
		}

		utilities.PrintTracks(f, tracklist)
	} else {
		utilities.PrintTracks(os.Stdout, tracklist)
	}
}
