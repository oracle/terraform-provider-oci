// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PathMatchType The type of matching to apply to incoming URIs.
type PathMatchType struct {

	// Specifies how the load balancing service compares a PathRoute
	// object's `path` string against the incoming URI.
	// *  **EXACT_MATCH** - Looks for a `path` string that exactly matches the incoming URI path.
	// *  **FORCE_LONGEST_PREFIX_MATCH** - Looks for the `path` string with the best, longest match of the beginning
	//    portion of the incoming URI path.
	// *  **PREFIX_MATCH** - Looks for a `path` string that matches the beginning portion of the incoming URI path.
	// *  **SUFFIX_MATCH** - Looks for a `path` string that matches the ending portion of the incoming URI path.
	// For a full description of how the system handles `matchType` in a path route set containing multiple rules, see
	// Managing Request Routing (https://docs.cloud.oracle.com/Content/Balance/Tasks/managingrequest.htm).
	MatchType PathMatchTypeMatchTypeEnum `mandatory:"true" json:"matchType"`
}

func (m PathMatchType) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PathMatchType) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPathMatchTypeMatchTypeEnum(string(m.MatchType)); !ok && m.MatchType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MatchType: %s. Supported values are: %s.", m.MatchType, strings.Join(GetPathMatchTypeMatchTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PathMatchTypeMatchTypeEnum Enum with underlying type: string
type PathMatchTypeMatchTypeEnum string

// Set of constants representing the allowable values for PathMatchTypeMatchTypeEnum
const (
	PathMatchTypeMatchTypeExactMatch              PathMatchTypeMatchTypeEnum = "EXACT_MATCH"
	PathMatchTypeMatchTypeForceLongestPrefixMatch PathMatchTypeMatchTypeEnum = "FORCE_LONGEST_PREFIX_MATCH"
	PathMatchTypeMatchTypePrefixMatch             PathMatchTypeMatchTypeEnum = "PREFIX_MATCH"
	PathMatchTypeMatchTypeSuffixMatch             PathMatchTypeMatchTypeEnum = "SUFFIX_MATCH"
)

var mappingPathMatchTypeMatchTypeEnum = map[string]PathMatchTypeMatchTypeEnum{
	"EXACT_MATCH":                PathMatchTypeMatchTypeExactMatch,
	"FORCE_LONGEST_PREFIX_MATCH": PathMatchTypeMatchTypeForceLongestPrefixMatch,
	"PREFIX_MATCH":               PathMatchTypeMatchTypePrefixMatch,
	"SUFFIX_MATCH":               PathMatchTypeMatchTypeSuffixMatch,
}

// GetPathMatchTypeMatchTypeEnumValues Enumerates the set of values for PathMatchTypeMatchTypeEnum
func GetPathMatchTypeMatchTypeEnumValues() []PathMatchTypeMatchTypeEnum {
	values := make([]PathMatchTypeMatchTypeEnum, 0)
	for _, v := range mappingPathMatchTypeMatchTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPathMatchTypeMatchTypeEnumStringValues Enumerates the set of values in String for PathMatchTypeMatchTypeEnum
func GetPathMatchTypeMatchTypeEnumStringValues() []string {
	return []string{
		"EXACT_MATCH",
		"FORCE_LONGEST_PREFIX_MATCH",
		"PREFIX_MATCH",
		"SUFFIX_MATCH",
	}
}

// GetMappingPathMatchTypeMatchTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPathMatchTypeMatchTypeEnum(val string) (PathMatchTypeMatchTypeEnum, bool) {
	mappingPathMatchTypeMatchTypeEnumIgnoreCase := make(map[string]PathMatchTypeMatchTypeEnum)
	for k, v := range mappingPathMatchTypeMatchTypeEnum {
		mappingPathMatchTypeMatchTypeEnumIgnoreCase[strings.ToLower(k)] = v
	}

	enum, ok := mappingPathMatchTypeMatchTypeEnumIgnoreCase[strings.ToLower(val)]
	return enum, ok
}
