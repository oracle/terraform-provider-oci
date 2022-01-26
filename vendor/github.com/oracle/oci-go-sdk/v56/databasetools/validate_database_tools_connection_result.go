// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Database Tools APIs to manage Connections and Private Endpoints.
//

package databasetools

import (
	"encoding/json"
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ValidateDatabaseToolsConnectionResult Connection validation result.
type ValidateDatabaseToolsConnectionResult interface {

	// A short code that defines the result of the validation, meant for programmatic parsing.
	GetCode() *string

	// A human-readable message that describes the result of the validation.
	GetMessage() *string

	// A human-readable message that describes possible causes for the validation error.
	GetCause() *string

	// A human-readable message that suggests a remedial action to resolve the validation error.
	GetAction() *string
}

type validatedatabasetoolsconnectionresult struct {
	JsonData []byte
	Code     *string `mandatory:"true" json:"code"`
	Message  *string `mandatory:"true" json:"message"`
	Cause    *string `mandatory:"false" json:"cause"`
	Action   *string `mandatory:"false" json:"action"`
	Type     string  `json:"type"`
}

// UnmarshalJSON unmarshals json
func (m *validatedatabasetoolsconnectionresult) UnmarshalJSON(data []byte) error {
	m.JsonData = data
	type Unmarshalervalidatedatabasetoolsconnectionresult validatedatabasetoolsconnectionresult
	s := struct {
		Model Unmarshalervalidatedatabasetoolsconnectionresult
	}{}
	err := json.Unmarshal(data, &s.Model)
	if err != nil {
		return err
	}
	m.Code = s.Model.Code
	m.Message = s.Model.Message
	m.Cause = s.Model.Cause
	m.Action = s.Model.Action
	m.Type = s.Model.Type

	return err
}

// UnmarshalPolymorphicJSON unmarshals polymorphic json
func (m *validatedatabasetoolsconnectionresult) UnmarshalPolymorphicJSON(data []byte) (interface{}, error) {

	if data == nil || string(data) == "null" {
		return nil, nil
	}

	var err error
	switch m.Type {
	case "ORACLE_DATABASE":
		mm := ValidateDatabaseToolsConnectionOracleDatabaseResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		return *m, nil
	}
}

//GetCode returns Code
func (m validatedatabasetoolsconnectionresult) GetCode() *string {
	return m.Code
}

//GetMessage returns Message
func (m validatedatabasetoolsconnectionresult) GetMessage() *string {
	return m.Message
}

//GetCause returns Cause
func (m validatedatabasetoolsconnectionresult) GetCause() *string {
	return m.Cause
}

//GetAction returns Action
func (m validatedatabasetoolsconnectionresult) GetAction() *string {
	return m.Action
}

func (m validatedatabasetoolsconnectionresult) String() string {
	return common.PointerString(m)
}
