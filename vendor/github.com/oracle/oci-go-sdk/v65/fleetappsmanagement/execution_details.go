// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management Service API. Use this API to for all FAMS related activities.
// To manage fleets,view complaince report for the Fleet,scedule patches and other lifecycle activities
//

package fleetappsmanagement

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecutionDetails Content Source Details
type ExecutionDetails interface {
}

type executiondetails struct {
	JsonData      []byte
	ExecutionType string `json:"executionType"`
}

// UnmarshalJSON unmarshals json
func (m *executiondetails) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalerexecutiondetails executiondetails
	s := struct {
		Model Unmarshalerexecutiondetails
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.ExecutionType = s.Model.ExecutionType

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *executiondetails) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.ExecutionType {
	case "SCRIPT":
		mm := ScriptBasedExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "API":
		mm := ApiBasedExecutionDetails{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ExecutionDetails: %s.", m.ExecutionType)
		return *m, nil
	}
}

func (m executiondetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m executiondetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
