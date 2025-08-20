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

// SubscriptionSku SKU information.
type SubscriptionSku struct {

	// Stock Keeping Unit (SKU) ID.
	Sku *string `mandatory:"true" json:"sku"`

	// Quantity of the stock units.
	Quantity *int `mandatory:"false" json:"quantity"`

	// Description of the stock units.
	Description *string `mandatory:"false" json:"description"`

	// Sales order line identifier.
	GsiOrderLineId *string `mandatory:"false" json:"gsiOrderLineId"`

	// Description of the covered product belonging to this SKU.
	LicensePartDescription *string `mandatory:"false" json:"licensePartDescription"`

	// Base metric for billing the service.
	MetricName *string `mandatory:"false" json:"metricName"`

	// Specifies if the SKU is considered as a parent or child.
	IsBaseServiceComponent *bool `mandatory:"false" json:"isBaseServiceComponent"`

	// Specifies if an additional test instance can be provisioned by the SaaS application.
	IsAdditionalInstance *bool `mandatory:"false" json:"isAdditionalInstance"`

	// Date and time when the SKU was created.
	StartDate *common.SDKTime `mandatory:"false" json:"startDate"`

	// Date and time when the SKU ended.
	EndDate *common.SDKTime `mandatory:"false" json:"endDate"`
}

func (m SubscriptionSku) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionSku) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
