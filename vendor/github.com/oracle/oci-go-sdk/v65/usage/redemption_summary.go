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

// RedemptionSummary The redemption summary for the requested subscription ID and date range.
type RedemptionSummary struct {

	// It provides redeem date.
	TimeRedeemed *common.SDKTime `mandatory:"false" json:"timeRedeemed"`

	// It provides the redemption email id.
	RedemptionEmail *string `mandatory:"false" json:"redemptionEmail"`

	// The redemption code used in the Billing Center during the reward redemption process.
	RedemptionCode *string `mandatory:"false" json:"redemptionCode"`

	// It provides the invoice number against the redemption.
	InvoiceNumber *string `mandatory:"false" json:"invoiceNumber"`

	// It provides the invoice total amount of given redemption.
	InvoiceTotalAmount *float64 `mandatory:"false" json:"invoiceTotalAmount"`

	// The currency associated with invoice.
	InvoiceCurrency *string `mandatory:"false" json:"invoiceCurrency"`

	// It provides the redeemed rewards in invoice currency.
	RedeemedRewards *float32 `mandatory:"false" json:"redeemedRewards"`

	// It provides the redeemed rewards in base/subscription currency.
	BaseRewards *float32 `mandatory:"false" json:"baseRewards"`

	// It provides the fxRate between invoice currency and subscription currency.
	FxRate *float64 `mandatory:"false" json:"fxRate"`

	// It provides the invoice date.
	TimeInvoiced *common.SDKTime `mandatory:"false" json:"timeInvoiced"`
}

func (m RedemptionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RedemptionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
