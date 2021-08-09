/*
 * Twilio - Api
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

// Optional parameters for the method 'CreateSipIpAddress'
type CreateSipIpAddressParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
	CidrPrefixLength *int `json:"CidrPrefixLength,omitempty"`
	// A human readable descriptive text for this resource, up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
	IpAddress *string `json:"IpAddress,omitempty"`
}

func (params *CreateSipIpAddressParams) SetPathAccountSid(PathAccountSid string) *CreateSipIpAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipIpAddressParams) SetCidrPrefixLength(CidrPrefixLength int) *CreateSipIpAddressParams {
	params.CidrPrefixLength = &CidrPrefixLength
	return params
}
func (params *CreateSipIpAddressParams) SetFriendlyName(FriendlyName string) *CreateSipIpAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateSipIpAddressParams) SetIpAddress(IpAddress string) *CreateSipIpAddressParams {
	params.IpAddress = &IpAddress
	return params
}

// Create a new IpAddress resource.
func (c *ApiService) CreateSipIpAddress(IpAccessControlListSid string, params *CreateSipIpAddressParams) (*ApiV2010SipIpAddress, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"IpAccessControlListSid"+"}", IpAccessControlListSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CidrPrefixLength != nil {
		data.Set("CidrPrefixLength", fmt.Sprint(*params.CidrPrefixLength))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IpAddress != nil {
		data.Set("IpAddress", *params.IpAddress)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipIpAddress'
type DeleteSipIpAddressParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipIpAddressParams) SetPathAccountSid(PathAccountSid string) *DeleteSipIpAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete an IpAddress resource.
func (c *ApiService) DeleteSipIpAddress(IpAccessControlListSid string, Sid string, params *DeleteSipIpAddressParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"IpAccessControlListSid"+"}", IpAccessControlListSid, -1)
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

// Optional parameters for the method 'FetchSipIpAddress'
type FetchSipIpAddressParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipIpAddressParams) SetPathAccountSid(PathAccountSid string) *FetchSipIpAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Read one IpAddress resource.
func (c *ApiService) FetchSipIpAddress(IpAccessControlListSid string, Sid string, params *FetchSipIpAddressParams) (*ApiV2010SipIpAddress, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"IpAccessControlListSid"+"}", IpAccessControlListSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipIpAddress'
type ListSipIpAddressParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSipIpAddressParams) SetPathAccountSid(PathAccountSid string) *ListSipIpAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipIpAddressParams) SetPageSize(PageSize int) *ListSipIpAddressParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSipIpAddressParams) SetLimit(Limit int) *ListSipIpAddressParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SipIpAddress records from the API. Request is executed immediately.
func (c *ApiService) PageSipIpAddress(IpAccessControlListSid string, params *ListSipIpAddressParams, pageToken string, pageNumber string) (*ListSipIpAddressResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"IpAccessControlListSid"+"}", IpAccessControlListSid, -1)

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

	ps := &ListSipIpAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipIpAddress records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipIpAddress(IpAccessControlListSid string, params *ListSipIpAddressParams) ([]ApiV2010SipIpAddress, error) {
	if params == nil {
		params = &ListSipIpAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipIpAddress(IpAccessControlListSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010SipIpAddress

	for response != nil {
		records = append(records, response.IpAddresses...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipIpAddressResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSipIpAddressResponse)
	}

	return records, err
}

// Streams SipIpAddress records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipIpAddress(IpAccessControlListSid string, params *ListSipIpAddressParams) (chan ApiV2010SipIpAddress, error) {
	if params == nil {
		params = &ListSipIpAddressParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipIpAddress(IpAccessControlListSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010SipIpAddress, 1)

	go func() {
		for response != nil {
			for item := range response.IpAddresses {
				channel <- response.IpAddresses[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipIpAddressResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSipIpAddressResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSipIpAddressResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipIpAddressResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSipIpAddress'
type UpdateSipIpAddressParams struct {
	// The unique id of the [Account](https://www.twilio.com/docs/iam/api/account) responsible for this resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// An integer representing the length of the CIDR prefix to use with this IP address when accepting traffic. By default the entire IP address is used.
	CidrPrefixLength *int `json:"CidrPrefixLength,omitempty"`
	// A human readable descriptive text for this resource, up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An IP address in dotted decimal notation from which you want to accept traffic. Any SIP requests from this IP address will be allowed by Twilio. IPv4 only supported today.
	IpAddress *string `json:"IpAddress,omitempty"`
}

func (params *UpdateSipIpAddressParams) SetPathAccountSid(PathAccountSid string) *UpdateSipIpAddressParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSipIpAddressParams) SetCidrPrefixLength(CidrPrefixLength int) *UpdateSipIpAddressParams {
	params.CidrPrefixLength = &CidrPrefixLength
	return params
}
func (params *UpdateSipIpAddressParams) SetFriendlyName(FriendlyName string) *UpdateSipIpAddressParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateSipIpAddressParams) SetIpAddress(IpAddress string) *UpdateSipIpAddressParams {
	params.IpAddress = &IpAddress
	return params
}

// Update an IpAddress resource.
func (c *ApiService) UpdateSipIpAddress(IpAccessControlListSid string, Sid string, params *UpdateSipIpAddressParams) (*ApiV2010SipIpAddress, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/IpAccessControlLists/{IpAccessControlListSid}/IpAddresses/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"IpAccessControlListSid"+"}", IpAccessControlListSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.CidrPrefixLength != nil {
		data.Set("CidrPrefixLength", fmt.Sprint(*params.CidrPrefixLength))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IpAddress != nil {
		data.Set("IpAddress", *params.IpAddress)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipIpAddress{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
