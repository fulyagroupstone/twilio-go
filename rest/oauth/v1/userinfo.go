/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Oauth
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"context"
	"encoding/json"
	"net/url"
)

// Retrieves the consented UserInfo and other claims about the logged-in subject (end-user).
func (c *ApiService) FetchUserInfo() (*OauthV1UserInfo, error) {
	return c.FetchUserInfoWithCtx(context.TODO())
}

// Retrieves the consented UserInfo and other claims about the logged-in subject (end-user).
func (c *ApiService) FetchUserInfoWithCtx(ctx context.Context) (*OauthV1UserInfo, error) {
	path := "/v1/userinfo"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(ctx, c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &OauthV1UserInfo{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}