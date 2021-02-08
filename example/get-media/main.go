package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/micheam/go-gmbapi"
)

func init() {
	// Overwrite with your credentials for GoogleApi.
	// os.Setenv("GMB_CLIENT_ID", "<YOUR_CLIENT_ID>")
	// os.Setenv("GMB_CLIENT_SECRET", "<YOUR_CLIENT_SECRET>")
	// os.Setenv("GMB_REFRESH_TOKEN", "<YOUR_REFRESH_TOKEN>")

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s [location_name]:\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Please specify parent 'location_name'.")
		os.Exit(1)
	}
	parent := flag.Arg(0)

	var (
		ctx    = context.Background()
		client *gmbapi.Client
		err    error
	)
	if client, err = gmbapi.New(); err != nil {
		fmt.Printf("failed to build gmbapi client: %s\n", err.Error())
		os.Exit(1)
	}
	media := client.MediaAccess(&gmbapi.Location{Name: &parent})
	mediaItems, err := media.List(ctx, nil)
	if err != nil {
		fmt.Printf("failed to get mediaItems: %s\n", err.Error())
		os.Exit(1)
	}
	for item := range mediaItems {
		b, _ := json.Marshal(item)
		fmt.Println(string(b))
	}
}
