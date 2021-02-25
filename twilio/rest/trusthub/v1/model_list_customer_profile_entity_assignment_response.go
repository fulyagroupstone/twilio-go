/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.10.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// ListCustomerProfileEntityAssignmentResponse struct for ListCustomerProfileEntityAssignmentResponse
type ListCustomerProfileEntityAssignmentResponse struct {
	Meta    ListCustomerProfileResponseMeta                            `json:"Meta,omitempty"`
	Results []TrusthubV1CustomerProfileCustomerProfileEntityAssignment `json:"Results,omitempty"`
}