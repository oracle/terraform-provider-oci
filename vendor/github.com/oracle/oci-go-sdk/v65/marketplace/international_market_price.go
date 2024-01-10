// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Use the Marketplace API to manage applications in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.cloud.oracle.com/Content/Marketplace/Concepts/marketoverview.htm)
//

package marketplace

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// InternationalMarketPrice The model for international market pricing.
type InternationalMarketPrice struct {

	// The currency of the pricing model.
	CurrencyCode PricingCurrencyEnumEnum `mandatory:"true" json:"currencyCode"`

	// The pricing rate.
	Rate *float64 `mandatory:"true" json:"rate"`

	// The symbol of the currency
	CurrencySymbol *string `mandatory:"false" json:"currencySymbol"`
}

func (m InternationalMarketPrice) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m InternationalMarketPrice) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingPricingCurrencyEnumEnum(string(m.CurrencyCode)); !ok && m.CurrencyCode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for CurrencyCode: %s. Supported values are: %s.", m.CurrencyCode, strings.Join(GetPricingCurrencyEnumEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
