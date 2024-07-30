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

// CreateTunnelInspectionRuleDetails Request for creating Tunnel Inspection Rule used in the firewall policy rules.
// Tunnel Inspection Rule determines whether tunnel inspection is applied on the traffic based on attributes
// such as Tunnel Inspection protocol, the source and destination IP address.
type CreateTunnelInspectionRuleDetails interface {

	// Name for the Tunnel Inspection Rule, must be unique within the policy.
	GetName() *string

	// Types of Inspect Action on the traffic flow.
	//   * INSPECT - Inspect the traffic.
	//   * INSPECT_AND_CAPTURE_LOG - Inspect and capture logs for the traffic.
	GetAction() InspectActionTypeEnum

	GetPosition() *RulePosition
}

type createtunnelinspectionruledetails struct {
	JsonData []byte
	Action   InspectActionTypeEnum `mandatory:"false" json:"action,omitempty"`
	Position *RulePosition         `mandatory:"false" json:"position"`
	Name     *string               `mandatory:"true" json:"name"`
	Protocol string                `json:"protocol"`
}

// UnmarshalJSON unmarshals json
func (m *createtunnelinspectionruledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatetunnelinspectionruledetails createtunnelinspectionruledetails
	s := struct {
		Model Unmarshalercreatetunnelinspectionruledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Action = s.Model.Action
	m.Position = s.Model.Position
	m.Protocol = s.Model.Protocol

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createtunnelinspectionruledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Protocol {
	case "VXLAN":
		mm := CreateVxlanInspectionRuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for CreateTunnelInspectionRuleDetails: %s.", m.Protocol)
		return *m, nil
	}
}

// GetAction returns Action
func (m createtunnelinspectionruledetails) GetAction() InspectActionTypeEnum {
	return m.Action
}

// GetPosition returns Position
func (m createtunnelinspectionruledetails) GetPosition() *RulePosition {
	return m.Position
}

// GetName returns Name
func (m createtunnelinspectionruledetails) GetName() *string {
	return m.Name
}

func (m createtunnelinspectionruledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createtunnelinspectionruledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInspectActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetInspectActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
