/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Bulkexports
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// BulkexportsV1ExportConfiguration struct for BulkexportsV1ExportConfiguration
type BulkexportsV1ExportConfiguration struct {
	// Whether files are automatically generated
	Enabled bool `json:"enabled,omitempty"`
	// URL targeted at export
	WebhookUrl string `json:"webhook_url,omitempty"`
	// Whether to GET or POST to the webhook url
	WebhookMethod string `json:"webhook_method,omitempty"`
	// The type of communication – Messages, Calls, Conferences, and Participants
	ResourceType string `json:"resource_type,omitempty"`
	// The URL of this resource.
	Url string `json:"url,omitempty"`
}
