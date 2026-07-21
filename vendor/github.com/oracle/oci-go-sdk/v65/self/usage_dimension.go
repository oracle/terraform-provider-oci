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

// UsageDimension A metered usage dimension associated with a usage-based or hybrid pricing plan.
type UsageDimension struct {

	// The name of the usage dimension.
	DimensionName *string `mandatory:"true" json:"dimensionName"`

	// The stable key used internally to map this usage dimension to billing details.
	DimensionKey *string `mandatory:"true" json:"dimensionKey"`

	// A detailed explanation of the usage dimension.
	DimensionDescription *string `mandatory:"true" json:"dimensionDescription"`

	// The metric type in which usage is measured.
	MetricType MetricTypeEnum `mandatory:"true" json:"metricType"`

	// Specifies the interval at which the usage dimension is billed.
	DimensionBillingFrequency BillingFrequencyEnum `mandatory:"true" json:"dimensionBillingFrequency"`

	// Dimension-level rates in various supported currencies.
	Rates []PricingRate `mandatory:"true" json:"rates"`

	// Quantity included in the base fee for hybrid plans.
	IncludedQuantity *float32 `mandatory:"false" json:"includedQuantity"`
}

func (m UsageDimension) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageDimension) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingMetricTypeEnum(string(m.MetricType)); !ok && m.MetricType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricType: %s. Supported values are: %s.", m.MetricType, strings.Join(GetMetricTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingBillingFrequencyEnum(string(m.DimensionBillingFrequency)); !ok && m.DimensionBillingFrequency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for DimensionBillingFrequency: %s. Supported values are: %s.", m.DimensionBillingFrequency, strings.Join(GetBillingFrequencyEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
