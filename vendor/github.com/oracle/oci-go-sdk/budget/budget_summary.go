// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts.
//

package budget

import (
	"github.com/oracle/oci-go-sdk/common"
)

// BudgetSummary A budget.
type BudgetSummary struct {

	// The OCID of the budget
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the budget.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The amount of the budget expressed in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget.
	ResetPeriod ResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// The current state of the budget.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Total number of alert rules in the budget
	AlertRuleCount *int `mandatory:"true" json:"alertRuleCount"`

	// Time budget was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time budget was updated
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// This is DEPRECATED. For backwards compatability, the property will be populated when
	// targetType is "COMPARTMENT" AND targets contains EXACT ONE target compartment ocid.
	// For all other scenarios, this property will be left empty.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// The type of target on which the budget is applied.
	TargetType TargetTypeEnum `mandatory:"false" json:"targetType,omitempty"`

	// The list of targets on which the budget is applied.
	//   If targetType is "COMPARTMENT", targets contains list of compartment OCIDs.
	//   If targetType is "TAG", targets contains list of cost tracking tag identifiers in the form of "{tagNamespace}.{tagKey}.{tagValue}".
	Targets []string `mandatory:"false" json:"targets"`

	// Version of the budget. Starts from 1 and increments by 1.
	Version *int `mandatory:"false" json:"version"`

	// The actual spend in currency for the current budget cycle
	ActualSpend *float32 `mandatory:"false" json:"actualSpend"`

	// The forecasted spend in currency by the end of the current budget cycle
	ForecastedSpend *float32 `mandatory:"false" json:"forecastedSpend"`

	// Time budget spend was last computed
	TimeSpendComputed *common.SDKTime `mandatory:"false" json:"timeSpendComputed"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m BudgetSummary) String() string {
	return common.PointerString(m)
}

// BudgetSummaryResetPeriodEnum is an alias to type: ResetPeriodEnum
// Consider using ResetPeriodEnum instead
// Deprecated
type BudgetSummaryResetPeriodEnum = ResetPeriodEnum

// Set of constants representing the allowable values for ResetPeriodEnum
// Deprecated
const (
	BudgetSummaryResetPeriodMonthly ResetPeriodEnum = "MONTHLY"
)

// GetBudgetSummaryResetPeriodEnumValues Enumerates the set of values for ResetPeriodEnum
// Consider using GetResetPeriodEnumValue
// Deprecated
var GetBudgetSummaryResetPeriodEnumValues = GetResetPeriodEnumValues

// BudgetSummaryTargetTypeEnum is an alias to type: TargetTypeEnum
// Consider using TargetTypeEnum instead
// Deprecated
type BudgetSummaryTargetTypeEnum = TargetTypeEnum

// Set of constants representing the allowable values for TargetTypeEnum
// Deprecated
const (
	BudgetSummaryTargetTypeCompartment TargetTypeEnum = "COMPARTMENT"
	BudgetSummaryTargetTypeTag         TargetTypeEnum = "TAG"
)

// GetBudgetSummaryTargetTypeEnumValues Enumerates the set of values for TargetTypeEnum
// Consider using GetTargetTypeEnumValue
// Deprecated
var GetBudgetSummaryTargetTypeEnumValues = GetTargetTypeEnumValues

// BudgetSummaryLifecycleStateEnum is an alias to type: LifecycleStateEnum
// Consider using LifecycleStateEnum instead
// Deprecated
type BudgetSummaryLifecycleStateEnum = LifecycleStateEnum

// Set of constants representing the allowable values for LifecycleStateEnum
// Deprecated
const (
	BudgetSummaryLifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	BudgetSummaryLifecycleStateInactive LifecycleStateEnum = "INACTIVE"
)

// GetBudgetSummaryLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
// Consider using GetLifecycleStateEnumValue
// Deprecated
var GetBudgetSummaryLifecycleStateEnumValues = GetLifecycleStateEnumValues
