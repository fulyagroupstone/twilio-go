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

// TaskrouterV1WorkflowStatistics struct for TaskrouterV1WorkflowStatistics
type TaskrouterV1WorkflowStatistics struct {
	// The SID of the Account that created the resource
	AccountSid string `json:"account_sid,omitempty"`
	// An object that contains the cumulative statistics for the Workflow
	Cumulative *interface{} `json:"cumulative,omitempty"`
	// An object that contains the real-time statistics for the Workflow
	Realtime *interface{} `json:"realtime,omitempty"`
	// Returns the list of Tasks that are being controlled by the Workflow with the specified SID value
	WorkflowSid string `json:"workflow_sid,omitempty"`
	// The SID of the Workspace that contains the Workflow
	WorkspaceSid string `json:"workspace_sid,omitempty"`
	// The absolute URL of the Workflow statistics resource
	Url string `json:"url,omitempty"`
}
