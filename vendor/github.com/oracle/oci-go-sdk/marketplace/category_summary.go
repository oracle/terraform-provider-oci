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

// CategorySummary The model for a summary of product categories for listings.
type CategorySummary struct {

	// Name of the product category.
	Name *string `mandatory:"false" json:"name"`
}

func (m CategorySummary) String() string {
	return common.PointerString(m)
}
