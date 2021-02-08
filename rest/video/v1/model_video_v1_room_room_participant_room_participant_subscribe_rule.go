/*
 * Twilio - Video
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// VideoV1RoomRoomParticipantRoomParticipantSubscribeRule struct for VideoV1RoomRoomParticipantRoomParticipantSubscribeRule
type VideoV1RoomRoomParticipantRoomParticipantSubscribeRule struct {
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	ParticipantSid string `json:"ParticipantSid,omitempty"`
	RoomSid string `json:"RoomSid,omitempty"`
	Rules []map[string]interface{} `json:"Rules,omitempty"`
}