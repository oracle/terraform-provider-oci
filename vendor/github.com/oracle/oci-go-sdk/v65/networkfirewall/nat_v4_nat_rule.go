// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// NatV4NatRule A Nat Rule is used to define to which traffic NAT should be applied by the firewall.
type NatV4NatRule struct {

	// Name for the NAT rule, must be unique within the policy.
	Name *string `mandatory:"true" json:"name"`

	// OCID of the Network Firewall Policy this decryption profile belongs to.
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`

	Condition *NatRuleMatchCriteria `mandatory:"true" json:"condition"`

	// Description of a NAT rule. This field can be used to add additional info.
	Description *string `mandatory:"false" json:"description"`

	// The priority order in which this rule should be evaluated
	PriorityOrder *int64 `mandatory:"false" json:"priorityOrder"`

	Position *RulePosition `mandatory:"false" json:"position"`

	// action:
	// * DIPP_SRC_NAT - Dynamic-ip-port source NAT.
	Action NatV4ActionTypeEnum `mandatory:"true" json:"action"`
}

// GetName returns Name
func (m NatV4NatRule) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m NatV4NatRule) GetDescription() *string {
	return m.Description
}

// GetPriorityOrder returns PriorityOrder
func (m NatV4NatRule) GetPriorityOrder() *int64 {
	return m.PriorityOrder
}

// GetPosition returns Position
func (m NatV4NatRule) GetPosition() *RulePosition {
	return m.Position
}

// GetParentResourceId returns ParentResourceId
func (m NatV4NatRule) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m NatV4NatRule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NatV4NatRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingNatV4ActionTypeEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetNatV4ActionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m NatV4NatRule) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNatV4NatRule NatV4NatRule
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNatV4NatRule
	}{
		"NATV4",
		(MarshalTypeNatV4NatRule)(m),
	}

	return json.Marshal(&s)
}
