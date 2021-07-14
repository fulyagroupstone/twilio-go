/*
 * Twilio - Flex
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
)

// Optional parameters for the method 'CreateChannel'
type CreateChannelParams struct {
	// The chat channel's friendly name.
	ChatFriendlyName *string `json:"ChatFriendlyName,omitempty"`
	// The chat channel's unique name.
	ChatUniqueName *string `json:"ChatUniqueName,omitempty"`
	// The chat participant's friendly name.
	ChatUserFriendlyName *string `json:"ChatUserFriendlyName,omitempty"`
	// The SID of the Flex Flow.
	FlexFlowSid *string `json:"FlexFlowSid,omitempty"`
	// The `identity` value that uniquely identifies the new resource's chat User.
	Identity *string `json:"Identity,omitempty"`
	// Whether to create the channel as long-lived.
	LongLived *bool `json:"LongLived,omitempty"`
	// The pre-engagement data.
	PreEngagementData *string `json:"PreEngagementData,omitempty"`
	// The Target Contact Identity, for example the phone number of an SMS.
	Target *string `json:"Target,omitempty"`
	// The Task attributes to be added for the TaskRouter Task.
	TaskAttributes *string `json:"TaskAttributes,omitempty"`
	// The SID of the TaskRouter Task. Only valid when integration type is `task`. `null` for integration types `studio` & `external`
	TaskSid *string `json:"TaskSid,omitempty"`
}

func (params *CreateChannelParams) SetChatFriendlyName(ChatFriendlyName string) *CreateChannelParams {
	params.ChatFriendlyName = &ChatFriendlyName
	return params
}
func (params *CreateChannelParams) SetChatUniqueName(ChatUniqueName string) *CreateChannelParams {
	params.ChatUniqueName = &ChatUniqueName
	return params
}
func (params *CreateChannelParams) SetChatUserFriendlyName(ChatUserFriendlyName string) *CreateChannelParams {
	params.ChatUserFriendlyName = &ChatUserFriendlyName
	return params
}
func (params *CreateChannelParams) SetFlexFlowSid(FlexFlowSid string) *CreateChannelParams {
	params.FlexFlowSid = &FlexFlowSid
	return params
}
func (params *CreateChannelParams) SetIdentity(Identity string) *CreateChannelParams {
	params.Identity = &Identity
	return params
}
func (params *CreateChannelParams) SetLongLived(LongLived bool) *CreateChannelParams {
	params.LongLived = &LongLived
	return params
}
func (params *CreateChannelParams) SetPreEngagementData(PreEngagementData string) *CreateChannelParams {
	params.PreEngagementData = &PreEngagementData
	return params
}
func (params *CreateChannelParams) SetTarget(Target string) *CreateChannelParams {
	params.Target = &Target
	return params
}
func (params *CreateChannelParams) SetTaskAttributes(TaskAttributes string) *CreateChannelParams {
	params.TaskAttributes = &TaskAttributes
	return params
}
func (params *CreateChannelParams) SetTaskSid(TaskSid string) *CreateChannelParams {
	params.TaskSid = &TaskSid
	return params
}

func (c *ApiService) CreateChannel(params *CreateChannelParams) (*FlexV1Channel, error) {
	path := "/v1/Channels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ChatFriendlyName != nil {
		data.Set("ChatFriendlyName", *params.ChatFriendlyName)
	}
	if params != nil && params.ChatUniqueName != nil {
		data.Set("ChatUniqueName", *params.ChatUniqueName)
	}
	if params != nil && params.ChatUserFriendlyName != nil {
		data.Set("ChatUserFriendlyName", *params.ChatUserFriendlyName)
	}
	if params != nil && params.FlexFlowSid != nil {
		data.Set("FlexFlowSid", *params.FlexFlowSid)
	}
	if params != nil && params.Identity != nil {
		data.Set("Identity", *params.Identity)
	}
	if params != nil && params.LongLived != nil {
		data.Set("LongLived", fmt.Sprint(*params.LongLived))
	}
	if params != nil && params.PreEngagementData != nil {
		data.Set("PreEngagementData", *params.PreEngagementData)
	}
	if params != nil && params.Target != nil {
		data.Set("Target", *params.Target)
	}
	if params != nil && params.TaskAttributes != nil {
		data.Set("TaskAttributes", *params.TaskAttributes)
	}
	if params != nil && params.TaskSid != nil {
		data.Set("TaskSid", *params.TaskSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteChannel(Sid string) error {
	path := "/v1/Channels/{Sid}"
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

func (c *ApiService) FetchChannel(Sid string) (*FlexV1Channel, error) {
	path := "/v1/Channels/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &FlexV1Channel{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListChannel'
type ListChannelParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
}

func (params *ListChannelParams) SetPageSize(PageSize int) *ListChannelParams {
	params.PageSize = &PageSize
	return params
}

func (c *ApiService) ListChannel(params *ListChannelParams) (*ListChannelResponse, error) {
	path := "/v1/Channels"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListChannelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}