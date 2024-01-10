// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage Proxy API
//
// Use the Usage Proxy API to list Oracle Support Rewards, view related detailed usage information, and manage users who redeem rewards. For more information, see Oracle Support Rewards Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/supportrewardsoverview.htm).
//

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UsageLimitSummary Encapsulates a collection of Hard and Soft Limits for a resource within a subscription.
type UsageLimitSummary struct {

	// Time when the usage limit was created
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// Entitlement ID of the usage limit
	EntitlementId *string `mandatory:"true" json:"entitlementId"`

	// The usage limit ID
	Id *string `mandatory:"true" json:"id"`

	// Time when the usage limit was modified
	TimeModified *common.SDKTime `mandatory:"true" json:"timeModified"`

	// The resource for which the limit is defined
	ResourceName *string `mandatory:"true" json:"resourceName"`

	// The service for which the limit is defined
	ServiceName *string `mandatory:"true" json:"serviceName"`

	// The limit value
	Limit *string `mandatory:"true" json:"limit"`

	// The user who created the limit
	CreatedBy *string `mandatory:"true" json:"createdBy"`

	// The user who modified the limit
	ModifiedBy *string `mandatory:"true" json:"modifiedBy"`

	// The action when usage limit is hit
	Action UsageLimitSummaryActionEnum `mandatory:"true" json:"action"`

	// The alert level of the usage limit
	AlertLevel *float32 `mandatory:"true" json:"alertLevel"`

	// The limit type of the usage limit
	LimitType UsageLimitSummaryLimitTypeEnum `mandatory:"true" json:"limitType"`

	// The value type of the usage limit
	ValueType UsageLimitSummaryValueTypeEnum `mandatory:"true" json:"valueType"`

	// The usage limit lifecycle state.
	LifecycleState UsageLimitSummaryLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The maximum hard limit set for the usage limit
	MaxHardLimit *string `mandatory:"false" json:"maxHardLimit"`

	// The SKU for which the usage limit is set
	SkuPartId *string `mandatory:"false" json:"skuPartId"`
}

