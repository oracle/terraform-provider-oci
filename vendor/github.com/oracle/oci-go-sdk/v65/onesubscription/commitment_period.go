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

// CommitmentPeriod Commitment period for the subscription.
type CommitmentPeriod struct {

	// Identifier for the commitment period.
	CommitmentPeriod *string `mandatory:"false" json:"commitmentPeriod"`

	// Start date and time of the commitment period.
	TimeCommitmentPeriodStart *common.SDKTime `mandatory:"false" json:"timeCommitmentPeriodStart"`

	// End date and time of the commitment period.
	TimeCommitmentPeriodEnd *common.SDKTime `mandatory:"false" json:"timeCommitmentPeriodEnd"`

	// Total committed amount for the commitment period.
	AmountCommitment *float64 `mandatory:"false" json:"amountCommitment"`

	// Amount consumed during the commitment period.
	AmountConsumed *float64 `mandatory:"false" json:"amountConsumed"`

	// Amount remaining and available for the commitment period.
	AmountAvailable *float64 `mandatory:"false" json:"amountAvailable"`
}

func (m CommitmentPeriod) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CommitmentPeriod) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
