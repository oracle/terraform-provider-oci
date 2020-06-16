// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// APIs for dynamically scaling Compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Overview of the Compute Service (https://docs.cloud.oracle.com/Content/Compute/Concepts/computeoverview.htm).
// **Note:** Autoscaling is not available in US Government Cloud tenancies. For more information, see
// Oracle Cloud Infrastructure US Government Cloud (https://docs.cloud.oracle.com/Content/General/Concepts/govoverview.htm).
//

package autoscaling

import (
	"github.com/oracle/oci-go-sdk/common"
)

// Threshold The representation of Threshold
type Threshold struct {

	// The comparison operator to use. Options are greater than (`GT`), greater than or equal to
	// (`GTE`), less than (`LT`), and less than or equal to (`LTE`).
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
