package gmbapi

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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
func (m *MediaAccess) List(ctx context.Context, params url.Values) ([]*MediaItem, error) {
	var list = make([]*MediaItem, 0)
	var next *string = nil
	for {
		mediaList, err := m.list(ctx, next, params)
		if err != nil {
			return nil, fmt.Errorf("failed to list locations: %w", err)
		}
		list = append(list, mediaList.Items...)
		next = mediaList.NextPageToken
		if next == nil {
			break
		}
	}
	return list, nil
}

func (m *MediaAccess) list(ctx context.Context, nextPageToken *string, params url.Values) (*MediaList, error) {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	if nextPageToken != nil {
		params.Add("pageToken", *nextPageToken)
	}
	_url := BaseEndpoint + "/" + m.parent.Name + "/media"
	b, err := m.client.doRequest(ctx, time.Now(), http.MethodGet, _url, nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest media.list: %w", err)
	}
	var dat = new(MediaList)
	if err := json.Unmarshal(b, dat); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return dat, nil
}

// Get ...
func (m *MediaAccess) Get(ctx context.Context, id string) (*MediaItem, error) {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	_url := BaseEndpoint + "/" + m.parent.Name + "/media/" + id
	b, err := m.client.doRequest(ctx, time.Now(), http.MethodGet, _url, nil, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest media.get: %w", err)
	}
	var media = new(MediaItem)
	if err := json.Unmarshal(b, media); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return media, nil
}

func (m *MediaAccess) Create(ctx context.Context, item *MediaItem, params url.Values) error {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	_url := BaseEndpoint + "/" + m.parent.Name + "/media"

	// Request Body
	b, err := json.Marshal(item)
	if err != nil {
		return fmt.Errorf("failed to marshal MediaItem into json: %w", err)
	}
	var body io.ReadSeeker = bytes.NewReader(b)
	res, err := m.client.doRequest(ctx, time.Now(), http.MethodPost, _url, body, params)
	if err != nil {
		return fmt.Errorf("failed to doRequest media.create: %w", err)
	}
	err = json.Unmarshal(res, item)
	if err != nil {
		return fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return nil
}

func (m *MediaAccess) Delete(ctx context.Context, item *MediaItem) error {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	_url := BaseEndpoint + "/" + item.Name
	res, err := m.client.doRequest(ctx, time.Now(), http.MethodDelete, _url, nil, url.Values{})
	if err != nil {
		return fmt.Errorf("failed to doRequest media.delete: %w", err)
	}
	err = json.Unmarshal(res, item)
	if err != nil {
		return fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return nil
}

// MediaList ...
type MediaList struct {
	Items               []*MediaItem `json:"mediaItems"`
	TotalMediaItemCount *int64       `json:"totalMediaItemCount"`
	NextPageToken       *string      `json:"nextPageToken"`
}

// MediaItemID is a identifier of Media
type MediaItemID string

func (m *MediaItem) ID() MediaItemID {
	s := strings.Split(m.Name, "/")
	return MediaItemID(s[len(s)-1])
}

// MediaItem is ...
//
// https://developers.google.com/my-business/reference/rest/v4/accounts.locations.media?hl=ja#MediaItem
type MediaItem struct {
	Name                string              `json:"name"`
	Format              *MediaFormat        `json:"mediaFormat"` // emun: MediaFormat
	LocationAssociation LocationAssociation `json:"locationAssociation"`
	GoogleUrl           *string             `json:"googleUrl"`
	ThumbnailUrl        *string             `json:"thumbnailUrl"`
	CreateTime          *string             `json:"createTime"`
	Dimensions          interface{}         `json:"dimensions"`
	Insights            MediaInsights       `json:"insights"`
	Attribution         Attribution         `json:"attribution"`
	Description         *string             `json:"description"`

	// Union field data can be only one of the following:
	Src     *string          `json:"sourceUrl,omitempty"`
	DataRef MediaItemDataRef `json:"dataRef,omitempty"`
}

type Dimensions interface{}
type LocationAssociation struct {
	Category        MediaItemCategory `json:"category"`
	PriceListItemID *string           `json:"priceListItemId"`
}
type MediaInsights interface{}
type Attribution interface{}
type MediaItemDataRef interface{}

type MediaFormat string
type MediaItemCategory string
