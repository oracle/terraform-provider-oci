// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// MigrationObjectSummary Database objects to include or exclude from migration
type MigrationObjectSummary struct {

	// Owner of the object (regular expression is allowed)
	Owner *string `mandatory:"true" json:"owner"`

	// Name of the object (regular expression is allowed)
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Type of object to exclude.
	// If not specified, matching owners and object names of type TABLE would be excluded.
	Type *string `mandatory:"false" json:"type"`

	// Object status.
	ObjectStatus ObjectStatusEnum `mandatory:"false" json:"objectStatus,omitempty"`

	// Whether an excluded table should be omitted from replication. Only valid for database objects that have are of type TABLE and object status EXCLUDE.
	IsOmitExcludedTableFromReplication *bool `mandatory:"false" json:"isOmitExcludedTableFromReplication"`
}

func (m MigrationObjectSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MigrationObjectSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingObjectStatusEnum(string(m.ObjectStatus)); !ok && m.ObjectStatus != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectStatus: %s. Supported values are: %s.", m.ObjectStatus, strings.Join(GetObjectStatusEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
