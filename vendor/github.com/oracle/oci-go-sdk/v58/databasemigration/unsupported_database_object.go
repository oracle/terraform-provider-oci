// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UnsupportedDatabaseObject) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUnsupportedDatabaseObjectTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUnsupportedDatabaseObjectTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UnsupportedDatabaseObjectTypeEnum Enum with underlying type: string
type UnsupportedDatabaseObjectTypeEnum string

// Set of constants representing the allowable values for UnsupportedDatabaseObjectTypeEnum
const (
	UnsupportedDatabaseObjectTypeGoldenGate UnsupportedDatabaseObjectTypeEnum = "GOLDEN_GATE"
)

var mappingUnsupportedDatabaseObjectTypeEnum = map[string]UnsupportedDatabaseObjectTypeEnum{
	"GOLDEN_GATE": UnsupportedDatabaseObjectTypeGoldenGate,
}

// GetUnsupportedDatabaseObjectTypeEnumValues Enumerates the set of values for UnsupportedDatabaseObjectTypeEnum
func GetUnsupportedDatabaseObjectTypeEnumValues() []UnsupportedDatabaseObjectTypeEnum {
	values := make([]UnsupportedDatabaseObjectTypeEnum, 0)
	for _, v := range mappingUnsupportedDatabaseObjectTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUnsupportedDatabaseObjectTypeEnumStringValues Enumerates the set of values in String for UnsupportedDatabaseObjectTypeEnum
func GetUnsupportedDatabaseObjectTypeEnumStringValues() []string {
	return []string{
		"GOLDEN_GATE",
	}
}

// GetMappingUnsupportedDatabaseObjectTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUnsupportedDatabaseObjectTypeEnum(val string) (UnsupportedDatabaseObjectTypeEnum, bool) {
	mappingUnsupportedDatabaseObjectTypeEnumIgnoreCase := make(map[string]UnsupportedDatabaseObjectTypeEnum)
	for k, v := range mappingUnsupportedDatabaseObjectTypeEnum {
		mappingUnsupportedDatabaseObjectTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingUnsupportedDatabaseObjectTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
