/*
 * Twilio - Numbers
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

// Optional parameters for the method 'CreateItemAssignment'
type CreateItemAssignmentParams struct {
	// The SID of an object bag that holds information of the different items.
	ObjectSid *string `json:"ObjectSid,omitempty"`
}

func (params *CreateItemAssignmentParams) SetObjectSid(ObjectSid string) *CreateItemAssignmentParams {
	params.ObjectSid = &ObjectSid
	return params
}

// Create a new Assigned Item.
func (c *ApiService) CreateItemAssignment(BundleSid string, params *CreateItemAssignmentParams) (*NumbersV2ItemAssignment, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ObjectSid != nil {
		data.Set("ObjectSid", *params.ObjectSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2ItemAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove an Assignment Item Instance.
func (c *ApiService) DeleteItemAssignment(BundleSid string, Sid string) error {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
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

// Fetch specific Assigned Item Instance.
func (c *ApiService) FetchItemAssignment(BundleSid string, Sid string) (*NumbersV2ItemAssignment, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments/{Sid}"
	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NumbersV2ItemAssignment{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListItemAssignment'
type ListItemAssignmentParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListItemAssignmentParams) SetPageSize(PageSize int) *ListItemAssignmentParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListItemAssignmentParams) SetLimit(Limit int) *ListItemAssignmentParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ItemAssignment records from the API. Request is executed immediately.
func (c *ApiService) PageItemAssignment(BundleSid string, params *ListItemAssignmentParams, pageToken, pageNumber string) (*ListItemAssignmentResponse, error) {
	path := "/v2/RegulatoryCompliance/Bundles/{BundleSid}/ItemAssignments"

	path = strings.Replace(path, "{"+"BundleSid"+"}", BundleSid, -1)

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

	ps := &ListItemAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ItemAssignment records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListItemAssignment(BundleSid string, params *ListItemAssignmentParams) ([]NumbersV2ItemAssignment, error) {
	response, errors := c.StreamItemAssignment(BundleSid, params)

	records := make([]NumbersV2ItemAssignment, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ItemAssignment records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamItemAssignment(BundleSid string, params *ListItemAssignmentParams) (chan NumbersV2ItemAssignment, chan error) {
	if params == nil {
		params = &ListItemAssignmentParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan NumbersV2ItemAssignment, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageItemAssignment(BundleSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamItemAssignment(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamItemAssignment(response *ListItemAssignmentResponse, params *ListItemAssignmentParams, recordChannel chan NumbersV2ItemAssignment, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Results
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListItemAssignmentResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListItemAssignmentResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListItemAssignmentResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListItemAssignmentResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
