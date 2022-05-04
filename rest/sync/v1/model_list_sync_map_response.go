/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncMapResponse struct for ListSyncMapResponse
type ListSyncMapResponse struct {
	Maps []SyncV1SyncMap         `json:"maps,omitempty"`
	Meta ListServiceResponseMeta `json:"meta,omitempty"`
}
