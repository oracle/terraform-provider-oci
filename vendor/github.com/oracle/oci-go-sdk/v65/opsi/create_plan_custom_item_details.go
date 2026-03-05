// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CreatePlanCustomItemDetails Custom configuration item details for a chargeback plan. Example items for Exadata Insights Chargeback are
// statistic(default value AVG), percentile, infrastructureCost, infrastructurePlanType, additionalServerCost and additionalServerPlanType.
type CreatePlanCustomItemDetails struct {

	// Name of chargeback plan customization item. Example items for Exadata Insights Chargeback are statistic, percentile, infrastructureCost, additionalServerCost etc.
	Name *string `mandatory:"false" json:"name"`

	// Value of chargeback plan customization item.
	Value *string `mandatory:"false" json:"value"`

	// Indicates whether the chargeback plan customization item can be customized.
	IsCustomizable *bool `mandatory:"false" json:"isCustomizable"`
}

func (m CreatePlanCustomItemDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CreatePlanCustomItemDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
