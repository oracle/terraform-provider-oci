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

// VxlanInspectionRuleSummary Vxlan Tunnel Inspection Rule used on the firewall policy rules.
type VxlanInspectionRuleSummary struct {

	// Name for the Tunnel Inspection Rule, must be unique within the policy.
	Name *string `mandatory:"true" json:"name"`

	// The priority order in which this rule should be evaluated
	PriorityOrder *int64 `mandatory:"true" json:"priorityOrder"`

	// OCID of the Network Firewall Policy this Tunnel Inspection Rule belongs to.
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`

	Profile *VxlanInspectionRuleProfile `mandatory:"true" json:"profile"`

	Condition *VxlanInspectionRuleMatchCriteria `mandatory:"false" json:"condition"`

	// Types of Inspect Action on the Traffic flow.
	//   * INSPECT - Inspect the traffic.
	//   * INSPECT_AND_CAPTURE_LOG - Inspect and capture logs for the traffic.
	Action InspectActionTypeEnum `mandatory:"false" json:"action,omitempty"`
}

// GetName returns Name
func (m VxlanInspectionRuleSummary) GetName() *string {
	return m.Name
}

// GetAction returns Action
func (m VxlanInspectionRuleSummary) GetAction() InspectActionTypeEnum {
	return m.Action
}

// GetPriorityOrder returns PriorityOrder
func (m VxlanInspectionRuleSummary) GetPriorityOrder() *int64 {
	return m.PriorityOrder
}

// GetParentResourceId returns ParentResourceId
func (m VxlanInspectionRuleSummary) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m VxlanInspectionRuleSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m VxlanInspectionRuleSummary) ValidateEnumValue() (bool, error) {
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
func (m VxlanInspectionRuleSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeVxlanInspectionRuleSummary VxlanInspectionRuleSummary
	s := struct {
		DiscriminatorParam string `json:"protocol"`
		MarshalTypeVxlanInspectionRuleSummary
	}{
		"VXLAN",
		(MarshalTypeVxlanInspectionRuleSummary)(m),
	}

	return json.Marshal(&s)
}
