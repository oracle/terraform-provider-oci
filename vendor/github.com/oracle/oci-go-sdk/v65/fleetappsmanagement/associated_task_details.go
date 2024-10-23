// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AssociatedTaskDetails The details of the task.
type AssociatedTaskDetails interface {
}

type associatedtaskdetails struct {
	JsonData []byte
	Scope    string `json:"scope"`
}

// UnmarshalJSON unmarshals json
func (m *associatedtaskdetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerassociatedtaskdetails associatedtaskdetails
	s := struct {
		Model Unmarshalerassociatedtaskdetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Scope = s.Model.Scope

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *associatedtaskdetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Scope {
	case "LOCAL":
		mm := AssociatedLocalTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "SHARED":
		mm := AssociatedSharedTaskDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for AssociatedTaskDetails: %s.", m.Scope)
		return *m, nil
	}
}

func (m associatedtaskdetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m associatedtaskdetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
