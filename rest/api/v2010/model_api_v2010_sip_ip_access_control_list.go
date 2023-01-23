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

// ApiV2010SipIpAccessControlList struct for ApiV2010SipIpAccessControlList
type ApiV2010SipIpAccessControlList struct {
	// A string that uniquely identifies this resource
	Sid string `json:"sid,omitempty"`
	// The unique sid that identifies this account
	AccountSid string `json:"account_sid,omitempty"`
	// A human readable description of this resource
	FriendlyName string `json:"friendly_name,omitempty"`
	// The date this resource was created
	DateCreated string `json:"date_created,omitempty"`
	// The date this resource was last updated
	DateUpdated string `json:"date_updated,omitempty"`
	// The IP addresses associated with this resource.
	SubresourceUris map[string]interface{} `json:"subresource_uris,omitempty"`
	// The URI for this resource
	Uri string `json:"uri,omitempty"`
}
