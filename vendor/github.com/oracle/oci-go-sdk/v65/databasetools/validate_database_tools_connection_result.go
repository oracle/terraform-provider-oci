// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools
//
// Use the Database Tools API to manage connections, private endpoints, and work requests in the Database Tools service.
//

package databasetools

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ValidateDatabaseToolsConnectionResult Connection validation result.
type ValidateDatabaseToolsConnectionResult interface {

	// A short code that defines the result of the validation, meant for programmatic parsing. The value OK indicates that the validation was successful.
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
	Cause    *string `mandatory:"false" json:"cause"`
	Action   *string `mandatory:"false" json:"action"`
	Code     *string `mandatory:"true" json:"code"`
	Message  *string `mandatory:"true" json:"message"`
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
	case "MYSQL":
		mm := ValidateDatabaseToolsConnectionMySqlResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	case "POSTGRESQL":
		mm := ValidateDatabaseToolsConnectionPostgresqlResult{}
		err = json.Unmarshal(data, &mm)
		return mm, err
	default:
		common.Logf("Recieved unsupported enum value for ValidateDatabaseToolsConnectionResult: %s.", m.Type)
		return *m, nil
	}
}

// GetCause returns Cause
func (m validatedatabasetoolsconnectionresult) GetCause() *string {
	return m.Cause
}

// GetAction returns Action
func (m validatedatabasetoolsconnectionresult) GetAction() *string {
	return m.Action
}

// GetCode returns Code
func (m validatedatabasetoolsconnectionresult) GetCode() *string {
	return m.Code
}

// GetMessage returns Message
func (m validatedatabasetoolsconnectionresult) GetMessage() *string {
	return m.Message
}

func (m validatedatabasetoolsconnectionresult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m validatedatabasetoolsconnectionresult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
