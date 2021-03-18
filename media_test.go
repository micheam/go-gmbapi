package gmbapi

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMediaItem_ID(t *testing.T) {
	m := MediaItem{Name: "accounts/11111/locations/222222/media/media-item-id"}
	want := MediaItemID("media-item-id")

	if diff := cmp.Diff(want, m.ID()); diff != "" {
		t.Errorf("ID() mismatch (-want, +got):%s\n", diff)
	}
}
