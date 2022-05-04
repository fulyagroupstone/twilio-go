/*
 * Twilio - Fax
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

// Delete a specific fax media instance.
func (c *ApiService) DeleteFaxMedia(FaxSid string, Sid string) error {
	path := "/v1/Faxes/{FaxSid}/Media/{Sid}"
	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific fax media instance.
func (c *ApiService) FetchFaxMedia(FaxSid string, Sid string) (*FaxV1FaxMedia, error) {
	path := "/v1/Faxes/{FaxSid}/Media/{Sid}"
	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FaxV1FaxMedia{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFaxMedia'
type ListFaxMediaParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFaxMediaParams) SetPageSize(PageSize int) *ListFaxMediaParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFaxMediaParams) SetLimit(Limit int) *ListFaxMediaParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of FaxMedia records from the API. Request is executed immediately.
func (c *ApiService) PageFaxMedia(FaxSid string, params *ListFaxMediaParams, pageToken, pageNumber string) (*ListFaxMediaResponse, error) {
	path := "/v1/Faxes/{FaxSid}/Media"

	path = strings.Replace(path, "{"+"FaxSid"+"}", FaxSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

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

	ps := &ListFaxMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FaxMedia records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFaxMedia(FaxSid string, params *ListFaxMediaParams) ([]FaxV1FaxMedia, error) {
	response, errors := c.StreamFaxMedia(FaxSid, params)

	records := make([]FaxV1FaxMedia, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams FaxMedia records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFaxMedia(FaxSid string, params *ListFaxMediaParams) (chan FaxV1FaxMedia, chan error) {
	if params == nil {
		params = &ListFaxMediaParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan FaxV1FaxMedia, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageFaxMedia(FaxSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamFaxMedia(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamFaxMedia(response *ListFaxMediaResponse, params *ListFaxMediaParams, recordChannel chan FaxV1FaxMedia, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Media
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListFaxMediaResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListFaxMediaResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListFaxMediaResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFaxMediaResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
