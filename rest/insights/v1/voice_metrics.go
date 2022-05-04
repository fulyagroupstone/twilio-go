/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'ListMetric'
type ListMetricParams struct {
	//
	Edge *string `json:"Edge,omitempty"`
	//
	Direction *string `json:"Direction,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListMetricParams) SetEdge(Edge string) *ListMetricParams {
	params.Edge = &Edge
	return params
}
func (params *ListMetricParams) SetDirection(Direction string) *ListMetricParams {
	params.Direction = &Direction
	return params
}
func (params *ListMetricParams) SetPageSize(PageSize int) *ListMetricParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListMetricParams) SetLimit(Limit int) *ListMetricParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Metric records from the API. Request is executed immediately.
func (c *ApiService) PageMetric(CallSid string, params *ListMetricParams, pageToken, pageNumber string) (*ListMetricResponse, error) {
	path := "/v1/Voice/{CallSid}/Metrics"

	path = strings.Replace(path, "{"+"CallSid"+"}", CallSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Edge != nil {
		data.Set("Edge", *params.Edge)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMetricResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Metric records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListMetric(CallSid string, params *ListMetricParams) ([]InsightsV1Metric, error) {
	response, errors := c.StreamMetric(CallSid, params)

	records := make([]InsightsV1Metric, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Metric records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamMetric(CallSid string, params *ListMetricParams) (chan InsightsV1Metric, chan error) {
	if params == nil {
		params = &ListMetricParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan InsightsV1Metric, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageMetric(CallSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamMetric(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamMetric(response *ListMetricResponse, params *ListMetricParams, recordChannel chan InsightsV1Metric, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Metrics
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListMetricResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListMetricResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListMetricResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListMetricResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
