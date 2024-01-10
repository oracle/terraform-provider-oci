// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResetDatabaseParametersDetails The details required to reset database parameter values.
type ResetDatabaseParametersDetails struct {
	Credentials *DatabaseCredentials `mandatory:"true" json:"credentials"`

	// The clause used to specify when the parameter change takes effect.
	// Use `MEMORY` to make the change in memory and ensure that it takes
	// effect immediately. Use `SPFILE` to make the change in the server
	// parameter file. The change takes effect when the database is next
	// shut down and started up again. Use `BOTH` to make the change in
	// memory and in the server parameter file. The change takes effect
	// immediately and persists after the database is shut down and
	// started up again.
	Scope ParameterScopeEnum `mandatory:"true" json:"scope"`

	// A list of database parameter names.
	Parameters []string `mandatory:"true" json:"parameters"`
}

func (m ResetDatabaseParametersDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResetDatabaseParametersDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingParameterScopeEnum(string(m.Scope)); !ok && m.Scope != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Scope: %s. Supported values are: %s.", m.Scope, strings.Join(GetParameterScopeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
