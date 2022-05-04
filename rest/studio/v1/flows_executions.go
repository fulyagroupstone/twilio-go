/*
 * Twilio - Studio
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateExecution'
type CreateExecutionParams struct {
	// The Twilio phone number to send messages or initiate calls from during the Flow's Execution. Available as variable `{{flow.channel.address}}`. For SMS, this can also be a Messaging Service SID.
	From *string `json:"From,omitempty"`
	// JSON data that will be added to the Flow's context and that can be accessed as variables inside your Flow. For example, if you pass in `Parameters={\\\"name\\\":\\\"Zeke\\\"}`, a widget in your Flow can reference the variable `{{flow.data.name}}`, which returns \\\"Zeke\\\". Note: the JSON value must explicitly be passed as a string, not as a hash object. Depending on your particular HTTP library, you may need to add quotes or URL encode the JSON string.
	Parameters *interface{} `json:"Parameters,omitempty"`
	// The Contact phone number to start a Studio Flow Execution, available as variable `{{contact.channel.address}}`.
	To *string `json:"To,omitempty"`
}

func (params *CreateExecutionParams) SetFrom(From string) *CreateExecutionParams {
	params.From = &From
	return params
}
func (params *CreateExecutionParams) SetParameters(Parameters interface{}) *CreateExecutionParams {
	params.Parameters = &Parameters
	return params
}
func (params *CreateExecutionParams) SetTo(To string) *CreateExecutionParams {
	params.To = &To
	return params
}

// Triggers a new Execution for the Flow
func (c *ApiService) CreateExecution(FlowSid string, params *CreateExecutionParams) (*StudioV1Execution, error) {
	path := "/v1/Flows/{FlowSid}/Executions"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.From != nil {
		data.Set("From", *params.From)
	}
	if params != nil && params.Parameters != nil {
		v, err := json.Marshal(params.Parameters)

		if err != nil {
			return nil, err
		}

		data.Set("Parameters", string(v))
	}
	if params != nil && params.To != nil {
		data.Set("To", *params.To)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Execution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete the Execution and all Steps relating to it.
func (c *ApiService) DeleteExecution(FlowSid string, Sid string) error {
	path := "/v1/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
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

// Retrieve an Execution
func (c *ApiService) FetchExecution(FlowSid string, Sid string) (*StudioV1Execution, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Execution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListExecution'
type ListExecutionParams struct {
	// Only show Execution resources starting on or after this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
	DateCreatedFrom *time.Time `json:"DateCreatedFrom,omitempty"`
	// Only show Execution resources starting before this [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date-time, given as `YYYY-MM-DDThh:mm:ss-hh:mm`.
	DateCreatedTo *time.Time `json:"DateCreatedTo,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListExecutionParams) SetDateCreatedFrom(DateCreatedFrom time.Time) *ListExecutionParams {
	params.DateCreatedFrom = &DateCreatedFrom
	return params
}
func (params *ListExecutionParams) SetDateCreatedTo(DateCreatedTo time.Time) *ListExecutionParams {
	params.DateCreatedTo = &DateCreatedTo
	return params
}
func (params *ListExecutionParams) SetPageSize(PageSize int) *ListExecutionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListExecutionParams) SetLimit(Limit int) *ListExecutionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Execution records from the API. Request is executed immediately.
func (c *ApiService) PageExecution(FlowSid string, params *ListExecutionParams, pageToken, pageNumber string) (*ListExecutionResponse, error) {
	path := "/v1/Flows/{FlowSid}/Executions"

	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateCreatedFrom != nil {
		data.Set("DateCreatedFrom", fmt.Sprint((*params.DateCreatedFrom).Format(time.RFC3339)))
	}
	if params != nil && params.DateCreatedTo != nil {
		data.Set("DateCreatedTo", fmt.Sprint((*params.DateCreatedTo).Format(time.RFC3339)))
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

	ps := &ListExecutionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Execution records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListExecution(FlowSid string, params *ListExecutionParams) ([]StudioV1Execution, error) {
	response, errors := c.StreamExecution(FlowSid, params)

	records := make([]StudioV1Execution, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Execution records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamExecution(FlowSid string, params *ListExecutionParams) (chan StudioV1Execution, chan error) {
	if params == nil {
		params = &ListExecutionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan StudioV1Execution, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageExecution(FlowSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamExecution(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamExecution(response *ListExecutionResponse, params *ListExecutionParams, recordChannel chan StudioV1Execution, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Executions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListExecutionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListExecutionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListExecutionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListExecutionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateExecution'
type UpdateExecutionParams struct {
	// The status of the Execution. Can only be `ended`.
	Status *string `json:"Status,omitempty"`
}

func (params *UpdateExecutionParams) SetStatus(Status string) *UpdateExecutionParams {
	params.Status = &Status
	return params
}

// Update the status of an Execution to &#x60;ended&#x60;.
func (c *ApiService) UpdateExecution(FlowSid string, Sid string, params *UpdateExecutionParams) (*StudioV1Execution, error) {
	path := "/v1/Flows/{FlowSid}/Executions/{Sid}"
	path = strings.Replace(path, "{"+"FlowSid"+"}", FlowSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &StudioV1Execution{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
