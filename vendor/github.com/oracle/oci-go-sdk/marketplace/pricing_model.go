// Copyright (c) 2016, 2018, 2019, Oracle and/or its affiliates. All rights reserved.
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

	// The pricing rate.
	Rate *float32 `mandatory:"false" json:"rate"`
}

func (m PricingModel) String() string {
	return common.PointerString(m)
}
