/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Insights
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// InsightsV1Annotation struct for InsightsV1Annotation
type InsightsV1Annotation struct {
	// Call SID.
	CallSid string `json:"call_sid,omitempty"`
	// Account SID.
	AccountSid        string  `json:"account_sid,omitempty"`
	AnsweredBy        *string `json:"answered_by,omitempty"`
	ConnectivityIssue *string `json:"connectivity_issue,omitempty"`
	// Indicates if the call had audio quality issues.
	QualityIssues []string `json:"quality_issues,omitempty"`
	// Call spam indicator
	Spam bool `json:"spam,omitempty"`
	// Call Score
	CallScore *int `json:"call_score,omitempty"`
	// User comments
	Comment string `json:"comment,omitempty"`
	// Call tag for incidents or support ticket
	Incident string `json:"incident,omitempty"`
	// The URL of this resource.
	Url string `json:"url,omitempty"`
}
