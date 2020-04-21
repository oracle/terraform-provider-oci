// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Acceleration and Security Services API
//
// OCI Web Application Acceleration and Security Services
//

package waas

import (
	"github.com/oracle/oci-go-sdk/common"
)

// WafLog A list of Web Application Firewall log entries. Each entry is a JSON object, including a timestamp property and other fields varying based on log type. Logs record what rules and countermeasures are triggered by requests and are used as a basis to move request handling into block mode. For more information about WAF logs, see Logs (https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/logs.htm).
type WafLog struct {

	// The action taken on the request, either `ALLOW`, `DETECT`, or `BLOCK`.
	Action *string `mandatory:"false" json:"action"`

	// The CAPTCHA action taken on the request, `ALLOW` or `BLOCK`. For more information about
	// CAPTCHAs, see `UpdateCaptchas`.
	CaptchaAction *string `mandatory:"false" json:"captchaAction"`

	// The CAPTCHA challenge answer that was expected.
	CaptchaExpected *string `mandatory:"false" json:"captchaExpected"`

	// The CAPTCHA challenge answer that was received.
	CaptchaReceived *string `mandatory:"false" json:"captchaReceived"`

	// The number of times the CAPTCHA challenge was failed.
	CaptchaFailCount *string `mandatory:"false" json:"captchaFailCount"`

	// The IPv4 address of the requesting client.
	ClientAddress *string `mandatory:"false" json:"clientAddress"`

	// The name of the country where the request originated.
	CountryName *string `mandatory:"false" json:"countryName"`

	// The value of the request's `User-Agent` header field.
	UserAgent *string `mandatory:"false" json:"userAgent"`

	// The `Host` header data of the request.
	Domain *string `mandatory:"false" json:"domain"`

	// A map of protection rule keys to detection message details. Detections are
	// requests that matched the criteria of a protection rule but the rule's
	// action was set to `DETECT`.
	ProtectionRuleDetections map[string]string `mandatory:"false" json:"protectionRuleDetections"`

	// The HTTP method of the request.
	HttpMethod *string `mandatory:"false" json:"httpMethod"`

	// The path and query string of the request.
	RequestUrl *string `mandatory:"false" json:"requestUrl"`

	// The map of the request's header names to their respective values.
	HttpHeaders map[string]string `mandatory:"false" json:"httpHeaders"`

	// The `Referrer` header value of the request.
	Referrer *string `mandatory:"false" json:"referrer"`

	// The status code of the response.
	ResponseCode *int `mandatory:"false" json:"responseCode"`

	// The size in bytes of the response.
	ResponseSize *int `mandatory:"false" json:"responseSize"`

	// The incident key of a request. An incident key is generated for
	// each request processed by the Web Application Firewall and is used to
	// idenitfy blocked requests in applicable logs.
	IncidentKey *string `mandatory:"false" json:"incidentKey"`

	// The hashed signature of the device's fingerprint. For more information,
	// see `DeviceFingerPrintChallenge`.
	Fingerprint *string `mandatory:"false" json:"fingerprint"`

	// The type of device that the request was made from.
	Device *string `mandatory:"false" json:"device"`

	// ISO 3166-1 alpha-2 code of the country from which the request originated.
	// For a list of codes, see ISO's website (https://www.iso.org/obp/ui/#search/code/).
	CountryCode *string `mandatory:"false" json:"countryCode"`

	// A map of header names to values of the request sent to the origin, including any headers
	// appended by the Web Application Firewall.
	RequestHeaders map[string]string `mandatory:"false" json:"requestHeaders"`

	// The `ThreatFeed` key that matched the request. For more information about
	// threat feeds, see `UpdateThreatFeeds`.
	ThreatFeedKey *string `mandatory:"false" json:"threatFeedKey"`

	// The `AccessRule` key that matched the request. For more information about
	// access rules, see `UpdateAccessRules`.
	AccessRuleKey *string `mandatory:"false" json:"accessRuleKey"`

	// The `AddressRateLimiting` key that matched the request. For more information
	// about address rate limiting, see `UpdateWafAddressRateLimiting`.
	AddressRateLimitingKey *string `mandatory:"false" json:"addressRateLimitingKey"`

	// The date and time the Web Application Firewall processed the request and logged it.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The type of log of the request. For more about log types, see Logs (https://docs.cloud.oracle.com/iaas/Content/WAF/Tasks/logs.htm).
	LogType *string `mandatory:"false" json:"logType"`

	// The address of the origin server where the request was sent.
	OriginAddress *string `mandatory:"false" json:"originAddress"`

	// The amount of time it took the origin server to respond to the request, in seconds.
	OriginResponseTime *string `mandatory:"false" json:"originResponseTime"`
}

func (m WafLog) String() string {
	return common.PointerString(m)
}
