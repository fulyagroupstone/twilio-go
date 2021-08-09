/*
 * Twilio - Sync
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

// Optional parameters for the method 'CreateSyncList'
type CreateSyncListParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// Alias for collection_ttl. If both are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
	// An application-defined string that uniquely identifies the resource. This value must be unique within its Service and it can be up to 320 characters long. The `unique_name` value can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateSyncListParams) SetCollectionTtl(CollectionTtl int) *CreateSyncListParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *CreateSyncListParams) SetTtl(Ttl int) *CreateSyncListParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreateSyncListParams) SetUniqueName(UniqueName string) *CreateSyncListParams {
	params.UniqueName = &UniqueName
	return params
}

func (c *ApiService) CreateSyncList(ServiceSid string, params *CreateSyncListParams) (*SyncV1SyncList, error) {
	path := "/v1/Services/{ServiceSid}/Lists"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteSyncList(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Lists/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

func (c *ApiService) FetchSyncList(ServiceSid string, Sid string) (*SyncV1SyncList, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSyncList'
type ListSyncListParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSyncListParams) SetPageSize(PageSize int) *ListSyncListParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSyncListParams) SetLimit(Limit int) *ListSyncListParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SyncList records from the API. Request is executed immediately.
func (c *ApiService) PageSyncList(ServiceSid string, params *ListSyncListParams, pageToken string, pageNumber string) (*ListSyncListResponse, error) {
	path := "/v1/Services/{ServiceSid}/Lists"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListSyncListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SyncList records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSyncList(ServiceSid string, params *ListSyncListParams) ([]SyncV1SyncList, error) {
	if params == nil {
		params = &ListSyncListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncList(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []SyncV1SyncList

	for response != nil {
		records = append(records, response.Lists...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncListResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSyncListResponse)
	}

	return records, err
}

// Streams SyncList records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSyncList(ServiceSid string, params *ListSyncListParams) (chan SyncV1SyncList, error) {
	if params == nil {
		params = &ListSyncListParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSyncList(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan SyncV1SyncList, 1)

	go func() {
		for response != nil {
			for item := range response.Lists {
				channel <- response.Lists[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSyncListResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSyncListResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSyncListResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSyncListResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSyncList'
type UpdateSyncListParams struct {
	// How long, [in seconds](https://www.twilio.com/docs/sync/limits#sync-payload-limits), before the Sync List expires (time-to-live) and is deleted.
	CollectionTtl *int `json:"CollectionTtl,omitempty"`
	// An alias for `collection_ttl`. If both are provided, this value is ignored.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *UpdateSyncListParams) SetCollectionTtl(CollectionTtl int) *UpdateSyncListParams {
	params.CollectionTtl = &CollectionTtl
	return params
}
func (params *UpdateSyncListParams) SetTtl(Ttl int) *UpdateSyncListParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) UpdateSyncList(ServiceSid string, Sid string, params *UpdateSyncListParams) (*SyncV1SyncList, error) {
	path := "/v1/Services/{ServiceSid}/Lists/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CollectionTtl != nil {
		data.Set("CollectionTtl", fmt.Sprint(*params.CollectionTtl))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SyncV1SyncList{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
