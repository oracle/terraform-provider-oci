// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Oracle Cloud Migrations API
//
// A description of the Oracle Cloud Migrations API.
//

package cloudmigrations

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// OlvmMigrationOptions Reference to cluster-wide configuration of migration of a running virtual machine to another host.
type OlvmMigrationOptions struct {

	// Enum representing the boolean value that can be either set, or inherited from a higher level. The inheritance order is virtual machine ??? cluster ??? engine-config.
	AutoConverge OlvmMigrationOptionsAutoConvergeEnum `mandatory:"false" json:"autoConverge,omitempty"`

	Bandwidth *MigrationBandwidth `mandatory:"false" json:"bandwidth"`

	// Enum representing the boolean value that can be either set, or inherited from a higher level. The inheritance order is virtual machine ??? cluster ??? engine-config.
	Compressed OlvmMigrationOptionsCompressedEnum `mandatory:"false" json:"compressed,omitempty"`

	// Specifies how many parallel migration connections to use.
	CustomParallelMigrations *int `mandatory:"false" json:"customParallelMigrations"`

	// Enum representing the boolean value that can be either set, or inherited from a higher level. The inheritance order is virtual machine ??? cluster ??? engine-config.
	Encrypted OlvmMigrationOptionsEncryptedEnum `mandatory:"false" json:"encrypted,omitempty"`

	// Type representing parallel migration connections policy.
	ParallelMigrationsPolicy OlvmMigrationOptionsParallelMigrationsPolicyEnum `mandatory:"false" json:"parallelMigrationsPolicy,omitempty"`
}

func (m OlvmMigrationOptions) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m OlvmMigrationOptions) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingOlvmMigrationOptionsAutoConvergeEnum(string(m.AutoConverge)); !ok && m.AutoConverge != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AutoConverge: %s. Supported values are: %s.", m.AutoConverge, strings.Join(GetOlvmMigrationOptionsAutoConvergeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmMigrationOptionsCompressedEnum(string(m.Compressed)); !ok && m.Compressed != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Compressed: %s. Supported values are: %s.", m.Compressed, strings.Join(GetOlvmMigrationOptionsCompressedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmMigrationOptionsEncryptedEnum(string(m.Encrypted)); !ok && m.Encrypted != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Encrypted: %s. Supported values are: %s.", m.Encrypted, strings.Join(GetOlvmMigrationOptionsEncryptedEnumStringValues(), ",")))
	}
	if _, ok := GetMappingOlvmMigrationOptionsParallelMigrationsPolicyEnum(string(m.ParallelMigrationsPolicy)); !ok && m.ParallelMigrationsPolicy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ParallelMigrationsPolicy: %s. Supported values are: %s.", m.ParallelMigrationsPolicy, strings.Join(GetOlvmMigrationOptionsParallelMigrationsPolicyEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// OlvmMigrationOptionsAutoConvergeEnum Enum with underlying type: string
type OlvmMigrationOptionsAutoConvergeEnum string

// Set of constants representing the allowable values for OlvmMigrationOptionsAutoConvergeEnum
const (
	OlvmMigrationOptionsAutoConvergeFalse   OlvmMigrationOptionsAutoConvergeEnum = "FALSE"
	OlvmMigrationOptionsAutoConvergeInherit OlvmMigrationOptionsAutoConvergeEnum = "INHERIT"
	OlvmMigrationOptionsAutoConvergeTrue    OlvmMigrationOptionsAutoConvergeEnum = "TRUE"
)

var mappingOlvmMigrationOptionsAutoConvergeEnum = map[string]OlvmMigrationOptionsAutoConvergeEnum{
	"FALSE":   OlvmMigrationOptionsAutoConvergeFalse,
	"INHERIT": OlvmMigrationOptionsAutoConvergeInherit,
	"TRUE":    OlvmMigrationOptionsAutoConvergeTrue,
}

var mappingOlvmMigrationOptionsAutoConvergeEnumLowerCase = map[string]OlvmMigrationOptionsAutoConvergeEnum{
	"false":   OlvmMigrationOptionsAutoConvergeFalse,
	"inherit": OlvmMigrationOptionsAutoConvergeInherit,
	"true":    OlvmMigrationOptionsAutoConvergeTrue,
}

