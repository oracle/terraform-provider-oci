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

// ExecuteSqlInputStandardDetails Contains the details for the Standard SQL statements to execute on the database connection.
type ExecuteSqlInputStandardDetails struct {

	// Statements to execute (Can be more than one).
	StatementText *string `mandatory:"true" json:"statementText"`

	// The maximum number of rows to return from the query (-1 disables pagination).
	Limit *int `mandatory:"false" json:"limit"`

	// The first row to return in the result set.
	Offset *int `mandatory:"false" json:"offset"`

	// Array of objects specifying the bind information.
	Binds []ExecuteSqlBind `mandatory:"false" json:"binds"`

	ResponseFormat *ExecuteSqlResponseFormat `mandatory:"false" json:"responseFormat"`

	// Client properties returned as-is in the response
	Properties *interface{} `mandatory:"false" json:"properties"`
}

func (m ExecuteSqlInputStandardDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlInputStandardDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlInputStandardDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlInputStandardDetails ExecuteSqlInputStandardDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlInputStandardDetails
	}{
		"STANDARD",
		(MarshalTypeExecuteSqlInputStandardDetails)(m),
	}

	return json.Marshal(&s)
}
