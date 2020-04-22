// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Marketplace Service API
//
// Manage applications in Oracle Cloud Infrastructure Marketplace.
//

package marketplace

import (
	"github.com/oracle/oci-go-sdk/common"
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
