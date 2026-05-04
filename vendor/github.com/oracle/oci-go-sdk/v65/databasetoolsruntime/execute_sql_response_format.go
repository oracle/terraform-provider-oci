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

// ExecuteSqlResponseFormat Specifies which major elements of the response are returned
type ExecuteSqlResponseFormat struct {

	// Set to false to exclude result set metadata from response
	ResultSetMetaData *bool `mandatory:"false" json:"resultSetMetaData"`

	// Set to false to exclude statement information from response
	StatementInformation *bool `mandatory:"false" json:"statementInformation"`

	// Set to false to exclude statement text from response
	StatementText *bool `mandatory:"false" json:"statementText"`

	// Set to false to exclude binds from response
	Binds *bool `mandatory:"false" json:"binds"`

	// Set to false to exclude result from response
	Result *bool `mandatory:"false" json:"result"`

	// Set to false to exclude response from response
	Response *bool `mandatory:"false" json:"response"`
}

func (m ExecuteSqlResponseFormat) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlResponseFormat) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
