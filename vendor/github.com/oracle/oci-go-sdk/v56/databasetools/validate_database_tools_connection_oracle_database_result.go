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

// ValidateDatabaseToolsConnectionOracleDatabaseResult Connection validaton result for the Oracle Database.
type ValidateDatabaseToolsConnectionOracleDatabaseResult struct {

	// A short code that defines the result of the validation, meant for programmatic parsing.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable message that describes the result of the validation.
	Message *string `mandatory:"true" json:"message"`

	// A human-readable message that describes possible causes for the validation error.
	Cause *string `mandatory:"false" json:"cause"`

	// A human-readable message that suggests a remedial action to resolve the validation error.
	Action *string `mandatory:"false" json:"action"`

	// The database name.
	DatabaseName *string `mandatory:"false" json:"databaseName"`

	// The database version.
	DatabaseVersion *string `mandatory:"false" json:"databaseVersion"`
}

//GetCode returns Code
func (m ValidateDatabaseToolsConnectionOracleDatabaseResult) GetCode() *string {
	return m.Code
}

//GetMessage returns Message
func (m ValidateDatabaseToolsConnectionOracleDatabaseResult) GetMessage() *string {
	return m.Message
}

//GetCause returns Cause
func (m ValidateDatabaseToolsConnectionOracleDatabaseResult) GetCause() *string {
	return m.Cause
}

//GetAction returns Action
func (m ValidateDatabaseToolsConnectionOracleDatabaseResult) GetAction() *string {
	return m.Action
}

func (m ValidateDatabaseToolsConnectionOracleDatabaseResult) String() string {
	return common.PointerString(m)
}

// MarshalJSON marshals to json representation
func (m ValidateDatabaseToolsConnectionOracleDatabaseResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValidateDatabaseToolsConnectionOracleDatabaseResult ValidateDatabaseToolsConnectionOracleDatabaseResult
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeValidateDatabaseToolsConnectionOracleDatabaseResult
	}{
		"ORACLE_DATABASE",
		(MarshalTypeValidateDatabaseToolsConnectionOracleDatabaseResult)(m),
	}

	return json.Marshal(&s)
}
