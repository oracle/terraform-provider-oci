// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// SELF Service API
//
// Use the SELF Service API to manage Subscriptions in Oracle Cloud Infrastructure Marketplace. For more information, see Overview of Marketplace (https://docs.oracle.com/iaas/Content/Marketplace/Concepts/marketoverview.htm)
//

package self

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscriptionDetails The details of a subscription
type SubscriptionDetails struct {
	PricingPlan *PricingPlan `mandatory:"true" json:"pricingPlan"`

	// The activation link given by the partner.
	PartnerRegistrationUrl *string `mandatory:"true" json:"partnerRegistrationUrl"`

	BillingDetails *BillingDetails `mandatory:"true" json:"billingDetails"`

	// The currency supported, in the format specified by ISO-4217
	Currency *string `mandatory:"false" json:"currency"`

	// Tha amount for the currency type.
	Amount *float32 `mandatory:"false" json:"amount"`

	// Whether subscription should be auto-renewed at the end of cycle.
	IsAutoRenew *bool `mandatory:"false" json:"isAutoRenew"`
}

func (m SubscriptionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
