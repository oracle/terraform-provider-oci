// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Database Autonomous Recovery Service API
//
// Use Oracle Database Autonomous Recovery Service API to manage Protected Databases.
//

package recovery

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RetentionPeriodValue Specifies the retention period for the long-term backup.
type RetentionPeriodValue struct {

	// Specifies the retention period type for the long-term backup. Allowed values are DAYS or YEARS.
	RetentionPeriodType RetentionPeriodValueRetentionPeriodTypeEnum `mandatory:"true" json:"retentionPeriodType"`

	// Specifies the duration (in days or years) to retain the long-term backup. If you have chosen the retentionPeriodType as 'DAYS', then specify a duration ranging from 90 days to 3650 days. If you have chosen the retentionPeriodType as 'YEARS', then specify a duration ranging from 1 year to 10 years.
	RetentionCount *int `mandatory:"true" json:"retentionCount"`
}

func (m RetentionPeriodValue) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RetentionPeriodValue) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingRetentionPeriodValueRetentionPeriodTypeEnum(string(m.RetentionPeriodType)); !ok && m.RetentionPeriodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RetentionPeriodType: %s. Supported values are: %s.", m.RetentionPeriodType, strings.Join(GetRetentionPeriodValueRetentionPeriodTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RetentionPeriodValueRetentionPeriodTypeEnum Enum with underlying type: string
type RetentionPeriodValueRetentionPeriodTypeEnum string

// Set of constants representing the allowable values for RetentionPeriodValueRetentionPeriodTypeEnum
const (
	RetentionPeriodValueRetentionPeriodTypeYear RetentionPeriodValueRetentionPeriodTypeEnum = "YEAR"
	RetentionPeriodValueRetentionPeriodTypeDay  RetentionPeriodValueRetentionPeriodTypeEnum = "DAY"
)

var mappingRetentionPeriodValueRetentionPeriodTypeEnum = map[string]RetentionPeriodValueRetentionPeriodTypeEnum{
	"YEAR": RetentionPeriodValueRetentionPeriodTypeYear,
	"DAY":  RetentionPeriodValueRetentionPeriodTypeDay,
}

var mappingRetentionPeriodValueRetentionPeriodTypeEnumLowerCase = map[string]RetentionPeriodValueRetentionPeriodTypeEnum{
	"year": RetentionPeriodValueRetentionPeriodTypeYear,
	"day":  RetentionPeriodValueRetentionPeriodTypeDay,
}

// GetRetentionPeriodValueRetentionPeriodTypeEnumValues Enumerates the set of values for RetentionPeriodValueRetentionPeriodTypeEnum
func GetRetentionPeriodValueRetentionPeriodTypeEnumValues() []RetentionPeriodValueRetentionPeriodTypeEnum {
	values := make([]RetentionPeriodValueRetentionPeriodTypeEnum, 0)
	for _, v := range mappingRetentionPeriodValueRetentionPeriodTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetRetentionPeriodValueRetentionPeriodTypeEnumStringValues Enumerates the set of values in String for RetentionPeriodValueRetentionPeriodTypeEnum
func GetRetentionPeriodValueRetentionPeriodTypeEnumStringValues() []string {
	return []string{
		"YEAR",
		"DAY",
	}
}

// GetMappingRetentionPeriodValueRetentionPeriodTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRetentionPeriodValueRetentionPeriodTypeEnum(val string) (RetentionPeriodValueRetentionPeriodTypeEnum, bool) {
	enum, ok := mappingRetentionPeriodValueRetentionPeriodTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
