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

// NatV4NatSummary Request for updating NATV4 type Nat Rule used in the firewall policy.
type NatV4NatSummary struct {

	// Name for the nat rule, must be unique within the policy.
	Name *string `mandatory:"true" json:"name"`

	// The priority order in which this rule should be evaluated.
	PriorityOrder *int64 `mandatory:"true" json:"priorityOrder"`

	// OCID of the Network Firewall Policy this application belongs to.
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`

	Condition *NatRuleMatchCriteria `mandatory:"true" json:"condition"`

	// Description of a NAT rule. This field can be used to add additional info.
	Description *string `mandatory:"false" json:"description"`

	// action:
	// * DIPP_SRC_NAT - Dynamic-ip-port source NAT.
	Action NatV4ActionTypeEnum `mandatory:"true" json:"action"`
}

// GetName returns Name
func (m NatV4NatSummary) GetName() *string {
	return m.Name
}

// GetDescription returns Description
func (m NatV4NatSummary) GetDescription() *string {
	return m.Description
}

// GetPriorityOrder returns PriorityOrder
func (m NatV4NatSummary) GetPriorityOrder() *int64 {
	return m.PriorityOrder
}

// GetParentResourceId returns ParentResourceId
func (m NatV4NatSummary) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m NatV4NatSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NatV4NatSummary) ValidateEnumValue() (bool, error) {
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
func (m NatV4NatSummary) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeNatV4NatSummary NatV4NatSummary
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeNatV4NatSummary
	}{
		"NATV4",
		(MarshalTypeNatV4NatSummary)(m),
	}

	return json.Marshal(&s)
}
