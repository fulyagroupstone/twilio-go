/*
 * Twilio - Supersim
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// SupersimV1IpCommand struct for SupersimV1IpCommand
type SupersimV1IpCommand struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The IP address of the device that the IP Command was sent to or received from
	DeviceIp *string `json:"device_ip,omitempty"`
	// The port that the IP Command either originated from or was sent to
	DevicePort *int `json:"device_port,omitempty"`
	// The direction of the IP Command
	Direction *string `json:"direction,omitempty"`
	// The payload of the IP Command sent to or from the Super SIM
	Payload *string `json:"payload,omitempty"`
	// The payload type of the IP Command
	PayloadType *string `json:"payload_type,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The ICCID of the Super SIM that this IP Command was sent to or from
	SimIccid *string `json:"sim_iccid,omitempty"`
	// The SID of the Super SIM that this IP Command was sent to or from
	SimSid *string `json:"sim_sid,omitempty"`
	// The status of the IP Command
	Status *string `json:"status,omitempty"`
	// The absolute URL of the IP Command resource
	Url *string `json:"url,omitempty"`
}
