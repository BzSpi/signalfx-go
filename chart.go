package signalfx

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/signalfx/signalfx-go/chart"
)

// ChartAPIURL is the base URL for interacting with charts.
const ChartAPIURL = "/v2/chart"

// CreateChart creates a chart.
func (c *Client) CreateChart(chartRequest *chart.CreateUpdateChartRequest) (*chart.CreateUpdateChartResponse, error) {
	payload, err := json.Marshal(chartRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("POST", ChartAPIURL, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	finalChart := &chart.CreateUpdateChartResponse{}

	err = json.NewDecoder(resp.Body).Decode(finalChart)

	return finalChart, err
}

// DeleteChart deletes a chart.
func (c *Client) DeleteChart(id string) error {
	resp, err := c.doRequest("DELETE", ChartAPIURL+"/"+id, nil, nil)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.New("Unexpected status code: " + resp.Status)
	}

	return nil
}

// GetChart gets a chart.
func (c *Client) GetChart(id string) (*chart.GetSingleChartResponse, error) {
	resp, err := c.doRequest("GET", ChartAPIURL+"/"+id, nil, nil)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	finalChart := &chart.GetSingleChartResponse{}

	err = json.NewDecoder(resp.Body).Decode(finalChart)

	return finalChart, err
}

// UpdateChart updates a chart.
func (c *Client) UpdateChart(id string, chartRequest *chart.CreateUpdateChartRequest) (*chart.CreateUpdateChartResponse, error) {
	payload, err := json.Marshal(chartRequest)
	if err != nil {
		return nil, err
	}

	resp, err := c.doRequest("PUT", ChartAPIURL+"/"+id, nil, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	finalChart := &chart.CreateUpdateChartResponse{}

	err = json.NewDecoder(resp.Body).Decode(finalChart)

	return finalChart, err
}

// SearchChart searches for charts, given a query string in `name`.
func (c *Client) SearchCharts(limit int, name string, offset int, tags string) (*chart.GetChartsResult, error) {
	params := url.Values{}
	params.Add("limit", strconv.Itoa(limit))
	params.Add("name", name)
	params.Add("offset", strconv.Itoa(offset))
	params.Add("tags", tags)

	resp, err := c.doRequest("GET", ChartAPIURL, params, nil)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	finalCharts := &chart.GetChartsResult{}

	err = json.NewDecoder(resp.Body).Decode(finalCharts)

	return finalCharts, err
}
