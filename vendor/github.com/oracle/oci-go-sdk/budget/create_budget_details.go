// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// CreateBudgetDetails The create budget details.
type CreateBudgetDetails struct {

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID of the compartment on which budget is applied
	TargetCompartmentId *string `mandatory:"true" json:"targetCompartmentId"`

	// The amount of the budget expressed as a whole number in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget.
	ResetPeriod ResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// The displayName of the budget.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateBudgetDetails) String() string {
	return common.PointerString(m)
}

// CreateBudgetDetailsResetPeriodEnum is an alias to type: ResetPeriodEnum
// Consider using ResetPeriodEnum instead
// Deprecated
type CreateBudgetDetailsResetPeriodEnum = ResetPeriodEnum

// Set of constants representing the allowable values for ResetPeriodEnum
// Deprecated
const (
	CreateBudgetDetailsResetPeriodMonthly ResetPeriodEnum = "MONTHLY"
)

// GetCreateBudgetDetailsResetPeriodEnumValues Enumerates the set of values for ResetPeriodEnum
// Consider using GetResetPeriodEnumValue
// Deprecated
var GetCreateBudgetDetailsResetPeriodEnumValues = GetResetPeriodEnumValues
