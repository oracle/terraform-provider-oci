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

// ValidateDatabaseToolsConnectionMySqlResult Connection validaton result for the MySQL Server.
type ValidateDatabaseToolsConnectionMySqlResult struct {

	// A short code that defines the result of the validation, meant for programmatic parsing. The value OK indicates that the validation was successful.
	Code *string `mandatory:"true" json:"code"`

	// A human-readable message that describes the result of the validation.
	Message *string `mandatory:"true" json:"message"`

	// The database version.
	DatabaseVersion *string `mandatory:"true" json:"databaseVersion"`

	// A human-readable message that describes possible causes for the validation error.
	Cause *string `mandatory:"false" json:"cause"`

	// A human-readable message that suggests a remedial action to resolve the validation error.
	Action *string `mandatory:"false" json:"action"`

	// The database name.
	DatabaseName *string `mandatory:"false" json:"databaseName"`
}

// GetCode returns Code
func (m ValidateDatabaseToolsConnectionMySqlResult) GetCode() *string {
	return m.Code
}

// GetMessage returns Message
func (m ValidateDatabaseToolsConnectionMySqlResult) GetMessage() *string {
	return m.Message
}

// GetCause returns Cause
func (m ValidateDatabaseToolsConnectionMySqlResult) GetCause() *string {
	return m.Cause
}

// GetAction returns Action
func (m ValidateDatabaseToolsConnectionMySqlResult) GetAction() *string {
	return m.Action
}

func (m ValidateDatabaseToolsConnectionMySqlResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ValidateDatabaseToolsConnectionMySqlResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ValidateDatabaseToolsConnectionMySqlResult) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeValidateDatabaseToolsConnectionMySqlResult ValidateDatabaseToolsConnectionMySqlResult
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeValidateDatabaseToolsConnectionMySqlResult
	}{
		"MYSQL",
		(MarshalTypeValidateDatabaseToolsConnectionMySqlResult)(m),
	}

	return json.Marshal(&s)
}
