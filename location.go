package gmbapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/micheam/gmbapi/internal/util/pointer"
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
func (a *LocationAccess) List(params url.Values) (*LocationList, error) {
	parentName := pointer.StringPtrDeref(a.parent.Name, "")
	_url := fmt.Sprintf("%s/%s/locations", BaseEndpoint, parentName)
	// FIXME(micheam): hundle NextPageToken
	b, err := a.client.doRequest(http.MethodGet, _url, nil, params)
	if err != nil {
		return nil, fmt.Errorf("failed to doRequest locations.list: %w", err)
	}
	var res = new(LocationList)
	if err := json.Unmarshal(b, res); err != nil {
		return nil, fmt.Errorf("failed to unmarshal api response: %w", err)
	}
	return res, nil
}

// LocationList ...
type LocationList struct {
	Locations     []*Location `json:"locations"`
	NextPageToken *string     `json:"nextPageToken"`
	TotalSize     *int        `json:"totalSize"`
}

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
