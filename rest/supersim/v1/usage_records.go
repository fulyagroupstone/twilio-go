/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"time"
)

// Optional parameters for the method 'ListUsageRecord'
type ListUsageRecordParams struct {
	// SID or unique name of a Sim resource. Only show UsageRecords representing usage incurred by this Super SIM.
	Sim *string `json:"Sim,omitempty"`
	// SID or unique name of a Fleet resource. Only show UsageRecords representing usage for Super SIMs belonging to this Fleet resource at the time the usage occurred.
	Fleet *string `json:"Fleet,omitempty"`
	// SID of a Network resource. Only show UsageRecords representing usage on this network.
	Network *string `json:"Network,omitempty"`
	// Alpha-2 ISO Country Code. Only show UsageRecords representing usage in this country.
	IsoCountry *string `json:"IsoCountry,omitempty"`
	// Dimension over which to aggregate usage records. Can be: `sim`, `fleet`, `network`, `isoCountry`. Default is to not aggregate across any of these dimensions, UsageRecords will be aggregated into the time buckets described by the `Granularity` parameter.
	Group *string `json:"Group,omitempty"`
	// Time-based grouping that UsageRecords should be aggregated by. Can be: `hour`, `day`, or `all`. Default is `all`. `all` returns one UsageRecord that describes the usage for the entire period.
	Granularity *string `json:"Granularity,omitempty"`
	// Only include usage that occurred at or after this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is one month before the `end_time`.
	StartTime *time.Time `json:"StartTime,omitempty"`
	// Only include usage that occurred before this time, specified in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format. Default is the current time.
	EndTime *time.Time `json:"EndTime,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListUsageRecordParams) SetSim(Sim string) *ListUsageRecordParams {
	params.Sim = &Sim
	return params
}
func (params *ListUsageRecordParams) SetFleet(Fleet string) *ListUsageRecordParams {
	params.Fleet = &Fleet
	return params
}
func (params *ListUsageRecordParams) SetNetwork(Network string) *ListUsageRecordParams {
	params.Network = &Network
	return params
}
func (params *ListUsageRecordParams) SetIsoCountry(IsoCountry string) *ListUsageRecordParams {
	params.IsoCountry = &IsoCountry
	return params
}
func (params *ListUsageRecordParams) SetGroup(Group string) *ListUsageRecordParams {
	params.Group = &Group
	return params
}
func (params *ListUsageRecordParams) SetGranularity(Granularity string) *ListUsageRecordParams {
	params.Granularity = &Granularity
	return params
}
func (params *ListUsageRecordParams) SetStartTime(StartTime time.Time) *ListUsageRecordParams {
	params.StartTime = &StartTime
	return params
}
func (params *ListUsageRecordParams) SetEndTime(EndTime time.Time) *ListUsageRecordParams {
	params.EndTime = &EndTime
	return params
}
func (params *ListUsageRecordParams) SetPageSize(PageSize int) *ListUsageRecordParams {
	params.PageSize = &PageSize
	return params
}

// List UsageRecords
func (c *ApiService) ListUsageRecord(params *ListUsageRecordParams) (*ListUsageRecordResponse, error) {
	path := "/v1/UsageRecords"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Fleet != nil {
		data.Set("Fleet", *params.Fleet)
	}
	if params != nil && params.Network != nil {
		data.Set("Network", *params.Network)
	}
	if params != nil && params.IsoCountry != nil {
		data.Set("IsoCountry", *params.IsoCountry)
	}
	if params != nil && params.Group != nil {
		data.Set("Group", *params.Group)
	}
	if params != nil && params.Granularity != nil {
		data.Set("Granularity", *params.Granularity)
	}
	if params != nil && params.StartTime != nil {
		data.Set("StartTime", fmt.Sprint((*params.StartTime).Format(time.RFC3339)))
	}
	if params != nil && params.EndTime != nil {
		data.Set("EndTime", fmt.Sprint((*params.EndTime).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUsageRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}