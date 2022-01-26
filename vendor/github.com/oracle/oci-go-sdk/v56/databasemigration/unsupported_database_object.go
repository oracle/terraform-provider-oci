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

// UnsupportedDatabaseObject Database objects to exclude from migration
type UnsupportedDatabaseObject struct {

	// Owner of the object (regular expression is allowed)
	Owner *string `mandatory:"true" json:"owner"`

	// Name of the object (regular expression is allowed)
	ObjectName *string `mandatory:"true" json:"objectName"`

	// Type of unsupported object
	Type UnsupportedDatabaseObjectTypeEnum `mandatory:"false" json:"type,omitempty"`
}

func (m UnsupportedDatabaseObject) String() string {
	return common.PointerString(m)
}

// UnsupportedDatabaseObjectTypeEnum Enum with underlying type: string
type UnsupportedDatabaseObjectTypeEnum string

// Set of constants representing the allowable values for UnsupportedDatabaseObjectTypeEnum
const (
	UnsupportedDatabaseObjectTypeGoldenGate UnsupportedDatabaseObjectTypeEnum = "GOLDEN_GATE"
)

var mappingUnsupportedDatabaseObjectType = map[string]UnsupportedDatabaseObjectTypeEnum{
	"GOLDEN_GATE": UnsupportedDatabaseObjectTypeGoldenGate,
}

// GetUnsupportedDatabaseObjectTypeEnumValues Enumerates the set of values for UnsupportedDatabaseObjectTypeEnum
func GetUnsupportedDatabaseObjectTypeEnumValues() []UnsupportedDatabaseObjectTypeEnum {
	values := make([]UnsupportedDatabaseObjectTypeEnum, 0)
	for _, v := range mappingUnsupportedDatabaseObjectType {
		values = append(values, v)
	}
	return values
}
