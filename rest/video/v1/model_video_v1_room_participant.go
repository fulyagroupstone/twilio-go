/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Video
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

// VideoV1RoomParticipant struct for VideoV1RoomParticipant
type VideoV1RoomParticipant struct {
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// The SID of the participant's room
	RoomSid string `json:"room_sid,omitempty"`
	// The SID of the Account that created the resource
	AccountSid string  `json:"account_sid,omitempty"`
	Status     *string `json:"status,omitempty"`
	// The string that identifies the resource's User
	Identity string `json:"identity,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// The time of participant connected to the room in ISO 8601 format
	StartTime time.Time `json:"start_time,omitempty"`
	// The time when the participant disconnected from the room in ISO 8601 format
	EndTime time.Time `json:"end_time,omitempty"`
	// Duration of time in seconds the participant was connected
	Duration *int `json:"duration,omitempty"`
	// The absolute URL of the resource
	Url string `json:"url,omitempty"`
	// The URLs of related resources
	Links map[string]interface{} `json:"links,omitempty"`
}
