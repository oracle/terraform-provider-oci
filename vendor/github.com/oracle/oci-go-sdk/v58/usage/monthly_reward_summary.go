// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// UsageApi API
//
// A description of the UsageApi API.
//

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// MonthlyRewardSummary Object describing the rewards summary for a month for the requested subscriptionId.
type MonthlyRewardSummary struct {

	// The number of rewards available for a specific usage period.
	AvailableRewards *float32 `mandatory:"false" json:"availableRewards"`

	// The number of rewards redeemed for a specific month.
	RedeemedRewards *float32 `mandatory:"false" json:"redeemedRewards"`

	// The number of rewards earned for the specific usage period.
	EarnedRewards *float32 `mandatory:"false" json:"earnedRewards"`

	// The boolean flag to tell if the available rewards are posted manually or not.
	IsManual *bool `mandatory:"false" json:"isManual"`

	// The date and time on which rewards are expired.
	TimeRewardsExpired *common.SDKTime `mandatory:"false" json:"timeRewardsExpired"`

	// The date and time on which rewards are accrued.
	TimeRewardsEarned *common.SDKTime `mandatory:"false" json:"timeRewardsEarned"`

	// The start date and time for the usage period.
	TimeUsageStarted *common.SDKTime `mandatory:"false" json:"timeUsageStarted"`

	// The end date and time for the usage period.
	TimeUsageEnded *common.SDKTime `mandatory:"false" json:"timeUsageEnded"`

	// The usage amount for the usage period.
	UsageAmount *float64 `mandatory:"false" json:"usageAmount"`

	// The eligible usage amount for the usage period.
	EligibleUsageAmount *float64 `mandatory:"false" json:"eligibleUsageAmount"`

	// The in eligible usage amount for the usage period.
	IneligibleUsageAmount *float64 `mandatory:"false" json:"ineligibleUsageAmount"`

	// The id for the usage period.
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
