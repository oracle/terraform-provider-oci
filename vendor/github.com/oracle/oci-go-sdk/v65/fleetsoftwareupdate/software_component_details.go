// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Exadata Fleet Update service API
//
// Use the Exadata Fleet Update service to patch large collections of components directly,
// as a single entity, orchestrating the maintenance actions to update all chosen components in the stack in a single cycle.
//

package fleetsoftwareupdate

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SoftwareComponentDetails Details of a component in an Exadata software stack.
type SoftwareComponentDetails interface {
}

type softwarecomponentdetails struct {
	JsonData      []byte
	ComponentType string `json:"componentType"`
}

// UnmarshalJSON unmarshals json
func (m *softwarecomponentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalersoftwarecomponentdetails softwarecomponentdetails
	s := struct {
		Model Unmarshalersoftwarecomponentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComponentType = s.Model.ComponentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *softwarecomponentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ComponentType {
	case "GI":
		mm := GiSoftwareComponentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GUEST_OS":
		mm := GuestOsSoftwareComponentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for SoftwareComponentDetails: %s.", m.ComponentType)
		return *m, nil
	}
}

func (m softwarecomponentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m softwarecomponentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
