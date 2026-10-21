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

// ProductAndService Authorization-order product or service that draws down from the subscription.
type ProductAndService struct {

	// Subscription identifier; may be blank for non-cloud transactions.
	SubId *string `mandatory:"false" json:"subId"`

	// Name of the product or service.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Start date and time for the product or service.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// End date and time for the product or service.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Amount authorized for the product or service.
	AmountAuthorized *float64 `mandatory:"false" json:"amountAuthorized"`

	// Amount provisioned for the product or service.
	AmountProvisioned *float64 `mandatory:"false" json:"amountProvisioned"`

	// Amount consumed by the product or service.
	AmountConsumed *float64 `mandatory:"false" json:"amountConsumed"`
}

func (m ProductAndService) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProductAndService) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
