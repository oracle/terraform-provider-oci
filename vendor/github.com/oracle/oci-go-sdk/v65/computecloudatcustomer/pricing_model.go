// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Compute Cloud@Customer API
//
// Use the Compute Cloud@Customer API to manage Compute Cloud@Customer infrastructures and upgrade schedules.
// For more information see Compute Cloud@Customer documentation (https://docs.oracle.com/iaas/compute-cloud-at-customer/home.htm).
//

package computecloudatcustomer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// PricingModel The model for pricing.
type PricingModel struct {

	// The type of pricing for a PAYGO model, eg PER_OCPU_LINEAR, PER_INSTANCE.
	// Null if type is not PAYGO.
	PayGoStrategy PricingModelPayGoStrategyEnum `mandatory:"false" json:"payGoStrategy,omitempty"`

	// The currency of the pricing model.
	Currency PricingCurrencyEnumEnum `mandatory:"false" json:"currency,omitempty"`

	// The pricing rate.
	Rate *float32 `mandatory:"false" json:"rate"`

	InternationalMarketPrice *InternationalMarketPrice `mandatory:"false" json:"internationalMarketPrice"`
}

func (m PricingModel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PricingModel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingPricingModelPayGoStrategyEnum(string(m.PayGoStrategy)); !ok && m.PayGoStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PayGoStrategy: %s. Supported values are: %s.", m.PayGoStrategy, strings.Join(GetPricingModelPayGoStrategyEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPricingCurrencyEnumEnum(string(m.Currency)); !ok && m.Currency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Currency: %s. Supported values are: %s.", m.Currency, strings.Join(GetPricingCurrencyEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// PricingModelPayGoStrategyEnum Enum with underlying type: string
type PricingModelPayGoStrategyEnum string

// Set of constants representing the allowable values for PricingModelPayGoStrategyEnum
const (
	PricingModelPayGoStrategyOcpuLinear PricingModelPayGoStrategyEnum = "PER_OCPU_LINEAR"
	PricingModelPayGoStrategyInstance   PricingModelPayGoStrategyEnum = "PER_INSTANCE"
)

var mappingPricingModelPayGoStrategyEnum = map[string]PricingModelPayGoStrategyEnum{
	"PER_OCPU_LINEAR": PricingModelPayGoStrategyOcpuLinear,
	"PER_INSTANCE":    PricingModelPayGoStrategyInstance,
}

var mappingPricingModelPayGoStrategyEnumLowerCase = map[string]PricingModelPayGoStrategyEnum{
	"per_ocpu_linear": PricingModelPayGoStrategyOcpuLinear,
	"per_instance":    PricingModelPayGoStrategyInstance,
}

// GetPricingModelPayGoStrategyEnumValues Enumerates the set of values for PricingModelPayGoStrategyEnum
func GetPricingModelPayGoStrategyEnumValues() []PricingModelPayGoStrategyEnum {
	values := make([]PricingModelPayGoStrategyEnum, 0)
	for _, v := range mappingPricingModelPayGoStrategyEnum {
		values = append(values, v)
	}
	return values
}

// GetPricingModelPayGoStrategyEnumStringValues Enumerates the set of values in String for PricingModelPayGoStrategyEnum
func GetPricingModelPayGoStrategyEnumStringValues() []string {
	return []string{
		"PER_OCPU_LINEAR",
		"PER_INSTANCE",
	}
}

// GetMappingPricingModelPayGoStrategyEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingPricingModelPayGoStrategyEnum(val string) (PricingModelPayGoStrategyEnum, bool) {
	enum, ok := mappingPricingModelPayGoStrategyEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
