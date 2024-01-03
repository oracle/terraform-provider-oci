// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage Proxy API
//
// Use the Usage Proxy API to list Oracle Support Rewards, view related detailed usage information, and manage users who redeem rewards. For more information, see Oracle Support Rewards Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/supportrewardsoverview.htm).
//

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProductSummary Provides details about product rewards and the usage amount.
type ProductSummary struct {

	// The rate card product number.
	ProductNumber *string `mandatory:"false" json:"productNumber"`

	// The rate card product name.
	ProductName *string `mandatory:"false" json:"productName"`

	// The rate card product usage amount.
	UsageAmount *float64 `mandatory:"false" json:"usageAmount"`

	// The earned rewards for the product.
	EarnedRewards *float32 `mandatory:"false" json:"earnedRewards"`

	// The boolean parameter to indicate if the product is eligible to earn rewards.
	IsEligibleToEarnRewards *bool `mandatory:"false" json:"isEligibleToEarnRewards"`
}

func (m ProductSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
