// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts. For more information, see Budgets Overview (https://docs.oracle.com/iaas/Content/Billing/Concepts/budgetsoverview.htm).
//

package budget

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CostAlertSubscriptionMap The mapping of cost monitor to alert subscription along with thresholds.
type CostAlertSubscriptionMap struct {

	// The filter operator. Example: 'AND', 'OR'.
	Operator CostAlertSubscriptionMapOperatorEnum `mandatory:"false" json:"operator,omitempty"`

	// The absolute threshold value.
	ThresholdAbsoluteValue *int `mandatory:"false" json:"thresholdAbsoluteValue"`

	// The relative percentage threshold value.
	ThresholdRelativePercent *int `mandatory:"false" json:"thresholdRelativePercent"`

	// The costAlertSubscription ocid which the cost monitor alert maps to.
	CostAlertSubscriptionId *string `mandatory:"false" json:"costAlertSubscriptionId"`
}

func (m CostAlertSubscriptionMap) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CostAlertSubscriptionMap) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCostAlertSubscriptionMapOperatorEnum(string(m.Operator)); !ok && m.Operator != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Operator: %s. Supported values are: %s.", m.Operator, strings.Join(GetCostAlertSubscriptionMapOperatorEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CostAlertSubscriptionMapOperatorEnum Enum with underlying type: string
type CostAlertSubscriptionMapOperatorEnum string

// Set of constants representing the allowable values for CostAlertSubscriptionMapOperatorEnum
const (
	CostAlertSubscriptionMapOperatorAnd CostAlertSubscriptionMapOperatorEnum = "AND"
	CostAlertSubscriptionMapOperatorOr  CostAlertSubscriptionMapOperatorEnum = "OR"
)

var mappingCostAlertSubscriptionMapOperatorEnum = map[string]CostAlertSubscriptionMapOperatorEnum{
	"AND": CostAlertSubscriptionMapOperatorAnd,
	"OR":  CostAlertSubscriptionMapOperatorOr,
}

var mappingCostAlertSubscriptionMapOperatorEnumLowerCase = map[string]CostAlertSubscriptionMapOperatorEnum{
	"and": CostAlertSubscriptionMapOperatorAnd,
	"or":  CostAlertSubscriptionMapOperatorOr,
}

// GetCostAlertSubscriptionMapOperatorEnumValues Enumerates the set of values for CostAlertSubscriptionMapOperatorEnum
func GetCostAlertSubscriptionMapOperatorEnumValues() []CostAlertSubscriptionMapOperatorEnum {
	values := make([]CostAlertSubscriptionMapOperatorEnum, 0)
	for _, v := range mappingCostAlertSubscriptionMapOperatorEnum {
		values = append(values, v)
	}
	return values
}

// GetCostAlertSubscriptionMapOperatorEnumStringValues Enumerates the set of values in String for CostAlertSubscriptionMapOperatorEnum
func GetCostAlertSubscriptionMapOperatorEnumStringValues() []string {
	return []string{
		"AND",
		"OR",
	}
}

// GetMappingCostAlertSubscriptionMapOperatorEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCostAlertSubscriptionMapOperatorEnum(val string) (CostAlertSubscriptionMapOperatorEnum, bool) {
	enum, ok := mappingCostAlertSubscriptionMapOperatorEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
