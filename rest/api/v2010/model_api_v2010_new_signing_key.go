/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010NewSigningKey struct for ApiV2010NewSigningKey
type ApiV2010NewSigningKey struct {
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName string `json:"friendly_name,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated string `json:"date_updated,omitempty"`
	// The secret your application uses to sign Access Tokens and to authenticate to the REST API.
	Secret string `json:"secret,omitempty"`
}
