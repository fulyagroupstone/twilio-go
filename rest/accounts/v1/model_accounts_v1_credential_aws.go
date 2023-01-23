/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Accounts
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

// AccountsV1CredentialAws struct for AccountsV1CredentialAws
type AccountsV1CredentialAws struct {
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid string `json:"account_sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName string `json:"friendly_name,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// The URI for this resource, relative to `https://accounts.twilio.com`
	Url string `json:"url,omitempty"`
}
