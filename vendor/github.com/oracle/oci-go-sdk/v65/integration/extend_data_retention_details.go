// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Integration API
//
// Oracle Integration API.
//

package integration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ExtendDataRetentionDetails Details for extending data retention for given integration instance
type ExtendDataRetentionDetails struct {

	// Data retention period set for given integration instance
	DataRetentionPeriod ExtendDataRetentionDetailsDataRetentionPeriodEnum `mandatory:"true" json:"dataRetentionPeriod"`
}

func (m ExtendDataRetentionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ExtendDataRetentionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingExtendDataRetentionDetailsDataRetentionPeriodEnum(string(m.DataRetentionPeriod)); !ok && m.DataRetentionPeriod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DataRetentionPeriod: %s. Supported values are: %s.", m.DataRetentionPeriod, strings.Join(GetExtendDataRetentionDetailsDataRetentionPeriodEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ExtendDataRetentionDetailsDataRetentionPeriodEnum Enum with underlying type: string
type ExtendDataRetentionDetailsDataRetentionPeriodEnum string

// Set of constants representing the allowable values for ExtendDataRetentionDetailsDataRetentionPeriodEnum
const (
	ExtendDataRetentionDetailsDataRetentionPeriod1 ExtendDataRetentionDetailsDataRetentionPeriodEnum = "MONTHS_1"
	ExtendDataRetentionDetailsDataRetentionPeriod3 ExtendDataRetentionDetailsDataRetentionPeriodEnum = "MONTHS_3"
	ExtendDataRetentionDetailsDataRetentionPeriod6 ExtendDataRetentionDetailsDataRetentionPeriodEnum = "MONTHS_6"
)

var mappingExtendDataRetentionDetailsDataRetentionPeriodEnum = map[string]ExtendDataRetentionDetailsDataRetentionPeriodEnum{
	"MONTHS_1": ExtendDataRetentionDetailsDataRetentionPeriod1,
	"MONTHS_3": ExtendDataRetentionDetailsDataRetentionPeriod3,
	"MONTHS_6": ExtendDataRetentionDetailsDataRetentionPeriod6,
}

var mappingExtendDataRetentionDetailsDataRetentionPeriodEnumLowerCase = map[string]ExtendDataRetentionDetailsDataRetentionPeriodEnum{
	"months_1": ExtendDataRetentionDetailsDataRetentionPeriod1,
	"months_3": ExtendDataRetentionDetailsDataRetentionPeriod3,
	"months_6": ExtendDataRetentionDetailsDataRetentionPeriod6,
}

// GetExtendDataRetentionDetailsDataRetentionPeriodEnumValues Enumerates the set of values for ExtendDataRetentionDetailsDataRetentionPeriodEnum
func GetExtendDataRetentionDetailsDataRetentionPeriodEnumValues() []ExtendDataRetentionDetailsDataRetentionPeriodEnum {
	values := make([]ExtendDataRetentionDetailsDataRetentionPeriodEnum, 0)
	for _, v := range mappingExtendDataRetentionDetailsDataRetentionPeriodEnum {
		values = append(values, v)
	}
	return values
}

// GetExtendDataRetentionDetailsDataRetentionPeriodEnumStringValues Enumerates the set of values in String for ExtendDataRetentionDetailsDataRetentionPeriodEnum
func GetExtendDataRetentionDetailsDataRetentionPeriodEnumStringValues() []string {
	return []string{
		"MONTHS_1",
		"MONTHS_3",
		"MONTHS_6",
	}
}

// GetMappingExtendDataRetentionDetailsDataRetentionPeriodEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingExtendDataRetentionDetailsDataRetentionPeriodEnum(val string) (ExtendDataRetentionDetailsDataRetentionPeriodEnum, bool) {
	enum, ok := mappingExtendDataRetentionDetailsDataRetentionPeriodEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
