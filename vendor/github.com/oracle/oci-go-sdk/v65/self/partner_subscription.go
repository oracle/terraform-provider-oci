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

// PartnerSubscription These api for partner to communicate to OCI.
type PartnerSubscription struct {

	// The unique identifier of the Subscription.
	SubscriptionId *string `mandatory:"true" json:"subscriptionId"`

	// The subscription name. Must be unique within the compartment. This value can be updated.
	DisplayName *string `mandatory:"false" json:"displayName"`

	PricingPlan *PricingPlan `mandatory:"false" json:"pricingPlan"`

	// The type of seller in SELF Service.
	SourceType SourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The unique OCID of the product, effectively functioning as the listing ID.
	ProductId *string `mandatory:"false" json:"productId"`

	// Whether subscription should be auto-renewed at the end of cycle.
	IsAutoRenew *bool `mandatory:"false" json:"isAutoRenew"`

	// Additional details that are specific for this subscription such as activation details.
	AdditionalDetails []ExtendedMetadata `mandatory:"false" json:"additionalDetails"`

	// The current state of the Subscription.
	LifecycleState LifecycleStateEnumEnum `mandatory:"false" json:"lifecycleState,omitempty"`

	// A message that describes the current state of the Subscription in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails LifecycleDetailsEnumEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// The date and time the Subscription was started, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the Subscription was ended, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// The date and time the Subscription was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// The date and time the Subscription was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`
}

func (m PartnerSubscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m PartnerSubscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetSourceTypeEnumStringValues(), ",")))
	}
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
