// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Service Mesh API
//
// Use the Service Mesh API to manage mesh, virtual service, access policy and other mesh related items.
//

package servicemesh

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// StringValueMatch Match criteria for an attribute.
type StringValueMatch struct {

	// Match type for the value.
	MatchType StringValueMatchMatchTypeEnum `mandatory:"false" json:"matchType,omitempty"`

	// Params value in request to be matched with value.
	Value *string `mandatory:"false" json:"value"`
}

func (m StringValueMatch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StringValueMatch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStringValueMatchMatchTypeEnum(string(m.MatchType)); !ok && m.MatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchType: %s. Supported values are: %s.", m.MatchType, strings.Join(GetStringValueMatchMatchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StringValueMatchMatchTypeEnum Enum with underlying type: string
type StringValueMatchMatchTypeEnum string

// Set of constants representing the allowable values for StringValueMatchMatchTypeEnum
const (
	StringValueMatchMatchTypePrefix    StringValueMatchMatchTypeEnum = "PREFIX"
	StringValueMatchMatchTypeRegex     StringValueMatchMatchTypeEnum = "REGEX"
	StringValueMatchMatchTypeExact     StringValueMatchMatchTypeEnum = "EXACT"
	StringValueMatchMatchTypeExists    StringValueMatchMatchTypeEnum = "EXISTS"
	StringValueMatchMatchTypeNotExists StringValueMatchMatchTypeEnum = "NOT_EXISTS"
)

var mappingStringValueMatchMatchTypeEnum = map[string]StringValueMatchMatchTypeEnum{
	"PREFIX":     StringValueMatchMatchTypePrefix,
	"REGEX":      StringValueMatchMatchTypeRegex,
	"EXACT":      StringValueMatchMatchTypeExact,
	"EXISTS":     StringValueMatchMatchTypeExists,
	"NOT_EXISTS": StringValueMatchMatchTypeNotExists,
}

var mappingStringValueMatchMatchTypeEnumLowerCase = map[string]StringValueMatchMatchTypeEnum{
	"prefix":     StringValueMatchMatchTypePrefix,
	"regex":      StringValueMatchMatchTypeRegex,
	"exact":      StringValueMatchMatchTypeExact,
	"exists":     StringValueMatchMatchTypeExists,
	"not_exists": StringValueMatchMatchTypeNotExists,
}

// GetStringValueMatchMatchTypeEnumValues Enumerates the set of values for StringValueMatchMatchTypeEnum
func GetStringValueMatchMatchTypeEnumValues() []StringValueMatchMatchTypeEnum {
	values := make([]StringValueMatchMatchTypeEnum, 0)
	for _, v := range mappingStringValueMatchMatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStringValueMatchMatchTypeEnumStringValues Enumerates the set of values in String for StringValueMatchMatchTypeEnum
func GetStringValueMatchMatchTypeEnumStringValues() []string {
	return []string{
		"PREFIX",
		"REGEX",
		"EXACT",
		"EXISTS",
		"NOT_EXISTS",
	}
}

// GetMappingStringValueMatchMatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStringValueMatchMatchTypeEnum(val string) (StringValueMatchMatchTypeEnum, bool) {
	enum, ok := mappingStringValueMatchMatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
