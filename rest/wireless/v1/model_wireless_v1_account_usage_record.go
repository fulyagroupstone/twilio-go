/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Wireless
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// WirelessV1AccountUsageRecord struct for WirelessV1AccountUsageRecord
type WirelessV1AccountUsageRecord struct {
	// The SID of the Account that created the resource
	AccountSid string `json:"account_sid,omitempty"`
	// The time period for which usage is reported
	Period *interface{} `json:"period,omitempty"`
	// An object that describes the aggregated Commands usage for all SIMs during the specified period
	Commands *interface{} `json:"commands,omitempty"`
	// An object that describes the aggregated Data usage for all SIMs over the period
	Data *interface{} `json:"data,omitempty"`
}
