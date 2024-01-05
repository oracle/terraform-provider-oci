// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Identity and Access Management Data Plane API
//
// APIs for managing identity data plane services. For example, use this API to create a scoped-access security token. To manage identity domains (for example, creating or deleting an identity domain) or to manage resources (for example, users and groups) within the default identity domain, see IAM API (https://docs.oracle.com/iaas/api/#/en/identity/).
//

package identitydataplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ContextVariable) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingContextVariableTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetContextVariableTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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

var mappingContextVariableTypeEnum = map[string]ContextVariableTypeEnum{
	"STRING":  ContextVariableTypeString,
	"NUMBER":  ContextVariableTypeNumber,
	"ENTITY":  ContextVariableTypeEntity,
	"BOOLEAN": ContextVariableTypeBoolean,
	"LIST":    ContextVariableTypeList,
}

var mappingContextVariableTypeEnumLowerCase = map[string]ContextVariableTypeEnum{
	"string":  ContextVariableTypeString,
	"number":  ContextVariableTypeNumber,
	"entity":  ContextVariableTypeEntity,
	"boolean": ContextVariableTypeBoolean,
	"list":    ContextVariableTypeList,
}

// GetContextVariableTypeEnumValues Enumerates the set of values for ContextVariableTypeEnum
func GetContextVariableTypeEnumValues() []ContextVariableTypeEnum {
	values := make([]ContextVariableTypeEnum, 0)
	for _, v := range mappingContextVariableTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetContextVariableTypeEnumStringValues Enumerates the set of values in String for ContextVariableTypeEnum
func GetContextVariableTypeEnumStringValues() []string {
	return []string{
		"STRING",
		"NUMBER",
		"ENTITY",
		"BOOLEAN",
		"LIST",
	}
}

// GetMappingContextVariableTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingContextVariableTypeEnum(val string) (ContextVariableTypeEnum, bool) {
	enum, ok := mappingContextVariableTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
