// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Web Application Firewall (WAF) API
//
// API for the Web Application Firewall service.
// Use this API to manage regional Web App Firewalls and corresponding policies for protecting HTTP services.
//

package waf

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ProtectionCapabilitySummary A summary of available OCI-managed protection capabilities in WebAppFirewallPolicy.
// Protection capabilies checks HTTP requests/responses if they are malicious.
type ProtectionCapabilitySummary struct {

	// Unique key of protection capability.
	Key *string `mandatory:"true" json:"key"`

	// The display name of protection capability.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The description of protection capability.
	Description *string `mandatory:"true" json:"description"`

	// The version of protection capability.
	Version *int `mandatory:"true" json:"version"`

	// The field that shows if this is the latest version of protection capability.
	IsLatestVersion *bool `mandatory:"true" json:"isLatestVersion"`

	// The type of protection capability.
	// * **REQUEST_PROTECTION_CAPABILITY** can only be used in `requestProtection` module of WebAppFirewallPolicy.
	// * **RESPONSE_PROTECTION_CAPABILITY** can only be used in `responseProtection` module of WebAppFirewallPolicy.
	Type ProtectionCapabilitySummaryTypeEnum `mandatory:"true" json:"type"`

	// The list of unique names protection capability group tags that are associated with this capability.
	// Example: ["PCI", "Recommended"]
	GroupTags []string `mandatory:"false" json:"groupTags"`

	// The default collaborative action threshold for OCI-managed collaborative protection capability.
	// Collaborative protection capabilities are made of several simple, non-collaborative protection capabilities
	// (referred to as `contributing capabilities` later on) which have weights assigned to them. These weights can
	// be found in the `collaborativeWeights` array.
	// For incoming/outgoing HTTP messages, all contributing capabilities are executed and the sum of all triggered
	// contributing capabilities weights is calculated. Only if this sum is greater than or equal to
	// `collaborativeActionThreshold` is the incoming/outgoing HTTP message marked as malicious.
	// This field is ignored for non-collaborative capabilities.
	CollaborativeActionThreshold *int `mandatory:"false" json:"collaborativeActionThreshold"`

	// The weights of contributing capabilities.
	// Defines how much each contributing capability contributes towards the action threshold of a collaborative protection capability.
	// This field is ignored for non-collaborative capabilities.
	CollaborativeWeights []CollaborativeCapabilityWeight `mandatory:"false" json:"collaborativeWeights"`
}

func (m ProtectionCapabilitySummary) String() string {
	return common.PointerString(m)
}

// ProtectionCapabilitySummaryTypeEnum Enum with underlying type: string
type ProtectionCapabilitySummaryTypeEnum string

// Set of constants representing the allowable values for ProtectionCapabilitySummaryTypeEnum
const (
	ProtectionCapabilitySummaryTypeRequestProtectionCapability  ProtectionCapabilitySummaryTypeEnum = "REQUEST_PROTECTION_CAPABILITY"
	ProtectionCapabilitySummaryTypeResponseProtectionCapability ProtectionCapabilitySummaryTypeEnum = "RESPONSE_PROTECTION_CAPABILITY"
)

var mappingProtectionCapabilitySummaryType = map[string]ProtectionCapabilitySummaryTypeEnum{
	"REQUEST_PROTECTION_CAPABILITY":  ProtectionCapabilitySummaryTypeRequestProtectionCapability,
	"RESPONSE_PROTECTION_CAPABILITY": ProtectionCapabilitySummaryTypeResponseProtectionCapability,
}

// GetProtectionCapabilitySummaryTypeEnumValues Enumerates the set of values for ProtectionCapabilitySummaryTypeEnum
func GetProtectionCapabilitySummaryTypeEnumValues() []ProtectionCapabilitySummaryTypeEnum {
	values := make([]ProtectionCapabilitySummaryTypeEnum, 0)
	for _, v := range mappingProtectionCapabilitySummaryType {
		values = append(values, v)
	}
	return values
}
