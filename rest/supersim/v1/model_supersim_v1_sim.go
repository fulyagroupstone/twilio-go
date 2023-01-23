/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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

// SupersimV1Sim struct for SupersimV1Sim
type SupersimV1Sim struct {
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// An application-defined string that uniquely identifies the resource
	UniqueName string `json:"unique_name,omitempty"`
	// The SID of the Account that the Super SIM belongs to
	AccountSid string `json:"account_sid,omitempty"`
	// The ICCID associated with the SIM
	Iccid  string  `json:"iccid,omitempty"`
	Status *string `json:"status,omitempty"`
	// The unique ID of the Fleet configured for this SIM
	FleetSid string `json:"fleet_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// The absolute URL of the Sim Resource
	Url   string                 `json:"url,omitempty"`
	Links map[string]interface{} `json:"links,omitempty"`
}
