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

// WafConfigDetails The Web Application Firewall configuration for the WAAS policy creation.
type WafConfigDetails struct {

	// The access rules applied to the Web Application Firewall. Access rules allow custom content access policies to be defined and `ALLOW`, `DETECT`, or `BLOCK` actions to be taken on a request when specified criteria are met.
	AccessRules []AccessRule `mandatory:"false" json:"accessRules"`

	// The settings used to limit the number of requests from an IP address.
	AddressRateLimiting *AddressRateLimiting `mandatory:"false" json:"addressRateLimiting"`

	// A list of CAPTCHA challenge settings. CAPTCHAs challenge requests to ensure a human is attempting to reach the specified URL and not a bot.
	Captchas []Captcha `mandatory:"false" json:"captchas"`

	// The device fingerprint challenge settings. Blocks bots based on unique device fingerprint information.
	DeviceFingerprintChallenge *DeviceFingerprintChallenge `mandatory:"false" json:"deviceFingerprintChallenge"`

	// The human interaction challenge settings. Detects natural human interactions such as mouse movements, time on site, and page scrolling to identify bots.
	HumanInteractionChallenge *HumanInteractionChallenge `mandatory:"false" json:"humanInteractionChallenge"`

	// The JavaScript challenge settings. Blocks bots by challenging requests from browsers that have no JavaScript support.
	JsChallenge *JsChallenge `mandatory:"false" json:"jsChallenge"`

	// The key in the map of origins referencing the origin used for the Web Application Firewall. The origin must already be included in `Origins`. Required when creating the `WafConfig` resource, but is not required upon updating the configuration.
	Origin *string `mandatory:"false" json:"origin"`

	// A list of caching rules applied to the web application.
	CachingRules []CachingRule `mandatory:"false" json:"cachingRules"`

	// A list of the custom protection rule OCIDs and their actions.
	CustomProtectionRules []CustomProtectionRuleSetting `mandatory:"false" json:"customProtectionRules"`

	// The map of origin groups and their keys used to associate origins to the `wafConfig`. Origin groups allow you to apply weights to groups of origins for load balancing purposes. Origins with higher weights will receive larger proportions of client requests.
	// To add additional origins to your WAAS policy, update the `origins` field of a `UpdateWaasPolicy` request.
	OriginGroups []string `mandatory:"false" json:"originGroups"`

	// The settings applied to protection rules.
	ProtectionSettings *ProtectionSettings `mandatory:"false" json:"protectionSettings"`

	// A list of IP addresses that bypass the Web Application Firewall.
	Whitelists []Whitelist `mandatory:"false" json:"whitelists"`
}

func (m WafConfigDetails) String() string {
	return common.PointerString(m)
}
