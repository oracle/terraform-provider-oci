// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Integration API
//
// Use the Data Integration API to organize your data integration projects, create data flows, pipelines and tasks, and then publish, schedule, and run tasks that extract, transform, and load data. For more information, see Data Integration (https://docs.oracle.com/iaas/data-integration/home.htm).
//

package dataintegration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// FlattenProjectionPreferences The preferences for the flatten operation.
type FlattenProjectionPreferences struct {

	// Property defining whether to create array indexes in flattened result.
	CreateArrayIndex FlattenProjectionPreferencesCreateArrayIndexEnum `mandatory:"true" json:"createArrayIndex"`

	// Property defining whether to retain all attributes in flattened result.
	RetainAllAttributes FlattenProjectionPreferencesRetainAllAttributesEnum `mandatory:"true" json:"retainAllAttributes"`

	// Property defining whether to ignore null values in flattened result.
	IgnoreNullValues FlattenProjectionPreferencesIgnoreNullValuesEnum `mandatory:"true" json:"ignoreNullValues"`

	// Property defining whether to retain parent name lineage.
	RetainParentNameLineage FlattenProjectionPreferencesRetainParentNameLineageEnum `mandatory:"true" json:"retainParentNameLineage"`
}

func (m FlattenProjectionPreferences) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m FlattenProjectionPreferences) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingFlattenProjectionPreferencesCreateArrayIndexEnum(string(m.CreateArrayIndex)); !ok && m.CreateArrayIndex != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CreateArrayIndex: %s. Supported values are: %s.", m.CreateArrayIndex, strings.Join(GetFlattenProjectionPreferencesCreateArrayIndexEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFlattenProjectionPreferencesRetainAllAttributesEnum(string(m.RetainAllAttributes)); !ok && m.RetainAllAttributes != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RetainAllAttributes: %s. Supported values are: %s.", m.RetainAllAttributes, strings.Join(GetFlattenProjectionPreferencesRetainAllAttributesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFlattenProjectionPreferencesIgnoreNullValuesEnum(string(m.IgnoreNullValues)); !ok && m.IgnoreNullValues != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for IgnoreNullValues: %s. Supported values are: %s.", m.IgnoreNullValues, strings.Join(GetFlattenProjectionPreferencesIgnoreNullValuesEnumStringValues(), ",")))
	}
	if _, ok := GetMappingFlattenProjectionPreferencesRetainParentNameLineageEnum(string(m.RetainParentNameLineage)); !ok && m.RetainParentNameLineage != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RetainParentNameLineage: %s. Supported values are: %s.", m.RetainParentNameLineage, strings.Join(GetFlattenProjectionPreferencesRetainParentNameLineageEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// FlattenProjectionPreferencesCreateArrayIndexEnum Enum with underlying type: string
type FlattenProjectionPreferencesCreateArrayIndexEnum string

// Set of constants representing the allowable values for FlattenProjectionPreferencesCreateArrayIndexEnum
const (
	FlattenProjectionPreferencesCreateArrayIndexAllow      FlattenProjectionPreferencesCreateArrayIndexEnum = "ALLOW"
	FlattenProjectionPreferencesCreateArrayIndexDoNotAllow FlattenProjectionPreferencesCreateArrayIndexEnum = "DO_NOT_ALLOW"
)

var mappingFlattenProjectionPreferencesCreateArrayIndexEnum = map[string]FlattenProjectionPreferencesCreateArrayIndexEnum{
	"ALLOW":        FlattenProjectionPreferencesCreateArrayIndexAllow,
	"DO_NOT_ALLOW": FlattenProjectionPreferencesCreateArrayIndexDoNotAllow,
}

var mappingFlattenProjectionPreferencesCreateArrayIndexEnumLowerCase = map[string]FlattenProjectionPreferencesCreateArrayIndexEnum{
	"allow":        FlattenProjectionPreferencesCreateArrayIndexAllow,
	"do_not_allow": FlattenProjectionPreferencesCreateArrayIndexDoNotAllow,
}

