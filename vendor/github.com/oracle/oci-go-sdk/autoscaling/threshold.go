// Copyright (c) 2016, 2018, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Auto Scaling API
//
// Auto Scaling API spec
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Threshold The representation of Threshold
type Threshold struct {

	// Support for the following operators
	// GT  - Greater than
	// GTE - Greater than equal to
	// LT  - Less than
	// LTE - Less than equal to
	Operator ThresholdOperatorEnum `mandatory:"true" json:"operator"`

	Value *int `mandatory:"true" json:"value"`
}

func (m Threshold) String() string {
	return common.PointerString(m)
}

// ThresholdOperatorEnum Enum with underlying type: string
type ThresholdOperatorEnum string

// Set of constants representing the allowable values for ThresholdOperatorEnum
const (
	ThresholdOperatorGt  ThresholdOperatorEnum = "GT"
	ThresholdOperatorGte ThresholdOperatorEnum = "GTE"
	ThresholdOperatorLt  ThresholdOperatorEnum = "LT"
	ThresholdOperatorLte ThresholdOperatorEnum = "LTE"
)

var mappingThresholdOperator = map[string]ThresholdOperatorEnum{
	"GT":  ThresholdOperatorGt,
	"GTE": ThresholdOperatorGte,
	"LT":  ThresholdOperatorLt,
	"LTE": ThresholdOperatorLte,
}

// GetThresholdOperatorEnumValues Enumerates the set of values for ThresholdOperatorEnum
func GetThresholdOperatorEnumValues() []ThresholdOperatorEnum {
	values := make([]ThresholdOperatorEnum, 0)
	for _, v := range mappingThresholdOperator {
		values = append(values, v)
	}
	return values
}
