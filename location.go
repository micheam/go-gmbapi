package gmbapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// LocationAccess ...
type LocationAccess struct {
	parent *Account
	client *Client
}

// LocationAccess ...
func (c *Client) LocationAccess(parent *Account) *LocationAccess {
	return &LocationAccess{parent: parent, client: c}
}

// List ...
func (l *LocationAccess) List(ctx context.Context, params url.Values) (<-chan *Location, error) {
	var stream = make(chan *Location, 100)
	go func() {
		defer close(stream)
		var next *string = nil
		for {
			accs, err := l.list(ctx, next, params)
			if err != nil {
				log.Printf("failed to list accounts: %v\n", err)
				return
			}
			for _, a := range accs.Locations {
				stream <- a
			}
			next = accs.NextPageToken
			if next == nil {
				break
			}
		}
	}()
	return stream, nil
}

// Get return the specified location. Returns ErrNotFound if the location does not exist.
func (l *LocationAccess) Get(ctx context.Context, id LocationID) (*Location, error) {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	_url := BaseEndpoint + "/" + l.parent.Name + "/locations/" + string(id)
	b, err := l.client.doRequest(ctx, time.Now(), http.MethodGet, _url, nil, url.Values{})
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest accounts.get: %w", err)
	}
	var loc = new(Location)
	if err := json.Unmarshal(b, loc); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return loc, nil
}

func (l *LocationAccess) list(ctx context.Context, nextPageToken *string, params url.Values) (*LocationList, error) {
	// TODO(micheam): QPS Limit
	//    maybe "golang.org/x/time/rate"
	if nextPageToken != nil {
		params.Add("pageToken", *nextPageToken)
	}
	_url := BaseEndpoint + "/" + l.parent.Name + "/locations"
	b, err := l.client.doRequest(ctx, time.Now(), http.MethodGet, _url, nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest accounts.list: %w", err)
	}
	var dat = new(LocationList)
	if err := json.Unmarshal(b, dat); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return dat, nil
}

// LocationList ...
type LocationList struct {
	Locations     []*Location `json:"locations"`
	NextPageToken *string     `json:"nextPageToken"`
	TotalSize     *int        `json:"totalSize"`
}

// LocationID is a identifier of Location
type LocationID string

// Location is ...
//
// TODO(mieahcm): add other fields
//    https://developers.google.com/my-business/reference/rest/v4/accounts.locations?hl=ja#Location
type Location struct {
	Name         *string `json:"name"`
	LocationName *string `json:"locationName"`
	StoreCode    *string `json:"storeCode"`
	PrimaryPhone *string `json:"primaryPhone"`
}

// StarRating ...
type StarRating int

// Definition of StarRating
const (
	StarRatingUNSPECIFIED StarRating = iota
	StarRatingONE
	StarRatingTWO
	StarRatingTHREE
	StarRatingFOUR
	StarRatingFIVE
)

func (s StarRating) String() string {
	return [...]string{
		"STAR_RATING_UNSPECIFIED",
		"ONE", "TWO", "THREE", "FOUR", "FIVE",
	}[s]
}

// UnmarshalJSON ...
func (s *StarRating) UnmarshalJSON(b []byte) error {
	var _s string
	if err := json.Unmarshal(b, &_s); err != nil {
		return err
	}
	switch strings.ToLower(_s) {
	default:
		*s = StarRatingUNSPECIFIED
	case "one":
		*s = StarRatingONE
	case "two":
		*s = StarRatingTWO
	case "three":
		*s = StarRatingTHREE
	case "four":
		*s = StarRatingFOUR
	case "five":
		*s = StarRatingFIVE
	}

	return nil
}

// MarshalJSON ...
func (s *StarRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
