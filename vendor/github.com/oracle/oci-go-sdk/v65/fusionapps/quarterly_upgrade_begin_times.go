// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fusion Applications Environment Management API
//
// Use the Fusion Applications Environment Management API to manage the environments where your Fusion Applications run. For more information, see the Fusion Applications Environment Management documentation (https://docs.cloud.oracle.com/iaas/Content/fusion-applications/home.htm).
//

package fusionapps

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QuarterlyUpgradeBeginTimes Determines the quarterly upgrade begin times (monthly maintenance group schedule ) of the Fusion environment.
type QuarterlyUpgradeBeginTimes struct {

	// Determines if the maintenance schedule of the Fusion environment is inherited from the Fusion environment family.
	OverrideType QuarterlyUpgradeBeginTimesOverrideTypeEnum `mandatory:"false" json:"overrideType,omitempty"`

	// The frequency and month when maintenance occurs for the Fusion environment.
	BeginTimesValue *string `mandatory:"false" json:"beginTimesValue"`
}

func (m QuarterlyUpgradeBeginTimes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QuarterlyUpgradeBeginTimes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingQuarterlyUpgradeBeginTimesOverrideTypeEnum(string(m.OverrideType)); !ok && m.OverrideType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for OverrideType: %s. Supported values are: %s.", m.OverrideType, strings.Join(GetQuarterlyUpgradeBeginTimesOverrideTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// QuarterlyUpgradeBeginTimesOverrideTypeEnum Enum with underlying type: string
type QuarterlyUpgradeBeginTimesOverrideTypeEnum string

// Set of constants representing the allowable values for QuarterlyUpgradeBeginTimesOverrideTypeEnum
const (
	QuarterlyUpgradeBeginTimesOverrideTypeOverridden QuarterlyUpgradeBeginTimesOverrideTypeEnum = "OVERRIDDEN"
	QuarterlyUpgradeBeginTimesOverrideTypeInherited  QuarterlyUpgradeBeginTimesOverrideTypeEnum = "INHERITED"
)

var mappingQuarterlyUpgradeBeginTimesOverrideTypeEnum = map[string]QuarterlyUpgradeBeginTimesOverrideTypeEnum{
	"OVERRIDDEN": QuarterlyUpgradeBeginTimesOverrideTypeOverridden,
	"INHERITED":  QuarterlyUpgradeBeginTimesOverrideTypeInherited,
}

var mappingQuarterlyUpgradeBeginTimesOverrideTypeEnumLowerCase = map[string]QuarterlyUpgradeBeginTimesOverrideTypeEnum{
	"overridden": QuarterlyUpgradeBeginTimesOverrideTypeOverridden,
	"inherited":  QuarterlyUpgradeBeginTimesOverrideTypeInherited,
}

// GetQuarterlyUpgradeBeginTimesOverrideTypeEnumValues Enumerates the set of values for QuarterlyUpgradeBeginTimesOverrideTypeEnum
func GetQuarterlyUpgradeBeginTimesOverrideTypeEnumValues() []QuarterlyUpgradeBeginTimesOverrideTypeEnum {
	values := make([]QuarterlyUpgradeBeginTimesOverrideTypeEnum, 0)
	for _, v := range mappingQuarterlyUpgradeBeginTimesOverrideTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetQuarterlyUpgradeBeginTimesOverrideTypeEnumStringValues Enumerates the set of values in String for QuarterlyUpgradeBeginTimesOverrideTypeEnum
func GetQuarterlyUpgradeBeginTimesOverrideTypeEnumStringValues() []string {
	return []string{
		"OVERRIDDEN",
		"INHERITED",
	}
}

// GetMappingQuarterlyUpgradeBeginTimesOverrideTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingQuarterlyUpgradeBeginTimesOverrideTypeEnum(val string) (QuarterlyUpgradeBeginTimesOverrideTypeEnum, bool) {
	enum, ok := mappingQuarterlyUpgradeBeginTimesOverrideTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
