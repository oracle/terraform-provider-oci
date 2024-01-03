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

// ApplicationSummary Summary object for application element in the network firewall policy.
type ApplicationSummary interface {

	// Name of the application.
	GetName() *string

	// OCID of the Network Firewall Policy this application belongs to.
	GetParentResourceId() *string
}

type applicationsummary struct {
	JsonData         []byte
	Name             *string `mandatory:"true" json:"name"`
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`
	Type             string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *applicationsummary) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerapplicationsummary applicationsummary
	s := struct {
		Model Unmarshalerapplicationsummary
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Name = s.Model.Name
	m.ParentResourceId = s.Model.ParentResourceId
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *applicationsummary) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ICMP_V6":
		mm := Icmp6ApplicationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ICMP":
		mm := IcmpApplicationSummary{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ApplicationSummary: %s.", m.Type)
		return *m, nil
	}
}

// GetName returns Name
func (m applicationsummary) GetName() *string {
	return m.Name
}

// GetParentResourceId returns ParentResourceId
func (m applicationsummary) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m applicationsummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m applicationsummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