// GetOlvmMigrationOptionsAutoConvergeEnumValues Enumerates the set of values for OlvmMigrationOptionsAutoConvergeEnum
func GetOlvmMigrationOptionsAutoConvergeEnumValues() []OlvmMigrationOptionsAutoConvergeEnum {
	values := make([]OlvmMigrationOptionsAutoConvergeEnum, 0)
	for _, v := range mappingOlvmMigrationOptionsAutoConvergeEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmMigrationOptionsAutoConvergeEnumStringValues Enumerates the set of values in String for OlvmMigrationOptionsAutoConvergeEnum
func GetOlvmMigrationOptionsAutoConvergeEnumStringValues() []string {
	return []string{
		"FALSE",
		"INHERIT",
		"TRUE",
	}
}

// GetMappingOlvmMigrationOptionsAutoConvergeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmMigrationOptionsAutoConvergeEnum(val string) (OlvmMigrationOptionsAutoConvergeEnum, bool) {
	enum, ok := mappingOlvmMigrationOptionsAutoConvergeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmMigrationOptionsCompressedEnum Enum with underlying type: string
type OlvmMigrationOptionsCompressedEnum string

// Set of constants representing the allowable values for OlvmMigrationOptionsCompressedEnum
const (
	OlvmMigrationOptionsCompressedFalse   OlvmMigrationOptionsCompressedEnum = "FALSE"
	OlvmMigrationOptionsCompressedInherit OlvmMigrationOptionsCompressedEnum = "INHERIT"
	OlvmMigrationOptionsCompressedTrue    OlvmMigrationOptionsCompressedEnum = "TRUE"
)

var mappingOlvmMigrationOptionsCompressedEnum = map[string]OlvmMigrationOptionsCompressedEnum{
	"FALSE":   OlvmMigrationOptionsCompressedFalse,
	"INHERIT": OlvmMigrationOptionsCompressedInherit,
	"TRUE":    OlvmMigrationOptionsCompressedTrue,
}

var mappingOlvmMigrationOptionsCompressedEnumLowerCase = map[string]OlvmMigrationOptionsCompressedEnum{
	"false":   OlvmMigrationOptionsCompressedFalse,
	"inherit": OlvmMigrationOptionsCompressedInherit,
	"true":    OlvmMigrationOptionsCompressedTrue,
}

// GetOlvmMigrationOptionsCompressedEnumValues Enumerates the set of values for OlvmMigrationOptionsCompressedEnum
func GetOlvmMigrationOptionsCompressedEnumValues() []OlvmMigrationOptionsCompressedEnum {
	values := make([]OlvmMigrationOptionsCompressedEnum, 0)
	for _, v := range mappingOlvmMigrationOptionsCompressedEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmMigrationOptionsCompressedEnumStringValues Enumerates the set of values in String for OlvmMigrationOptionsCompressedEnum
func GetOlvmMigrationOptionsCompressedEnumStringValues() []string {
	return []string{
		"FALSE",
		"INHERIT",
		"TRUE",
	}
}

// GetMappingOlvmMigrationOptionsCompressedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmMigrationOptionsCompressedEnum(val string) (OlvmMigrationOptionsCompressedEnum, bool) {
	enum, ok := mappingOlvmMigrationOptionsCompressedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmMigrationOptionsEncryptedEnum Enum with underlying type: string
type OlvmMigrationOptionsEncryptedEnum string

// Set of constants representing the allowable values for OlvmMigrationOptionsEncryptedEnum
const (
	OlvmMigrationOptionsEncryptedFalse   OlvmMigrationOptionsEncryptedEnum = "FALSE"
	OlvmMigrationOptionsEncryptedInherit OlvmMigrationOptionsEncryptedEnum = "INHERIT"
	OlvmMigrationOptionsEncryptedTrue    OlvmMigrationOptionsEncryptedEnum = "TRUE"
)

var mappingOlvmMigrationOptionsEncryptedEnum = map[string]OlvmMigrationOptionsEncryptedEnum{
	"FALSE":   OlvmMigrationOptionsEncryptedFalse,
	"INHERIT": OlvmMigrationOptionsEncryptedInherit,
	"TRUE":    OlvmMigrationOptionsEncryptedTrue,
}

var mappingOlvmMigrationOptionsEncryptedEnumLowerCase = map[string]OlvmMigrationOptionsEncryptedEnum{
	"false":   OlvmMigrationOptionsEncryptedFalse,
	"inherit": OlvmMigrationOptionsEncryptedInherit,
	"true":    OlvmMigrationOptionsEncryptedTrue,
}

// GetOlvmMigrationOptionsEncryptedEnumValues Enumerates the set of values for OlvmMigrationOptionsEncryptedEnum
func GetOlvmMigrationOptionsEncryptedEnumValues() []OlvmMigrationOptionsEncryptedEnum {
	values := make([]OlvmMigrationOptionsEncryptedEnum, 0)
	for _, v := range mappingOlvmMigrationOptionsEncryptedEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmMigrationOptionsEncryptedEnumStringValues Enumerates the set of values in String for OlvmMigrationOptionsEncryptedEnum
func GetOlvmMigrationOptionsEncryptedEnumStringValues() []string {
	return []string{
		"FALSE",
		"INHERIT",
		"TRUE",
	}
}

// GetMappingOlvmMigrationOptionsEncryptedEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmMigrationOptionsEncryptedEnum(val string) (OlvmMigrationOptionsEncryptedEnum, bool) {
	enum, ok := mappingOlvmMigrationOptionsEncryptedEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// OlvmMigrationOptionsParallelMigrationsPolicyEnum Enum with underlying type: string
type OlvmMigrationOptionsParallelMigrationsPolicyEnum string

// Set of constants representing the allowable values for OlvmMigrationOptionsParallelMigrationsPolicyEnum
const (
	OlvmMigrationOptionsParallelMigrationsPolicyAuto         OlvmMigrationOptionsParallelMigrationsPolicyEnum = "AUTO"
	OlvmMigrationOptionsParallelMigrationsPolicyAutoParallel OlvmMigrationOptionsParallelMigrationsPolicyEnum = "AUTO_PARALLEL"
	OlvmMigrationOptionsParallelMigrationsPolicyCustom       OlvmMigrationOptionsParallelMigrationsPolicyEnum = "CUSTOM"
	OlvmMigrationOptionsParallelMigrationsPolicyDisabled     OlvmMigrationOptionsParallelMigrationsPolicyEnum = "DISABLED"
	OlvmMigrationOptionsParallelMigrationsPolicyInherit      OlvmMigrationOptionsParallelMigrationsPolicyEnum = "INHERIT"
)

var mappingOlvmMigrationOptionsParallelMigrationsPolicyEnum = map[string]OlvmMigrationOptionsParallelMigrationsPolicyEnum{
	"AUTO":          OlvmMigrationOptionsParallelMigrationsPolicyAuto,
	"AUTO_PARALLEL": OlvmMigrationOptionsParallelMigrationsPolicyAutoParallel,
	"CUSTOM":        OlvmMigrationOptionsParallelMigrationsPolicyCustom,
	"DISABLED":      OlvmMigrationOptionsParallelMigrationsPolicyDisabled,
	"INHERIT":       OlvmMigrationOptionsParallelMigrationsPolicyInherit,
}

var mappingOlvmMigrationOptionsParallelMigrationsPolicyEnumLowerCase = map[string]OlvmMigrationOptionsParallelMigrationsPolicyEnum{
	"auto":          OlvmMigrationOptionsParallelMigrationsPolicyAuto,
	"auto_parallel": OlvmMigrationOptionsParallelMigrationsPolicyAutoParallel,
	"custom":        OlvmMigrationOptionsParallelMigrationsPolicyCustom,
	"disabled":      OlvmMigrationOptionsParallelMigrationsPolicyDisabled,
	"inherit":       OlvmMigrationOptionsParallelMigrationsPolicyInherit,
}

// GetOlvmMigrationOptionsParallelMigrationsPolicyEnumValues Enumerates the set of values for OlvmMigrationOptionsParallelMigrationsPolicyEnum
func GetOlvmMigrationOptionsParallelMigrationsPolicyEnumValues() []OlvmMigrationOptionsParallelMigrationsPolicyEnum {
	values := make([]OlvmMigrationOptionsParallelMigrationsPolicyEnum, 0)
	for _, v := range mappingOlvmMigrationOptionsParallelMigrationsPolicyEnum {
		values = append(values, v)
	}
	return values
}

// GetOlvmMigrationOptionsParallelMigrationsPolicyEnumStringValues Enumerates the set of values in String for OlvmMigrationOptionsParallelMigrationsPolicyEnum
func GetOlvmMigrationOptionsParallelMigrationsPolicyEnumStringValues() []string {
	return []string{
		"AUTO",
		"AUTO_PARALLEL",
		"CUSTOM",
		"DISABLED",
		"INHERIT",
	}
}

// GetMappingOlvmMigrationOptionsParallelMigrationsPolicyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingOlvmMigrationOptionsParallelMigrationsPolicyEnum(val string) (OlvmMigrationOptionsParallelMigrationsPolicyEnum, bool) {
	enum, ok := mappingOlvmMigrationOptionsParallelMigrationsPolicyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
