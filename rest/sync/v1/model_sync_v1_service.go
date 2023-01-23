/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Sync
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// SyncV1Service struct for SyncV1Service
type SyncV1Service struct {
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName string `json:"unique_name,omitempty"`
	// The SID of the Account that created the resource
	AccountSid string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName string `json:"friendly_name,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Service resource
	Url string `json:"url,omitempty"`
	// The URL we call when Sync objects are manipulated
	WebhookUrl string `json:"webhook_url,omitempty"`
	// Whether the Service instance should call webhook_url when the REST API is used to update Sync objects
	WebhooksFromRestEnabled bool `json:"webhooks_from_rest_enabled,omitempty"`
	// Whether the service instance calls webhook_url when client endpoints connect to Sync
	ReachabilityWebhooksEnabled bool `json:"reachability_webhooks_enabled,omitempty"`
	// Whether token identities in the Service must be granted access to Sync objects by using the Permissions resource
	AclEnabled bool `json:"acl_enabled,omitempty"`
	// Whether every endpoint_disconnected event occurs after a configurable delay
	ReachabilityDebouncingEnabled bool `json:"reachability_debouncing_enabled,omitempty"`
	// The reachability event delay in milliseconds
	ReachabilityDebouncingWindow int `json:"reachability_debouncing_window,omitempty"`
	// The URLs of related resources
	Links map[string]interface{} `json:"links,omitempty"`
}
