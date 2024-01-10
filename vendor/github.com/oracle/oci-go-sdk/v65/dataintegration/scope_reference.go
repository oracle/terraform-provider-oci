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

// ScopeReference The `ScopeReference` class is a base class for any model object that wraps a scope reference to a TypedObject.
type ScopeReference struct {

	// A key or shallow reference to an object.  For direct reference, it points to the actual scope object.  For BOUND_ENTITY_SHAPE or BOUND_ENTITY_SHAPE_FIELD, it points to the source or target operator.   For OCI_FUNCTION_INPUT_SHAPE or OCI_FUNCTION_OUTPUT_SHAPE, it points to the OCI Function object.
	ReferenceObject *string `mandatory:"true" json:"referenceObject"`

	// The reference type for this reference.  Set to null for a direct reference, for indirect references set to a type of association such as "BOUND_ENTITY_SHAPE".   Current known reference type values are "BOUND_ENTITY_SHAPE", "BOUND_ENTITY_SHAPE_FIELD", "OCI_FUNCTION_INPUT_SHAPE", "OCI_FUNCTION_OUTPUT_SHAPE"
	ReferenceType ScopeReferenceReferenceTypeEnum `mandatory:"false" json:"referenceType,omitempty"`

	// The referenced object name for this reference.  Set to the field name if the referenceType is BOUND_ENTITY_SHAPE_FIELD, else set to null.
	RefObjectName *string `mandatory:"false" json:"refObjectName"`
}

func (m ScopeReference) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ScopeReference) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingScopeReferenceReferenceTypeEnum(string(m.ReferenceType)); !ok && m.ReferenceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ReferenceType: %s. Supported values are: %s.", m.ReferenceType, strings.Join(GetScopeReferenceReferenceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// ScopeReferenceReferenceTypeEnum Enum with underlying type: string
type ScopeReferenceReferenceTypeEnum string

// Set of constants representing the allowable values for ScopeReferenceReferenceTypeEnum
const (
	ScopeReferenceReferenceTypeDirectRef              ScopeReferenceReferenceTypeEnum = "DIRECT_REF"
	ScopeReferenceReferenceTypeBoundEntityShape       ScopeReferenceReferenceTypeEnum = "BOUND_ENTITY_SHAPE"
	ScopeReferenceReferenceTypeBoundEntityShapeField  ScopeReferenceReferenceTypeEnum = "BOUND_ENTITY_SHAPE_FIELD"
	ScopeReferenceReferenceTypeOciFunctionInputShape  ScopeReferenceReferenceTypeEnum = "OCI_FUNCTION_INPUT_SHAPE"
	ScopeReferenceReferenceTypeOciFunctionOutputShape ScopeReferenceReferenceTypeEnum = "OCI_FUNCTION_OUTPUT_SHAPE"
)

var mappingScopeReferenceReferenceTypeEnum = map[string]ScopeReferenceReferenceTypeEnum{
	"DIRECT_REF":                ScopeReferenceReferenceTypeDirectRef,
	"BOUND_ENTITY_SHAPE":        ScopeReferenceReferenceTypeBoundEntityShape,
	"BOUND_ENTITY_SHAPE_FIELD":  ScopeReferenceReferenceTypeBoundEntityShapeField,
	"OCI_FUNCTION_INPUT_SHAPE":  ScopeReferenceReferenceTypeOciFunctionInputShape,
	"OCI_FUNCTION_OUTPUT_SHAPE": ScopeReferenceReferenceTypeOciFunctionOutputShape,
}

var mappingScopeReferenceReferenceTypeEnumLowerCase = map[string]ScopeReferenceReferenceTypeEnum{
	"direct_ref":                ScopeReferenceReferenceTypeDirectRef,
	"bound_entity_shape":        ScopeReferenceReferenceTypeBoundEntityShape,
	"bound_entity_shape_field":  ScopeReferenceReferenceTypeBoundEntityShapeField,
	"oci_function_input_shape":  ScopeReferenceReferenceTypeOciFunctionInputShape,
	"oci_function_output_shape": ScopeReferenceReferenceTypeOciFunctionOutputShape,
}

// GetScopeReferenceReferenceTypeEnumValues Enumerates the set of values for ScopeReferenceReferenceTypeEnum
func GetScopeReferenceReferenceTypeEnumValues() []ScopeReferenceReferenceTypeEnum {
	values := make([]ScopeReferenceReferenceTypeEnum, 0)
	for _, v := range mappingScopeReferenceReferenceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetScopeReferenceReferenceTypeEnumStringValues Enumerates the set of values in String for ScopeReferenceReferenceTypeEnum
func GetScopeReferenceReferenceTypeEnumStringValues() []string {
	return []string{
		"DIRECT_REF",
		"BOUND_ENTITY_SHAPE",
		"BOUND_ENTITY_SHAPE_FIELD",
		"OCI_FUNCTION_INPUT_SHAPE",
		"OCI_FUNCTION_OUTPUT_SHAPE",
	}
}

// GetMappingScopeReferenceReferenceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingScopeReferenceReferenceTypeEnum(val string) (ScopeReferenceReferenceTypeEnum, bool) {
	enum, ok := mappingScopeReferenceReferenceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
