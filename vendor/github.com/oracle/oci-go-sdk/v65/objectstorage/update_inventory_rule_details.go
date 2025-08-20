// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Object Storage Service API
//
// Use Object Storage and Archive Storage APIs to manage buckets, objects, and related resources.
// For more information, see Overview of Object Storage (https://docs.oracle.com/iaas/Content/Object/Concepts/objectstorageoverview.htm) and
// Overview of Archive Storage (https://docs.oracle.com/iaas/Content/Archive/Concepts/archivestorageoverview.htm).
//

package objectstorage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateInventoryRuleDetails The details to update a inventory rule.
type UpdateInventoryRuleDetails struct {

	// The user-defined name for the inventory rule.
	// - Must be unique within the tenancy's Object Storage namespace.
	// - Must be between 1 and 256 characters in length.
	// - Valid characters are uppercase and lowercase letters, numbers, hyphens (-), underscores (_), and periods (.).
	RuleName *string `mandatory:"true" json:"ruleName"`

	Filter *RuleFilter `mandatory:"true" json:"filter"`

	Report *Report `mandatory:"true" json:"report"`

	// The inventory rule type.
	RuleType UpdateInventoryRuleDetailsRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`

	// A Boolean that determines whether this rule is currently enabled.
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Defines the recurring times for report generation using iCal RFC 5545 format.
	// The report will be generated based on the provided frequency (e.g., daily, weekly, monthly).
	// Supported values:
	// -  FREQ=DAILY (Daily recurrence)
	// -  FREQ=WEEKLY (Weekly recurrence without specifying days)
	// -  FREQ=MONTHLY (Monthly recurrence without specifying days, default value)
	// The recurrence is based purely on the frequency (e.g., daily, weekly, monthly),
	// with no support for finer details like specific days or intervals.
	ReportRecurrences *string `mandatory:"false" json:"reportRecurrences"`
}

func (m UpdateInventoryRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateInventoryRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingUpdateInventoryRuleDetailsRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetUpdateInventoryRuleDetailsRuleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateInventoryRuleDetailsRuleTypeEnum Enum with underlying type: string
type UpdateInventoryRuleDetailsRuleTypeEnum string

// Set of constants representing the allowable values for UpdateInventoryRuleDetailsRuleTypeEnum
const (
	UpdateInventoryRuleDetailsRuleTypeBucketRule UpdateInventoryRuleDetailsRuleTypeEnum = "BUCKET_RULE"
	UpdateInventoryRuleDetailsRuleTypeObjectRule UpdateInventoryRuleDetailsRuleTypeEnum = "OBJECT_RULE"
)

var mappingUpdateInventoryRuleDetailsRuleTypeEnum = map[string]UpdateInventoryRuleDetailsRuleTypeEnum{
	"BUCKET_RULE": UpdateInventoryRuleDetailsRuleTypeBucketRule,
	"OBJECT_RULE": UpdateInventoryRuleDetailsRuleTypeObjectRule,
}

var mappingUpdateInventoryRuleDetailsRuleTypeEnumLowerCase = map[string]UpdateInventoryRuleDetailsRuleTypeEnum{
	"bucket_rule": UpdateInventoryRuleDetailsRuleTypeBucketRule,
	"object_rule": UpdateInventoryRuleDetailsRuleTypeObjectRule,
}

// GetUpdateInventoryRuleDetailsRuleTypeEnumValues Enumerates the set of values for UpdateInventoryRuleDetailsRuleTypeEnum
func GetUpdateInventoryRuleDetailsRuleTypeEnumValues() []UpdateInventoryRuleDetailsRuleTypeEnum {
	values := make([]UpdateInventoryRuleDetailsRuleTypeEnum, 0)
	for _, v := range mappingUpdateInventoryRuleDetailsRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateInventoryRuleDetailsRuleTypeEnumStringValues Enumerates the set of values in String for UpdateInventoryRuleDetailsRuleTypeEnum
func GetUpdateInventoryRuleDetailsRuleTypeEnumStringValues() []string {
	return []string{
		"BUCKET_RULE",
		"OBJECT_RULE",
	}
}

// GetMappingUpdateInventoryRuleDetailsRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateInventoryRuleDetailsRuleTypeEnum(val string) (UpdateInventoryRuleDetailsRuleTypeEnum, bool) {
	enum, ok := mappingUpdateInventoryRuleDetailsRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
