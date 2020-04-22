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

// CreateBudgetDetails The create budget details.
// Client should use 'targetType' & 'targets' to specify the target type and list of targets on which the budget is applied.
// For backwards compatibility, 'targetCompartmentId' will still be supported for all existing clients.
// However, this is considered deprecreated and all clients be upgraded to use 'targetType' & 'targets'.
// Specifying both 'targetCompartmentId' and 'targets' will cause a Bad Request.
type CreateBudgetDetails struct {

	// The OCID of the compartment
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The amount of the budget expressed as a whole number in the currency of the customer's rate card.
	Amount *float32 `mandatory:"true" json:"amount"`

	// The reset period for the budget.
	ResetPeriod ResetPeriodEnum `mandatory:"true" json:"resetPeriod"`

	// This is DEPRECTAED. Set the target compartment id in targets instead.
	TargetCompartmentId *string `mandatory:"false" json:"targetCompartmentId"`

	// The displayName of the budget.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The description of the budget.
	Description *string `mandatory:"false" json:"description"`

	// The type of target on which the budget is applied.
	TargetType TargetTypeEnum `mandatory:"false" json:"targetType,omitempty"`

	// The list of targets on which the budget is applied.
	//   If targetType is "COMPARTMENT", targets contains list of compartment OCIDs.
	//   If targetType is "TAG", targets contains list of cost tracking tag identifiers in the form of "{tagNamespace}.{tagKey}.{tagValue}".
	// Curerntly, the array should contain EXACT ONE item.
	Targets []string `mandatory:"false" json:"targets"`

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

// CreateBudgetDetailsTargetTypeEnum is an alias to type: TargetTypeEnum
// Consider using TargetTypeEnum instead
// Deprecated
type CreateBudgetDetailsTargetTypeEnum = TargetTypeEnum

// Set of constants representing the allowable values for TargetTypeEnum
// Deprecated
const (
	CreateBudgetDetailsTargetTypeCompartment TargetTypeEnum = "COMPARTMENT"
	CreateBudgetDetailsTargetTypeTag         TargetTypeEnum = "TAG"
)

// GetCreateBudgetDetailsTargetTypeEnumValues Enumerates the set of values for TargetTypeEnum
// Consider using GetTargetTypeEnumValue
// Deprecated
var GetCreateBudgetDetailsTargetTypeEnumValues = GetTargetTypeEnumValues
