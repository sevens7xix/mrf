package mrf

import (
	"fmt"
	"log"
)

func GetTracks(bestFlag bool, printFlag bool, args []string) {
	service, err := GetService(args, bestFlag)

	if err != nil {
		log.Fatalf("Something happened geting the cli Service '%s'", err)
		return
	}

	album_names, err := service.GetProcessedItems()

	if err != nil {
		log.Fatalf("Something happened parsing the items result '%s'", err)
		return
	}

	fmt.Print(album_names)
}

/*
func PrintTracks(w io.Writer, request string) {

}
*/
