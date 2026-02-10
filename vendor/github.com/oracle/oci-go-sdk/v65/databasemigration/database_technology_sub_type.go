// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DatabaseTechnologySubType Technology sub-type e.g. ADW_SHARED and database versions corresponding to the sub-type.
type DatabaseTechnologySubType struct {

	// Technology sub-type e.g. ADW_SHARED.
	TechnologySubType *string `mandatory:"false" json:"technologySubType"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	TechnologySubTypeDisplayName *string `mandatory:"false" json:"technologySubTypeDisplayName"`

	// Array of database versions
	DatabaseVersions []string `mandatory:"false" json:"databaseVersions"`
}

func (m DatabaseTechnologySubType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseTechnologySubType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
