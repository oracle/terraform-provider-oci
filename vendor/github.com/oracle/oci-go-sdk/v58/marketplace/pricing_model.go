// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// PricingModel The model for pricing.
type PricingModel struct {

	// The type of the pricing model.
	Type PricingTypeEnumEnum `mandatory:"true" json:"type"`

	// The type of pricing for a PAYGO model, eg PER_OCPU_LINEAR, PER_OCPU_MIN_BILLING, PER_INSTANCE.  Null if type is not PAYGO.
	PayGoStrategy PricingStrategyEnumEnum `mandatory:"false" json:"payGoStrategy,omitempty"`

	// The currency of the pricing model.
	Currency PricingCurrencyEnumEnum `mandatory:"false" json:"currency,omitempty"`

	// The pricing rate.
	Rate *float32 `mandatory:"false" json:"rate"`
}

func (m PricingModel) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PricingModel) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPricingTypeEnumEnum(string(m.Type)); !ok && m.Type != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Type: %s. Supported values are: %s.", m.Type, strings.Join(GetPricingTypeEnumEnumStringValues(), ",")))
	}

	if _, ok := GetMappingPricingStrategyEnumEnum(string(m.PayGoStrategy)); !ok && m.PayGoStrategy != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for PayGoStrategy: %s. Supported values are: %s.", m.PayGoStrategy, strings.Join(GetPricingStrategyEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingPricingCurrencyEnumEnum(string(m.Currency)); !ok && m.Currency != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Currency: %s. Supported values are: %s.", m.Currency, strings.Join(GetPricingCurrencyEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
