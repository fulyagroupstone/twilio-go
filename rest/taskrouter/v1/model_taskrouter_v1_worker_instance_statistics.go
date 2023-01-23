/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Taskrouter
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// TaskrouterV1WorkerInstanceStatistics struct for TaskrouterV1WorkerInstanceStatistics
type TaskrouterV1WorkerInstanceStatistics struct {
	// The SID of the Account that created the resource
	AccountSid string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the Worker
	Cumulative *interface{} `json:"cumulative,omitempty"`
	// The SID of the Worker that contains the WorkerChannel
	WorkerSid string `json:"worker_sid,omitempty"`
	// The SID of the Workspace that contains the WorkerChannel
	WorkspaceSid string `json:"workspace_sid,omitempty"`
	// The absolute URL of the WorkerChannel statistics resource
	Url string `json:"url,omitempty"`
}
