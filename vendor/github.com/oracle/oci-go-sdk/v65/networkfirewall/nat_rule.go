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

// NatRule A Nat Rule is used to define to which traffic NAT should be applied by the firewall.
type NatRule interface {

	// Name for the NAT rule, must be unique within the policy.
	GetName() *string

	// OCID of the Network Firewall Policy this decryption profile belongs to.
	GetParentResourceId() *string

	// Description of a NAT rule. This field can be used to add additional info.
	GetDescription() *string

	// The priority order in which this rule should be evaluated
	GetPriorityOrder() *int64

	GetPosition() *RulePosition
}

type natrule struct {
	JsonData         []byte
	Description      *string       `mandatory:"false" json:"description"`
	PriorityOrder    *int64        `mandatory:"false" json:"priorityOrder"`
	Position         *RulePosition `mandatory:"false" json:"position"`
	Name             *string       `mandatory:"true" json:"name"`
	ParentResourceId *string       `mandatory:"true" json:"parentResourceId"`
	Type             string        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *natrule) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernatrule natrule
	s := struct {
		Model Unmarshalernatrule
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.ParentResourceId = s.Model.ParentResourceId
	m.Description = s.Model.Description
	m.PriorityOrder = s.Model.PriorityOrder
	m.Position = s.Model.Position
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *natrule) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NATV4":
		mm := NatV4NatRule{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for NatRule: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m natrule) GetDescription() *string {
	return m.Description
}

// GetPriorityOrder returns PriorityOrder
func (m natrule) GetPriorityOrder() *int64 {
	return m.PriorityOrder
}

// GetPosition returns Position
func (m natrule) GetPosition() *RulePosition {
	return m.Position
}

// GetName returns Name
func (m natrule) GetName() *string {
	return m.Name
}

// GetParentResourceId returns ParentResourceId
func (m natrule) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m natrule) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m natrule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
