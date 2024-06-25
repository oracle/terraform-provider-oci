// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Resource Scheduler API
//
// Use the Resource scheduler API to manage schedules, to perform actions on a collection of resources.
//

package resourcescheduler

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TimeCreatedResourceFilter This is a resource filter for filtering resources based on their creation time.
type TimeCreatedResourceFilter struct {

	// This is the date and time as the value of the filter.
	Value *string `mandatory:"false" json:"value"`

	// This is the condition for the filter in comparison to its creation time.
	Condition TimeCreatedResourceFilterConditionEnum `mandatory:"false" json:"condition,omitempty"`
}

func (m TimeCreatedResourceFilter) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TimeCreatedResourceFilter) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingTimeCreatedResourceFilterConditionEnum(string(m.Condition)); !ok && m.Condition != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Condition: %s. Supported values are: %s.", m.Condition, strings.Join(GetTimeCreatedResourceFilterConditionEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TimeCreatedResourceFilter) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTimeCreatedResourceFilter TimeCreatedResourceFilter
	s := struct {
		DiscriminatorParam string `json:"attribute"`
		MarshalTypeTimeCreatedResourceFilter
	}{
		"TIME_CREATED",
		(MarshalTypeTimeCreatedResourceFilter)(m),
	}

	return json.Marshal(&s)
}

// TimeCreatedResourceFilterConditionEnum Enum with underlying type: string
type TimeCreatedResourceFilterConditionEnum string

// Set of constants representing the allowable values for TimeCreatedResourceFilterConditionEnum
const (
	TimeCreatedResourceFilterConditionEqual  TimeCreatedResourceFilterConditionEnum = "EQUAL"
	TimeCreatedResourceFilterConditionBefore TimeCreatedResourceFilterConditionEnum = "BEFORE"
	TimeCreatedResourceFilterConditionAfter  TimeCreatedResourceFilterConditionEnum = "AFTER"
)

var mappingTimeCreatedResourceFilterConditionEnum = map[string]TimeCreatedResourceFilterConditionEnum{
	"EQUAL":  TimeCreatedResourceFilterConditionEqual,
	"BEFORE": TimeCreatedResourceFilterConditionBefore,
	"AFTER":  TimeCreatedResourceFilterConditionAfter,
}

var mappingTimeCreatedResourceFilterConditionEnumLowerCase = map[string]TimeCreatedResourceFilterConditionEnum{
	"equal":  TimeCreatedResourceFilterConditionEqual,
	"before": TimeCreatedResourceFilterConditionBefore,
	"after":  TimeCreatedResourceFilterConditionAfter,
}

// GetTimeCreatedResourceFilterConditionEnumValues Enumerates the set of values for TimeCreatedResourceFilterConditionEnum
func GetTimeCreatedResourceFilterConditionEnumValues() []TimeCreatedResourceFilterConditionEnum {
	values := make([]TimeCreatedResourceFilterConditionEnum, 0)
	for _, v := range mappingTimeCreatedResourceFilterConditionEnum {
		values = append(values, v)
	}
	return values
}

// GetTimeCreatedResourceFilterConditionEnumStringValues Enumerates the set of values in String for TimeCreatedResourceFilterConditionEnum
func GetTimeCreatedResourceFilterConditionEnumStringValues() []string {
	return []string{
		"EQUAL",
		"BEFORE",
		"AFTER",
	}
}

// GetMappingTimeCreatedResourceFilterConditionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTimeCreatedResourceFilterConditionEnum(val string) (TimeCreatedResourceFilterConditionEnum, bool) {
	enum, ok := mappingTimeCreatedResourceFilterConditionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
