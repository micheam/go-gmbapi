package gmbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// MediaAccess ...
type MediaAccess struct {
	parent *Location
	client *Client
}

// MediaAccess ...
func (c *Client) MediaAccess(parent *Location) *MediaAccess {
	return &MediaAccess{parent: parent, client: c}
}

// List ...
func (l *MediaAccess) List(ctx context.Context, params url.Values) (<-chan *MediaItem, error) {
	var stream = make(chan *MediaItem, 100)
	go func() {
		defer close(stream)
		var next *string = nil
		for {
			mediaList, err := l.list(ctx, next, params)
			if err != nil {
				log.Printf("failed to list mediaItems: %v\n", err)
				return
			}
			for _, a := range mediaList.Items {
				stream <- a
			}
			next = mediaList.NextPageToken
			if next == nil {
				break
			}
		}
	}()
	return stream, nil
}

func (l *MediaAccess) list(ctx context.Context, nextPageToken *string, params url.Values) (*MediaList, error) {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	if nextPageToken != nil {
		params.Add("pageToken", *nextPageToken)
	}
	_url := BaseEndpoint + "/" + *l.parent.Name + "/media"
	b, err := l.client.doRequest(ctx, http.MethodGet, _url, nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest media.list: %w", err)
	}
	var dat = new(MediaList)
	if err := json.Unmarshal(b, dat); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return dat, nil
}

// MediaList ...
type MediaList struct {
	Items               []*MediaItem `json:"mediaItems"`
	TotalMediaItemCount *int64       `json:"totalMediaItemCount"`
	NextPageToken       *string      `json:"nextPageToken"`
}

// MediaID is a identifier of Media
type MediaID string

// MediaItem is ...
//
// https://developers.google.com/my-business/reference/rest/v4/accounts.locations.media?hl=ja#MediaItem
type MediaItem struct {
	Name                *string             `json:"name"`
	Format              *string             `json:"mediaFormat"` // emun: MediaFormat
	LocationAssociation LocationAssociation `json:"locationAssociation"`
	GoogleUrl           *string             `json:"googleUrl"`
	ThumbnailUrl        *string             `json:"thumbnailUrl"`
	CreateTime          *string             `json:"createTime"`
	Dimensions          interface{}         `json:"dimensions"`
	Insights            MediaInsights       `json:"insights"`
	Attribution         Attribution         `json:"attribution"`
	Description         *string             `json:"description"`

	// Union field data can be only one of the following:
	Src     *string          `json:"sourceUrl"`
	DataRef MediaItemDataRef `json:"dataRef"`
}

type Dimensions interface{}
type LocationAssociation interface{}
type MediaInsights interface{}
type Attribution interface{}
type MediaItemDataRef interface{}
