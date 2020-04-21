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

// AlertRule The alert rule.
type AlertRule struct {

	// The OCID of the alert rule
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the budget
	BudgetId *string `mandatory:"true" json:"budgetId"`

	// The name of the alert rule.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of alert. Valid values are ACTUAL (the alert will trigger based on actual usage) or
	// FORECAST (the alert will trigger based on predicted usage).
	Type AlertTypeEnum `mandatory:"true" json:"type"`

	// The threshold for triggering the alert. If thresholdType is PERCENTAGE, the maximum value is 10000.
	Threshold *float32 `mandatory:"true" json:"threshold"`

	// The type of threshold.
	ThresholdType ThresholdTypeEnum `mandatory:"true" json:"thresholdType"`

	// The current state of the alert rule.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// Delimited list of email addresses to receive the alert when it triggers.
	// Delimiter character can be comma, space, TAB, or semicolon.
	Recipients *string `mandatory:"true" json:"recipients"`

	// Time budget was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Time budget was updated
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Custom message sent when alert is triggered
	Message *string `mandatory:"false" json:"message"`

	// The description of the alert rule.
	Description *string `mandatory:"false" json:"description"`

	// Version of the alert rule. Starts from 1 and increments by 1.
	Version *int `mandatory:"false" json:"version"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m AlertRule) String() string {
	return common.PointerString(m)
}

// AlertRuleTypeEnum is an alias to type: AlertTypeEnum
// Consider using AlertTypeEnum instead
// Deprecated
type AlertRuleTypeEnum = AlertTypeEnum

// Set of constants representing the allowable values for AlertTypeEnum
// Deprecated
const (
	AlertRuleTypeActual   AlertTypeEnum = "ACTUAL"
	AlertRuleTypeForecast AlertTypeEnum = "FORECAST"
)

// GetAlertRuleTypeEnumValues Enumerates the set of values for AlertTypeEnum
// Consider using GetAlertTypeEnumValue
// Deprecated
var GetAlertRuleTypeEnumValues = GetAlertTypeEnumValues

// AlertRuleThresholdTypeEnum is an alias to type: ThresholdTypeEnum
// Consider using ThresholdTypeEnum instead
// Deprecated
type AlertRuleThresholdTypeEnum = ThresholdTypeEnum

// Set of constants representing the allowable values for ThresholdTypeEnum
// Deprecated
const (
	AlertRuleThresholdTypePercentage ThresholdTypeEnum = "PERCENTAGE"
	AlertRuleThresholdTypeAbsolute   ThresholdTypeEnum = "ABSOLUTE"
)

// GetAlertRuleThresholdTypeEnumValues Enumerates the set of values for ThresholdTypeEnum
// Consider using GetThresholdTypeEnumValue
// Deprecated
var GetAlertRuleThresholdTypeEnumValues = GetThresholdTypeEnumValues

// AlertRuleLifecycleStateEnum is an alias to type: LifecycleStateEnum
// Consider using LifecycleStateEnum instead
// Deprecated
type AlertRuleLifecycleStateEnum = LifecycleStateEnum

// Set of constants representing the allowable values for LifecycleStateEnum
// Deprecated
const (
	AlertRuleLifecycleStateActive   LifecycleStateEnum = "ACTIVE"
	AlertRuleLifecycleStateInactive LifecycleStateEnum = "INACTIVE"
)

// GetAlertRuleLifecycleStateEnumValues Enumerates the set of values for LifecycleStateEnum
// Consider using GetLifecycleStateEnumValue
// Deprecated
var GetAlertRuleLifecycleStateEnumValues = GetLifecycleStateEnumValues
