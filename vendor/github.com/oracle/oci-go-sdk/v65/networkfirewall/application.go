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

// Application A protocol identifier (such as TCP, UDP, or ICMP)
// and protocol-specific parameters (such as a port range).
type Application interface {

	// Name of the application.
	GetName() *string

	// OCID of the Network Firewall Policy this application belongs to.
	GetParentResourceId() *string
}

type application struct {
	JsonData         []byte
	Name             *string `mandatory:"true" json:"name"`
	ParentResourceId *string `mandatory:"true" json:"parentResourceId"`
	Type             string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *application) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerapplication application
	s := struct {
		Model Unmarshalerapplication
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
func (m *application) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ICMP":
		mm := IcmpApplication{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "ICMP_V6":
		mm := Icmp6Application{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for Application: %s.", m.Type)
		return *m, nil
	}
}

// GetName returns Name
func (m application) GetName() *string {
	return m.Name
}

// GetParentResourceId returns ParentResourceId
func (m application) GetParentResourceId() *string {
	return m.ParentResourceId
}

func (m application) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m application) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
