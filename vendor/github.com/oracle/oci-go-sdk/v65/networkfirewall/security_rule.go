// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SecurityRule Security Rule used in the firewall policy rules.
// Security Rules determine whether to block or allow a session based on traffic attributes,
// such as  the source and destination IP address, protocol/port, and the HTTP(S) target URL.
type SecurityRule struct {

	// Name for the Security rule, must be unique within the policy.
	Name *string `mandatory:"true" json:"name"`

	Condition *SecurityRuleMatchCriteria `mandatory:"true" json:"condition"`

	// Types of Action on the Traffic flow.
	//   * ALLOW - Allows the traffic.
	//   * DROP - Silently drops the traffic, e.g. without sending a TCP reset.
	//   * REJECT - Rejects the traffic, sending a TCP reset to client and/or server as applicable.
	//   * INSPECT - Inspects traffic for vulnerability as specified in `inspection`, which may result in rejection.
	Action SecurityRuleActionEnum `mandatory:"true" json:"action"`

	// Type of inspection to affect the Traffic flow. This is only applicable if action is INSPECT.
	//   * INTRUSION_DETECTION - Intrusion Detection.
	//   * INTRUSION_PREVENTION - Intrusion Detection and Prevention. Traffic classified as potentially malicious will be rejected as described in `type`.
	Inspection SecurityRuleInspectionEnum `mandatory:"false" json:"inspection,omitempty"`
}

func (m SecurityRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SecurityRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSecurityRuleActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetSecurityRuleActionEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSecurityRuleInspectionEnum(string(m.Inspection)); !ok && m.Inspection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Inspection: %s. Supported values are: %s.", m.Inspection, strings.Join(GetSecurityRuleInspectionEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// SecurityRuleActionEnum Enum with underlying type: string
type SecurityRuleActionEnum string

// Set of constants representing the allowable values for SecurityRuleActionEnum
const (
	SecurityRuleActionAllow   SecurityRuleActionEnum = "ALLOW"
	SecurityRuleActionDrop    SecurityRuleActionEnum = "DROP"
	SecurityRuleActionReject  SecurityRuleActionEnum = "REJECT"
	SecurityRuleActionInspect SecurityRuleActionEnum = "INSPECT"
)

var mappingSecurityRuleActionEnum = map[string]SecurityRuleActionEnum{
	"ALLOW":   SecurityRuleActionAllow,
	"DROP":    SecurityRuleActionDrop,
	"REJECT":  SecurityRuleActionReject,
	"INSPECT": SecurityRuleActionInspect,
}

var mappingSecurityRuleActionEnumLowerCase = map[string]SecurityRuleActionEnum{
	"allow":   SecurityRuleActionAllow,
	"drop":    SecurityRuleActionDrop,
	"reject":  SecurityRuleActionReject,
	"inspect": SecurityRuleActionInspect,
}

// GetSecurityRuleActionEnumValues Enumerates the set of values for SecurityRuleActionEnum
func GetSecurityRuleActionEnumValues() []SecurityRuleActionEnum {
	values := make([]SecurityRuleActionEnum, 0)
	for _, v := range mappingSecurityRuleActionEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityRuleActionEnumStringValues Enumerates the set of values in String for SecurityRuleActionEnum
func GetSecurityRuleActionEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DROP",
		"REJECT",
		"INSPECT",
	}
}

// GetMappingSecurityRuleActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityRuleActionEnum(val string) (SecurityRuleActionEnum, bool) {
	enum, ok := mappingSecurityRuleActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// SecurityRuleInspectionEnum Enum with underlying type: string
type SecurityRuleInspectionEnum string

// Set of constants representing the allowable values for SecurityRuleInspectionEnum
const (
	SecurityRuleInspectionDetection  SecurityRuleInspectionEnum = "INTRUSION_DETECTION"
	SecurityRuleInspectionPrevention SecurityRuleInspectionEnum = "INTRUSION_PREVENTION"
)

var mappingSecurityRuleInspectionEnum = map[string]SecurityRuleInspectionEnum{
	"INTRUSION_DETECTION":  SecurityRuleInspectionDetection,
	"INTRUSION_PREVENTION": SecurityRuleInspectionPrevention,
}

var mappingSecurityRuleInspectionEnumLowerCase = map[string]SecurityRuleInspectionEnum{
	"intrusion_detection":  SecurityRuleInspectionDetection,
	"intrusion_prevention": SecurityRuleInspectionPrevention,
}

// GetSecurityRuleInspectionEnumValues Enumerates the set of values for SecurityRuleInspectionEnum
func GetSecurityRuleInspectionEnumValues() []SecurityRuleInspectionEnum {
	values := make([]SecurityRuleInspectionEnum, 0)
	for _, v := range mappingSecurityRuleInspectionEnum {
		values = append(values, v)
	}
	return values
}

// GetSecurityRuleInspectionEnumStringValues Enumerates the set of values in String for SecurityRuleInspectionEnum
func GetSecurityRuleInspectionEnumStringValues() []string {
	return []string{
		"INTRUSION_DETECTION",
		"INTRUSION_PREVENTION",
	}
}

// GetMappingSecurityRuleInspectionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingSecurityRuleInspectionEnum(val string) (SecurityRuleInspectionEnum, bool) {
	enum, ok := mappingSecurityRuleInspectionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
