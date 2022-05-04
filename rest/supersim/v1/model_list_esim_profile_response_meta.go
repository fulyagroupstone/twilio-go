/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListEsimProfileResponseMeta struct for ListEsimProfileResponseMeta
type ListEsimProfileResponseMeta struct {
	FirstPageUrl    string `json:"first_page_url,omitempty"`
	Key             string `json:"key,omitempty"`
	NextPageUrl     string `json:"next_page_url,omitempty"`
	Page            int    `json:"page,omitempty"`
	PageSize        int    `json:"page_size,omitempty"`
	PreviousPageUrl string `json:"previous_page_url,omitempty"`
	Url             string `json:"url,omitempty"`
}
