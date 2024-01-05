// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PluggableDatabaseNodeLevelDetails Pluggable Database Node Level Details.
type PluggableDatabaseNodeLevelDetails struct {

	// The Node name of the Database Instance.
	NodeName *string `mandatory:"true" json:"nodeName"`

	// The mode that pluggable database is in. Open mode can only be changed to READ_ONLY or MIGRATE directly from the backend (within the Oracle Database software).
	OpenMode PluggableDatabaseNodeLevelDetailsOpenModeEnum `mandatory:"true" json:"openMode"`
}

func (m PluggableDatabaseNodeLevelDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PluggableDatabaseNodeLevelDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPluggableDatabaseNodeLevelDetailsOpenModeEnum(string(m.OpenMode)); !ok && m.OpenMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OpenMode: %s. Supported values are: %s.", m.OpenMode, strings.Join(GetPluggableDatabaseNodeLevelDetailsOpenModeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PluggableDatabaseNodeLevelDetailsOpenModeEnum Enum with underlying type: string
type PluggableDatabaseNodeLevelDetailsOpenModeEnum string

// Set of constants representing the allowable values for PluggableDatabaseNodeLevelDetailsOpenModeEnum
const (
	PluggableDatabaseNodeLevelDetailsOpenModeReadOnly  PluggableDatabaseNodeLevelDetailsOpenModeEnum = "READ_ONLY"
	PluggableDatabaseNodeLevelDetailsOpenModeReadWrite PluggableDatabaseNodeLevelDetailsOpenModeEnum = "READ_WRITE"
	PluggableDatabaseNodeLevelDetailsOpenModeMounted   PluggableDatabaseNodeLevelDetailsOpenModeEnum = "MOUNTED"
	PluggableDatabaseNodeLevelDetailsOpenModeMigrate   PluggableDatabaseNodeLevelDetailsOpenModeEnum = "MIGRATE"
)

var mappingPluggableDatabaseNodeLevelDetailsOpenModeEnum = map[string]PluggableDatabaseNodeLevelDetailsOpenModeEnum{
	"READ_ONLY":  PluggableDatabaseNodeLevelDetailsOpenModeReadOnly,
	"READ_WRITE": PluggableDatabaseNodeLevelDetailsOpenModeReadWrite,
	"MOUNTED":    PluggableDatabaseNodeLevelDetailsOpenModeMounted,
	"MIGRATE":    PluggableDatabaseNodeLevelDetailsOpenModeMigrate,
}

var mappingPluggableDatabaseNodeLevelDetailsOpenModeEnumLowerCase = map[string]PluggableDatabaseNodeLevelDetailsOpenModeEnum{
	"read_only":  PluggableDatabaseNodeLevelDetailsOpenModeReadOnly,
	"read_write": PluggableDatabaseNodeLevelDetailsOpenModeReadWrite,
	"mounted":    PluggableDatabaseNodeLevelDetailsOpenModeMounted,
	"migrate":    PluggableDatabaseNodeLevelDetailsOpenModeMigrate,
}

// GetPluggableDatabaseNodeLevelDetailsOpenModeEnumValues Enumerates the set of values for PluggableDatabaseNodeLevelDetailsOpenModeEnum
func GetPluggableDatabaseNodeLevelDetailsOpenModeEnumValues() []PluggableDatabaseNodeLevelDetailsOpenModeEnum {
	values := make([]PluggableDatabaseNodeLevelDetailsOpenModeEnum, 0)
	for _, v := range mappingPluggableDatabaseNodeLevelDetailsOpenModeEnum {
		values = append(values, v)
	}
	return values
}

// GetPluggableDatabaseNodeLevelDetailsOpenModeEnumStringValues Enumerates the set of values in String for PluggableDatabaseNodeLevelDetailsOpenModeEnum
func GetPluggableDatabaseNodeLevelDetailsOpenModeEnumStringValues() []string {
	return []string{
		"READ_ONLY",
		"READ_WRITE",
		"MOUNTED",
		"MIGRATE",
	}
}

// GetMappingPluggableDatabaseNodeLevelDetailsOpenModeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPluggableDatabaseNodeLevelDetailsOpenModeEnum(val string) (PluggableDatabaseNodeLevelDetailsOpenModeEnum, bool) {
	enum, ok := mappingPluggableDatabaseNodeLevelDetailsOpenModeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
