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

// ExecuteSqlResponseItemResultSetMetadata Metadata of the column
type ExecuteSqlResponseItemResultSetMetadata struct {

	// Name of the column in the Oracle Database.
	DatabaseColumnName *string `mandatory:"false" json:"databaseColumnName"`

	// Name of the column in the response Result Set.
	UniqueColumnName *string `mandatory:"false" json:"uniqueColumnName"`

	// Oracle Database data type of the column.
	ColumnTypeName *string `mandatory:"false" json:"columnTypeName"`

	// Precision of the column.
	Precision *int `mandatory:"false" json:"precision"`

	// Scale of the column.
	Scale *int `mandatory:"false" json:"scale"`

	// Specifies if the column is nullable (0 if the column is not nullable)
	IsNullable *bool `mandatory:"false" json:"isNullable"`
}

func (m ExecuteSqlResponseItemResultSetMetadata) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlResponseItemResultSetMetadata) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
