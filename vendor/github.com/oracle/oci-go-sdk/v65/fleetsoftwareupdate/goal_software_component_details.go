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

// GoalSoftwareComponentDetails Details of goal version for a component in an 'EXADB_STACK' type Exadata Fleet Update Collection.
type GoalSoftwareComponentDetails interface {
}

type goalsoftwarecomponentdetails struct {
	JsonData      []byte
	ComponentType string `json:"componentType"`
}

// UnmarshalJSON unmarshals json
func (m *goalsoftwarecomponentdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalergoalsoftwarecomponentdetails goalsoftwarecomponentdetails
	s := struct {
		Model Unmarshalergoalsoftwarecomponentdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ComponentType = s.Model.ComponentType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *goalsoftwarecomponentdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ComponentType {
	case "GUEST_OS":
		mm := GuestOsGoalSoftwareComponentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "GI":
		mm := GiGoalSoftwareComponentDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Received unsupported enum value for GoalSoftwareComponentDetails: %s.", m.ComponentType)
		return *m, nil
	}
}

func (m goalsoftwarecomponentdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m goalsoftwarecomponentdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
