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

// VideoV1RoomRecording struct for VideoV1RoomRecording
type VideoV1RoomRecording struct {
	// The SID of the Account that created the resource
	AccountSid string  `json:"account_sid,omitempty"`
	Status     *string `json:"status,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated time.Time `json:"date_created,omitempty"`
	// The unique string that identifies the resource
	Sid string `json:"sid,omitempty"`
	// The SID of the recording source
	SourceSid string `json:"source_sid,omitempty"`
	// The size of the recorded track in bytes
	Size int64 `json:"size,omitempty"`
	// The absolute URL of the resource
	Url  string  `json:"url,omitempty"`
	Type *string `json:"type,omitempty"`
	// The duration of the recording in seconds
	Duration        *int    `json:"duration,omitempty"`
	ContainerFormat *string `json:"container_format,omitempty"`
	Codec           *string `json:"codec,omitempty"`
	// A list of SIDs related to the Recording
	GroupingSids *interface{} `json:"grouping_sids,omitempty"`
	// The name that was given to the source track of the recording
	TrackName string `json:"track_name,omitempty"`
	// The number of milliseconds between a point in time that is common to all rooms in a group and when the source room of the recording started
	Offset int64 `json:"offset,omitempty"`
	// The URL of the media file associated with the recording when stored externally
	MediaExternalLocation string `json:"media_external_location,omitempty"`
	// The SID of the Room resource the recording is associated with
	RoomSid string `json:"room_sid,omitempty"`
	// The URLs of related resources
	Links map[string]interface{} `json:"links,omitempty"`
}
