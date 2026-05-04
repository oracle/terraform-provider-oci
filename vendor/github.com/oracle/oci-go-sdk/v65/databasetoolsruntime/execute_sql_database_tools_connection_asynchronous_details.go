// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlDatabaseToolsConnectionAsynchronousDetails Asynchronous request.
type ExecuteSqlDatabaseToolsConnectionAsynchronousDetails struct {
	Input ExecuteSqlAsynchronousInputDetails `mandatory:"true" json:"input"`

	Output ExecuteSqlOutputDetails `mandatory:"false" json:"output"`

	// Maximum time in seconds allowed for the request to complete, measured from submission.
	TimeoutInSeconds *int `mandatory:"false" json:"timeoutInSeconds"`
}

// GetOutput returns Output
func (m ExecuteSqlDatabaseToolsConnectionAsynchronousDetails) GetOutput() ExecuteSqlOutputDetails {
	return m.Output
}

func (m ExecuteSqlDatabaseToolsConnectionAsynchronousDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlDatabaseToolsConnectionAsynchronousDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlDatabaseToolsConnectionAsynchronousDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlDatabaseToolsConnectionAsynchronousDetails ExecuteSqlDatabaseToolsConnectionAsynchronousDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlDatabaseToolsConnectionAsynchronousDetails
	}{
		"ASYNCHRONOUS",
		(MarshalTypeExecuteSqlDatabaseToolsConnectionAsynchronousDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExecuteSqlDatabaseToolsConnectionAsynchronousDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Output           executesqloutputdetails            `json:"output"`
		TimeoutInSeconds *int                               `json:"timeoutInSeconds"`
		Input            executesqlasynchronousinputdetails `json:"input"`
	}{}

	e = json.Unmarshal(data, &model)
	if e != nil {
		return
	}
	var nn interface{}
	nn, e = model.Output.UnmarshalPolymorphicJSON(model.Output.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Output = nn.(ExecuteSqlOutputDetails)
	} else {
		m.Output = nil
	}

	m.TimeoutInSeconds = model.TimeoutInSeconds

	nn, e = model.Input.UnmarshalPolymorphicJSON(model.Input.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Input = nn.(ExecuteSqlAsynchronousInputDetails)
	} else {
		m.Input = nil
	}

	return
}
