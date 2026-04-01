// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PricingPlan A pricing plan details provided by the Publisher.
type PricingPlan struct {

	// The type of the subscription plan.
	PlanType PricingPlanPlanTypeEnum `mandatory:"true" json:"planType"`

	// The name of the subscription plan used to identify the plan.
	PlanName *string `mandatory:"true" json:"planName"`

	// Specifies the interval at which billing occurs for the subscription plan.
	BillingFrequency PricingPlanBillingFrequencyEnum `mandatory:"true" json:"billingFrequency"`

	// The pricing details of the subscription plan in various supported currencies.
	Rates []PricingRate `mandatory:"true" json:"rates"`

	// A detailed explanation of the subscription plan.
	PlanDescription *string `mandatory:"false" json:"planDescription"`

	// Specifies the interval at which billing occurs for the subscription plan.
	PlanDuration PricingPlanPlanDurationEnum `mandatory:"false" json:"planDuration,omitempty"`
}

func (m PricingPlan) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PricingPlan) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPricingPlanPlanTypeEnum(string(m.PlanType)); !ok && m.PlanType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanType: %s. Supported values are: %s.", m.PlanType, strings.Join(GetPricingPlanPlanTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPricingPlanBillingFrequencyEnum(string(m.BillingFrequency)); !ok && m.BillingFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingFrequency: %s. Supported values are: %s.", m.BillingFrequency, strings.Join(GetPricingPlanBillingFrequencyEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPricingPlanPlanDurationEnum(string(m.PlanDuration)); !ok && m.PlanDuration != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PlanDuration: %s. Supported values are: %s.", m.PlanDuration, strings.Join(GetPricingPlanPlanDurationEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PricingPlanPlanTypeEnum Enum with underlying type: string
type PricingPlanPlanTypeEnum string

// Set of constants representing the allowable values for PricingPlanPlanTypeEnum
const (
	PricingPlanPlanTypeFixed PricingPlanPlanTypeEnum = "FIXED"
)

var mappingPricingPlanPlanTypeEnum = map[string]PricingPlanPlanTypeEnum{
	"FIXED": PricingPlanPlanTypeFixed,
}

var mappingPricingPlanPlanTypeEnumLowerCase = map[string]PricingPlanPlanTypeEnum{
	"fixed": PricingPlanPlanTypeFixed,
}

// GetPricingPlanPlanTypeEnumValues Enumerates the set of values for PricingPlanPlanTypeEnum
func GetPricingPlanPlanTypeEnumValues() []PricingPlanPlanTypeEnum {
	values := make([]PricingPlanPlanTypeEnum, 0)
	for _, v := range mappingPricingPlanPlanTypeEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingPlanPlanTypeEnumStringValues Enumerates the set of values in String for PricingPlanPlanTypeEnum
func GetPricingPlanPlanTypeEnumStringValues() []string {
	return []string{
		"FIXED",
	}
}

// GetMappingPricingPlanPlanTypeEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingPlanPlanTypeEnum(val string) (PricingPlanPlanTypeEnum, bool) {
	enum, ok := mappingPricingPlanPlanTypeEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PricingPlanBillingFrequencyEnum Enum with underlying type: string
type PricingPlanBillingFrequencyEnum string

// Set of constants representing the allowable values for PricingPlanBillingFrequencyEnum
const (
	PricingPlanBillingFrequencyYearly PricingPlanBillingFrequencyEnum = "YEARLY"
)

var mappingPricingPlanBillingFrequencyEnum = map[string]PricingPlanBillingFrequencyEnum{
	"YEARLY": PricingPlanBillingFrequencyYearly,
}

var mappingPricingPlanBillingFrequencyEnumLowerCase = map[string]PricingPlanBillingFrequencyEnum{
	"yearly": PricingPlanBillingFrequencyYearly,
}

// GetPricingPlanBillingFrequencyEnumValues Enumerates the set of values for PricingPlanBillingFrequencyEnum
func GetPricingPlanBillingFrequencyEnumValues() []PricingPlanBillingFrequencyEnum {
	values := make([]PricingPlanBillingFrequencyEnum, 0)
	for _, v := range mappingPricingPlanBillingFrequencyEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingPlanBillingFrequencyEnumStringValues Enumerates the set of values in String for PricingPlanBillingFrequencyEnum
func GetPricingPlanBillingFrequencyEnumStringValues() []string {
	return []string{
		"YEARLY",
	}
}

// GetMappingPricingPlanBillingFrequencyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingPlanBillingFrequencyEnum(val string) (PricingPlanBillingFrequencyEnum, bool) {
	enum, ok := mappingPricingPlanBillingFrequencyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// PricingPlanPlanDurationEnum Enum with underlying type: string
type PricingPlanPlanDurationEnum string

// Set of constants representing the allowable values for PricingPlanPlanDurationEnum
const (
	PricingPlanPlanDurationAnnual PricingPlanPlanDurationEnum = "ANNUAL"
)

var mappingPricingPlanPlanDurationEnum = map[string]PricingPlanPlanDurationEnum{
	"ANNUAL": PricingPlanPlanDurationAnnual,
}

var mappingPricingPlanPlanDurationEnumLowerCase = map[string]PricingPlanPlanDurationEnum{
	"annual": PricingPlanPlanDurationAnnual,
}

// GetPricingPlanPlanDurationEnumValues Enumerates the set of values for PricingPlanPlanDurationEnum
func GetPricingPlanPlanDurationEnumValues() []PricingPlanPlanDurationEnum {
	values := make([]PricingPlanPlanDurationEnum, 0)
	for _, v := range mappingPricingPlanPlanDurationEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingPlanPlanDurationEnumStringValues Enumerates the set of values in String for PricingPlanPlanDurationEnum
func GetPricingPlanPlanDurationEnumStringValues() []string {
	return []string{
		"ANNUAL",
	}
}

// GetMappingPricingPlanPlanDurationEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingPlanPlanDurationEnum(val string) (PricingPlanPlanDurationEnum, bool) {
	enum, ok := mappingPricingPlanPlanDurationEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
