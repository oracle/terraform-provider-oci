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

// ExecuteSqlResponseItemProperties Property information
type ExecuteSqlResponseItemProperties struct {

	// Client properties returned as-is in the response
	Query *interface{} `mandatory:"false" json:"query"`

	// Client properties returned as-is in the response
	Statement *interface{} `mandatory:"false" json:"statement"`

	// Client properties returned as-is in the response.
	// For asynchronous requests, the "displayName" property can be set to provide a user-friendly name that will propagate to the resulting Work Request to allow for easier identification and tracking of the operation's progress and outcome.
	DisplayName *interface{} `mandatory:"false" json:"displayName"`
}

func (m ExecuteSqlResponseItemProperties) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExecuteSqlResponseItemProperties) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
