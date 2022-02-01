package sitesim_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/maxatome/go-testdeep/td"

	"github.com/baez90/windpark-challenge/sitesim"
)

func TestClient_ListSites(t *testing.T) {
	t.Parallel()
	client := sitesim.Client{
		Client:  http.DefaultClient,
		BaseURL: "http://renewables-codechallenge.azurewebsites.net",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	t.Cleanup(cancel)
	sites, err := client.ListSites(ctx)
	if err != nil {
		t.Errorf("client.ListSites(ctx) error = %v", err)
		return
	}

	td.Cmp(t, sites, td.NotEmpty())
}

func TestClient_GetSite(t *testing.T) {
	t.Parallel()
	client := sitesim.Client{
		Client:  http.DefaultClient,
		BaseURL: "http://renewables-codechallenge.azurewebsites.net",
	}

	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	t.Cleanup(cancel)
	sites, err := client.GetSite(ctx, 1)
	if err != nil {
		t.Errorf("client.ListSites(ctx) error = %v", err)
		return
	}

	td.Cmp(t, sites, td.Struct(sitesim.Site{}, td.StructFields{
		"Turbines": td.NotEmpty(),
	}))
}
