/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Proxy
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateSession'
type CreateSessionParams struct {
	// An application-defined string that uniquely identifies the resource. This value must be 191 characters or fewer in length and be unique. **This value should not have PII.**
	UniqueName *string `json:"UniqueName,omitempty"`
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value.
	DateExpiry *time.Time `json:"DateExpiry,omitempty"`
	// The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction.
	Ttl *int `json:"Ttl,omitempty"`
	//
	Mode *string `json:"Mode,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
	// The Participant objects to include in the new session.
	Participants *[]interface{} `json:"Participants,omitempty"`
	// [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to reject a Session create (with Participants) request that could cause the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. Depending on the context, this could be a 409 error (Twilio error code 80623) or a 400 error (Twilio error code 80604). If not provided, requests will be allowed to succeed and a Debugger notification (80802) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
	FailOnParticipantConflict *bool `json:"FailOnParticipantConflict,omitempty"`
}

func (params *CreateSessionParams) SetUniqueName(UniqueName string) *CreateSessionParams {
	params.UniqueName = &UniqueName
	return params
}
func (params *CreateSessionParams) SetDateExpiry(DateExpiry time.Time) *CreateSessionParams {
	params.DateExpiry = &DateExpiry
	return params
}
func (params *CreateSessionParams) SetTtl(Ttl int) *CreateSessionParams {
	params.Ttl = &Ttl
	return params
}
func (params *CreateSessionParams) SetMode(Mode string) *CreateSessionParams {
	params.Mode = &Mode
	return params
}
func (params *CreateSessionParams) SetStatus(Status string) *CreateSessionParams {
	params.Status = &Status
	return params
}
func (params *CreateSessionParams) SetParticipants(Participants []interface{}) *CreateSessionParams {
	params.Participants = &Participants
	return params
}
func (params *CreateSessionParams) SetFailOnParticipantConflict(FailOnParticipantConflict bool) *CreateSessionParams {
	params.FailOnParticipantConflict = &FailOnParticipantConflict
	return params
}

// Create a new Session
func (c *ApiService) CreateSession(ServiceSid string, params *CreateSessionParams) (*ProxyV1Session, error) {
	path := "/v1/Services/{ServiceSid}/Sessions"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}
	if params != nil && params.DateExpiry != nil {
		data.Set("DateExpiry", fmt.Sprint((*params.DateExpiry).Format(time.RFC3339)))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.Mode != nil {
		data.Set("Mode", *params.Mode)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Participants != nil {
		for _, item := range *params.Participants {
			v, err := json.Marshal(item)

			if err != nil {
				return nil, err
			}

			data.Add("Participants", string(v))
		}
	}
	if params != nil && params.FailOnParticipantConflict != nil {
		data.Set("FailOnParticipantConflict", fmt.Sprint(*params.FailOnParticipantConflict))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Session{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a specific Session.
func (c *ApiService) DeleteSession(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/Sessions/{Sid}"
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

// Fetch a specific Session.
func (c *ApiService) FetchSession(ServiceSid string, Sid string) (*ProxyV1Session, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Session{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSession'
type ListSessionParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSessionParams) SetPageSize(PageSize int) *ListSessionParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSessionParams) SetLimit(Limit int) *ListSessionParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Session records from the API. Request is executed immediately.
func (c *ApiService) PageSession(ServiceSid string, params *ListSessionParams, pageToken, pageNumber string) (*ListSessionResponse, error) {
	path := "/v1/Services/{ServiceSid}/Sessions"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Session records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSession(ServiceSid string, params *ListSessionParams) ([]ProxyV1Session, error) {
	response, errors := c.StreamSession(ServiceSid, params)

	records := make([]ProxyV1Session, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams Session records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSession(ServiceSid string, params *ListSessionParams) (chan ProxyV1Session, chan error) {
	if params == nil {
		params = &ListSessionParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan ProxyV1Session, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageSession(ServiceSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamSession(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamSession(response *ListSessionResponse, params *ListSessionParams, recordChannel chan ProxyV1Session, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Sessions
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListSessionResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListSessionResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListSessionResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSessionResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSession'
type UpdateSessionParams struct {
	// The [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) date when the Session should expire. If this is value is present, it overrides the `ttl` value.
	DateExpiry *time.Time `json:"DateExpiry,omitempty"`
	// The time, in seconds, when the session will expire. The time is measured from the last Session create or the Session's last Interaction.
	Ttl *int `json:"Ttl,omitempty"`
	//
	Status *string `json:"Status,omitempty"`
	// [Experimental] For accounts with the ProxyAllowParticipantConflict account flag, setting to true enables per-request opt-in to allowing Proxy to return a 400 error (Twilio error code 80604) when a request to set a Session to in-progress would cause Participants with the same Identifier/ProxyIdentifier pair to be active in multiple Sessions. If not provided, requests will be allowed to succeed, and a Debugger notification (80801) will be emitted. Having multiple, active Participants with the same Identifier/ProxyIdentifier pair causes calls and messages from affected Participants to be routed incorrectly. Please note, the default behavior for accounts without the ProxyAllowParticipantConflict flag is to reject the request as described.  This will eventually be the default for all accounts.
	FailOnParticipantConflict *bool `json:"FailOnParticipantConflict,omitempty"`
}

func (params *UpdateSessionParams) SetDateExpiry(DateExpiry time.Time) *UpdateSessionParams {
	params.DateExpiry = &DateExpiry
	return params
}
func (params *UpdateSessionParams) SetTtl(Ttl int) *UpdateSessionParams {
	params.Ttl = &Ttl
	return params
}
func (params *UpdateSessionParams) SetStatus(Status string) *UpdateSessionParams {
	params.Status = &Status
	return params
}
func (params *UpdateSessionParams) SetFailOnParticipantConflict(FailOnParticipantConflict bool) *UpdateSessionParams {
	params.FailOnParticipantConflict = &FailOnParticipantConflict
	return params
}

// Update a specific Session.
func (c *ApiService) UpdateSession(ServiceSid string, Sid string, params *UpdateSessionParams) (*ProxyV1Session, error) {
	path := "/v1/Services/{ServiceSid}/Sessions/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.DateExpiry != nil {
		data.Set("DateExpiry", fmt.Sprint((*params.DateExpiry).Format(time.RFC3339)))
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.FailOnParticipantConflict != nil {
		data.Set("FailOnParticipantConflict", fmt.Sprint(*params.FailOnParticipantConflict))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ProxyV1Session{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
