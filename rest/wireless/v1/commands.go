/*
 * Twilio - Wireless
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

// Optional parameters for the method 'CreateCommand'
type CreateCommandParams struct {
	// The HTTP method we use to call `callback_url`. Can be: `POST` or `GET`, and the default is `POST`.
	CallbackMethod *string `json:"CallbackMethod,omitempty"`
	// The URL we call using the `callback_url` when the Command has finished sending, whether the command was delivered or it failed.
	CallbackUrl *string `json:"CallbackUrl,omitempty"`
	// The message body of the Command. Can be plain text in text mode or a Base64 encoded byte string in binary mode.
	Command *string `json:"Command,omitempty"`
	// The mode to use when sending the SMS message. Can be: `text` or `binary`. The default SMS mode is `text`.
	CommandMode *string `json:"CommandMode,omitempty"`
	// Whether to request delivery receipt from the recipient. For Commands that request delivery receipt, the Command state transitions to 'delivered' once the server has received a delivery receipt from the device. The default value is `true`.
	DeliveryReceiptRequested *bool `json:"DeliveryReceiptRequested,omitempty"`
	// Whether to include the SID of the command in the message body. Can be: `none`, `start`, or `end`, and the default behavior is `none`. When sending a Command to a SIM in text mode, we can automatically include the SID of the Command in the message body, which could be used to ensure that the device does not process the same Command more than once.  A value of `start` will prepend the message with the Command SID, and `end` will append it to the end, separating the Command SID from the message body with a space. The length of the Command SID is included in the 160 character limit so the SMS body must be 128 characters or less before the Command SID is included.
	IncludeSid *string `json:"IncludeSid,omitempty"`
	// The `sid` or `unique_name` of the [SIM](https://www.twilio.com/docs/wireless/api/sim-resource) to send the Command to.
	Sim *string `json:"Sim,omitempty"`
}

func (params *CreateCommandParams) SetCallbackMethod(CallbackMethod string) *CreateCommandParams {
	params.CallbackMethod = &CallbackMethod
	return params
}
func (params *CreateCommandParams) SetCallbackUrl(CallbackUrl string) *CreateCommandParams {
	params.CallbackUrl = &CallbackUrl
	return params
}
func (params *CreateCommandParams) SetCommand(Command string) *CreateCommandParams {
	params.Command = &Command
	return params
}
func (params *CreateCommandParams) SetCommandMode(CommandMode string) *CreateCommandParams {
	params.CommandMode = &CommandMode
	return params
}
func (params *CreateCommandParams) SetDeliveryReceiptRequested(DeliveryReceiptRequested bool) *CreateCommandParams {
	params.DeliveryReceiptRequested = &DeliveryReceiptRequested
	return params
}
func (params *CreateCommandParams) SetIncludeSid(IncludeSid string) *CreateCommandParams {
	params.IncludeSid = &IncludeSid
	return params
}
func (params *CreateCommandParams) SetSim(Sim string) *CreateCommandParams {
	params.Sim = &Sim
	return params
}

// Send a Command to a Sim.
func (c *ApiService) CreateCommand(params *CreateCommandParams) (*WirelessV1Command, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CallbackMethod != nil {
		data.Set("CallbackMethod", *params.CallbackMethod)
	}
	if params != nil && params.CallbackUrl != nil {
		data.Set("CallbackUrl", *params.CallbackUrl)
	}
	if params != nil && params.Command != nil {
		data.Set("Command", *params.Command)
	}
	if params != nil && params.CommandMode != nil {
		data.Set("CommandMode", *params.CommandMode)
	}
	if params != nil && params.DeliveryReceiptRequested != nil {
		data.Set("DeliveryReceiptRequested", fmt.Sprint(*params.DeliveryReceiptRequested))
	}
	if params != nil && params.IncludeSid != nil {
		data.Set("IncludeSid", *params.IncludeSid)
	}
	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Delete a Command instance from your account.
func (c *ApiService) DeleteCommand(Sid string) error {
	path := "/v1/Commands/{Sid}"
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

// Fetch a Command instance from your account.
func (c *ApiService) FetchCommand(Sid string) (*WirelessV1Command, error) {
	path := "/v1/Commands/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &WirelessV1Command{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListCommand'
type ListCommandParams struct {
	// The `sid` or `unique_name` of the [Sim resources](https://www.twilio.com/docs/wireless/api/sim-resource) to read.
	Sim *string `json:"Sim,omitempty"`
	// The status of the resources to read. Can be: `queued`, `sent`, `delivered`, `received`, or `failed`.
	Status *string `json:"Status,omitempty"`
	// Only return Commands with this direction value.
	Direction *string `json:"Direction,omitempty"`
	// Only return Commands with this transport value. Can be: `sms` or `ip`.
	Transport *string `json:"Transport,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListCommandParams) SetSim(Sim string) *ListCommandParams {
	params.Sim = &Sim
	return params
}
func (params *ListCommandParams) SetStatus(Status string) *ListCommandParams {
	params.Status = &Status
	return params
}
func (params *ListCommandParams) SetDirection(Direction string) *ListCommandParams {
	params.Direction = &Direction
	return params
}
func (params *ListCommandParams) SetTransport(Transport string) *ListCommandParams {
	params.Transport = &Transport
	return params
}
func (params *ListCommandParams) SetPageSize(PageSize int) *ListCommandParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListCommandParams) SetLimit(Limit int) *ListCommandParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Command records from the API. Request is executed immediately.
func (c *ApiService) PageCommand(params *ListCommandParams, pageToken string, pageNumber string) (*ListCommandResponse, error) {
	path := "/v1/Commands"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Sim != nil {
		data.Set("Sim", *params.Sim)
	}
	if params != nil && params.Status != nil {
		data.Set("Status", *params.Status)
	}
	if params != nil && params.Direction != nil {
		data.Set("Direction", *params.Direction)
	}
	if params != nil && params.Transport != nil {
		data.Set("Transport", *params.Transport)
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

	ps := &ListCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Command records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListCommand(params *ListCommandParams) ([]WirelessV1Command, error) {
	if params == nil {
		params = &ListCommandParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCommand(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []WirelessV1Command

	for response != nil {
		records = append(records, response.Commands...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCommandResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListCommandResponse)
	}

	return records, err
}

// Streams Command records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamCommand(params *ListCommandParams) (chan WirelessV1Command, error) {
	if params == nil {
		params = &ListCommandParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageCommand(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan WirelessV1Command, 1)

	go func() {
		for response != nil {
			for item := range response.Commands {
				channel <- response.Commands[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListCommandResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListCommandResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListCommandResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListCommandResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
