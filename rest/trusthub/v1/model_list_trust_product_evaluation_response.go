/*
 * Twilio - Trusthub
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTrustProductEvaluationResponse struct for ListTrustProductEvaluationResponse
type ListTrustProductEvaluationResponse struct {
	Meta    ListCustomerProfileResponseMeta    `json:"meta,omitempty"`
	Results []TrusthubV1TrustProductEvaluation `json:"results,omitempty"`
}
