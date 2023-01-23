/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Messaging
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

// MessagingV1BrandRegistrations struct for MessagingV1BrandRegistrations
type MessagingV1BrandRegistrations struct {
	// A2P BrandRegistration Sid
	Sid string `json:"sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid string `json:"account_sid,omitempty"`
	// A2P Messaging Profile Bundle BundleSid
	CustomerProfileBundleSid string `json:"customer_profile_bundle_sid,omitempty"`
	// A2P Messaging Profile Bundle BundleSid
	A2pProfileBundleSid string `json:"a2p_profile_bundle_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// Type of brand. One of: \"STANDARD\", \"STARTER\".
	BrandType string  `json:"brand_type,omitempty"`
	Status    *string `json:"status,omitempty"`
	// Campaign Registry (TCR) Brand ID
	TcrId string `json:"tcr_id,omitempty"`
	// A reason why brand registration has failed
	FailureReason string `json:"failure_reason,omitempty"`
	// The absolute URL of the Brand Registration
	Url string `json:"url,omitempty"`
	// Brand score
	BrandScore *int `json:"brand_score,omitempty"`
	// Brand feedback
	BrandFeedback  *[]string `json:"brand_feedback,omitempty"`
	IdentityStatus *string   `json:"identity_status,omitempty"`
	// Russell 3000
	Russell3000 bool `json:"russell_3000,omitempty"`
	// Government Entity
	GovernmentEntity bool `json:"government_entity,omitempty"`
	// Tax Exempt Status
	TaxExemptStatus *string `json:"tax_exempt_status,omitempty"`
	// Skip Automatic Secondary Vetting
	SkipAutomaticSecVet bool `json:"skip_automatic_sec_vet,omitempty"`
	// A boolean that specifies whether brand should be a mock or not. If true, brand will be registered as a mock brand. Defaults to false if no value is provided.
	Mock  bool                   `json:"mock,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
}
