// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Service API
//
// APIs for managing users, groups, compartments, and policies.
//

package identity

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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

// StandardTagDefinitionTemplateTypeEnum Enum with underlying type: string
type StandardTagDefinitionTemplateTypeEnum string

// Set of constants representing the allowable values for StandardTagDefinitionTemplateTypeEnum
const (
	StandardTagDefinitionTemplateTypeEnumvalue StandardTagDefinitionTemplateTypeEnum = "ENUM"
	StandardTagDefinitionTemplateTypeString    StandardTagDefinitionTemplateTypeEnum = "STRING"
)

var mappingStandardTagDefinitionTemplateType = map[string]StandardTagDefinitionTemplateTypeEnum{
	"ENUM":   StandardTagDefinitionTemplateTypeEnumvalue,
	"STRING": StandardTagDefinitionTemplateTypeString,
}

// GetStandardTagDefinitionTemplateTypeEnumValues Enumerates the set of values for StandardTagDefinitionTemplateTypeEnum
func GetStandardTagDefinitionTemplateTypeEnumValues() []StandardTagDefinitionTemplateTypeEnum {
	values := make([]StandardTagDefinitionTemplateTypeEnum, 0)
	for _, v := range mappingStandardTagDefinitionTemplateType {
		values = append(values, v)
	}
	return values
}

// StandardTagDefinitionTemplateEnumMutabilityEnum Enum with underlying type: string
type StandardTagDefinitionTemplateEnumMutabilityEnum string

// Set of constants representing the allowable values for StandardTagDefinitionTemplateEnumMutabilityEnum
const (
	StandardTagDefinitionTemplateEnumMutabilityImmutable  StandardTagDefinitionTemplateEnumMutabilityEnum = "IMMUTABLE"
	StandardTagDefinitionTemplateEnumMutabilityMutable    StandardTagDefinitionTemplateEnumMutabilityEnum = "MUTABLE"
	StandardTagDefinitionTemplateEnumMutabilityAppendable StandardTagDefinitionTemplateEnumMutabilityEnum = "APPENDABLE"
)

var mappingStandardTagDefinitionTemplateEnumMutability = map[string]StandardTagDefinitionTemplateEnumMutabilityEnum{
	"IMMUTABLE":  StandardTagDefinitionTemplateEnumMutabilityImmutable,
	"MUTABLE":    StandardTagDefinitionTemplateEnumMutabilityMutable,
	"APPENDABLE": StandardTagDefinitionTemplateEnumMutabilityAppendable,
}

// GetStandardTagDefinitionTemplateEnumMutabilityEnumValues Enumerates the set of values for StandardTagDefinitionTemplateEnumMutabilityEnum
func GetStandardTagDefinitionTemplateEnumMutabilityEnumValues() []StandardTagDefinitionTemplateEnumMutabilityEnum {
	values := make([]StandardTagDefinitionTemplateEnumMutabilityEnum, 0)
	for _, v := range mappingStandardTagDefinitionTemplateEnumMutability {
		values = append(values, v)
	}
	return values
}
