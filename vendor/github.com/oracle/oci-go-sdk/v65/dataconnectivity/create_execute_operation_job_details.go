// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Connectivity Management API
//
// Use the DCMS APIs to perform Metadata/Data operations.
//

package dataconnectivity

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreateExecuteOperationJobDetails Input details to execute operation.
type CreateExecuteOperationJobDetails struct {
	Operation Operation `mandatory:"false" json:"operation"`

	CallOperationConfig *CallOperationConfig `mandatory:"false" json:"callOperationConfig"`

	// Collection of input parameters supplied.
	InputRecords []OperationInputRecord `mandatory:"false" json:"inputRecords"`
}

func (m CreateExecuteOperationJobDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateExecuteOperationJobDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnmarshalJSON unmarshals from json
func (m *CreateExecuteOperationJobDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Operation           operation              `json:"operation"`
		CallOperationConfig *CallOperationConfig   `json:"callOperationConfig"`
		InputRecords        []OperationInputRecord `json:"inputRecords"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Operation.UnmarshalPolymorphicJSON(model.Operation.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Operation = nn.(Operation)
	} else {
		m.Operation = nil
	}

	m.CallOperationConfig = model.CallOperationConfig

	m.InputRecords = make([]OperationInputRecord, len(model.InputRecords))
	for i, n := range model.InputRecords {
		m.InputRecords[i] = n
	}

	return
}