// GetFlattenProjectionPreferencesCreateArrayIndexEnumValues Enumerates the set of values for FlattenProjectionPreferencesCreateArrayIndexEnum
func GetFlattenProjectionPreferencesCreateArrayIndexEnumValues() []FlattenProjectionPreferencesCreateArrayIndexEnum {
	values := make([]FlattenProjectionPreferencesCreateArrayIndexEnum, 0)
	for _, v := range mappingFlattenProjectionPreferencesCreateArrayIndexEnum {
		values = append(values, v)
	}
	return values
}

// GetFlattenProjectionPreferencesCreateArrayIndexEnumStringValues Enumerates the set of values in String for FlattenProjectionPreferencesCreateArrayIndexEnum
func GetFlattenProjectionPreferencesCreateArrayIndexEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DO_NOT_ALLOW",
	}
}

// GetMappingFlattenProjectionPreferencesCreateArrayIndexEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlattenProjectionPreferencesCreateArrayIndexEnum(val string) (FlattenProjectionPreferencesCreateArrayIndexEnum, bool) {
	enum, ok := mappingFlattenProjectionPreferencesCreateArrayIndexEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FlattenProjectionPreferencesRetainAllAttributesEnum Enum with underlying type: string
type FlattenProjectionPreferencesRetainAllAttributesEnum string

// Set of constants representing the allowable values for FlattenProjectionPreferencesRetainAllAttributesEnum
const (
	FlattenProjectionPreferencesRetainAllAttributesAllow      FlattenProjectionPreferencesRetainAllAttributesEnum = "ALLOW"
	FlattenProjectionPreferencesRetainAllAttributesDoNotAllow FlattenProjectionPreferencesRetainAllAttributesEnum = "DO_NOT_ALLOW"
)

var mappingFlattenProjectionPreferencesRetainAllAttributesEnum = map[string]FlattenProjectionPreferencesRetainAllAttributesEnum{
	"ALLOW":        FlattenProjectionPreferencesRetainAllAttributesAllow,
	"DO_NOT_ALLOW": FlattenProjectionPreferencesRetainAllAttributesDoNotAllow,
}

var mappingFlattenProjectionPreferencesRetainAllAttributesEnumLowerCase = map[string]FlattenProjectionPreferencesRetainAllAttributesEnum{
	"allow":        FlattenProjectionPreferencesRetainAllAttributesAllow,
	"do_not_allow": FlattenProjectionPreferencesRetainAllAttributesDoNotAllow,
}

// GetFlattenProjectionPreferencesRetainAllAttributesEnumValues Enumerates the set of values for FlattenProjectionPreferencesRetainAllAttributesEnum
func GetFlattenProjectionPreferencesRetainAllAttributesEnumValues() []FlattenProjectionPreferencesRetainAllAttributesEnum {
	values := make([]FlattenProjectionPreferencesRetainAllAttributesEnum, 0)
	for _, v := range mappingFlattenProjectionPreferencesRetainAllAttributesEnum {
		values = append(values, v)
	}
	return values
}

// GetFlattenProjectionPreferencesRetainAllAttributesEnumStringValues Enumerates the set of values in String for FlattenProjectionPreferencesRetainAllAttributesEnum
func GetFlattenProjectionPreferencesRetainAllAttributesEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DO_NOT_ALLOW",
	}
}

// GetMappingFlattenProjectionPreferencesRetainAllAttributesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlattenProjectionPreferencesRetainAllAttributesEnum(val string) (FlattenProjectionPreferencesRetainAllAttributesEnum, bool) {
	enum, ok := mappingFlattenProjectionPreferencesRetainAllAttributesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FlattenProjectionPreferencesIgnoreNullValuesEnum Enum with underlying type: string
type FlattenProjectionPreferencesIgnoreNullValuesEnum string

// Set of constants representing the allowable values for FlattenProjectionPreferencesIgnoreNullValuesEnum
const (
	FlattenProjectionPreferencesIgnoreNullValuesAllow      FlattenProjectionPreferencesIgnoreNullValuesEnum = "ALLOW"
	FlattenProjectionPreferencesIgnoreNullValuesDoNotAllow FlattenProjectionPreferencesIgnoreNullValuesEnum = "DO_NOT_ALLOW"
)

var mappingFlattenProjectionPreferencesIgnoreNullValuesEnum = map[string]FlattenProjectionPreferencesIgnoreNullValuesEnum{
	"ALLOW":        FlattenProjectionPreferencesIgnoreNullValuesAllow,
	"DO_NOT_ALLOW": FlattenProjectionPreferencesIgnoreNullValuesDoNotAllow,
}

