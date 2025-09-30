// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Fleet Application Management Service API
//
// Fleet Application Management provides a centralized platform to help you automate resource management tasks, validate patch compliance, and enhance operational efficiency across an enterprise.
//

package fleetappsmanagement

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StackSource Object representing the source information for the stack, indicating origin type and a reference string.
type StackSource struct {

	// The source type of the stack (e.g. MARKETPLACE, QUICKSTART, or WEB).
	Type StackSourceTypeEnum `mandatory:"false" json:"type,omitempty"`

	// Reference string providing a pointer or identifier for the source.
	Reference *string `mandatory:"false" json:"reference"`
}

func (m StackSource) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StackSource) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStackSourceTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetStackSourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StackSourceTypeEnum Enum with underlying type: string
type StackSourceTypeEnum string

// Set of constants representing the allowable values for StackSourceTypeEnum
const (
	StackSourceTypeMarketplace StackSourceTypeEnum = "MARKETPLACE"
	StackSourceTypeQuickstart  StackSourceTypeEnum = "QUICKSTART"
	StackSourceTypeWeb         StackSourceTypeEnum = "WEB"
)

var mappingStackSourceTypeEnum = map[string]StackSourceTypeEnum{
	"MARKETPLACE": StackSourceTypeMarketplace,
	"QUICKSTART":  StackSourceTypeQuickstart,
	"WEB":         StackSourceTypeWeb,
}

var mappingStackSourceTypeEnumLowerCase = map[string]StackSourceTypeEnum{
	"marketplace": StackSourceTypeMarketplace,
	"quickstart":  StackSourceTypeQuickstart,
	"web":         StackSourceTypeWeb,
}

// GetStackSourceTypeEnumValues Enumerates the set of values for StackSourceTypeEnum
func GetStackSourceTypeEnumValues() []StackSourceTypeEnum {
	values := make([]StackSourceTypeEnum, 0)
	for _, v := range mappingStackSourceTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStackSourceTypeEnumStringValues Enumerates the set of values in String for StackSourceTypeEnum
func GetStackSourceTypeEnumStringValues() []string {
	return []string{
		"MARKETPLACE",
		"QUICKSTART",
		"WEB",
	}
}

// GetMappingStackSourceTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStackSourceTypeEnum(val string) (StackSourceTypeEnum, bool) {
	enum, ok := mappingStackSourceTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
