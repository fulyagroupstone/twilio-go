/*
 * Twilio - Taskrouter
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListWorkspaceResponse struct for ListWorkspaceResponse
type ListWorkspaceResponse struct {
	Meta ListWorkspaceResponseMeta `json:"Meta,omitempty"`
	Workspaces []TaskrouterV1Workspace `json:"Workspaces,omitempty"`
}