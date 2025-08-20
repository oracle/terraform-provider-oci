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

// RollbackListenerSpecification Specification for rollback listener
type RollbackListenerSpecification struct {

	// If not true, listener will not be rolled back if the configuration file is residing in a non default location (This is not a recommended configuration).
	IsForce *bool `mandatory:"false" json:"isForce"`

	Schedule *Schedule `mandatory:"false" json:"schedule"`

	Resources *Resources `mandatory:"false" json:"resources"`
}

func (m RollbackListenerSpecification) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RollbackListenerSpecification) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
