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

// UpdateNatRuleDetails Request for updating NAT Rule used in the firewall policy.
// A Nat Rule is used to define to which traffic NAT should be applied by the firewall, and how it should do so.
type UpdateNatRuleDetails interface {

	// Description of a NAT rule. This field can be used to add additional info.
	GetDescription() *string

	GetPosition() *RulePosition
}

type updatenatruledetails struct {
	JsonData    []byte
	Description *string       `mandatory:"false" json:"description"`
	Position    *RulePosition `mandatory:"false" json:"position"`
	Type        string        `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *updatenatruledetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerupdatenatruledetails updatenatruledetails
	s := struct {
		Model Unmarshalerupdatenatruledetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Description = s.Model.Description
	m.Position = s.Model.Position
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *updatenatruledetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NATV4":
		mm := UpdateNatV4RuleDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for UpdateNatRuleDetails: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m updatenatruledetails) GetDescription() *string {
	return m.Description
}

// GetPosition returns Position
func (m updatenatruledetails) GetPosition() *RulePosition {
	return m.Position
}

func (m updatenatruledetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m updatenatruledetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
