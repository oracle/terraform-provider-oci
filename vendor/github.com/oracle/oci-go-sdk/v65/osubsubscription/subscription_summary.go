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

// SubscriptionSummary Subscription summary
type SubscriptionSummary struct {

	// Status of the plan
	Status *string `mandatory:"false" json:"status"`

	// Represents the date when the first service of the subscription was activated
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// Represents the date when the last service of the subscription ends
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	Currency *Currency `mandatory:"false" json:"currency"`

	// Customer friendly service name provided by PRG
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// List of Subscribed Services of the plan
	SubscribedServices []SubscribedServiceSummary `mandatory:"false" json:"subscribedServices"`
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
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
