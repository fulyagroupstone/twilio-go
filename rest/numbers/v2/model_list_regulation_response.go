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

// ListRegulationResponse struct for ListRegulationResponse
type ListRegulationResponse struct {
	Meta    ListBundleResponseMeta `json:"meta,omitempty"`
	Results []NumbersV2Regulation  `json:"results,omitempty"`
}
