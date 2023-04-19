// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
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

// StringMatch Match criteria for an attribute for specified Key.
type StringMatch struct {

	// name of the parameter
	Key *string `mandatory:"true" json:"key"`

	// Match type for the value.
	MatchType StringMatchMatchTypeEnum `mandatory:"false" json:"matchType,omitempty"`

	// Params value in request to be matched with value.
	Value *string `mandatory:"false" json:"value"`
}

func (m StringMatch) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m StringMatch) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingStringMatchMatchTypeEnum(string(m.MatchType)); !ok && m.MatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchType: %s. Supported values are: %s.", m.MatchType, strings.Join(GetStringMatchMatchTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// StringMatchMatchTypeEnum Enum with underlying type: string
type StringMatchMatchTypeEnum string

// Set of constants representing the allowable values for StringMatchMatchTypeEnum
const (
	StringMatchMatchTypePrefix    StringMatchMatchTypeEnum = "PREFIX"
	StringMatchMatchTypeRegex     StringMatchMatchTypeEnum = "REGEX"
	StringMatchMatchTypeExact     StringMatchMatchTypeEnum = "EXACT"
	StringMatchMatchTypeExists    StringMatchMatchTypeEnum = "EXISTS"
	StringMatchMatchTypeNotExists StringMatchMatchTypeEnum = "NOT_EXISTS"
)

var mappingStringMatchMatchTypeEnum = map[string]StringMatchMatchTypeEnum{
	"PREFIX":     StringMatchMatchTypePrefix,
	"REGEX":      StringMatchMatchTypeRegex,
	"EXACT":      StringMatchMatchTypeExact,
	"EXISTS":     StringMatchMatchTypeExists,
	"NOT_EXISTS": StringMatchMatchTypeNotExists,
}

var mappingStringMatchMatchTypeEnumLowerCase = map[string]StringMatchMatchTypeEnum{
	"prefix":     StringMatchMatchTypePrefix,
	"regex":      StringMatchMatchTypeRegex,
	"exact":      StringMatchMatchTypeExact,
	"exists":     StringMatchMatchTypeExists,
	"not_exists": StringMatchMatchTypeNotExists,
}

// GetStringMatchMatchTypeEnumValues Enumerates the set of values for StringMatchMatchTypeEnum
func GetStringMatchMatchTypeEnumValues() []StringMatchMatchTypeEnum {
	values := make([]StringMatchMatchTypeEnum, 0)
	for _, v := range mappingStringMatchMatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetStringMatchMatchTypeEnumStringValues Enumerates the set of values in String for StringMatchMatchTypeEnum
func GetStringMatchMatchTypeEnumStringValues() []string {
	return []string{
		"PREFIX",
		"REGEX",
		"EXACT",
		"EXISTS",
		"NOT_EXISTS",
	}
}

// GetMappingStringMatchMatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingStringMatchMatchTypeEnum(val string) (StringMatchMatchTypeEnum, bool) {
	enum, ok := mappingStringMatchMatchTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
