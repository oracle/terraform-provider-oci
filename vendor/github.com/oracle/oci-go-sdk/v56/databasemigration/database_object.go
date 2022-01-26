// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// DatabaseObject Database objects to include or exclude from migration
type DatabaseObject struct {

	// Owner of the object (regular expression is allowed)
	Owner *string `mandatory:"true" json:"owner"`

	// Name of the object (regular expression is allowed)
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Type of object to exclude.
	// If not specified, matching owners and object names of type TABLE would be excluded.
	Type *string `mandatory:"false" json:"type"`
}

func (m DatabaseObject) String() string {
	return common.PointerString(m)
}
