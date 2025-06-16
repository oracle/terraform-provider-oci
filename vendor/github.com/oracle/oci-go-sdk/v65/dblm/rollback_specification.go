// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// RollbackSpecification Specification for rollback (moving) resources to their previous Oracle homes in case of patching failures or performance issues in the patched home.
type RollbackSpecification struct {

	// If true, run the CVU validations.
	IsSkipCvuChecks *bool `mandatory:"false" json:"isSkipCvuChecks"`

	Schedule *Schedule `mandatory:"false" json:"schedule"`

	Resources *Resources `mandatory:"false" json:"resources"`

	// If true, the rollback will happen in a rolling mode.
	IsRollingMode *bool `mandatory:"false" json:"isRollingMode"`

	// Wait time for connections to drain out
	DrainTimeout *int `mandatory:"false" json:"drainTimeout"`
}

func (m RollbackSpecification) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RollbackSpecification) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
