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

// DatabaseTool Summary of database tools of autonomous database.
type DatabaseTool struct {

	// Name of database tool.
	Name DatabaseToolNameEnum `mandatory:"true" json:"name"`

	// Indicates whether tool is enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Compute used by database tools.
	ComputeCount *float32 `mandatory:"false" json:"computeCount"`

	// The max idle time, in minutes, after which the VM used by database tools will be terminated.
	MaxIdleTimeInMinutes *int `mandatory:"false" json:"maxIdleTimeInMinutes"`
}

func (m DatabaseTool) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DatabaseTool) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingDatabaseToolNameEnum(string(m.Name)); !ok && m.Name != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Name: %s. Supported values are: %s.", m.Name, strings.Join(GetDatabaseToolNameEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// DatabaseToolNameEnum Enum with underlying type: string
type DatabaseToolNameEnum string

// Set of constants representing the allowable values for DatabaseToolNameEnum
const (
	DatabaseToolNameApex            DatabaseToolNameEnum = "APEX"
	DatabaseToolNameDatabaseActions DatabaseToolNameEnum = "DATABASE_ACTIONS"
	DatabaseToolNameGraphStudio     DatabaseToolNameEnum = "GRAPH_STUDIO"
	DatabaseToolNameOml             DatabaseToolNameEnum = "OML"
	DatabaseToolNameDataTransforms  DatabaseToolNameEnum = "DATA_TRANSFORMS"
	DatabaseToolNameOrds            DatabaseToolNameEnum = "ORDS"
	DatabaseToolNameMongodbApi      DatabaseToolNameEnum = "MONGODB_API"
)

var mappingDatabaseToolNameEnum = map[string]DatabaseToolNameEnum{
	"APEX":             DatabaseToolNameApex,
	"DATABASE_ACTIONS": DatabaseToolNameDatabaseActions,
	"GRAPH_STUDIO":     DatabaseToolNameGraphStudio,
	"OML":              DatabaseToolNameOml,
	"DATA_TRANSFORMS":  DatabaseToolNameDataTransforms,
	"ORDS":             DatabaseToolNameOrds,
	"MONGODB_API":      DatabaseToolNameMongodbApi,
}

var mappingDatabaseToolNameEnumLowerCase = map[string]DatabaseToolNameEnum{
	"apex":             DatabaseToolNameApex,
	"database_actions": DatabaseToolNameDatabaseActions,
	"graph_studio":     DatabaseToolNameGraphStudio,
	"oml":              DatabaseToolNameOml,
	"data_transforms":  DatabaseToolNameDataTransforms,
	"ords":             DatabaseToolNameOrds,
	"mongodb_api":      DatabaseToolNameMongodbApi,
}

// GetDatabaseToolNameEnumValues Enumerates the set of values for DatabaseToolNameEnum
func GetDatabaseToolNameEnumValues() []DatabaseToolNameEnum {
	values := make([]DatabaseToolNameEnum, 0)
	for _, v := range mappingDatabaseToolNameEnum {
		values = append(values, v)
	}
	return values
}

// GetDatabaseToolNameEnumStringValues Enumerates the set of values in String for DatabaseToolNameEnum
func GetDatabaseToolNameEnumStringValues() []string {
	return []string{
		"APEX",
		"DATABASE_ACTIONS",
		"GRAPH_STUDIO",
		"OML",
		"DATA_TRANSFORMS",
		"ORDS",
		"MONGODB_API",
	}
}

// GetMappingDatabaseToolNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingDatabaseToolNameEnum(val string) (DatabaseToolNameEnum, bool) {
	enum, ok := mappingDatabaseToolNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
