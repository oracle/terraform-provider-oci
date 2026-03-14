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

// Subscription The model for a listing subscription.
type Subscription struct {

	// The unique identifier for the subscription within a specific compartment.
	Id *string `mandatory:"true" json:"id"`

	// The subscription name. Must be unique within the compartment. This value can be updated.
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The unique identifier for the compartment where the subscription was purchased.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The unique identifier for the tenant where the subscription was purchased.
	TenantId *string `mandatory:"true" json:"tenantId"`

	// The date and time the Subscription was created, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The current lifecycle state of the Subscription.
	LifecycleState LifecycleStateEnumEnum `mandatory:"true" json:"lifecycleState"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"true" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"true" json:"definedTags"`

	SubscriptionDetails *SubscriptionDetails `mandatory:"false" json:"subscriptionDetails"`

	// The OCID that identifies the seller within the platform.
	SellerId *string `mandatory:"false" json:"sellerId"`

	// The type of seller in SELF Service.
	SourceType SourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// The unique OCID of the product, effectively functioning as the listing ID.
	ProductId *string `mandatory:"false" json:"productId"`

	// Additional details that are specific for this subscription such as activation details.
	AdditionalDetails []ExtendedMetadata `mandatory:"false" json:"additionalDetails"`

	// The realm from where customer is buying the subscription.
	Realm *string `mandatory:"false" json:"realm"`

	// The region from where customer is buying the subscription.
	Region *string `mandatory:"false" json:"region"`

	// The date and time the Subscription was updated, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"false" json:"timeUpdated"`

	// The date and time the Subscription was started, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The date and time the Subscription was ended, in the format defined by RFC 3339 (https://tools.ietf.org/html/rfc3339).
	// Example: `2016-08-25T21:10:29.600Z`
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// A message that describes the current state of the Subscription in more detail. For example,
	// can be used to provide actionable information for a resource in the Failed state.
	LifecycleDetails LifecycleDetailsEnumEnum `mandatory:"false" json:"lifecycleDetails,omitempty"`

	// System tags for this resource. Each key is predefined and scoped to a namespace.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m Subscription) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Subscription) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingLifecycleStateEnumEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetLifecycleStateEnumEnumStringValues(), ",")))
	}

	if _, ok := GetMappingSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetSourceTypeEnumStringValues(), ",")))
	}
	if _, ok := GetMappingLifecycleDetailsEnumEnum(string(m.LifecycleDetails)); !ok && m.LifecycleDetails != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleDetails: %s. Supported values are: %s.", m.LifecycleDetails, strings.Join(GetLifecycleDetailsEnumEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
