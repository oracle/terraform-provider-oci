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

// NatRuleSummary Summary of NAT Rules used in the firewall policy.
// A Nat Rule is used to define which traffic NAT should be applied by the firewall, and how it should do so.
type NatRuleSummary interface {

	// Name for the nat rule, must be unique within the policy.
	GetName() *string

	// The priority order in which this rule should be evaluated.
	GetPriorityOrder() *int64

	// OCID of the Network Firewall Policy this application belongs to.
	GetParentResourceId() *string

	// Description of a NAT rule. This field can be used to add additional info.
	GetDescription() *string
}

type natrulesummary struct {
	JsonData         []byte
	Description      *string `mandatory:"false" json:"description"`
	Name             *string `mandatory:"true" json:"name"`
	PriorityOrder    *int64  `mandatory:"true" json:"priorityOrder"`
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`
	Type             string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *natrulesummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalernatrulesummary natrulesummary
	s := struct {
		Model Unmarshalernatrulesummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.PriorityOrder = s.Model.PriorityOrder
	m.ParentResourceId = s.Model.ParentResourceId
	m.Description = s.Model.Description
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *natrulesummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "NATV4":
		mm := NatV4NatSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for NatRuleSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetDescription returns Description
func (m natrulesummary) GetDescription() *string {
	return m.Description
}

// GetName returns Name
func (m natrulesummary) GetName() *string {
	return m.Name
}

// GetPriorityOrder returns PriorityOrder
func (m natrulesummary) GetPriorityOrder() *int64 {
	return m.PriorityOrder
}

// GetParentResourceId returns ParentResourceId
func (m natrulesummary) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m natrulesummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m natrulesummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
