/*
 * Twilio - Notify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.29.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
)

// Optional parameters for the method 'CreateNotification'
type CreateNotificationParams struct {
	// The actions to display for the notification. For APNS, translates to the `aps.category` value. For GCM, translates to the `data.twi_action` value. For SMS, this parameter is not supported and is omitted from deliveries to those channels.
	Action *string `json:"Action,omitempty"`
	// Deprecated.
	Alexa *interface{} `json:"Alexa,omitempty"`
	// The APNS-specific payload that overrides corresponding attributes in the generic payload for APNS Bindings. This property maps to the APNS `Payload` item, therefore the `aps` key must be used to change standard attributes. Adds custom key-value pairs to the root of the dictionary. See the [APNS documentation](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html) for more details. We reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed.
	Apn *interface{} `json:"Apn,omitempty"`
	// The notification text. For FCM and GCM, translates to `data.twi_body`. For APNS, translates to `aps.alert.body`. For SMS, translates to `body`. SMS requires either this `body` value, or `media_urls` attribute defined in the `sms` parameter of the notification.
	Body *string `json:"Body,omitempty"`
	// The custom key-value pairs of the notification's payload. For FCM and GCM, this value translates to `data` in the FCM and GCM payloads. FCM and GCM [reserve certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref) that cannot be used in those channels. For APNS, attributes of `data` are inserted into the APNS payload as custom properties outside of the `aps` dictionary. In all channels, we reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed and are rejected as 400 Bad request with no delivery attempted. For SMS, this parameter is not supported and is omitted from deliveries to those channels.
	Data *interface{} `json:"Data,omitempty"`
	// URL to send webhooks.
	DeliveryCallbackUrl *string `json:"DeliveryCallbackUrl,omitempty"`
	// Deprecated.
	FacebookMessenger *interface{} `json:"FacebookMessenger,omitempty"`
	// The FCM-specific payload that overrides corresponding attributes in the generic payload for FCM Bindings. This property maps to the root JSON dictionary. See the [FCM documentation](https://firebase.google.com/docs/cloud-messaging/http-server-ref#downstream) for more details. Target parameters `to`, `registration_ids`, `condition`, and `notification_key` are not allowed in this parameter. We reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed. FCM also [reserves certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref), which cannot be used in that channel.
	Fcm *interface{} `json:"Fcm,omitempty"`
	// The GCM-specific payload that overrides corresponding attributes in the generic payload for GCM Bindings.  This property maps to the root JSON dictionary. See the [GCM documentation](https://firebase.google.com/docs/cloud-messaging/http-server-ref) for more details. Target parameters `to`, `registration_ids`, and `notification_key` are not allowed. We reserve keys that start with `twi_` for future use. Custom keys that start with `twi_` are not allowed. GCM also [reserves certain keys](https://firebase.google.com/docs/cloud-messaging/http-server-ref).
	Gcm *interface{} `json:"Gcm,omitempty"`
	// The `identity` value that uniquely identifies the new resource's [User](https://www.twilio.com/docs/chat/rest/user-resource) within the [Service](https://www.twilio.com/docs/notify/api/service-resource). Delivery will be attempted only to Bindings with an Identity in this list. No more than 20 items are allowed in this list.
	Identity *[]string `json:"Identity,omitempty"`
	// The priority of the notification. Can be: `low` or `high` and the default is `high`. A value of `low` optimizes the client app's battery consumption; however, notifications may be delivered with unspecified delay. For FCM and GCM, `low` priority is the same as `Normal` priority. For APNS `low` priority is the same as `5`. A value of `high` sends the notification immediately, and can wake up a sleeping device. For FCM and GCM, `high` is the same as `High` priority. For APNS, `high` is a priority `10`. SMS does not support this property.
	Priority *string `json:"Priority,omitempty"`
	// The Segment resource is deprecated. Use the `tag` parameter, instead.
	Segment *[]string `json:"Segment,omitempty"`
	// The SMS-specific payload that overrides corresponding attributes in the generic payload for SMS Bindings.  Each attribute in this value maps to the corresponding `form` parameter of the Twilio [Message](https://www.twilio.com/docs/sms/send-messages) resource.  These parameters of the Message resource are supported in snake case format: `body`, `media_urls`, `status_callback`, and `max_price`.  The `status_callback` parameter overrides the corresponding parameter in the messaging service, if configured. The `media_urls` property expects a JSON array.
	Sms *interface{} `json:"Sms,omitempty"`
	// The name of the sound to be played for the notification. For FCM and GCM, this Translates to `data.twi_sound`.  For APNS, this translates to `aps.sound`.  SMS does not support this property.
	Sound *string `json:"Sound,omitempty"`
	// A tag that selects the Bindings to notify. Repeat this parameter to specify more than one tag, up to a total of 5 tags. The implicit tag `all` is available to notify all Bindings in a Service instance. Similarly, the implicit tags `apn`, `fcm`, `gcm`, `sms` and `facebook-messenger` are available to notify all Bindings in a specific channel.
	Tag *[]string `json:"Tag,omitempty"`
	// The notification title. For FCM and GCM, this translates to the `data.twi_title` value. For APNS, this translates to the `aps.alert.title` value. SMS does not support this property. This field is not visible on iOS phones and tablets but appears on Apple Watch and Android devices.
	Title *string `json:"Title,omitempty"`
	// The destination address specified as a JSON string.  Multiple `to_binding` parameters can be included but the total size of the request entity should not exceed 1MB. This is typically sufficient for 10,000 phone numbers.
	ToBinding *[]string `json:"ToBinding,omitempty"`
	// How long, in seconds, the notification is valid. Can be an integer between 0 and 2,419,200, which is 4 weeks, the default and the maximum supported time to live (TTL). Delivery should be attempted if the device is offline until the TTL elapses. Zero means that the notification delivery is attempted immediately, only once, and is not stored for future delivery. SMS does not support this property.
	Ttl *int `json:"Ttl,omitempty"`
}

func (params *CreateNotificationParams) SetAction(Action string) *CreateNotificationParams {
	params.Action = &Action
	return params
}
func (params *CreateNotificationParams) SetAlexa(Alexa interface{}) *CreateNotificationParams {
	params.Alexa = &Alexa
	return params
}
func (params *CreateNotificationParams) SetApn(Apn interface{}) *CreateNotificationParams {
	params.Apn = &Apn
	return params
}
func (params *CreateNotificationParams) SetBody(Body string) *CreateNotificationParams {
	params.Body = &Body
	return params
}
func (params *CreateNotificationParams) SetData(Data interface{}) *CreateNotificationParams {
	params.Data = &Data
	return params
}
func (params *CreateNotificationParams) SetDeliveryCallbackUrl(DeliveryCallbackUrl string) *CreateNotificationParams {
	params.DeliveryCallbackUrl = &DeliveryCallbackUrl
	return params
}
func (params *CreateNotificationParams) SetFacebookMessenger(FacebookMessenger interface{}) *CreateNotificationParams {
	params.FacebookMessenger = &FacebookMessenger
	return params
}
func (params *CreateNotificationParams) SetFcm(Fcm interface{}) *CreateNotificationParams {
	params.Fcm = &Fcm
	return params
}
func (params *CreateNotificationParams) SetGcm(Gcm interface{}) *CreateNotificationParams {
	params.Gcm = &Gcm
	return params
}
func (params *CreateNotificationParams) SetIdentity(Identity []string) *CreateNotificationParams {
	params.Identity = &Identity
	return params
}
func (params *CreateNotificationParams) SetPriority(Priority string) *CreateNotificationParams {
	params.Priority = &Priority
	return params
}
func (params *CreateNotificationParams) SetSegment(Segment []string) *CreateNotificationParams {
	params.Segment = &Segment
	return params
}
func (params *CreateNotificationParams) SetSms(Sms interface{}) *CreateNotificationParams {
	params.Sms = &Sms
	return params
}
func (params *CreateNotificationParams) SetSound(Sound string) *CreateNotificationParams {
	params.Sound = &Sound
	return params
}
func (params *CreateNotificationParams) SetTag(Tag []string) *CreateNotificationParams {
	params.Tag = &Tag
	return params
}
func (params *CreateNotificationParams) SetTitle(Title string) *CreateNotificationParams {
	params.Title = &Title
	return params
}
func (params *CreateNotificationParams) SetToBinding(ToBinding []string) *CreateNotificationParams {
	params.ToBinding = &ToBinding
	return params
}
func (params *CreateNotificationParams) SetTtl(Ttl int) *CreateNotificationParams {
	params.Ttl = &Ttl
	return params
}

func (c *ApiService) CreateNotification(ServiceSid string, params *CreateNotificationParams) (*NotifyV1Notification, error) {
	path := "/v1/Services/{ServiceSid}/Notifications"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Action != nil {
		data.Set("Action", *params.Action)
	}
	if params != nil && params.Alexa != nil {
		v, err := json.Marshal(params.Alexa)

		if err != nil {
			return nil, err
		}

		data.Set("Alexa", string(v))
	}
	if params != nil && params.Apn != nil {
		v, err := json.Marshal(params.Apn)

		if err != nil {
			return nil, err
		}

		data.Set("Apn", string(v))
	}
	if params != nil && params.Body != nil {
		data.Set("Body", *params.Body)
	}
	if params != nil && params.Data != nil {
		v, err := json.Marshal(params.Data)

		if err != nil {
			return nil, err
		}

		data.Set("Data", string(v))
	}
	if params != nil && params.DeliveryCallbackUrl != nil {
		data.Set("DeliveryCallbackUrl", *params.DeliveryCallbackUrl)
	}
	if params != nil && params.FacebookMessenger != nil {
		v, err := json.Marshal(params.FacebookMessenger)

		if err != nil {
			return nil, err
		}

		data.Set("FacebookMessenger", string(v))
	}
	if params != nil && params.Fcm != nil {
		v, err := json.Marshal(params.Fcm)

		if err != nil {
			return nil, err
		}

		data.Set("Fcm", string(v))
	}
	if params != nil && params.Gcm != nil {
		v, err := json.Marshal(params.Gcm)

		if err != nil {
			return nil, err
		}

		data.Set("Gcm", string(v))
	}
	if params != nil && params.Identity != nil {
		for _, item := range *params.Identity {
			data.Add("Identity", item)
		}
	}
	if params != nil && params.Priority != nil {
		data.Set("Priority", *params.Priority)
	}
	if params != nil && params.Segment != nil {
		for _, item := range *params.Segment {
			data.Add("Segment", item)
		}
	}
	if params != nil && params.Sms != nil {
		v, err := json.Marshal(params.Sms)

		if err != nil {
			return nil, err
		}

		data.Set("Sms", string(v))
	}
	if params != nil && params.Sound != nil {
		data.Set("Sound", *params.Sound)
	}
	if params != nil && params.Tag != nil {
		for _, item := range *params.Tag {
			data.Add("Tag", item)
		}
	}
	if params != nil && params.Title != nil {
		data.Set("Title", *params.Title)
	}
	if params != nil && params.ToBinding != nil {
		for _, item := range *params.ToBinding {
			data.Add("ToBinding", item)
		}
	}
	if params != nil && params.Ttl != nil {
		data.Set("Ttl", fmt.Sprint(*params.Ttl))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &NotifyV1Notification{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
