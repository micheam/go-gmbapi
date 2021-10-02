package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/url"
	"os"

	"github.com/micheam/go-gmbapi"
)

func init() {
	// Overwrite with your credentials for GoogleApi.
	// os.Setenv("GMB_CLIENT_ID", "<YOUR_CLIENT_ID>")
	// os.Setenv("GMB_CLIENT_SECRET", "<YOUR_CLIENT_SECRET>")
	// os.Setenv("GMB_REFRESH_TOKEN", "<YOUR_REFRESH_TOKEN>")

	// switch io.Stderr or something if you want to watch logs
	log.SetOutput(io.Discard)

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s [location_name]:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	var (
		ctx    = context.Background()
		client *gmbapi.Client
		err    error
	)
	if client, err = gmbapi.New(); err != nil {
		fmt.Printf("failed to build gmbapi client: %s\n", err.Error())
		os.Exit(1)
	}

	if len(flag.Args()) != 3 {
		fmt.Fprintf(os.Stderr, "illegal num of args\n")
		os.Exit(1)
	}

	locName := flag.Arg(0)
	loc := &gmbapi.Location{Name: &locName}
	category := flag.Arg(1)
	srcURL := flag.Arg(2)

	log.Printf("location name: %s", locName)
	log.Printf("category: %s", category)
	log.Printf("source url: %s", srcURL)

	mediaFormat := gmbapi.MediaFormat("PHOTO")
	item := &gmbapi.MediaItem{
		Format: &mediaFormat,
		LocationAssociation: gmbapi.LocationAssociation{
			Category: gmbapi.MediaItemCategory(category),
		},
		Src: &srcURL,
	}

	err = client.MediaAccess(loc).Create(ctx, item, url.Values{})
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to create media-item: %s\n", err.Error())
		os.Exit(1)
	}

	b, _ := json.Marshal(item)
	fmt.Println(string(b))
}