func (m UsageLimitSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageLimitSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingUsageLimitSummaryActionEnum(string(m.Action)); !ok && m.Action != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Action: %s. Supported values are: %s.", m.Action, strings.Join(GetUsageLimitSummaryActionEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageLimitSummaryLimitTypeEnum(string(m.LimitType)); !ok && m.LimitType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LimitType: %s. Supported values are: %s.", m.LimitType, strings.Join(GetUsageLimitSummaryLimitTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageLimitSummaryValueTypeEnum(string(m.ValueType)); !ok && m.ValueType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for ValueType: %s. Supported values are: %s.", m.ValueType, strings.Join(GetUsageLimitSummaryValueTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUsageLimitSummaryLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetUsageLimitSummaryLifecycleStateEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UsageLimitSummaryActionEnum Enum with underlying type: string
type UsageLimitSummaryActionEnum string

// Set of constants representing the allowable values for UsageLimitSummaryActionEnum
const (
	UsageLimitSummaryActionBreach UsageLimitSummaryActionEnum = "QUOTA_BREACH"
	UsageLimitSummaryActionAlert  UsageLimitSummaryActionEnum = "QUOTA_ALERT"
)

var mappingUsageLimitSummaryActionEnum = map[string]UsageLimitSummaryActionEnum{
	"QUOTA_BREACH": UsageLimitSummaryActionBreach,
	"QUOTA_ALERT":  UsageLimitSummaryActionAlert,
}

var mappingUsageLimitSummaryActionEnumLowerCase = map[string]UsageLimitSummaryActionEnum{
	"quota_breach": UsageLimitSummaryActionBreach,
	"quota_alert":  UsageLimitSummaryActionAlert,
}

// GetUsageLimitSummaryActionEnumValues Enumerates the set of values for UsageLimitSummaryActionEnum
func GetUsageLimitSummaryActionEnumValues() []UsageLimitSummaryActionEnum {
	values := make([]UsageLimitSummaryActionEnum, 0)
	for _, v := range mappingUsageLimitSummaryActionEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageLimitSummaryActionEnumStringValues Enumerates the set of values in String for UsageLimitSummaryActionEnum
func GetUsageLimitSummaryActionEnumStringValues() []string {
	return []string{
		"QUOTA_BREACH",
		"QUOTA_ALERT",
	}
}

// GetMappingUsageLimitSummaryActionEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageLimitSummaryActionEnum(val string) (UsageLimitSummaryActionEnum, bool) {
	enum, ok := mappingUsageLimitSummaryActionEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UsageLimitSummaryLimitTypeEnum Enum with underlying type: string
type UsageLimitSummaryLimitTypeEnum string

// Set of constants representing the allowable values for UsageLimitSummaryLimitTypeEnum
const (
	UsageLimitSummaryLimitTypeHard UsageLimitSummaryLimitTypeEnum = "HARD"
	UsageLimitSummaryLimitTypeSoft UsageLimitSummaryLimitTypeEnum = "SOFT"
)

var mappingUsageLimitSummaryLimitTypeEnum = map[string]UsageLimitSummaryLimitTypeEnum{
	"HARD": UsageLimitSummaryLimitTypeHard,
	"SOFT": UsageLimitSummaryLimitTypeSoft,
}

var mappingUsageLimitSummaryLimitTypeEnumLowerCase = map[string]UsageLimitSummaryLimitTypeEnum{
	"hard": UsageLimitSummaryLimitTypeHard,
	"soft": UsageLimitSummaryLimitTypeSoft,
}

// GetUsageLimitSummaryLimitTypeEnumValues Enumerates the set of values for UsageLimitSummaryLimitTypeEnum
func GetUsageLimitSummaryLimitTypeEnumValues() []UsageLimitSummaryLimitTypeEnum {
	values := make([]UsageLimitSummaryLimitTypeEnum, 0)
	for _, v := range mappingUsageLimitSummaryLimitTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageLimitSummaryLimitTypeEnumStringValues Enumerates the set of values in String for UsageLimitSummaryLimitTypeEnum
func GetUsageLimitSummaryLimitTypeEnumStringValues() []string {
	return []string{
		"HARD",
		"SOFT",
	}
}

// GetMappingUsageLimitSummaryLimitTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageLimitSummaryLimitTypeEnum(val string) (UsageLimitSummaryLimitTypeEnum, bool) {
	enum, ok := mappingUsageLimitSummaryLimitTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UsageLimitSummaryValueTypeEnum Enum with underlying type: string
type UsageLimitSummaryValueTypeEnum string

// Set of constants representing the allowable values for UsageLimitSummaryValueTypeEnum
const (
	UsageLimitSummaryValueTypeAbsolute   UsageLimitSummaryValueTypeEnum = "ABSOLUTE"
	UsageLimitSummaryValueTypePercentage UsageLimitSummaryValueTypeEnum = "PERCENTAGE"
)

var mappingUsageLimitSummaryValueTypeEnum = map[string]UsageLimitSummaryValueTypeEnum{
	"ABSOLUTE":   UsageLimitSummaryValueTypeAbsolute,
	"PERCENTAGE": UsageLimitSummaryValueTypePercentage,
}

var mappingUsageLimitSummaryValueTypeEnumLowerCase = map[string]UsageLimitSummaryValueTypeEnum{
	"absolute":   UsageLimitSummaryValueTypeAbsolute,
	"percentage": UsageLimitSummaryValueTypePercentage,
}

// GetUsageLimitSummaryValueTypeEnumValues Enumerates the set of values for UsageLimitSummaryValueTypeEnum
func GetUsageLimitSummaryValueTypeEnumValues() []UsageLimitSummaryValueTypeEnum {
	values := make([]UsageLimitSummaryValueTypeEnum, 0)
	for _, v := range mappingUsageLimitSummaryValueTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageLimitSummaryValueTypeEnumStringValues Enumerates the set of values in String for UsageLimitSummaryValueTypeEnum
func GetUsageLimitSummaryValueTypeEnumStringValues() []string {
	return []string{
		"ABSOLUTE",
		"PERCENTAGE",
	}
}

// GetMappingUsageLimitSummaryValueTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageLimitSummaryValueTypeEnum(val string) (UsageLimitSummaryValueTypeEnum, bool) {
	enum, ok := mappingUsageLimitSummaryValueTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// UsageLimitSummaryLifecycleStateEnum Enum with underlying type: string
type UsageLimitSummaryLifecycleStateEnum string

// Set of constants representing the allowable values for UsageLimitSummaryLifecycleStateEnum
const (
	UsageLimitSummaryLifecycleStateActive UsageLimitSummaryLifecycleStateEnum = "ACTIVE"
)

var mappingUsageLimitSummaryLifecycleStateEnum = map[string]UsageLimitSummaryLifecycleStateEnum{
	"ACTIVE": UsageLimitSummaryLifecycleStateActive,
}

var mappingUsageLimitSummaryLifecycleStateEnumLowerCase = map[string]UsageLimitSummaryLifecycleStateEnum{
	"active": UsageLimitSummaryLifecycleStateActive,
}

// GetUsageLimitSummaryLifecycleStateEnumValues Enumerates the set of values for UsageLimitSummaryLifecycleStateEnum
func GetUsageLimitSummaryLifecycleStateEnumValues() []UsageLimitSummaryLifecycleStateEnum {
	values := make([]UsageLimitSummaryLifecycleStateEnum, 0)
	for _, v := range mappingUsageLimitSummaryLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetUsageLimitSummaryLifecycleStateEnumStringValues Enumerates the set of values in String for UsageLimitSummaryLifecycleStateEnum
func GetUsageLimitSummaryLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
	}
}

// GetMappingUsageLimitSummaryLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUsageLimitSummaryLifecycleStateEnum(val string) (UsageLimitSummaryLifecycleStateEnum, bool) {
	enum, ok := mappingUsageLimitSummaryLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
