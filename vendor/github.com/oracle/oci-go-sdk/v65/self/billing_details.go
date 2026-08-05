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

// BillingDetails Billing detail entry associated with a subscription.
type BillingDetails struct {

	// Unique key used to map this SKU to the pricing plan.
	PricingPlanKey *string `mandatory:"true" json:"pricingPlanKey"`

	// The billing model this billing detail applies to.
	BillingModel BillingDetailsBillingModelEnum `mandatory:"true" json:"billingModel"`

	// Sku for service.
	Sku *string `mandatory:"true" json:"sku"`

	// The part's metric.
	MetricType MetricTypeEnum `mandatory:"true" json:"metricType"`

	// Tha rate of this sku meter.
	RateAllocation *float32 `mandatory:"true" json:"rateAllocation"`

	// The meters associated with sku.
	Meters []Meter `mandatory:"true" json:"meters"`

	// Whether this sku is assign to gov product.
	HasGovSku *bool `mandatory:"false" json:"hasGovSku"`
}

func (m BillingDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m BillingDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBillingDetailsBillingModelEnum(string(m.BillingModel)); !ok && m.BillingModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingModel: %s. Supported values are: %s.", m.BillingModel, strings.Join(GetBillingDetailsBillingModelEnumStringValues(), ",")))
	}
	if _, ok := GetMappingMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetMetricTypeEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// BillingDetailsBillingModelEnum Enum with underlying type: string
type BillingDetailsBillingModelEnum string

// Set of constants representing the allowable values for BillingDetailsBillingModelEnum
const (
	BillingDetailsBillingModelFlatRate   BillingDetailsBillingModelEnum = "FLAT_RATE"
	BillingDetailsBillingModelUsageBased BillingDetailsBillingModelEnum = "USAGE_BASED"
)

var mappingBillingDetailsBillingModelEnum = map[string]BillingDetailsBillingModelEnum{
	"FLAT_RATE":   BillingDetailsBillingModelFlatRate,
	"USAGE_BASED": BillingDetailsBillingModelUsageBased,
}

var mappingBillingDetailsBillingModelEnumLowerCase = map[string]BillingDetailsBillingModelEnum{
	"flat_rate":   BillingDetailsBillingModelFlatRate,
	"usage_based": BillingDetailsBillingModelUsageBased,
}

// GetBillingDetailsBillingModelEnumValues Enumerates the set of values for BillingDetailsBillingModelEnum
func GetBillingDetailsBillingModelEnumValues() []BillingDetailsBillingModelEnum {
	values := make([]BillingDetailsBillingModelEnum, 0)
	for _, v := range mappingBillingDetailsBillingModelEnum {
		values = append(values, v)
	}
	return values
}

// GetBillingDetailsBillingModelEnumStringValues Enumerates the set of values in String for BillingDetailsBillingModelEnum
func GetBillingDetailsBillingModelEnumStringValues() []string {
	return []string{
		"FLAT_RATE",
		"USAGE_BASED",
	}
}

// GetMappingBillingDetailsBillingModelEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingBillingDetailsBillingModelEnum(val string) (BillingDetailsBillingModelEnum, bool) {
	enum, ok := mappingBillingDetailsBillingModelEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
