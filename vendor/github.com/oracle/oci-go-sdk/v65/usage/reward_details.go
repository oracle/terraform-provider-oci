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

// RewardDetails The overall monthly reward summary.
type RewardDetails struct {

	// The OCID of the target tenancy.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The entitlement ID from MQS, which is the same as the subcription ID.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The currency unit for the reward amount.
	Currency *string `mandatory:"false" json:"currency"`

	// The current Rewards percentage in decimal format.
	RewardsRate *float64 `mandatory:"false" json:"rewardsRate"`

	// The total number of available rewards for a given subscription ID.
	TotalRewardsAvailable *float32 `mandatory:"false" json:"totalRewardsAvailable"`

	// The redemption code used in the Billing Center during the reward redemption process.
	RedemptionCode *string `mandatory:"false" json:"redemptionCode"`
}

func (m RewardDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RewardDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
