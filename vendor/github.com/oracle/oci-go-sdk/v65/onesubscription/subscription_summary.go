// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription APIs
//
// OneSubscription APIs
//

package onesubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscriptionSummary Subscription summary.
type SubscriptionSummary struct {

	// Plan status.
	Status *string `mandatory:"false" json:"status"`

	// The date when the first service of the subscription was activated.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The date when the last service of the subscription ends.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	Currency *SubscriptionCurrency `mandatory:"false" json:"currency"`

	// Customer-friendly service name provided by the PRG (Product Release Group).
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Hold reason of the plan.
	HoldReason *string `mandatory:"false" json:"holdReason"`

	// The date of the hold release.
	TimeHoldReleaseEta *common.SDKTime `mandatory:"false" json:"timeHoldReleaseEta"`

	// List of plan's subscribed services.
	SubscribedServices []SubscriptionSubscribedService `mandatory:"false" json:"subscribedServices"`
}

func (m SubscriptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
