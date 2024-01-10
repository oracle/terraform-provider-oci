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

// AlertRule The alert rule.
type AlertRule struct {

	// The OCID of the alert rule.
	Id *string `mandatory:"true" json:"id"`

	// The OCID of the budget.
	BudgetId *string `mandatory:"true" json:"budgetId"`

	// The name of the alert rule. Avoid entering confidential information.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The type of the alert. Valid values are ACTUAL (the alert triggers based on actual usage), or
	// FORECAST (the alert triggers based on predicted usage).
	Type AlertTypeEnum `mandatory:"true" json:"type"`

	// The threshold for triggering the alert. If the thresholdType is PERCENTAGE, the maximum value is 10000.
	Threshold *float32 `mandatory:"true" json:"threshold"`

	// The type of threshold.
	ThresholdType ThresholdTypeEnum `mandatory:"true" json:"thresholdType"`

	// The current state of the alert rule.
	LifecycleState LifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The delimited list of email addresses to receive the alert when it triggers.
	// Delimiter characters can be a comma, space, TAB, or semicolon.
	Recipients *string `mandatory:"true" json:"recipients"`

	// The time the budget was created.
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The time the budget was updated.
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// Custom message sent when an alert is triggered.
	Message *string `mandatory:"false" json:"message"`

	// The description of the alert rule.
	Description *string `mandatory:"false" json:"description"`

	// The version of the alert rule. Starts from 1 and increments by 1.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlertRule) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlertTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAlertTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingThresholdTypeEnum(string(m.ThresholdType)); !ok && m.ThresholdType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ThresholdType: %s. Supported values are: %s.", m.ThresholdType, strings.Join(GetThresholdTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
