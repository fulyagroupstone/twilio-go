/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// SupersimV1UsageRecord struct for SupersimV1UsageRecord
type SupersimV1UsageRecord struct {
	AccountSid string `json:"AccountSid,omitempty"`
	DataDownload int32 `json:"DataDownload,omitempty"`
	DataTotal int32 `json:"DataTotal,omitempty"`
	DataUpload int32 `json:"DataUpload,omitempty"`
	FleetSid string `json:"FleetSid,omitempty"`
	IsoCountry string `json:"IsoCountry,omitempty"`
	NetworkSid string `json:"NetworkSid,omitempty"`
	Period map[string]interface{} `json:"Period,omitempty"`
	SimSid string `json:"SimSid,omitempty"`
}