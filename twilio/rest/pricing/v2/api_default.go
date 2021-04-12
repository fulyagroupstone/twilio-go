/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.13.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL string
	client  twilio.BaseClient
}

func NewDefaultApiService(client twilio.BaseClient) *DefaultApiService {
	return &DefaultApiService{
		client:  client,
		baseURL: "https://pricing.twilio.com",
	}
}

/*
* FetchVoiceCountry Method for FetchVoiceCountry
* Fetch a specific Country.
* @param IsoCountry The [ISO country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) of the origin-based voice pricing information to fetch.
* @return PricingV2VoiceVoiceCountryInstance
 */
func (c *DefaultApiService) FetchVoiceCountry(IsoCountry string) (*PricingV2VoiceVoiceCountryInstance, error) {
	path := "/v2/Voice/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV2VoiceVoiceCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// FetchVoiceNumberParams Optional parameters for the method 'FetchVoiceNumber'
type FetchVoiceNumberParams struct {
	OriginationNumber *string `json:"OriginationNumber,omitempty"`
}

/*
* FetchVoiceNumber Method for FetchVoiceNumber
* Fetch pricing information for a specific destination and, optionally, origination phone number.
* @param DestinationNumber The destination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.
* @param optional nil or *FetchVoiceNumberParams - Optional Parameters:
* @param "OriginationNumber" (string) - The origination phone number, in [E.164](https://www.twilio.com/docs/glossary/what-e164) format, for which to fetch the origin-based voice pricing information. E.164 format consists of a + followed by the country code and subscriber number.
* @return PricingV2VoiceVoiceNumber
 */
func (c *DefaultApiService) FetchVoiceNumber(DestinationNumber string, params *FetchVoiceNumberParams) (*PricingV2VoiceVoiceNumber, error) {
	path := "/v2/Voice/Numbers/{DestinationNumber}"
	path = strings.Replace(path, "{"+"DestinationNumber"+"}", DestinationNumber, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.OriginationNumber != nil {
		data.Set("OriginationNumber", *params.OriginationNumber)
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV2VoiceVoiceNumber{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// ListVoiceCountryParams Optional parameters for the method 'ListVoiceCountry'
type ListVoiceCountryParams struct {
	PageSize *int32 `json:"PageSize,omitempty"`
}

/*
* ListVoiceCountry Method for ListVoiceCountry
* @param optional nil or *ListVoiceCountryParams - Optional Parameters:
* @param "PageSize" (int32) - How many resources to return in each list page. The default is 50, and the maximum is 1000.
* @return ListVoiceCountryResponse
 */
func (c *DefaultApiService) ListVoiceCountry(params *ListVoiceCountryParams) (*ListVoiceCountryResponse, error) {
	path := "/v2/Voice/Countries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.client.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListVoiceCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
