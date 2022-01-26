// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity Service
//
// API for the Identity Dataplane
//

package identitydataplane

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ContextVariable The representation of ContextVariable
type ContextVariable struct {

	// The name of the variable.
	Name *string `mandatory:"true" json:"name"`

	// The value of the variable.
	Value *string `mandatory:"true" json:"value"`

	// The type of the variable.
	Type ContextVariableTypeEnum `mandatory:"true" json:"type"`
}

func (m ContextVariable) String() string {
	return common.PointerString(m)
}

// ContextVariableTypeEnum Enum with underlying type: string
type ContextVariableTypeEnum string

// Set of constants representing the allowable values for ContextVariableTypeEnum
const (
	ContextVariableTypeString  ContextVariableTypeEnum = "STRING"
	ContextVariableTypeNumber  ContextVariableTypeEnum = "NUMBER"
	ContextVariableTypeEntity  ContextVariableTypeEnum = "ENTITY"
	ContextVariableTypeBoolean ContextVariableTypeEnum = "BOOLEAN"
	ContextVariableTypeList    ContextVariableTypeEnum = "LIST"
)

var mappingContextVariableType = map[string]ContextVariableTypeEnum{
	"STRING":  ContextVariableTypeString,
	"NUMBER":  ContextVariableTypeNumber,
	"ENTITY":  ContextVariableTypeEntity,
	"BOOLEAN": ContextVariableTypeBoolean,
	"LIST":    ContextVariableTypeList,
}

// GetContextVariableTypeEnumValues Enumerates the set of values for ContextVariableTypeEnum
func GetContextVariableTypeEnumValues() []ContextVariableTypeEnum {
	values := make([]ContextVariableTypeEnum, 0)
	for _, v := range mappingContextVariableType {
		values = append(values, v)
	}
	return values
}
