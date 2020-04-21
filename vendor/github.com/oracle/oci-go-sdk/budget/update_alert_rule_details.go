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

// UpdateAlertRuleDetails The update alert rule details.
type UpdateAlertRuleDetails struct {

	// The name of the alert rule.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// Type of alert. Valid values are ACTUAL (the alert will trigger based on actual usage) or
	// FORECAST (the alert will trigger based on predicted usage).
	Type AlertTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The threshold for triggering the alert expressed as a whole number or decimal value.
	// If thresholdType is ABSOLUTE, threshold can have at most 12 digits before the decimal point and up to 2 digits after the decimal point.
	// If thresholdType is PERCENTAGE, the maximum value is 10000 and can have up to 2 digits after the decimal point.
	Threshold *float32 `mandatory:"false" json:"threshold"`

	// The type of threshold.
	ThresholdType ThresholdTypeEnum `mandatory:"false" json:"thresholdType,omitempty"`

	// The audience that will receive the alert when it triggers. If you need to clear out this value, please pass in an empty string instead of null.
	Recipients *string `mandatory:"false" json:"recipients"`

	// The description of the alert rule
	Description *string `mandatory:"false" json:"description"`

	// The message to be delivered to the recipients when alert is triggered
	Message *string `mandatory:"false" json:"message"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.cloud.oracle.com/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAlertRuleDetails) String() string {
	return common.PointerString(m)
}

// UpdateAlertRuleDetailsTypeEnum is an alias to type: AlertTypeEnum
// Consider using AlertTypeEnum instead
// Deprecated
type UpdateAlertRuleDetailsTypeEnum = AlertTypeEnum

// Set of constants representing the allowable values for AlertTypeEnum
// Deprecated
const (
	UpdateAlertRuleDetailsTypeActual   AlertTypeEnum = "ACTUAL"
	UpdateAlertRuleDetailsTypeForecast AlertTypeEnum = "FORECAST"
)

// GetUpdateAlertRuleDetailsTypeEnumValues Enumerates the set of values for AlertTypeEnum
// Consider using GetAlertTypeEnumValue
// Deprecated
var GetUpdateAlertRuleDetailsTypeEnumValues = GetAlertTypeEnumValues

// UpdateAlertRuleDetailsThresholdTypeEnum is an alias to type: ThresholdTypeEnum
// Consider using ThresholdTypeEnum instead
// Deprecated
type UpdateAlertRuleDetailsThresholdTypeEnum = ThresholdTypeEnum

// Set of constants representing the allowable values for ThresholdTypeEnum
// Deprecated
const (
	UpdateAlertRuleDetailsThresholdTypePercentage ThresholdTypeEnum = "PERCENTAGE"
	UpdateAlertRuleDetailsThresholdTypeAbsolute   ThresholdTypeEnum = "ABSOLUTE"
)

// GetUpdateAlertRuleDetailsThresholdTypeEnumValues Enumerates the set of values for ThresholdTypeEnum
// Consider using GetThresholdTypeEnumValue
// Deprecated
var GetUpdateAlertRuleDetailsThresholdTypeEnumValues = GetThresholdTypeEnumValues
