/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// VerifyV2Factor struct for VerifyV2Factor
type VerifyV2Factor struct {
	// Account Sid.
	AccountSid *string `json:"account_sid,omitempty"`
	// Configurations for a `factor_type`.
	Config *interface{} `json:"config,omitempty"`
	// The date this Factor was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date this Factor was updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Entity Sid.
	EntitySid *string `json:"entity_sid,omitempty"`
	// The Type of this Factor
	FactorType *string `json:"factor_type,omitempty"`
	// A human readable description of this resource.
	FriendlyName *string `json:"friendly_name,omitempty"`
	// Unique external identifier of the Entity
	Identity *string `json:"identity,omitempty"`
	// Metadata of the factor.
	Metadata *interface{} `json:"metadata,omitempty"`
	// Service Sid.
	ServiceSid *string `json:"service_sid,omitempty"`
	// A string that uniquely identifies this Factor.
	Sid *string `json:"sid,omitempty"`
	// The Status of this Factor
	Status *string `json:"status,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
