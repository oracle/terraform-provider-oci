// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProcessSetSpecificationDetails Details of a single regular expression specification in a Process Set.
type ProcessSetSpecificationDetails struct {

	// Optional label used to identify a single filter.
	Label *string `mandatory:"false" json:"label"`

	// String literal used for exact matching on process name.
	ProcessCommand *string `mandatory:"false" json:"processCommand"`

	// String literal used for exact matching on process user.
	ProcessUser *string `mandatory:"false" json:"processUser"`

	// Regex pattern matching on process arguments.
	ProcessLineRegexPattern *string `mandatory:"false" json:"processLineRegexPattern"`
}

func (m ProcessSetSpecificationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProcessSetSpecificationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
