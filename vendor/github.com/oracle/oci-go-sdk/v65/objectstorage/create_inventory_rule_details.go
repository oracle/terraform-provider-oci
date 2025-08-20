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

// CreateInventoryRuleDetails The details to create a inventory rule.
type CreateInventoryRuleDetails struct {

	// The user-defined name for the inventory rule.
	// - Must be unique within the tenancy's Object Storage namespace.
	// - Must be between 1 and 256 characters in length.
	// - Valid characters are uppercase and lowercase letters, numbers, hyphens (-), underscores (_), and periods (.).
	RuleName *string `mandatory:"true" json:"ruleName"`

	Filter *RuleFilter `mandatory:"true" json:"filter"`

	Report *Report `mandatory:"true" json:"report"`

	// The inventory rule type.
	RuleType CreateInventoryRuleDetailsRuleTypeEnum `mandatory:"false" json:"ruleType,omitempty"`

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

func (m CreateInventoryRuleDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateInventoryRuleDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingCreateInventoryRuleDetailsRuleTypeEnum(string(m.RuleType)); !ok && m.RuleType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for RuleType: %s. Supported values are: %s.", m.RuleType, strings.Join(GetCreateInventoryRuleDetailsRuleTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// CreateInventoryRuleDetailsRuleTypeEnum Enum with underlying type: string
type CreateInventoryRuleDetailsRuleTypeEnum string

// Set of constants representing the allowable values for CreateInventoryRuleDetailsRuleTypeEnum
const (
	CreateInventoryRuleDetailsRuleTypeBucketRule CreateInventoryRuleDetailsRuleTypeEnum = "BUCKET_RULE"
	CreateInventoryRuleDetailsRuleTypeObjectRule CreateInventoryRuleDetailsRuleTypeEnum = "OBJECT_RULE"
)

var mappingCreateInventoryRuleDetailsRuleTypeEnum = map[string]CreateInventoryRuleDetailsRuleTypeEnum{
	"BUCKET_RULE": CreateInventoryRuleDetailsRuleTypeBucketRule,
	"OBJECT_RULE": CreateInventoryRuleDetailsRuleTypeObjectRule,
}

var mappingCreateInventoryRuleDetailsRuleTypeEnumLowerCase = map[string]CreateInventoryRuleDetailsRuleTypeEnum{
	"bucket_rule": CreateInventoryRuleDetailsRuleTypeBucketRule,
	"object_rule": CreateInventoryRuleDetailsRuleTypeObjectRule,
}

// GetCreateInventoryRuleDetailsRuleTypeEnumValues Enumerates the set of values for CreateInventoryRuleDetailsRuleTypeEnum
func GetCreateInventoryRuleDetailsRuleTypeEnumValues() []CreateInventoryRuleDetailsRuleTypeEnum {
	values := make([]CreateInventoryRuleDetailsRuleTypeEnum, 0)
	for _, v := range mappingCreateInventoryRuleDetailsRuleTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetCreateInventoryRuleDetailsRuleTypeEnumStringValues Enumerates the set of values in String for CreateInventoryRuleDetailsRuleTypeEnum
func GetCreateInventoryRuleDetailsRuleTypeEnumStringValues() []string {
	return []string{
		"BUCKET_RULE",
		"OBJECT_RULE",
	}
}

// GetMappingCreateInventoryRuleDetailsRuleTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingCreateInventoryRuleDetailsRuleTypeEnum(val string) (CreateInventoryRuleDetailsRuleTypeEnum, bool) {
	enum, ok := mappingCreateInventoryRuleDetailsRuleTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
