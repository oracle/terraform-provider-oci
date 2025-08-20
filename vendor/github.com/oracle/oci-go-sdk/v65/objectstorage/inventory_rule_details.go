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

// InventoryRuleDetails The details of a inventory rule.
type InventoryRuleDetails struct {

	// The user-defined name for the inventory rule.
	// - Must be unique within the tenancy's Object Storage namespace.
	// - Must be between 1 and 256 characters in length.
	// - Valid characters are uppercase and lowercase letters, numbers, hyphens (-), underscores (_), and periods (.).
	RuleName *string `mandatory:"true" json:"ruleName"`

	Filter *RuleFilter `mandatory:"true" json:"filter"`

	Report *Report `mandatory:"true" json:"report"`

	// The inventory rule type.
	RuleType InventoryRuleDetailsRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`

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

func (m InventoryRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InventoryRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingInventoryRuleDetailsRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetInventoryRuleDetailsRuleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// InventoryRuleDetailsRuleTypeEnum Enum with underlying type: string
type InventoryRuleDetailsRuleTypeEnum string

// Set of constants representing the allowable values for InventoryRuleDetailsRuleTypeEnum
const (
	InventoryRuleDetailsRuleTypeBucketRule InventoryRuleDetailsRuleTypeEnum = "BUCKET_RULE"
	InventoryRuleDetailsRuleTypeObjectRule InventoryRuleDetailsRuleTypeEnum = "OBJECT_RULE"
)

var mappingInventoryRuleDetailsRuleTypeEnum = map[string]InventoryRuleDetailsRuleTypeEnum{
	"BUCKET_RULE": InventoryRuleDetailsRuleTypeBucketRule,
	"OBJECT_RULE": InventoryRuleDetailsRuleTypeObjectRule,
}

var mappingInventoryRuleDetailsRuleTypeEnumLowerCase = map[string]InventoryRuleDetailsRuleTypeEnum{
	"bucket_rule": InventoryRuleDetailsRuleTypeBucketRule,
	"object_rule": InventoryRuleDetailsRuleTypeObjectRule,
}

// GetInventoryRuleDetailsRuleTypeEnumValues Enumerates the set of values for InventoryRuleDetailsRuleTypeEnum
func GetInventoryRuleDetailsRuleTypeEnumValues() []InventoryRuleDetailsRuleTypeEnum {
	values := make([]InventoryRuleDetailsRuleTypeEnum, 0)
	for _, v := range mappingInventoryRuleDetailsRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetInventoryRuleDetailsRuleTypeEnumStringValues Enumerates the set of values in String for InventoryRuleDetailsRuleTypeEnum
func GetInventoryRuleDetailsRuleTypeEnumStringValues() []string {
	return []string{
		"BUCKET_RULE",
		"OBJECT_RULE",
	}
}

// GetMappingInventoryRuleDetailsRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingInventoryRuleDetailsRuleTypeEnum(val string) (InventoryRuleDetailsRuleTypeEnum, bool) {
	enum, ok := mappingInventoryRuleDetailsRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
