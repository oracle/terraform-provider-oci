// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the dimension of your choosing. The Usage API is used by the Cost Analysis tool in the Console. Also see Using the Usage API (https://docs.cloud.oracle.com/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UsageCarbonEmissionSummary The usage carbon emission store result.
type UsageCarbonEmissionSummary struct {

	// The usage start time.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The usage end time.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// The carbon emission in MTCO2 unit.
	ComputedCarbonEmission *float64 `mandatory:"true" json:"computedCarbonEmission"`

	// The method used to calculate carbon emission.
	EmissionCalculationMethod *string `mandatory:"true" json:"emissionCalculationMethod"`

	// The tenancy OCID.
	TenantId *string `mandatory:"false" json:"tenantId"`

	// The tenancy name.
	TenantName *string `mandatory:"false" json:"tenantName"`

	// The compartment OCID.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The compartment path, starting from root.
	CompartmentPath *string `mandatory:"false" json:"compartmentPath"`

	// The compartment name.
	CompartmentName *string `mandatory:"false" json:"compartmentName"`

	// The service name that is incurring the cost.
	Service *string `mandatory:"false" json:"service"`

	// The resource name that is incurring the cost.
	ResourceName *string `mandatory:"false" json:"resourceName"`

	// The resource OCID that is incurring the cost.
	ResourceId *string `mandatory:"false" json:"resourceId"`

	// The region of the usage.
	Region *string `mandatory:"false" json:"region"`

	// The availability domain of the usage.
	Ad *string `mandatory:"false" json:"ad"`

	// The SKU part number.
	SkuPartNumber *string `mandatory:"false" json:"skuPartNumber"`

	// The SKU friendly name.
	SkuName *string `mandatory:"false" json:"skuName"`

	// Platform for the cost.
	Platform *string `mandatory:"false" json:"platform"`

	// The subscription ID.
	SubscriptionId *string `mandatory:"false" json:"subscriptionId"`

	// For grouping, a tag definition. For filtering, a definition and key.
	Tags []Tag `mandatory:"false" json:"tags"`
}

func (m UsageCarbonEmissionSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UsageCarbonEmissionSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
