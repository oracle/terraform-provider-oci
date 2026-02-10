// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Migration API
//
// Use the Oracle Cloud Infrastructure Database Migration APIs to perform database migration operations.
//

package databasemigration

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UserDefinedProperty User Defined Property
type UserDefinedProperty struct {

	// The property name.
	Name *string `mandatory:"true" json:"name"`

	// A user-friendly name. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the user defined property.
	Type UserDefinedPropertyTypeEnum `mandatory:"true" json:"type"`

	// A user-friendly description. Does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	Description *string `mandatory:"false" json:"description"`

	// True if the property is required, false otherwise
	IsRequired *bool `mandatory:"false" json:"isRequired"`

	// Minimum length of the text
	MinLength *int `mandatory:"false" json:"minLength"`

	// Maximum length of the text
	MaxLength *int `mandatory:"false" json:"maxLength"`

	// The default value of the property.
	DefaultValue *string `mandatory:"false" json:"defaultValue"`

	// The value of the property.
	Value *string `mandatory:"false" json:"value"`

	// User defined property options.
	Options []UserDefinedPropertyOption `mandatory:"false" json:"options"`
}

func (m UserDefinedProperty) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UserDefinedProperty) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUserDefinedPropertyTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetUserDefinedPropertyTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UserDefinedPropertyTypeEnum Enum with underlying type: string
type UserDefinedPropertyTypeEnum string

// Set of constants representing the allowable values for UserDefinedPropertyTypeEnum
const (
	UserDefinedPropertyTypeText     UserDefinedPropertyTypeEnum = "TEXT"
	UserDefinedPropertyTypePassword UserDefinedPropertyTypeEnum = "PASSWORD"
	UserDefinedPropertyTypeRadio    UserDefinedPropertyTypeEnum = "RADIO"
)

var mappingUserDefinedPropertyTypeEnum = map[string]UserDefinedPropertyTypeEnum{
	"TEXT":     UserDefinedPropertyTypeText,
	"PASSWORD": UserDefinedPropertyTypePassword,
	"RADIO":    UserDefinedPropertyTypeRadio,
}

var mappingUserDefinedPropertyTypeEnumLowerCase = map[string]UserDefinedPropertyTypeEnum{
	"text":     UserDefinedPropertyTypeText,
	"password": UserDefinedPropertyTypePassword,
	"radio":    UserDefinedPropertyTypeRadio,
}

// GetUserDefinedPropertyTypeEnumValues Enumerates the set of values for UserDefinedPropertyTypeEnum
func GetUserDefinedPropertyTypeEnumValues() []UserDefinedPropertyTypeEnum {
	values := make([]UserDefinedPropertyTypeEnum, 0)
	for _, v := range mappingUserDefinedPropertyTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUserDefinedPropertyTypeEnumStringValues Enumerates the set of values in String for UserDefinedPropertyTypeEnum
func GetUserDefinedPropertyTypeEnumStringValues() []string {
	return []string{
		"TEXT",
		"PASSWORD",
		"RADIO",
	}
}

// GetMappingUserDefinedPropertyTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUserDefinedPropertyTypeEnum(val string) (UserDefinedPropertyTypeEnum, bool) {
	enum, ok := mappingUserDefinedPropertyTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
