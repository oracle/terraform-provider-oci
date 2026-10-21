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

// UsageLedgerSummary Revenue invoice consumption detail for a subscription.
type UsageLedgerSummary struct {

	// Commitment period from the subscription details summary.
	CommitmentPeriod *string `mandatory:"false" json:"commitmentPeriod"`

	// Revenue invoice consumption date.
	TimeConsumption *common.SDKTime `mandatory:"false" json:"timeConsumption"`

	// Authorization-order subscription identifier.
	SubId *string `mandatory:"false" json:"subId"`

	// Name of the consumed product or service.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Amount consumed on the consumption date.
	ConsumptionAmount *float64 `mandatory:"false" json:"consumptionAmount"`

	// Consumption category. Known values include Consumption, Commitment Expiration, and Overage.
	Type *string `mandatory:"false" json:"type"`
}

func (m UsageLedgerSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageLedgerSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
