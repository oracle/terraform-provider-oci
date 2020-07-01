// Copyright (c) 2016, 2018, 2020, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// A description of the UsageApi API.
//

package usageapi

import (
	"github.com/oracle/oci-go-sdk/common"
)

// UsageSummary The result from usage store.
type UsageSummary struct {

	// The start time of the usage.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The end time of the usage.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// The OCID of the compartment.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The path of the compartment, starting from root.
	CompartmentPath *string `mandatory:"false" json:"compartmentPath"`

	// The name of the compartment.
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// The name of the service that is incurring the cost.
	Service *string `mandatory:"false" json:"service"`

	// The name of the resource that is incurring the cost.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The Ocid of the resource that is incurring the cost.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The region of the usage.
	Region *string `mandatory:"false" json:"region"`

	// The availability domain of the usage.
	Ad *string `mandatory:"false" json:"ad"`

	// The size of resource being metered.
	Weight *float32 `mandatory:"false" json:"weight"`

	// The shape of the resource.
	Shape *string `mandatory:"false" json:"shape"`

	// The part number of the SKU.
	SkuPartNumber *string `mandatory:"false" json:"skuPartNumber"`

	// The friendly name for the SKU.
	SkuName *string `mandatory:"false" json:"skuName"`

	// The unit of the usage.
	Unit *string `mandatory:"false" json:"unit"`

	// The discretionary discount applied to the SKU.
	Discount *float32 `mandatory:"false" json:"discount"`

	// The list rate for the SKU (not discount).
	ListRate *float32 `mandatory:"false" json:"listRate"`

	// Platform for the cost.
	Platform *string `mandatory:"false" json:"platform"`

	// The computed cost.
	ComputedAmount *float32 `mandatory:"false" json:"computedAmount"`

	// The usage number.
	ComputedQuantity *float32 `mandatory:"false" json:"computedQuantity"`

	// The SPM OverageFlag.
	OveragesFlag *string `mandatory:"false" json:"overagesFlag"`

	// The price per unit.
	UnitPrice *float32 `mandatory:"false" json:"unitPrice"`

	// The currency for the price.
	Currency *string `mandatory:"false" json:"currency"`

	// The subscription Id.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// The overage usage.
	Overage *string `mandatory:"false" json:"overage"`

	// For grouping, a tag definition. For filtering, a definition and key
	Tags []Tag `mandatory:"false" json:"tags"`
}

func (m UsageSummary) String() string {
	return common.PointerString(m)
}
