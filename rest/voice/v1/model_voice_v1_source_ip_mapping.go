/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Voice
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

// VoiceV1SourceIpMapping struct for VoiceV1SourceIpMapping
type VoiceV1SourceIpMapping struct {
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// The unique string that identifies an IP Record
	IpRecordSid string `json:"ip_record_sid,omitempty"`
	// The unique string that identifies a SIP Domain
	SipDomainSid string `json:"sip_domain_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the resource
	Url string `json:"url,omitempty"`
}
