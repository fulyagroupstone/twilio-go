/*
 * Twilio - Media
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.23.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// MediaV1PlayerStreamerPlaybackGrant struct for MediaV1PlayerStreamerPlaybackGrant
type MediaV1PlayerStreamerPlaybackGrant struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The grant that authorizes the player sdk to connect to the livestream
	Grant *map[string]interface{} `json:"grant,omitempty"`
	// The unique string that identifies the PlayerStreamer associated with this PlaybackGrant.
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}