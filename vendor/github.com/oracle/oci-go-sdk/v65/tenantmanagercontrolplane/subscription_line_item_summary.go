// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Organizations API
//
// Use the Organizations API to consolidate multiple OCI tenancies into an organization, and centrally manage your tenancies and organization resources. For more information, see Organization Management Overview (https://docs.oracle.com/iaas/Content/General/Concepts/organization_management_overview.htm).
//

package tenantmanagercontrolplane

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SubscriptionLineItemSummary Summary of line items in a subscription.
type SubscriptionLineItemSummary struct {

	// Subscription line item identifier.
	Id *string `mandatory:"true" json:"id"`

	// Product code.
	ProductCode *string `mandatory:"true" json:"productCode"`

	// Product number.
	Quantity *float32 `mandatory:"true" json:"quantity"`

	// Billing model supported by the associated line item.
	BillingModel BillingModelEnum `mandatory:"true" json:"billingModel"`

	// The time the subscription item and associated products should start. An RFC 3339 formatted date and time string.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// The time the subscription item and associated products should end. An RFC 3339 formatted date and time string.
	TimeEnded *common.SDKTime `mandatory:"true" json:"timeEnded"`

	// Usage of system tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"orcl-cloud": {"free-tier-retained": "true"}}`
	SystemTags map[string]map[string]interface{} `mandatory:"false" json:"systemTags"`
}

func (m SubscriptionLineItemSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionLineItemSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingBillingModelEnum(string(m.BillingModel)); !ok && m.BillingModel != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for BillingModel: %s. Supported values are: %s.", m.BillingModel, strings.Join(GetBillingModelEnumStringValues(), ",")))
	}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
