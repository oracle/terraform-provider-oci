// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// UpdateSecurityRuleDetails Update Request for Security Rule used in the firewall policy rules.
// Security Rules determine whether to block or allow a session based on traffic attributes,
// such as  the source and destination IP address, protocol/port, and the HTTP(S) target URL.
type UpdateSecurityRuleDetails struct {
	Condition *SecurityRuleMatchCriteria `mandatory:"true" json:"condition"`

	// Types of Action on the Traffic flow.
	//   * ALLOW - Allows the traffic.
	//   * DROP - Silently drops the traffic, e.g. without sending a TCP reset.
	//   * REJECT - Rejects the traffic, sending a TCP reset to client and/or server as applicable.
	//   * INSPECT - Inspects traffic for vulnerability as specified in `inspection`, which may result in rejection.
	Action TrafficActionTypeEnum `mandatory:"true" json:"action"`

	// Type of inspection to affect the Traffic flow. This is only applicable if action is INSPECT.
	//   * INTRUSION_DETECTION - Intrusion Detection.
	//   * INTRUSION_PREVENTION - Intrusion Detection and Prevention. Traffic classified as potentially malicious will be rejected as described in `type`.
	Inspection TrafficInspectionTypeEnum `mandatory:"false" json:"inspection,omitempty"`

	Position *RulePosition `mandatory:"false" json:"position"`
}

func (m UpdateSecurityRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateSecurityRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTrafficActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetTrafficActionTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingTrafficInspectionTypeEnum(string(m.Inspection)); !ok && m.Inspection != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Inspection: %s. Supported values are: %s.", m.Inspection, strings.Join(GetTrafficInspectionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
