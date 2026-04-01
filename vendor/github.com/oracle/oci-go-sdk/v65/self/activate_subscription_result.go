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

// ActivateSubscriptionResult Response model for the activation of a subscription. Includes the details of the activated subscription.
type ActivateSubscriptionResult struct {

	// The unique identifier of the activated subscription.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// The subscription name. Must be unique within the compartment. This value can be updated.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The start date of the subscription.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// The end date of the subscription.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The current state of the Subscription.
	LifecycleState LifecycleStateEnumEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message that describes the current state of the Subscription in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails LifecycleDetailsEnumEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`
}

func (m ActivateSubscriptionResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ActivateSubscriptionResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingLifecycleStateEnumEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleDetailsEnumEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetLifecycleDetailsEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
