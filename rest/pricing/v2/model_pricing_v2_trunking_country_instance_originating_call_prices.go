/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.28.2
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"

	"github.com/twilio/twilio-go/client"
)

// PricingV2TrunkingCountryInstanceOriginatingCallPrices struct for PricingV2TrunkingCountryInstanceOriginatingCallPrices
type PricingV2TrunkingCountryInstanceOriginatingCallPrices struct {
	BasePrice    float32 `json:"base_price,omitempty"`
	CurrentPrice float32 `json:"current_price,omitempty"`
	NumberType   string  `json:"number_type,omitempty"`
}

func (response *PricingV2TrunkingCountryInstanceOriginatingCallPrices) UnmarshalJSON(bytes []byte) (err error) {
	raw := struct {
		BasePrice    interface{} `json:"base_price"`
		CurrentPrice interface{} `json:"current_price"`
		NumberType   string      `json:"number_type"`
	}{}

	if err = json.Unmarshal(bytes, &raw); err != nil {
		return err
	}

	*response = PricingV2TrunkingCountryInstanceOriginatingCallPrices{
		NumberType: raw.NumberType,
	}

	responseBasePrice, err := client.UnmarshalFloat32(&raw.BasePrice)
	if err != nil {
		return err
	}
	response.BasePrice = *responseBasePrice

	responseCurrentPrice, err := client.UnmarshalFloat32(&raw.CurrentPrice)
	if err != nil {
		return err
	}
	response.CurrentPrice = *responseCurrentPrice

	return
}
