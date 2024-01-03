// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// OneSubscription API Subscription, Commitment and and Rate Card Details
//
// Set of APIs that return the Subscription Details, Commitment and Effective Rate Card Details
//

package osubsubscription

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CommitmentSummary Subscribed Service commitment summary
type CommitmentSummary struct {

	// SPM internal Commitment ID
	Id *string `mandatory:"true" json:"id"`

	// Commitment start date
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Commitment end date
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// Commitment quantity
	Quantity *string `mandatory:"false" json:"quantity"`

	// Commitment used amount
	UsedAmount *string `mandatory:"false" json:"usedAmount"`

	// Commitment available amount
	AvailableAmount *string `mandatory:"false" json:"availableAmount"`

	// Funded Allocation line value
	// example: 12000.00
	FundedAllocationValue *string `mandatory:"false" json:"fundedAllocationValue"`
}

func (m CommitmentSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CommitmentSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
