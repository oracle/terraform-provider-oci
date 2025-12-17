// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// ConvertInstanceDetails Details for converting integration instance to Disaster Recovery Enabled instance type
type ConvertInstanceDetails struct {

	// Convert given instance to specified DR instance
	ConversionType ConvertInstanceDetailsConversionTypeEnum `mandatory:"true" json:"conversionType"`
}

func (m ConvertInstanceDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ConvertInstanceDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingConvertInstanceDetailsConversionTypeEnum(string(m.ConversionType)); !ok && m.ConversionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ConversionType: %s. Supported values are: %s.", m.ConversionType, strings.Join(GetConvertInstanceDetailsConversionTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ConvertInstanceDetailsConversionTypeEnum Enum with underlying type: string
type ConvertInstanceDetailsConversionTypeEnum string

// Set of constants representing the allowable values for ConvertInstanceDetailsConversionTypeEnum
const (
	ConvertInstanceDetailsConversionTypeDisasterRecovery ConvertInstanceDetailsConversionTypeEnum = "DISASTER_RECOVERY"
)

var mappingConvertInstanceDetailsConversionTypeEnum = map[string]ConvertInstanceDetailsConversionTypeEnum{
	"DISASTER_RECOVERY": ConvertInstanceDetailsConversionTypeDisasterRecovery,
}

var mappingConvertInstanceDetailsConversionTypeEnumLowerCase = map[string]ConvertInstanceDetailsConversionTypeEnum{
	"disaster_recovery": ConvertInstanceDetailsConversionTypeDisasterRecovery,
}

// GetConvertInstanceDetailsConversionTypeEnumValues Enumerates the set of values for ConvertInstanceDetailsConversionTypeEnum
func GetConvertInstanceDetailsConversionTypeEnumValues() []ConvertInstanceDetailsConversionTypeEnum {
	values := make([]ConvertInstanceDetailsConversionTypeEnum, 0)
	for _, v := range mappingConvertInstanceDetailsConversionTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetConvertInstanceDetailsConversionTypeEnumStringValues Enumerates the set of values in String for ConvertInstanceDetailsConversionTypeEnum
func GetConvertInstanceDetailsConversionTypeEnumStringValues() []string {
	return []string{
		"DISASTER_RECOVERY",
	}
}

// GetMappingConvertInstanceDetailsConversionTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingConvertInstanceDetailsConversionTypeEnum(val string) (ConvertInstanceDetailsConversionTypeEnum, bool) {
	enum, ok := mappingConvertInstanceDetailsConversionTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
