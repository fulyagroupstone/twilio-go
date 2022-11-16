/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Flex
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// FlexV1Gooddata struct for FlexV1Gooddata
type FlexV1Gooddata struct {
	// Unique workspace ID in gooddata
	WorkspaceId *string `json:"workspace_id,omitempty"`
	// The session expiry date and time
	SessionExpiry *string `json:"session_expiry,omitempty"`
	// Unique session ID
	SessionId *string `json:"session_id,omitempty"`
	// GoodData login base URL
	GdBaseUrl *string `json:"gd_base_url,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}
