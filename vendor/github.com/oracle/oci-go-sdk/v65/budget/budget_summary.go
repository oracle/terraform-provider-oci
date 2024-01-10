// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Budgets API
//
// Use the Budgets API to manage budgets and budget alerts. For more information, see Budgets Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/budgetsoverview.htm).
//

package budget

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// BudgetSummary A budget.
type BudgetSummary struct {

	// The OCID of the budget.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The display name of the budget. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The amount of the budget, expressed in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget.
	ResetPeriod ResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// The current state of the budget.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The total number of alert rules in the budget.
	AlertRuleCount *int `mandatory:"true" json:"alertRuleCount"`

	// The time the budget was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the budget was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// This is DEPRECATED. For backwards compatability, the property is populated when
	// the targetType is "COMPARTMENT", and the targets contain the specific target compartment OCID.
	// For all other scenarios, this property is left empty.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// The number of days offset from the first day of the month, at which the budget processing period starts. In months that have fewer days than this value, processing will begin on the last day of that month. For example, for a value of 12, processing starts every month on the 12th at midnight.
	BudgetProcessingPeriodStartOffset *int `mandatory:"false" json:"budgetProcessingPeriodStartOffset"`

	// The type of the budget processing period. Valid values are INVOICE, MONTH, and SINGLE_USE.
	ProcessingPeriodType ProcessingPeriodTypeEnum `mandatory:"false" json:"processingPeriodType,omitempty"`

	// The date when the one-time budget begins. For example, `2023-03-23`. The date-time format conforms to RFC 3339, and will be truncated to the starting point of the date provided after being converted to UTC time.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// The time when the one-time budget concludes. For example, - `2023-03-23`. The date-time format conforms to RFC 3339, and will be truncated to the starting point of the date provided after being converted to UTC time.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`

	// The type of target on which the budget is applied.
	TargetType TargetTypeEnum `mandatory:"false" json:"targetType,omitempty"`

	// The list of targets on which the budget is applied.
	//   If the targetType is "COMPARTMENT", the targets contain the list of compartment OCIDs.
	//   If the targetType is "TAG", the targets contain the list of cost tracking tag identifiers in the form of "{tagNamespace}.{tagKey}.{tagValue}".
	Targets []string `mandatory:"false" json:"targets"`

	// The version of the budget. Starts from 1 and increments by 1.
	Version *int `mandatory:"false" json:"version"`

	// The actual spend in currency for the current budget cycle.
	ActualSpend *float32 `mandatory:"false" json:"actualSpend"`

	// The forecasted spend in currency by the end of the current budget cycle.
	ForecastedSpend *float32 `mandatory:"false" json:"forecastedSpend"`

	// The time the budget spend was last computed.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BudgetSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingResetPeriodEnum(string(m.ResetPeriod)); !ok && m.ResetPeriod != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ResetPeriod: %s. Supported values are: %s.", m.ResetPeriod, strings.Join(GetResetPeriodEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingProcessingPeriodTypeEnum(string(m.ProcessingPeriodType)); !ok && m.ProcessingPeriodType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ProcessingPeriodType: %s. Supported values are: %s.", m.ProcessingPeriodType, strings.Join(GetProcessingPeriodTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingTargetTypeEnum(string(m.TargetType)); !ok && m.TargetType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for TargetType: %s. Supported values are: %s.", m.TargetType, strings.Join(GetTargetTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
