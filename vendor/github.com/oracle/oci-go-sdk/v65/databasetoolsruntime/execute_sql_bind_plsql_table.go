// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Tools Runtime API
//
// Use the Database Tools Runtime API to connect to databases through Database Tools Connections.
//

package databasetoolsruntime

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExecuteSqlBindPlsqlTable PL/SQL table type details
type ExecuteSqlBindPlsqlTable struct {

	// Only an empty string is accepted as the value.
	TypeName *string `mandatory:"false" json:"typeName"`

	// Only an empty string is accepted as the value.
	TypeSubname *string `mandatory:"false" json:"typeSubname"`

	// Specifies an array of data types in the PL/SQL table.
	TypeComponents []ExecuteSqlBindTypeComponent `mandatory:"false" json:"typeComponents"`
}

func (m ExecuteSqlBindPlsqlTable) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlBindPlsqlTable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
