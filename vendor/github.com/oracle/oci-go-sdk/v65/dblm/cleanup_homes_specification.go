// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Dblm API
//
// A description of the Dblm API
//

package dblm

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CleanupHomesSpecification Specification for cleaning up the old Oracle homes or newly deployed Oracle home. As a prerequisite, the homes to be cleaned up, must have been created through a deploy home(s) specification.
type CleanupHomesSpecification struct {
	Resources *Resources `mandatory:"true" json:"resources"`

	Schedule *Schedule `mandatory:"false" json:"schedule"`

	// If true, run the CVU validations.
	IsSkipCvuChecks *bool `mandatory:"false" json:"isSkipCvuChecks"`

	// Location of the ORACLE_HOME to be cleaned up on the subscriber database host. If value is passed then cleanup the old Oracle home and if empty then cleanup the latest successful or failed deployed Oracle home.
	OracleHomePath *string `mandatory:"false" json:"oracleHomePath"`
}

func (m CleanupHomesSpecification) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CleanupHomesSpecification) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
