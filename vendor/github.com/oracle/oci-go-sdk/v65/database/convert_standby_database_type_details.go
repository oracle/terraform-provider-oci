// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ConvertStandbyDatabaseTypeDetails The convertStandbyDatabaseType request parameters.
type ConvertStandbyDatabaseTypeDetails struct {

	// The administrator password of the primary database in this Data Guard association.
	// **The password MUST be the same as the primary admin password.**
	DatabaseAdminPassword *string `mandatory:"true" json:"databaseAdminPassword"`

	// Defines the conversion type of the standby database. Specify this to convert a physical standby to a snapshot standby and vice versa.
	// Valid standbyConversionType:
	//     - SNAPSHOT
	//     - PHYSICAL
	StandbyConversionType ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum `mandatory:"true" json:"standbyConversionType"`

	// SnapshotDurationInDays is the duration in day(s) after which the Snapshot Standby Database will get converted back to Physical Standby.
	// The minimum value of snapshotDurationInDays is 3 days and maximum value is 14 days. Default value will be 7 days if not provided in the Request.
	// This field is only applicable if the requested database role is snapshot standby.
	SnapshotDurationInDays *int `mandatory:"false" json:"snapshotDurationInDays"`
}

func (m ConvertStandbyDatabaseTypeDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConvertStandbyDatabaseTypeDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum(string(m.StandbyConversionType)); !ok && m.StandbyConversionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for StandbyConversionType: %s. Supported values are: %s.", m.StandbyConversionType, strings.Join(GetConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum Enum with underlying type: string
type ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum string

// Set of constants representing the allowable values for ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum
const (
	ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeSnapshot ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum = "SNAPSHOT"
	ConvertStandbyDatabaseTypeDetailsStandbyConversionTypePhysical ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum = "PHYSICAL"
)

var mappingConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum = map[string]ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum{
	"SNAPSHOT": ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeSnapshot,
	"PHYSICAL": ConvertStandbyDatabaseTypeDetailsStandbyConversionTypePhysical,
}

var mappingConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnumLowerCase = map[string]ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum{
	"snapshot": ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeSnapshot,
	"physical": ConvertStandbyDatabaseTypeDetailsStandbyConversionTypePhysical,
}

// GetConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnumValues Enumerates the set of values for ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum
func GetConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnumValues() []ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum {
	values := make([]ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum, 0)
	for _, v := range mappingConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnumStringValues Enumerates the set of values in String for ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum
func GetConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnumStringValues() []string {
	return []string{
		"SNAPSHOT",
		"PHYSICAL",
	}
}

// GetMappingConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum(val string) (ConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnum, bool) {
	enum, ok := mappingConvertStandbyDatabaseTypeDetailsStandbyConversionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
