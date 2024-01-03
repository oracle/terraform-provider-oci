// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// Use the Identity and Access Management Service API to manage users, groups, identity domains, compartments, policies, tagging, and limits. For information about managing users, groups, compartments, and policies, see Identity and Access Management (without identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm). For information about tagging and service limits, see Tagging (https://docs.cloud.oracle.com/iaas/Content/Tagging/Concepts/taggingoverview.htm) and Service Limits (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/servicelimits.htm). For information about creating, modifying, and deleting identity domains, see Identity and Access Management (with identity domains) (https://docs.cloud.oracle.com/iaas/Content/Identity/home.htm).
//

package identity

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StandardTagDefinitionTemplate The template of the tag definition. This object includes necessary details to create the provided standard tag definition.
type StandardTagDefinitionTemplate struct {

	// The default description of the tag namespace that users can use to create the tag definition
	Description *string `mandatory:"true" json:"description"`

	// The name of this standard tag definition
	TagDefinitionName *string `mandatory:"true" json:"tagDefinitionName"`

	// The type of tag definition. Enum or string.
	Type StandardTagDefinitionTemplateTypeEnum `mandatory:"true" json:"type"`

	// Is the tag a cost tracking tag. Default will be false as cost tracking tags have been deprecated
	IsCostTracking *bool `mandatory:"true" json:"isCostTracking"`

	// List of possible values. An optional parameter that will be present if the type of definition is enum.
	PossibleValues []string `mandatory:"false" json:"possibleValues"`

	// The mutability of the possible values list for enum tags. This will default to IMMUTABLE for string value tags
	EnumMutability StandardTagDefinitionTemplateEnumMutabilityEnum `mandatory:"false" json:"enumMutability,omitempty"`
}

func (m StandardTagDefinitionTemplate) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StandardTagDefinitionTemplate) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingStandardTagDefinitionTemplateTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetStandardTagDefinitionTemplateTypeEnumStringValues(), ",")))
	}

	if _, ok := GetMappingStandardTagDefinitionTemplateEnumMutabilityEnum(string(m.EnumMutability)); !ok && m.EnumMutability != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EnumMutability: %s. Supported values are: %s.", m.EnumMutability, strings.Join(GetStandardTagDefinitionTemplateEnumMutabilityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StandardTagDefinitionTemplateTypeEnum Enum with underlying type: string
type StandardTagDefinitionTemplateTypeEnum string

// Set of constants representing the allowable values for StandardTagDefinitionTemplateTypeEnum
const (
	StandardTagDefinitionTemplateTypeEnumvalue StandardTagDefinitionTemplateTypeEnum = "ENUM"
	StandardTagDefinitionTemplateTypeString    StandardTagDefinitionTemplateTypeEnum = "STRING"
)

var mappingStandardTagDefinitionTemplateTypeEnum = map[string]StandardTagDefinitionTemplateTypeEnum{
	"ENUM":   StandardTagDefinitionTemplateTypeEnumvalue,
	"STRING": StandardTagDefinitionTemplateTypeString,
}

var mappingStandardTagDefinitionTemplateTypeEnumLowerCase = map[string]StandardTagDefinitionTemplateTypeEnum{
	"enum":   StandardTagDefinitionTemplateTypeEnumvalue,
	"string": StandardTagDefinitionTemplateTypeString,
}

// GetStandardTagDefinitionTemplateTypeEnumValues Enumerates the set of values for StandardTagDefinitionTemplateTypeEnum
func GetStandardTagDefinitionTemplateTypeEnumValues() []StandardTagDefinitionTemplateTypeEnum {
	values := make([]StandardTagDefinitionTemplateTypeEnum, 0)
	for _, v := range mappingStandardTagDefinitionTemplateTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStandardTagDefinitionTemplateTypeEnumStringValues Enumerates the set of values in String for StandardTagDefinitionTemplateTypeEnum
func GetStandardTagDefinitionTemplateTypeEnumStringValues() []string {
	return []string{
		"ENUM",
		"STRING",
	}
}

// GetMappingStandardTagDefinitionTemplateTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStandardTagDefinitionTemplateTypeEnum(val string) (StandardTagDefinitionTemplateTypeEnum, bool) {
	enum, ok := mappingStandardTagDefinitionTemplateTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// StandardTagDefinitionTemplateEnumMutabilityEnum Enum with underlying type: string
type StandardTagDefinitionTemplateEnumMutabilityEnum string

// Set of constants representing the allowable values for StandardTagDefinitionTemplateEnumMutabilityEnum
const (
	StandardTagDefinitionTemplateEnumMutabilityImmutable  StandardTagDefinitionTemplateEnumMutabilityEnum = "IMMUTABLE"
	StandardTagDefinitionTemplateEnumMutabilityMutable    StandardTagDefinitionTemplateEnumMutabilityEnum = "MUTABLE"
	StandardTagDefinitionTemplateEnumMutabilityAppendable StandardTagDefinitionTemplateEnumMutabilityEnum = "APPENDABLE"
)

var mappingStandardTagDefinitionTemplateEnumMutabilityEnum = map[string]StandardTagDefinitionTemplateEnumMutabilityEnum{
	"IMMUTABLE":  StandardTagDefinitionTemplateEnumMutabilityImmutable,
	"MUTABLE":    StandardTagDefinitionTemplateEnumMutabilityMutable,
	"APPENDABLE": StandardTagDefinitionTemplateEnumMutabilityAppendable,
}

var mappingStandardTagDefinitionTemplateEnumMutabilityEnumLowerCase = map[string]StandardTagDefinitionTemplateEnumMutabilityEnum{
	"immutable":  StandardTagDefinitionTemplateEnumMutabilityImmutable,
	"mutable":    StandardTagDefinitionTemplateEnumMutabilityMutable,
	"appendable": StandardTagDefinitionTemplateEnumMutabilityAppendable,
}

// GetStandardTagDefinitionTemplateEnumMutabilityEnumValues Enumerates the set of values for StandardTagDefinitionTemplateEnumMutabilityEnum
func GetStandardTagDefinitionTemplateEnumMutabilityEnumValues() []StandardTagDefinitionTemplateEnumMutabilityEnum {
	values := make([]StandardTagDefinitionTemplateEnumMutabilityEnum, 0)
	for _, v := range mappingStandardTagDefinitionTemplateEnumMutabilityEnum {
		values = append(values, v)
	}
	return values
}

// GetStandardTagDefinitionTemplateEnumMutabilityEnumStringValues Enumerates the set of values in String for StandardTagDefinitionTemplateEnumMutabilityEnum
func GetStandardTagDefinitionTemplateEnumMutabilityEnumStringValues() []string {
	return []string{
		"IMMUTABLE",
		"MUTABLE",
		"APPENDABLE",
	}
}

// GetMappingStandardTagDefinitionTemplateEnumMutabilityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStandardTagDefinitionTemplateEnumMutabilityEnum(val string) (StandardTagDefinitionTemplateEnumMutabilityEnum, bool) {
	enum, ok := mappingStandardTagDefinitionTemplateEnumMutabilityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
