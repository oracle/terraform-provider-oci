// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Usage API
//
// Use the Usage API to view your Oracle Cloud usage and costs. The API allows you to request data that meets the specified filter criteria, and to group that data by the chosen dimension. The Usage API is used by Cost Analysis (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm), Scheduled Reports (https://docs.oracle.com/iaas/Content/Billing/Concepts/scheduledreportoverview.htm), and Carbon Emissions Analysis (https://docs.oracle.com/iaas/Content/General/Concepts/emissions-management.htm) in the Console. Also see Using the Usage API (https://docs.oracle.com/iaas/Content/Billing/Concepts/costanalysisoverview.htm#cost_analysis_using_the_api) for more information.
//

package usageapi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UsageCarbonEmissionSummary The carbon emission usage store result.
type UsageCarbonEmissionSummary struct {

	// The usage start time.
	TimeUsageStarted *common.SDKTime `mandatory:"true" json:"timeUsageStarted"`

	// The usage end time.
	TimeUsageEnded *common.SDKTime `mandatory:"true" json:"timeUsageEnded"`

	// The carbon emission usage in MTCO2 units.
	ComputedCarbonEmission *float64 `mandatory:"true" json:"computedCarbonEmission"`

	// Specifies the approach for calculating carbon emissions, supports both SPEND_BASED (based on expenditure data) and POWER_BASED (based on power consumption, newly introduced in the metering pipeline)
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

	// The emission type, such as MARKET_BASED or LOCATION_BASED.
	EmissionType RequestUsageCarbonEmissionsDetailsEmissionTypeEnum `mandatory:"false" json:"emissionType,omitempty"`

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

	if _, ok := GetMappingRequestUsageCarbonEmissionsDetailsEmissionTypeEnum(string(m.EmissionType)); !ok && m.EmissionType != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for EmissionType: %s. Supported values are: %s.", m.EmissionType, strings.Join(GetRequestUsageCarbonEmissionsDetailsEmissionTypeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
