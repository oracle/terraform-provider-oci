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

// ExecuteSqlDatabaseToolsConnectionSynchronousDetails Synchronous request.
type ExecuteSqlDatabaseToolsConnectionSynchronousDetails struct {
	Input ExecuteSqlInputDetails `mandatory:"true" json:"input"`

	Output ExecuteSqlOutputDetails `mandatory:"false" json:"output"`
}

// GetOutput returns Output
func (m ExecuteSqlDatabaseToolsConnectionSynchronousDetails) GetOutput() ExecuteSqlOutputDetails {
	return m.Output
}

func (m ExecuteSqlDatabaseToolsConnectionSynchronousDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlDatabaseToolsConnectionSynchronousDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlDatabaseToolsConnectionSynchronousDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlDatabaseToolsConnectionSynchronousDetails ExecuteSqlDatabaseToolsConnectionSynchronousDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlDatabaseToolsConnectionSynchronousDetails
	}{
		"SYNCHRONOUS",
		(MarshalTypeExecuteSqlDatabaseToolsConnectionSynchronousDetails)(m),
	}

	return json.Marshal(&s)
}

// UnmarshalJSON unmarshals from json
func (m *ExecuteSqlDatabaseToolsConnectionSynchronousDetails) UnmarshalJSON(data []byte) (e error) {
	model := struct {
		Output executesqloutputdetails `json:"output"`
		Input  executesqlinputdetails  `json:"input"`
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

	nn, e = model.Input.UnmarshalPolymorphicJSON(model.Input.JsonData)
	if e != nil {
		return
	}
	if nn != nil {
		m.Input = nn.(ExecuteSqlInputDetails)
	} else {
		m.Input = nil
	}

	return
}
