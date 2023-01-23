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

// ApiV2010RecordingAddOnResultPayload struct for ApiV2010RecordingAddOnResultPayload
type ApiV2010RecordingAddOnResultPayload struct {
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// The SID of the AddOnResult to which the payload belongs
	AddOnResultSid string `json:"add_on_result_sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid string `json:"account_sid,omitempty"`
	// The string that describes the payload
	Label string `json:"label,omitempty"`
	// The SID of the Add-on to which the result belongs
	AddOnSid string `json:"add_on_sid,omitempty"`
	// The SID of the Add-on configuration
	AddOnConfigurationSid string `json:"add_on_configuration_sid,omitempty"`
	// The MIME type of the payload
	ContentType string `json:"content_type,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated string `json:"date_updated,omitempty"`
	// The SID of the recording to which the AddOnResult resource that contains the payload belongs
	ReferenceSid string `json:"reference_sid,omitempty"`
	// A list of related resources identified by their relative URIs
	SubresourceUris map[string]interface{} `json:"subresource_uris,omitempty"`
}
