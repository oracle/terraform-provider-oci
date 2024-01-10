// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Management API
//
// Use the Database Management API to perform tasks such as obtaining performance and resource usage metrics
// for a fleet of Managed Databases or a specific Managed Database, creating Managed Database Groups, and
// running a SQL job on a Managed Database or Managed Database Group.
//

package databasemanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TablespaceStorageSize Storage size.
type TablespaceStorageSize struct {

	// Storage size number in bytes, kilobytes, megabytes, gigabytes, or terabytes.
	Size *float32 `mandatory:"true" json:"size"`

	// Storage size unit: bytes, kilobytes, megabytes, gigabytes, or terabytes.
	Unit TablespaceStorageSizeUnitEnum `mandatory:"false" json:"unit,omitempty"`
}

func (m TablespaceStorageSize) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TablespaceStorageSize) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTablespaceStorageSizeUnitEnum(string(m.Unit)); !ok && m.Unit != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Unit: %s. Supported values are: %s.", m.Unit, strings.Join(GetTablespaceStorageSizeUnitEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TablespaceStorageSizeUnitEnum Enum with underlying type: string
type TablespaceStorageSizeUnitEnum string

// Set of constants representing the allowable values for TablespaceStorageSizeUnitEnum
const (
	TablespaceStorageSizeUnitBytes     TablespaceStorageSizeUnitEnum = "BYTES"
	TablespaceStorageSizeUnitKilobytes TablespaceStorageSizeUnitEnum = "KILOBYTES"
	TablespaceStorageSizeUnitMegabytes TablespaceStorageSizeUnitEnum = "MEGABYTES"
	TablespaceStorageSizeUnitGigabytes TablespaceStorageSizeUnitEnum = "GIGABYTES"
	TablespaceStorageSizeUnitTerabytes TablespaceStorageSizeUnitEnum = "TERABYTES"
)

var mappingTablespaceStorageSizeUnitEnum = map[string]TablespaceStorageSizeUnitEnum{
	"BYTES":     TablespaceStorageSizeUnitBytes,
	"KILOBYTES": TablespaceStorageSizeUnitKilobytes,
	"MEGABYTES": TablespaceStorageSizeUnitMegabytes,
	"GIGABYTES": TablespaceStorageSizeUnitGigabytes,
	"TERABYTES": TablespaceStorageSizeUnitTerabytes,
}

var mappingTablespaceStorageSizeUnitEnumLowerCase = map[string]TablespaceStorageSizeUnitEnum{
	"bytes":     TablespaceStorageSizeUnitBytes,
	"kilobytes": TablespaceStorageSizeUnitKilobytes,
	"megabytes": TablespaceStorageSizeUnitMegabytes,
	"gigabytes": TablespaceStorageSizeUnitGigabytes,
	"terabytes": TablespaceStorageSizeUnitTerabytes,
}

// GetTablespaceStorageSizeUnitEnumValues Enumerates the set of values for TablespaceStorageSizeUnitEnum
func GetTablespaceStorageSizeUnitEnumValues() []TablespaceStorageSizeUnitEnum {
	values := make([]TablespaceStorageSizeUnitEnum, 0)
	for _, v := range mappingTablespaceStorageSizeUnitEnum {
		values = append(values, v)
	}
	return values
}

// GetTablespaceStorageSizeUnitEnumStringValues Enumerates the set of values in String for TablespaceStorageSizeUnitEnum
func GetTablespaceStorageSizeUnitEnumStringValues() []string {
	return []string{
		"BYTES",
		"KILOBYTES",
		"MEGABYTES",
		"GIGABYTES",
		"TERABYTES",
	}
}

// GetMappingTablespaceStorageSizeUnitEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTablespaceStorageSizeUnitEnum(val string) (TablespaceStorageSizeUnitEnum, bool) {
	enum, ok := mappingTablespaceStorageSizeUnitEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
