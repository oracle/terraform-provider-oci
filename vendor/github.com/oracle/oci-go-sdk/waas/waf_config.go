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

// WafConfig The Web Application Firewall configuration for the WAAS policy.
type WafConfig struct {

	// The access rules applied to the Web Application Firewall. Used for defining custom access policies with the combination of `ALLOW`, `DETECT`, and `BLOCK` rules, based on different criteria.
	AccessRules []AccessRule `mandatory:"false" json:"accessRules"`

	// The IP address rate limiting settings used to limit the number of requests from an address.
	AddressRateLimiting *AddressRateLimiting `mandatory:"false" json:"addressRateLimiting"`

	// A list of CAPTCHA challenge settings. These are used to challenge requests with a CAPTCHA to block bots.
	Captchas []Captcha `mandatory:"false" json:"captchas"`

	// The device fingerprint challenge settings. Used to detect unique devices based on the device fingerprint information collected in order to block bots.
	DeviceFingerprintChallenge *DeviceFingerprintChallenge `mandatory:"false" json:"deviceFingerprintChallenge"`

	// A list of bots allowed to access the web application.
	GoodBots []GoodBot `mandatory:"false" json:"goodBots"`

	// The human interaction challenge settings. Used to look for natural human interactions such as mouse movements, time on site, and page scrolling to identify bots.
	HumanInteractionChallenge *HumanInteractionChallenge `mandatory:"false" json:"humanInteractionChallenge"`

	// The JavaScript challenge settings. Used to challenge requests with a JavaScript challenge and take the action if a browser has no JavaScript support in order to block bots.
	JsChallenge *JsChallenge `mandatory:"false" json:"jsChallenge"`

	// The key in the map of origins referencing the origin used for the Web Application Firewall. The origin must already be included in `Origins`. Required when creating the `WafConfig` resource, but not on update.
	Origin *string `mandatory:"false" json:"origin"`

	// A list of caching rules applied to the web application.
	CachingRules []CachingRule `mandatory:"false" json:"cachingRules"`

	// A list of the custom protection rule OCIDs and their actions.
	CustomProtectionRules []CustomProtectionRuleSetting `mandatory:"false" json:"customProtectionRules"`

	// The map of origin groups and their keys used to associate origins to the `wafConfig`. Origin groups allow you to apply weights to groups of origins for load balancing purposes. Origins with higher weights will receive larger proportions of client requests.
	// To add additional origins to your WAAS policy, update the `origins` field of a `UpdateWaasPolicy` request.
	OriginGroups []string `mandatory:"false" json:"originGroups"`

	// A list of the protection rules and their details.
	ProtectionRules []ProtectionRule `mandatory:"false" json:"protectionRules"`

	// The settings to apply to protection rules.
	ProtectionSettings *ProtectionSettings `mandatory:"false" json:"protectionSettings"`

	// A list of threat intelligence feeds and the actions to apply to known malicious traffic based on internet intelligence.
	ThreatFeeds []ThreatFeed `mandatory:"false" json:"threatFeeds"`

	// A list of IP addresses that bypass the Web Application Firewall.
	Whitelists []Whitelist `mandatory:"false" json:"whitelists"`
}

func (m WafConfig) String() string {
	return common.PointerString(m)
}
