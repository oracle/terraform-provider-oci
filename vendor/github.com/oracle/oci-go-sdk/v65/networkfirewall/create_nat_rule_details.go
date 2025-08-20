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

// CreateNatRuleDetails Request for creating Nat Rule used in the firewall policy.
// A Nat Rule is used to define to which traffic NAT should be applied by the firewall, and how it should do so.
type CreateNatRuleDetails interface {

	// Name for the NAT rule, must be unique within the policy.
	GetName() *string

	// Description of a NAT rule. This field can be used to add additional info.
	GetDescription() *string

	GetPosition() *RulePosition
}

type createnatruledetails struct {
	JsonData    []byte
	Description *string       `mandatory:"false" json:"description"`
	Position    *RulePosition `mandatory:"false" json:"position"`
	Name        *string       `mandatory:"true" json:"name"`
	Type        string        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *createnatruledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalercreatenatruledetails createnatruledetails
	s := struct {
		Model Unmarshalercreatenatruledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.Description = s.Model.Description
	m.Position = s.Model.Position
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *createnatruledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NATV4":
		mm := CreateNatV4RuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for CreateNatRuleDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m createnatruledetails) GetDescription() *string {
	return m.Description
}

// GetPosition returns Position
func (m createnatruledetails) GetPosition() *RulePosition {
	return m.Position
}

// GetName returns Name
func (m createnatruledetails) GetName() *string {
	return m.Name
}

func (m createnatruledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m createnatruledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
