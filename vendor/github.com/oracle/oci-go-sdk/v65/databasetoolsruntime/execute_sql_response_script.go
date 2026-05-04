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

// ExecuteSqlResponseScript Contains the details for the Script SQL statements to execute on the database connection.
type ExecuteSqlResponseScript struct {

	// The execution result of a statement.
	Items []ExecuteSqlResponseItemScript `mandatory:"true" json:"items"`

	Env *ExecuteSqlResponseEnv `mandatory:"false" json:"env"`

	// Script version
	Version *string `mandatory:"false" json:"version"`
}

// GetEnv returns Env
func (m ExecuteSqlResponseScript) GetEnv() *ExecuteSqlResponseEnv {
	return m.Env
}

// GetVersion returns Version
func (m ExecuteSqlResponseScript) GetVersion() *string {
	return m.Version
}

func (m ExecuteSqlResponseScript) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlResponseScript) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m ExecuteSqlResponseScript) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeExecuteSqlResponseScript ExecuteSqlResponseScript
	s := struct {
		DiscriminatorParam string `json:"type"`
		MarshalTypeExecuteSqlResponseScript
	}{
		"SCRIPT",
		(MarshalTypeExecuteSqlResponseScript)(m),
	}

	return json.Marshal(&s)
}
