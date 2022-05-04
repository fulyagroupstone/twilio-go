/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2Voice struct for PricingV2Voice
type PricingV2Voice struct {
	// The URLs of the related Countries and Numbers resources
	Links *map[string]interface{} `json:"links,omitempty"`
	// The resource name
	Name *string `json:"name,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
