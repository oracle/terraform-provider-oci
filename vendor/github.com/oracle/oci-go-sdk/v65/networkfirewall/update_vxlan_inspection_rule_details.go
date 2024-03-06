// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Network Firewall API
//
// Use the Network Firewall API to create network firewalls and configure policies that regulates network traffic in and across VCNs.
//

package networkfirewall

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateVxlanInspectionRuleDetails Update Request for creating Vxlan Tunnel Inspection Rule used in the firewall policy rules.
type UpdateVxlanInspectionRuleDetails struct {
	Condition *VxlanInspectionRuleMatchCriteria `mandatory:"true" json:"condition"`

	Position *RulePosition `mandatory:"false" json:"position"`

	Profile *VxlanInspectionRuleProfile `mandatory:"false" json:"profile"`

	// Types of Inspect Action on the Traffic flow.
	//   * INSPECT - Inspect the traffic.
	//   * INSPECT_AND_CAPTURE_LOG - Inspect and capture logs for the traffic.
	Action InspectActionTypeEnum `mandatory:"false" json:"action,omitempty"`
}

// GetAction returns Action
func (m UpdateVxlanInspectionRuleDetails) GetAction() InspectActionTypeEnum {
	return m.Action
}

// GetPosition returns Position
func (m UpdateVxlanInspectionRuleDetails) GetPosition() *RulePosition {
	return m.Position
}

func (m UpdateVxlanInspectionRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateVxlanInspectionRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInspectActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetInspectActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m UpdateVxlanInspectionRuleDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeUpdateVxlanInspectionRuleDetails UpdateVxlanInspectionRuleDetails
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeUpdateVxlanInspectionRuleDetails
	}{
		"VXLAN",
		(MarshalTypeUpdateVxlanInspectionRuleDetails)(m),
	}

	return json.Marshal(&s)
}
