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

// TunnelInspectionRuleSummary Summary for the Tunnel Inspection Rule used in the firewall policy rules.
// Tunnel Inspection Rule determines whether tunnel inspection is applied on the traffic based on attributes
// such as Tunnel Inspect protocol, the source and destination IP address.
type TunnelInspectionRuleSummary interface {

	// Name for the Tunnel Inspection Rule, must be unique within the policy.
	GetName() *string

	// The priority order in which this rule should be evaluated
	GetPriorityOrder() *int64

	// OCID of the Network Firewall Policy this Tunnel Inspection Rule belongs to.
	GetParentResourceId() *string

	// Types of Inspect Action on the Traffic flow.
	//   * INSPECT - Inspect the traffic.
	//   * INSPECT_AND_CAPTURE_LOG - Inspect and capture logs for the traffic.
	GetAction() InspectActionTypeEnum
}

type tunnelinspectionrulesummary struct {
	JsonData         []byte
	Action           InspectActionTypeEnum `mandatory:"false" json:"action,omitempty"`
	Name             *string               `mandatory:"true" json:"name"`
	PriorityOrder    *int64                `mandatory:"true" json:"priorityOrder"`
	ParentResourceId *string               `mandatory:"true" json:"parentResourceId"`
	Protocol         string                `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *tunnelinspectionrulesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalertunnelinspectionrulesummary tunnelinspectionrulesummary
	s := struct {
		Model Unmarshalertunnelinspectionrulesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.PriorityOrder = s.Model.PriorityOrder
	m.ParentResourceId = s.Model.ParentResourceId
	m.Action = s.Model.Action
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *tunnelinspectionrulesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "VXLAN":
		mm := VxlanInspectionRuleSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for TunnelInspectionRuleSummary: %s.", m.Protocol)
		return *m, nil
	}
}

// GetAction returns Action
func (m tunnelinspectionrulesummary) GetAction() InspectActionTypeEnum {
	return m.Action
}

// GetName returns Name
func (m tunnelinspectionrulesummary) GetName() *string {
	return m.Name
}

// GetPriorityOrder returns PriorityOrder
func (m tunnelinspectionrulesummary) GetPriorityOrder() *int64 {
	return m.PriorityOrder
}

// GetParentResourceId returns ParentResourceId
func (m tunnelinspectionrulesummary) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m tunnelinspectionrulesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m tunnelinspectionrulesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInspectActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetInspectActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
