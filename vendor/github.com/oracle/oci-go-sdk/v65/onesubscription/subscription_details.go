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

// SubscriptionDetails Commitment summary for a subscription.
type SubscriptionDetails struct {

	// Contract identifier.
	ContractId *string `mandatory:"false" json:"contractId"`

	// Subscription identifier.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// Sold-to customer name.
	CustomerName *string `mandatory:"false" json:"customerName"`

	// The currency of the subscription, in the format specified by ISO-4217 (https://www.iso.org/iso-4217-currency-codes.html).
	CurrencyCode *string `mandatory:"false" json:"currencyCode"`

	// Start date and time of the subscription commitment.
	TimeCommitmentStart *common.SDKTime `mandatory:"false" json:"timeCommitmentStart"`

	// End date and time of the subscription commitment.
	TimeCommitmentEnd *common.SDKTime `mandatory:"false" json:"timeCommitmentEnd"`

	// Total committed amount for the subscription.
	AmountCommitment *float64 `mandatory:"false" json:"amountCommitment"`

	// Amount authorized for the subscription.
	AmountAuthorized *float64 `mandatory:"false" json:"amountAuthorized"`

	// Amount provisioned for the subscription.
	AmountProvisioned *float64 `mandatory:"false" json:"amountProvisioned"`

	// Amount consumed from the subscription commitment.
	AmountConsumed *float64 `mandatory:"false" json:"amountConsumed"`

	// Amount remaining and available for the subscription.
	AmountAvailable *float64 `mandatory:"false" json:"amountAvailable"`

	// Commitment periods associated with the subscription.
	CommitmentPeriods []CommitmentPeriod `mandatory:"false" json:"commitmentPeriods"`

	// Products and services that draw down from the subscription.
	ProductsAndServices []ProductAndService `mandatory:"false" json:"productsAndServices"`
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
