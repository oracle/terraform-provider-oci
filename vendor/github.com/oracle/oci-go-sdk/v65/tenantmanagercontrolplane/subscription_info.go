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

// SubscriptionInfo A single subscription's details.
type SubscriptionInfo struct {

	// Subscription ID.
	SpmSubscriptionId *string `mandatory:"true" json:"spmSubscriptionId"`

	// Subscription service name.
	Service *string `mandatory:"true" json:"service"`

	// Subscription start date. An RFC 3339-formatted date and time string.
	StartDate *common.SDKTime `mandatory:"true" json:"startDate"`

	// Subscription end date. An RFC 3339-formatted date and time string.
	EndDate *common.SDKTime `mandatory:"true" json:"endDate"`

	// List of SKUs the subscription contains.
	Skus []Sku `mandatory:"true" json:"skus"`
}

func (m SubscriptionInfo) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SubscriptionInfo) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
