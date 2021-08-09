/*
 * Twilio - Autopilot
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

// Optional parameters for the method 'CreateFieldValue'
type CreateFieldValueParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US`
	Language *string `json:"Language,omitempty"`
	// The string value that indicates which word the field value is a synonym of.
	SynonymOf *string `json:"SynonymOf,omitempty"`
	// The Field Value data.
	Value *string `json:"Value,omitempty"`
}

func (params *CreateFieldValueParams) SetLanguage(Language string) *CreateFieldValueParams {
	params.Language = &Language
	return params
}
func (params *CreateFieldValueParams) SetSynonymOf(SynonymOf string) *CreateFieldValueParams {
	params.SynonymOf = &SynonymOf
	return params
}
func (params *CreateFieldValueParams) SetValue(Value string) *CreateFieldValueParams {
	params.Value = &Value
	return params
}

func (c *ApiService) CreateFieldValue(AssistantSid string, FieldTypeSid string, params *CreateFieldValueParams) (*AutopilotV1FieldValue, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"FieldTypeSid"+"}", FieldTypeSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
	}
	if params != nil && params.SynonymOf != nil {
		data.Set("SynonymOf", *params.SynonymOf)
	}
	if params != nil && params.Value != nil {
		data.Set("Value", *params.Value)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1FieldValue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteFieldValue(AssistantSid string, FieldTypeSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"FieldTypeSid"+"}", FieldTypeSid, -1)
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

func (c *ApiService) FetchFieldValue(AssistantSid string, FieldTypeSid string, Sid string) (*AutopilotV1FieldValue, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"FieldTypeSid"+"}", FieldTypeSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1FieldValue{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListFieldValue'
type ListFieldValueParams struct {
	// The [ISO language-country](https://docs.oracle.com/cd/E13214_01/wli/docs92/xref/xqisocodes.html) tag that specifies the language of the value. Currently supported tags: `en-US`
	Language *string `json:"Language,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListFieldValueParams) SetLanguage(Language string) *ListFieldValueParams {
	params.Language = &Language
	return params
}
func (params *ListFieldValueParams) SetPageSize(PageSize int) *ListFieldValueParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListFieldValueParams) SetLimit(Limit int) *ListFieldValueParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of FieldValue records from the API. Request is executed immediately.
func (c *ApiService) PageFieldValue(AssistantSid string, FieldTypeSid string, params *ListFieldValueParams, pageToken string, pageNumber string) (*ListFieldValueResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/FieldTypes/{FieldTypeSid}/FieldValues"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"FieldTypeSid"+"}", FieldTypeSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Language != nil {
		data.Set("Language", *params.Language)
	}
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

	ps := &ListFieldValueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists FieldValue records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListFieldValue(AssistantSid string, FieldTypeSid string, params *ListFieldValueParams) ([]AutopilotV1FieldValue, error) {
	if params == nil {
		params = &ListFieldValueParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFieldValue(AssistantSid, FieldTypeSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []AutopilotV1FieldValue

	for response != nil {
		records = append(records, response.FieldValues...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFieldValueResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListFieldValueResponse)
	}

	return records, err
}

// Streams FieldValue records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamFieldValue(AssistantSid string, FieldTypeSid string, params *ListFieldValueParams) (chan AutopilotV1FieldValue, error) {
	if params == nil {
		params = &ListFieldValueParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageFieldValue(AssistantSid, FieldTypeSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan AutopilotV1FieldValue, 1)

	go func() {
		for response != nil {
			for item := range response.FieldValues {
				channel <- response.FieldValues[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListFieldValueResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListFieldValueResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListFieldValueResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListFieldValueResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
