/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
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

// Creates an evaluation for a bundle
func (c *ApiService) CreateEvaluation(BundleSid string) (*NumbersV2Evaluation, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Evaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Fetch specific Evaluation Instance.
func (c *ApiService) FetchEvaluation(BundleSid string, Sid string) (*NumbersV2Evaluation, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2Evaluation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEvaluation'
type ListEvaluationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListEvaluationParams) SetPageSize(PageSize int) *ListEvaluationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListEvaluationParams) SetLimit(Limit int) *ListEvaluationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Evaluation records from the API. Request is executed immediately.
func (c *ApiService) PageEvaluation(BundleSid string, params *ListEvaluationParams, pageToken string, pageNumber string) (*ListEvaluationResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/Evaluations"

	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Evaluation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListEvaluation(BundleSid string, params *ListEvaluationParams) ([]NumbersV2Evaluation, error) {
	if params == nil {
		params = &ListEvaluationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEvaluation(BundleSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []NumbersV2Evaluation

	for response != nil {
		records = append(records, response.Results...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEvaluationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListEvaluationResponse)
	}

	return records, err
}

// Streams Evaluation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamEvaluation(BundleSid string, params *ListEvaluationParams) (chan NumbersV2Evaluation, error) {
	if params == nil {
		params = &ListEvaluationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageEvaluation(BundleSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan NumbersV2Evaluation, 1)

	go func() {
		for response != nil {
			for item := range response.Results {
				channel <- response.Results[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListEvaluationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListEvaluationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListEvaluationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEvaluationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
