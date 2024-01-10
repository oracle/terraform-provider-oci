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

// UpdateAlertRuleDetails The update alert rule details.
type UpdateAlertRuleDetails struct {

	// The name of the alert rule. Avoid entering confidential information.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of the alert. Valid values are ACTUAL (the alert triggers based on actual usage), or
	// FORECAST (the alert triggers based on predicted usage).
	Type AlertTypeEnum `mandatory:"false" json:"type,omitempty"`

	// The threshold for triggering the alert, expressed as a whole number or decimal value.
	// If the thresholdType is ABSOLUTE, the threshold can have at most 12 digits before the decimal point, and up to two digits after the decimal point.
	// If the thresholdType is PERCENTAGE, the maximum value is 10000 and can have up to two digits after the decimal point.
	Threshold *float32 `mandatory:"false" json:"threshold"`

	// The type of threshold.
	ThresholdType ThresholdTypeEnum `mandatory:"false" json:"thresholdType,omitempty"`

	// The audience that receives the alert when it triggers. If you need to clear out this value, pass in an empty string instead of a null value.
	Recipients *string `mandatory:"false" json:"recipients"`

	// The description of the alert rule.
	Description *string `mandatory:"false" json:"description"`

	// The message to be delivered to the recipients when an alert is triggered.
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

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAlertRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAlertTypeEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetAlertTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingThresholdTypeEnum(string(m.ThresholdType)); !ok && m.ThresholdType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ThresholdType: %s. Supported values are: %s.", m.ThresholdType, strings.Join(GetThresholdTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
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
