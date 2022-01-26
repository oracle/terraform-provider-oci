// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// LabelPriority The label priority.
type LabelPriority struct {

	// The label priority. Default value is NONE.
	Priority LabelPriorityPriorityEnum `mandatory:"false" json:"priority,omitempty"`
}

func (m LabelPriority) String() string {
	return common.PointerString(m)
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

var mappingLabelPriorityPriority = map[string]LabelPriorityPriorityEnum{
	"NONE":   LabelPriorityPriorityNone,
	"LOW":    LabelPriorityPriorityLow,
	"MEDIUM": LabelPriorityPriorityMedium,
	"HIGH":   LabelPriorityPriorityHigh,
}

// GetLabelPriorityPriorityEnumValues Enumerates the set of values for LabelPriorityPriorityEnum
func GetLabelPriorityPriorityEnumValues() []LabelPriorityPriorityEnum {
	values := make([]LabelPriorityPriorityEnum, 0)
	for _, v := range mappingLabelPriorityPriority {
		values = append(values, v)
	}
	return values
}
