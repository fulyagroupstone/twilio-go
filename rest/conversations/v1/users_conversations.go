/*
 * Twilio - Conversations
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
	"time"

	"github.com/twilio/twilio-go/client"
)

// Delete a specific User Conversation.
func (c *ApiService) DeleteUserConversation(UserSid string, ConversationSid string) error {
	path := "/v1/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a specific User Conversation.
func (c *ApiService) FetchUserConversation(UserSid string, ConversationSid string) (*ConversationsV1UserConversation, error) {
	path := "/v1/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1UserConversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListUserConversation'
type ListUserConversationParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListUserConversationParams) SetPageSize(PageSize int) *ListUserConversationParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListUserConversationParams) SetLimit(Limit int) *ListUserConversationParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of UserConversation records from the API. Request is executed immediately.
func (c *ApiService) PageUserConversation(UserSid string, params *ListUserConversationParams, pageToken string, pageNumber string) (*ListUserConversationResponse, error) {
	path := "/v1/Users/{UserSid}/Conversations"

	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)

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

	ps := &ListUserConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists UserConversation records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListUserConversation(UserSid string, params *ListUserConversationParams) ([]ConversationsV1UserConversation, error) {
	if params == nil {
		params = &ListUserConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUserConversation(UserSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1UserConversation

	for response != nil {
		records = append(records, response.Conversations...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUserConversationResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListUserConversationResponse)
	}

	return records, err
}

// Streams UserConversation records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamUserConversation(UserSid string, params *ListUserConversationParams) (chan ConversationsV1UserConversation, error) {
	if params == nil {
		params = &ListUserConversationParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageUserConversation(UserSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1UserConversation, 1)

	go func() {
		for response != nil {
			for item := range response.Conversations {
				channel <- response.Conversations[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListUserConversationResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListUserConversationResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListUserConversationResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListUserConversationResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateUserConversation'
type UpdateUserConversationParams struct {
	// The index of the last Message in the Conversation that the Participant has read.
	LastReadMessageIndex *int `json:"LastReadMessageIndex,omitempty"`
	// The date of the last message read in conversation by the user, given in ISO 8601 format.
	LastReadTimestamp *time.Time `json:"LastReadTimestamp,omitempty"`
	// The Notification Level of this User Conversation. One of `default` or `muted`.
	NotificationLevel *string `json:"NotificationLevel,omitempty"`
}

func (params *UpdateUserConversationParams) SetLastReadMessageIndex(LastReadMessageIndex int) *UpdateUserConversationParams {
	params.LastReadMessageIndex = &LastReadMessageIndex
	return params
}
func (params *UpdateUserConversationParams) SetLastReadTimestamp(LastReadTimestamp time.Time) *UpdateUserConversationParams {
	params.LastReadTimestamp = &LastReadTimestamp
	return params
}
func (params *UpdateUserConversationParams) SetNotificationLevel(NotificationLevel string) *UpdateUserConversationParams {
	params.NotificationLevel = &NotificationLevel
	return params
}

// Update a specific User Conversation.
func (c *ApiService) UpdateUserConversation(UserSid string, ConversationSid string, params *UpdateUserConversationParams) (*ConversationsV1UserConversation, error) {
	path := "/v1/Users/{UserSid}/Conversations/{ConversationSid}"
	path = strings.Replace(path, "{"+"UserSid"+"}", UserSid, -1)
	path = strings.Replace(path, "{"+"ConversationSid"+"}", ConversationSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LastReadMessageIndex != nil {
		data.Set("LastReadMessageIndex", fmt.Sprint(*params.LastReadMessageIndex))
	}
	if params != nil && params.LastReadTimestamp != nil {
		data.Set("LastReadTimestamp", fmt.Sprint((*params.LastReadTimestamp).Format(time.RFC3339)))
	}
	if params != nil && params.NotificationLevel != nil {
		data.Set("NotificationLevel", *params.NotificationLevel)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1UserConversation{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
