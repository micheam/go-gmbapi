package config

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func ExampleLoad() {
	os.Setenv("GMB_CLIENT_ID", "THIS_IS_YOUR_CLIENT_ID")
	os.Setenv("GMB_CLIENT_SECRET", "THIS_IS_YOUR_CLIENT_SECRET")
	os.Setenv("GMB_REFRESH_TOKEN", "THIS_IS_YOUR_REFRESH_TOKEN")

	got, err := Load()
	if err != nil {
		panic(err)
	}
	fmt.Printf("ClientID=%q\nClientSecret=%q\nRefreshToken=%q\n",
		got.ClientID, got.ClientSecret, got.RefreshToken)

	// Output:
	// ClientID="THIS_IS_YOUR_CLIENT_ID"
	// ClientSecret="THIS_IS_YOUR_CLIENT_SECRET"
	// RefreshToken="THIS_IS_YOUR_REFRESH_TOKEN"
}
