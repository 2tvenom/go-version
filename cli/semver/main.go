package main

import (
	"os"
	"fmt"
	"log"
	"github.com/2tvenom/go-version"
)

const (
	GET_RELEASE = "getrelease"
	GET_VERSION = "getversion"
	GET_META = "getmeta"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		fmt.Println("Usage: semver <COMMAND> <VERSION>")
		fmt.Println("Commands:")
		fmt.Println(GET_VERSION + "	Get version X.Y.Z")
		fmt.Println(GET_RELEASE + "	Get version release")
		fmt.Println(GET_META + "		Get version metadata")
		os.Exit(2)
	}

	version, err := version.NewVersion(args[1])
	if err != nil {
		log.Fatalf(err.Error())
	}


	switch args[0] {
	case GET_RELEASE:
		fmt.Println(version.Prerelease())
	case GET_VERSION:
		segments:= version.Segments()
		fmt.Printf("%d.%d.%d\n", segments[0], segments[1], segments[2])
	case GET_META:
		fmt.Println(version.Metadata())
	default:
		log.Fatalln("Incorrect command")
	}
}