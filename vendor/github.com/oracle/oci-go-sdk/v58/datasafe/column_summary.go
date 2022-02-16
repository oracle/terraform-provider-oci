// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Safe API
//
// APIs for using Oracle Data Safe.
//

package datasafe

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// ColumnSummary Details of a column in a table fetched from the database.
type ColumnSummary struct {

	// Name of the column.
	ColumnName *string `mandatory:"true" json:"columnName"`

	// Data type of the column.
	DataType *string `mandatory:"true" json:"dataType"`

	// Length of the data represented by the column.
	Length *int64 `mandatory:"true" json:"length"`

	// Name of the table.
	TableName *string `mandatory:"true" json:"tableName"`

	// Name of the schema.
	SchemaName *string `mandatory:"true" json:"schemaName"`

	// Precision of the column.
	Precision *int `mandatory:"false" json:"precision"`

	// Scale of the column.
	Scale *int `mandatory:"false" json:"scale"`

	// Character length.
	CharacterLength *int `mandatory:"false" json:"characterLength"`
}

func (m ColumnSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ColumnSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
