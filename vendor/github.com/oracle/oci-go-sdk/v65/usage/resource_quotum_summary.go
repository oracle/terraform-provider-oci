// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage Proxy API
//
// Use the Usage Proxy API to list Oracle Support Rewards, view related detailed usage information, and manage users who redeem rewards. For more information, see Oracle Support Rewards Overview (https://docs.cloud.oracle.com/iaas/Content/Billing/Concepts/supportrewardsoverview.htm).
//

package usage

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ResourceQuotumSummary The resource quota balance details.
type ResourceQuotumSummary struct {

	// The resource name.
	Name *string `mandatory:"false" json:"name"`

	// Used to indicate if further quota consumption isAllowed.
	IsAllowed *bool `mandatory:"false" json:"isAllowed"`

	// The quota limit.
	Limit *float64 `mandatory:"false" json:"limit"`

	// The quota balance.
	Balance *float64 `mandatory:"false" json:"balance"`

	// Used to indicate if overages are incurred.
	IsOverage *bool `mandatory:"false" json:"isOverage"`

	// The purchased quota limit.
	PurchasedLimit *float64 `mandatory:"false" json:"purchasedLimit"`

	// The service name.
	Service *string `mandatory:"false" json:"service"`

	// Used to indicate any resource dependencies.
	IsDependency *bool `mandatory:"false" json:"isDependency"`

	// The affected resource name.
	AffectedResource *string `mandatory:"false" json:"affectedResource"`
}

func (m ResourceQuotumSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ResourceQuotumSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
