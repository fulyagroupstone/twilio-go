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

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// ApiV2010CallFeedbackSummary struct for ApiV2010CallFeedbackSummary
type ApiV2010CallFeedbackSummary struct {
	// The unique sid that identifies this account
	AccountSid *string `json:"account_sid,omitempty"`
	// The total number of calls
	CallCount *int `json:"call_count,omitempty"`
	// The total number of calls with a feedback entry
	CallFeedbackCount *int `json:"call_feedback_count,omitempty"`
	// The date this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The date this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// The latest feedback entry date in the summary
	EndDate *string `json:"end_date,omitempty"`
	// Whether the feedback summary includes subaccounts
	IncludeSubaccounts *bool `json:"include_subaccounts,omitempty"`
	// Issues experienced during the call
	Issues *[]interface{} `json:"issues,omitempty"`
	// The average QualityScore of the feedback entries
	QualityScoreAverage *float32 `json:"quality_score_average,omitempty"`
	// The median QualityScore of the feedback entries
	QualityScoreMedian *float32 `json:"quality_score_median,omitempty"`
	// The standard deviation of the quality scores
	QualityScoreStandardDeviation *float32 `json:"quality_score_standard_deviation,omitempty"`
	// A string that uniquely identifies this feedback entry
	Sid *string `json:"sid,omitempty"`
	// The earliest feedback entry date in the summary
	StartDate *string `json:"start_date,omitempty"`
	Status    *string `json:"status,omitempty"`
}

func (response *ApiV2010CallFeedbackSummary) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		AccountSid                    *string        `json:"account_sid"`
		CallCount                     *int           `json:"call_count"`
		CallFeedbackCount             *int           `json:"call_feedback_count"`
		DateCreated                   *string        `json:"date_created"`
		DateUpdated                   *string        `json:"date_updated"`
		EndDate                       *string        `json:"end_date"`
		IncludeSubaccounts            *bool          `json:"include_subaccounts"`
		Issues                        *[]interface{} `json:"issues"`
		QualityScoreAverage           *interface{}   `json:"quality_score_average"`
		QualityScoreMedian            *interface{}   `json:"quality_score_median"`
		QualityScoreStandardDeviation *interface{}   `json:"quality_score_standard_deviation"`
		Sid                           *string        `json:"sid"`
		StartDate                     *string        `json:"start_date"`
		Status                        *string        `json:"status"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = ApiV2010CallFeedbackSummary{
		AccountSid:         raw.AccountSid,
		CallCount:          raw.CallCount,
		CallFeedbackCount:  raw.CallFeedbackCount,
		DateCreated:        raw.DateCreated,
		DateUpdated:        raw.DateUpdated,
		EndDate:            raw.EndDate,
		IncludeSubaccounts: raw.IncludeSubaccounts,
		Issues:             raw.Issues,
		Sid:                raw.Sid,
		StartDate:          raw.StartDate,
		Status:             raw.Status,
	}

	responseQualityScoreAverage, err := client.UnmarshalFloat32(raw.QualityScoreAverage)
	if err != nil {
		return err
	}
	response.QualityScoreAverage = responseQualityScoreAverage

	responseQualityScoreMedian, err := client.UnmarshalFloat32(raw.QualityScoreMedian)
	if err != nil {
		return err
	}
	response.QualityScoreMedian = responseQualityScoreMedian

	responseQualityScoreStandardDeviation, err := client.UnmarshalFloat32(raw.QualityScoreStandardDeviation)
	if err != nil {
		return err
	}
	response.QualityScoreStandardDeviation = responseQualityScoreStandardDeviation

	return
}
