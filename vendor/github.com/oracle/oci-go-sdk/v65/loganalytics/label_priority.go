// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LabelPriority The label priority.
type LabelPriority struct {

	// The label priority. Default value is NONE.
	Priority LabelPriorityPriorityEnum `mandatory:"false" json:"priority,omitempty"`
}

func (m LabelPriority) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LabelPriority) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLabelPriorityPriorityEnum(string(m.Priority)); !ok && m.Priority != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Priority: %s. Supported values are: %s.", m.Priority, strings.Join(GetLabelPriorityPriorityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// LabelPriorityPriorityEnum Enum with underlying type: string
type LabelPriorityPriorityEnum string

// Set of constants representing the allowable values for LabelPriorityPriorityEnum
const (
	LabelPriorityPriorityNone   LabelPriorityPriorityEnum = "NONE"
	LabelPriorityPriorityLow    LabelPriorityPriorityEnum = "LOW"
	LabelPriorityPriorityMedium LabelPriorityPriorityEnum = "MEDIUM"
	LabelPriorityPriorityHigh   LabelPriorityPriorityEnum = "HIGH"
)

var mappingLabelPriorityPriorityEnum = map[string]LabelPriorityPriorityEnum{
	"NONE":   LabelPriorityPriorityNone,
	"LOW":    LabelPriorityPriorityLow,
	"MEDIUM": LabelPriorityPriorityMedium,
	"HIGH":   LabelPriorityPriorityHigh,
}

var mappingLabelPriorityPriorityEnumLowerCase = map[string]LabelPriorityPriorityEnum{
	"none":   LabelPriorityPriorityNone,
	"low":    LabelPriorityPriorityLow,
	"medium": LabelPriorityPriorityMedium,
	"high":   LabelPriorityPriorityHigh,
}

// GetLabelPriorityPriorityEnumValues Enumerates the set of values for LabelPriorityPriorityEnum
func GetLabelPriorityPriorityEnumValues() []LabelPriorityPriorityEnum {
	values := make([]LabelPriorityPriorityEnum, 0)
	for _, v := range mappingLabelPriorityPriorityEnum {
		values = append(values, v)
	}
	return values
}

// GetLabelPriorityPriorityEnumStringValues Enumerates the set of values in String for LabelPriorityPriorityEnum
func GetLabelPriorityPriorityEnumStringValues() []string {
	return []string{
		"NONE",
		"LOW",
		"MEDIUM",
		"HIGH",
	}
}

// GetMappingLabelPriorityPriorityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingLabelPriorityPriorityEnum(val string) (LabelPriorityPriorityEnum, bool) {
	enum, ok := mappingLabelPriorityPriorityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
