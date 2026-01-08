// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ConvertStandbyTypeDetails The convertStandbyType request parameters.
type ConvertStandbyTypeDetails struct {

	// The administrator password of the primary database in this Data Guard association.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword"`

	// Defines the conversion type of the standby database. Specify this to convert a physical standby to a snapshot standby and vice versa.
	// Valid standbyConversionType:
	//     - SNAPSHOT
	//     - PHYSICAL
	StandbyConversionType ConvertStandbyTypeDetailsStandbyConversionTypeEnum `mandatory:"true" json:"standbyConversionType"`
}

func (m ConvertStandbyTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConvertStandbyTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConvertStandbyTypeDetailsStandbyConversionTypeEnum(string(m.StandbyConversionType)); !ok && m.StandbyConversionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StandbyConversionType: %s. Supported values are: %s.", m.StandbyConversionType, strings.Join(GetConvertStandbyTypeDetailsStandbyConversionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConvertStandbyTypeDetailsStandbyConversionTypeEnum Enum with underlying type: string
type ConvertStandbyTypeDetailsStandbyConversionTypeEnum string

// Set of constants representing the allowable values for ConvertStandbyTypeDetailsStandbyConversionTypeEnum
const (
	ConvertStandbyTypeDetailsStandbyConversionTypeSnapshot ConvertStandbyTypeDetailsStandbyConversionTypeEnum = "SNAPSHOT"
	ConvertStandbyTypeDetailsStandbyConversionTypePhysical ConvertStandbyTypeDetailsStandbyConversionTypeEnum = "PHYSICAL"
)

var mappingConvertStandbyTypeDetailsStandbyConversionTypeEnum = map[string]ConvertStandbyTypeDetailsStandbyConversionTypeEnum{
	"SNAPSHOT": ConvertStandbyTypeDetailsStandbyConversionTypeSnapshot,
	"PHYSICAL": ConvertStandbyTypeDetailsStandbyConversionTypePhysical,
}

var mappingConvertStandbyTypeDetailsStandbyConversionTypeEnumLowerCase = map[string]ConvertStandbyTypeDetailsStandbyConversionTypeEnum{
	"snapshot": ConvertStandbyTypeDetailsStandbyConversionTypeSnapshot,
	"physical": ConvertStandbyTypeDetailsStandbyConversionTypePhysical,
}

// GetConvertStandbyTypeDetailsStandbyConversionTypeEnumValues Enumerates the set of values for ConvertStandbyTypeDetailsStandbyConversionTypeEnum
func GetConvertStandbyTypeDetailsStandbyConversionTypeEnumValues() []ConvertStandbyTypeDetailsStandbyConversionTypeEnum {
	values := make([]ConvertStandbyTypeDetailsStandbyConversionTypeEnum, 0)
	for _, v := range mappingConvertStandbyTypeDetailsStandbyConversionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertStandbyTypeDetailsStandbyConversionTypeEnumStringValues Enumerates the set of values in String for ConvertStandbyTypeDetailsStandbyConversionTypeEnum
func GetConvertStandbyTypeDetailsStandbyConversionTypeEnumStringValues() []string {
	return []string{
		"SNAPSHOT",
		"PHYSICAL",
	}
}

// GetMappingConvertStandbyTypeDetailsStandbyConversionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertStandbyTypeDetailsStandbyConversionTypeEnum(val string) (ConvertStandbyTypeDetailsStandbyConversionTypeEnum, bool) {
	enum, ok := mappingConvertStandbyTypeDetailsStandbyConversionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
