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

// ExecuteSqlInputScriptDetails Contains the details for the Script SQL statements to execute on the database connection.
type ExecuteSqlInputScriptDetails struct {

	// The collection of scripts to execute.
	Scripts []ExecuteSqlInputScriptSqlRequestDetails `mandatory:"true" json:"scripts"`

	// Request payload version, returned as-is in the response
	Version *string `mandatory:"false" json:"version"`

	// Client properties returned as-is in the response
	Properties []map[string]interface{} `mandatory:"false" json:"properties"`
}

func (m ExecuteSqlInputScriptDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlInputScriptDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlInputScriptDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlInputScriptDetails ExecuteSqlInputScriptDetails
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlInputScriptDetails
	}{
		"SCRIPT",
		(MarshalTypeExecuteSqlInputScriptDetails)(m),
	}

	return json.Marshal(&s)
}
