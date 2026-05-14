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

// ExecuteSqlResponseItemResultSet Result set generated from the query.
type ExecuteSqlResponseItemResultSet struct {

	// Metadata of the column.
	Metadata []ExecuteSqlResponseItemResultSetMetadata `mandatory:"false" json:"metadata"`

	// All rows in the result set.
	Items []map[string]interface{} `mandatory:"false" json:"items"`

	// Specifies whether the result set has more rows.
	HasMore *bool `mandatory:"false" json:"hasMore"`

	// The number of rows returned.
	Count *int `mandatory:"false" json:"count"`

	// The first row returned in the result set.
	Offset *int `mandatory:"false" json:"offset"`

	// Maximum number of rows returned from the query.
	Limit *int `mandatory:"false" json:"limit"`
}

func (m ExecuteSqlResponseItemResultSet) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlResponseItemResultSet) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