var mappingFlattenProjectionPreferencesIgnoreNullValuesEnumLowerCase = map[string]FlattenProjectionPreferencesIgnoreNullValuesEnum{
	"allow":        FlattenProjectionPreferencesIgnoreNullValuesAllow,
	"do_not_allow": FlattenProjectionPreferencesIgnoreNullValuesDoNotAllow,
}

// GetFlattenProjectionPreferencesIgnoreNullValuesEnumValues Enumerates the set of values for FlattenProjectionPreferencesIgnoreNullValuesEnum
func GetFlattenProjectionPreferencesIgnoreNullValuesEnumValues() []FlattenProjectionPreferencesIgnoreNullValuesEnum {
	values := make([]FlattenProjectionPreferencesIgnoreNullValuesEnum, 0)
	for _, v := range mappingFlattenProjectionPreferencesIgnoreNullValuesEnum {
		values = append(values, v)
	}
	return values
}

// GetFlattenProjectionPreferencesIgnoreNullValuesEnumStringValues Enumerates the set of values in String for FlattenProjectionPreferencesIgnoreNullValuesEnum
func GetFlattenProjectionPreferencesIgnoreNullValuesEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DO_NOT_ALLOW",
	}
}

// GetMappingFlattenProjectionPreferencesIgnoreNullValuesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlattenProjectionPreferencesIgnoreNullValuesEnum(val string) (FlattenProjectionPreferencesIgnoreNullValuesEnum, bool) {
	enum, ok := mappingFlattenProjectionPreferencesIgnoreNullValuesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// FlattenProjectionPreferencesRetainParentNameLineageEnum Enum with underlying type: string
type FlattenProjectionPreferencesRetainParentNameLineageEnum string

// Set of constants representing the allowable values for FlattenProjectionPreferencesRetainParentNameLineageEnum
const (
	FlattenProjectionPreferencesRetainParentNameLineageAllow      FlattenProjectionPreferencesRetainParentNameLineageEnum = "ALLOW"
	FlattenProjectionPreferencesRetainParentNameLineageDoNotAllow FlattenProjectionPreferencesRetainParentNameLineageEnum = "DO_NOT_ALLOW"
)

var mappingFlattenProjectionPreferencesRetainParentNameLineageEnum = map[string]FlattenProjectionPreferencesRetainParentNameLineageEnum{
	"ALLOW":        FlattenProjectionPreferencesRetainParentNameLineageAllow,
	"DO_NOT_ALLOW": FlattenProjectionPreferencesRetainParentNameLineageDoNotAllow,
}

var mappingFlattenProjectionPreferencesRetainParentNameLineageEnumLowerCase = map[string]FlattenProjectionPreferencesRetainParentNameLineageEnum{
	"allow":        FlattenProjectionPreferencesRetainParentNameLineageAllow,
	"do_not_allow": FlattenProjectionPreferencesRetainParentNameLineageDoNotAllow,
}

// GetFlattenProjectionPreferencesRetainParentNameLineageEnumValues Enumerates the set of values for FlattenProjectionPreferencesRetainParentNameLineageEnum
func GetFlattenProjectionPreferencesRetainParentNameLineageEnumValues() []FlattenProjectionPreferencesRetainParentNameLineageEnum {
	values := make([]FlattenProjectionPreferencesRetainParentNameLineageEnum, 0)
	for _, v := range mappingFlattenProjectionPreferencesRetainParentNameLineageEnum {
		values = append(values, v)
	}
	return values
}

// GetFlattenProjectionPreferencesRetainParentNameLineageEnumStringValues Enumerates the set of values in String for FlattenProjectionPreferencesRetainParentNameLineageEnum
func GetFlattenProjectionPreferencesRetainParentNameLineageEnumStringValues() []string {
	return []string{
		"ALLOW",
		"DO_NOT_ALLOW",
	}
}

// GetMappingFlattenProjectionPreferencesRetainParentNameLineageEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingFlattenProjectionPreferencesRetainParentNameLineageEnum(val string) (FlattenProjectionPreferencesRetainParentNameLineageEnum, bool) {
	enum, ok := mappingFlattenProjectionPreferencesRetainParentNameLineageEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
