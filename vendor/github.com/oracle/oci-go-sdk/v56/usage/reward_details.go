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

// RewardDetails The overrall reward summary of the monthly summary rewards.
type RewardDetails struct {

	// The OCID of the target tenancy.
	TenancyId *string `mandatory:"false" json:"tenancyId"`

	// The entitlement id from MQS and it is same as subcription id.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The currency unit for the reward amount.
	Currency *string `mandatory:"false" json:"currency"`

	// The current Rewards percentage in decimal format.
	RewardsRate *float64 `mandatory:"false" json:"rewardsRate"`

	// The total number of available rewards for a given subscription Id.
	TotalRewardsAvailable *float32 `mandatory:"false" json:"totalRewardsAvailable"`
}

func (m RewardDetails) String() string {
	return common.PointerString(m)
}
