// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// WindowTypeDescription A key-value pair where the key will hold the window type and value will be a list of window details from all the active execution windows of that window type.
type WindowTypeDescription struct {

	// The execution window is of PLANNED or UNPLANNED type.
	WindowType WindowTypeDescriptionWindowTypeEnum `mandatory:"false" json:"windowType,omitempty"`

	// A list of window detail messages from all the active execution windows based on the window type.
	Messages []string `mandatory:"false" json:"messages"`
}

func (m WindowTypeDescription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m WindowTypeDescription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingWindowTypeDescriptionWindowTypeEnum(string(m.WindowType)); !ok && m.WindowType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for WindowType: %s. Supported values are: %s.", m.WindowType, strings.Join(GetWindowTypeDescriptionWindowTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// WindowTypeDescriptionWindowTypeEnum Enum with underlying type: string
type WindowTypeDescriptionWindowTypeEnum string

// Set of constants representing the allowable values for WindowTypeDescriptionWindowTypeEnum
const (
	WindowTypeDescriptionWindowTypePlanned   WindowTypeDescriptionWindowTypeEnum = "PLANNED"
	WindowTypeDescriptionWindowTypeUnplanned WindowTypeDescriptionWindowTypeEnum = "UNPLANNED"
)

var mappingWindowTypeDescriptionWindowTypeEnum = map[string]WindowTypeDescriptionWindowTypeEnum{
	"PLANNED":   WindowTypeDescriptionWindowTypePlanned,
	"UNPLANNED": WindowTypeDescriptionWindowTypeUnplanned,
}

var mappingWindowTypeDescriptionWindowTypeEnumLowerCase = map[string]WindowTypeDescriptionWindowTypeEnum{
	"planned":   WindowTypeDescriptionWindowTypePlanned,
	"unplanned": WindowTypeDescriptionWindowTypeUnplanned,
}

// GetWindowTypeDescriptionWindowTypeEnumValues Enumerates the set of values for WindowTypeDescriptionWindowTypeEnum
func GetWindowTypeDescriptionWindowTypeEnumValues() []WindowTypeDescriptionWindowTypeEnum {
	values := make([]WindowTypeDescriptionWindowTypeEnum, 0)
	for _, v := range mappingWindowTypeDescriptionWindowTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetWindowTypeDescriptionWindowTypeEnumStringValues Enumerates the set of values in String for WindowTypeDescriptionWindowTypeEnum
func GetWindowTypeDescriptionWindowTypeEnumStringValues() []string {
	return []string{
		"PLANNED",
		"UNPLANNED",
	}
}

// GetMappingWindowTypeDescriptionWindowTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingWindowTypeDescriptionWindowTypeEnum(val string) (WindowTypeDescriptionWindowTypeEnum, bool) {
	enum, ok := mappingWindowTypeDescriptionWindowTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
