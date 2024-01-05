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

// MonthlyRewardSummary Object describing the monthly rewards summary for the requested subscription ID.
type MonthlyRewardSummary struct {

	// The number of rewards available for a specific usage period.
	AvailableRewards *float32 `mandatory:"false" json:"availableRewards"`

	// The number of rewards redeemed for a specific month.
	RedeemedRewards *float32 `mandatory:"false" json:"redeemedRewards"`

	// The number of rewards earned for the specific usage period.
	EarnedRewards *float32 `mandatory:"false" json:"earnedRewards"`

	// The boolean parameter to indicate whether or not the available rewards are manually posted.
	IsManual *bool `mandatory:"false" json:"isManual"`

	// The date and time when rewards expire.
	TimeRewardsExpired *common.SDKTime `mandatory:"false" json:"timeRewardsExpired"`

	// The date and time when rewards accrue.
	TimeRewardsEarned *common.SDKTime `mandatory:"false" json:"timeRewardsEarned"`

	// The start date and time for the usage period.
	TimeUsageStarted *common.SDKTime `mandatory:"false" json:"timeUsageStarted"`

	// The end date and time for the usage period.
	TimeUsageEnded *common.SDKTime `mandatory:"false" json:"timeUsageEnded"`

	// The usage amount for the usage period.
	UsageAmount *float64 `mandatory:"false" json:"usageAmount"`

	// The eligible usage amount for the usage period.
	EligibleUsageAmount *float64 `mandatory:"false" json:"eligibleUsageAmount"`

	// The ineligible usage amount for the usage period.
	IneligibleUsageAmount *float64 `mandatory:"false" json:"ineligibleUsageAmount"`

	// The usage period ID.
	UsagePeriodKey *string `mandatory:"false" json:"usagePeriodKey"`
}

func (m MonthlyRewardSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonthlyRewardSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
