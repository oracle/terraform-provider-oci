// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// ProductSummary It provides details about a product rewards and usage amount.
type ProductSummary struct {

	// The ratecard product number.
	ProductNumber *string `mandatory:"false" json:"productNumber"`

	// The ratecard product Name.
	ProductName *string `mandatory:"false" json:"productName"`

	// The ratecard product usage amount.
	UsageAmount *float64 `mandatory:"false" json:"usageAmount"`

	// The earned rewards for the product.
	EarnedRewards *float32 `mandatory:"false" json:"earnedRewards"`

	// The boolean flag to tell if the product is eligible for earning rewards.
	IsEligibleToEarnRewards *bool `mandatory:"false" json:"isEligibleToEarnRewards"`
}

func (m ProductSummary) String() string {
	return common.PointerString(m)
}
