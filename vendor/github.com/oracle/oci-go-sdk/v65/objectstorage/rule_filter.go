// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RuleFilter Defines rule filter.
type RuleFilter struct {

	// An array of prefixes to which the rule will apply.
	// If the array is not provided, the rule applies to all prefixes.
	// If one or more prefixes are specified, only buckets/objects matching any of the specified prefixes will be included.
	// Objects matching any Include filter will be included unless excluded by an Exclude filter.
	// In case of overlap, Exclude filters take precedence over Include filters.
	InclusionPrefixes []string `mandatory:"false" json:"inclusionPrefixes"`

	// An array of prefixes that the rule will exclude.
	// If the array is not provided, no exclusions will be applied, and all objects will be considered.
	// Any object matching any Exclude filter will be excluded, regardless of matching Include filters.
	// Exclude filters take precedence over Include filters.
	ExclusionPrefixes []string `mandatory:"false" json:"exclusionPrefixes"`

	// An array of object types to be included in report.
	ObjectTypes []RuleFilterObjectTypesEnum `mandatory:"false" json:"objectTypes,omitempty"`

	// Objects created within the last createdWithinDays days (inclusive) will be considered.
	// For example, a value of 3 includes objects created within the last three days.
	// During report generation, the calculation is based on the report generation start time, rounded up to the next full day.
	// The object age is calculated relative to this adjusted start time.
	// This parameter is optional. If not provided, the value is assumed to be -1,
	// meaning no filtering based on object age will be applied, and all objects will be included in the report regardless of their creation date.
	// A value of -1 explicitly indicates that all objects should be included.
	CreatedWithinDays *int `mandatory:"false" json:"createdWithinDays"`
}

func (m RuleFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RuleFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	for _, val := range m.ObjectTypes {
		if _, ok := GetMappingRuleFilterObjectTypesEnum(string(val)); !ok && val != "" {
			errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ObjectTypes: %s. Supported values are: %s.", val, strings.Join(GetRuleFilterObjectTypesEnumStringValues(), ",")))
		}
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// RuleFilterObjectTypesEnum Enum with underlying type: string
type RuleFilterObjectTypesEnum string

// Set of constants representing the allowable values for RuleFilterObjectTypesEnum
const (
	RuleFilterObjectTypesObject RuleFilterObjectTypesEnum = "OBJECT"
)

var mappingRuleFilterObjectTypesEnum = map[string]RuleFilterObjectTypesEnum{
	"OBJECT": RuleFilterObjectTypesObject,
}

var mappingRuleFilterObjectTypesEnumLowerCase = map[string]RuleFilterObjectTypesEnum{
	"object": RuleFilterObjectTypesObject,
}

// GetRuleFilterObjectTypesEnumValues Enumerates the set of values for RuleFilterObjectTypesEnum
func GetRuleFilterObjectTypesEnumValues() []RuleFilterObjectTypesEnum {
	values := make([]RuleFilterObjectTypesEnum, 0)
	for _, v := range mappingRuleFilterObjectTypesEnum {
		values = append(values, v)
	}
	return values
}

// GetRuleFilterObjectTypesEnumStringValues Enumerates the set of values in String for RuleFilterObjectTypesEnum
func GetRuleFilterObjectTypesEnumStringValues() []string {
	return []string{
		"OBJECT",
	}
}

// GetMappingRuleFilterObjectTypesEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingRuleFilterObjectTypesEnum(val string) (RuleFilterObjectTypesEnum, bool) {
	enum, ok := mappingRuleFilterObjectTypesEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
