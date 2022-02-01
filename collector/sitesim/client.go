package sitesim

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.uber.org/multierr"
)

type Turbine struct {
	ID                int     `json:"Id"`
	Name              string  `json:"Name"`
	Manufacturer      string  `json:"Manufacturer"`
	Version           int     `json:"Version"`
	MaxProduction     int     `json:"MaxProduction"`
	CurrentProduction float64 `json:"CurrentProduction"`
	WindSpeed         float64 `json:"Windspeed"`
	WindDirection     string  `json:"WindDirection"`
}

type Site struct {
	ID          int       `json:"Id"`
	Name        string    `json:"Name"`
	Description string    `json:"Description"`
	Region      string    `json:"Region"`
	Country     string    `json:"Country"`
	Turbines    []Turbine `json:"Turbines"`
}

type Client struct {
	Client  *http.Client
	BaseURL string
}

func (c Client) ListSites(ctx context.Context) (sites []Site, err error) {
	url := fmt.Sprintf("%s/api/Site", c.BaseURL)

	if err = c.do(ctx, url, &sites); err != nil {
		return nil, err
	}
	return sites, nil
}

func (c Client) GetSite(ctx context.Context, siteID int) (site Site, err error) {
	url := fmt.Sprintf("%s/api/Site/%d", c.BaseURL, siteID)

	if err = c.do(ctx, url, &site); err != nil {
		return Site{}, err
	}
	return site, nil
}

func (c Client) do(ctx context.Context, url string, into interface{}) (err error) {
	var req *http.Request
	if req, err = http.NewRequestWithContext(ctx, http.MethodGet, url, nil); err != nil {
		return err
	}

	var resp *http.Response
	if resp, err = c.Client.Do(req); err != nil {
		return err
	}
	defer multierr.AppendInvoke(&err, multierr.Close(resp.Body))

	decoder := json.NewDecoder(resp.Body)
	if err = decoder.Decode(into); err != nil {
		return err
	}
	return nil
}
