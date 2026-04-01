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

// CreateSubscriptionDetails The data to create a Subscription.
type CreateSubscriptionDetails struct {

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the subscription in.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenant to create the subscription in.
	TenantId *string `mandatory:"true" json:"tenantId"`

	SubscriptionDetails *SubscriptionDetails `mandatory:"true" json:"subscriptionDetails"`

	// The OCID for the seller in SELF Service.
	SellerId *string `mandatory:"true" json:"sellerId"`

	// The unique identifier of the marketplace listing in Oracle Cloud Infrastructure.
	ProductId *string `mandatory:"true" json:"productId"`

	// The subscription name. Must be unique within the compartment. This value can be updated.
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The type of seller in SELF Service.
	SourceType SourceTypeEnum `mandatory:"false" json:"sourceType,omitempty"`

	// Additional details that are specific for this subscription such as activation details.
	AdditionalDetails []ExtendedMetadata `mandatory:"false" json:"additionalDetails"`

	// The realm from where customer is buying the subscription.
	Realm *string `mandatory:"false" json:"realm"`

	// The region from where customer is buying the subscription.
	Region *string `mandatory:"false" json:"region"`

	// Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Defined tags for this resource. Each key is predefined and scoped to a namespace.
	// For more information, see Resource Tags (https://docs.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m CreateSubscriptionDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreateSubscriptionDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingSourceTypeEnum(string(m.SourceType)); !ok && m.SourceType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceType: %s. Supported values are: %s.", m.SourceType, strings.Join(GetSourceTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
