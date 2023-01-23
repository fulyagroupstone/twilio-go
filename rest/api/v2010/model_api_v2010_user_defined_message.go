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

// ApiV2010UserDefinedMessage struct for ApiV2010UserDefinedMessage
type ApiV2010UserDefinedMessage struct {
	// Account SID.
	AccountSid string `json:"account_sid,omitempty"`
	// Call SID.
	CallSid string `json:"call_sid,omitempty"`
	// User Defined Message SID.
	Sid string `json:"sid,omitempty"`
	// The date this User Defined Message was created.
	DateCreated string `json:"date_created,omitempty"`
}
