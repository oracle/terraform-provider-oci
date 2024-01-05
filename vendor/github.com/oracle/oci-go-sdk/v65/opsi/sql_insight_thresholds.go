// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SqlInsightThresholds Inventory details.
type SqlInsightThresholds struct {

	// Degradation Percent Threshold is used to derive degrading SQLs.
	DegradationInPct *int `mandatory:"true" json:"degradationInPct"`

	// Variability Percent Threshold is used to derive variant SQLs.
	Variability *float32 `mandatory:"true" json:"variability"`

	// Inefficiency Percent Threshold is used to derive inefficient SQLs.
	InefficiencyInPct *int `mandatory:"true" json:"inefficiencyInPct"`

	// PctIncreaseInIO is used for deriving insights for SQLs which are degrading or
	// variant or inefficient. And these SQLs should also have increasing change in IO Time
	// beyond threshold. Insights are derived using linear regression.
	IncreaseInIOInPct *int `mandatory:"true" json:"increaseInIOInPct"`

	// PctIncreaseInCPU is used for deriving insights for SQLs which are degrading or
	// variant or inefficient. And these SQLs should also have increasing change in CPU Time
	// beyond threshold. Insights are derived using linear regression.
	IncreaseInCPUInPct *int `mandatory:"true" json:"increaseInCPUInPct"`

	// PctIncreaseInIO is used for deriving insights for SQLs which are degrading or
	// variant or inefficient. And these SQLs should also have increasing change in
	// Other Wait Time beyond threshold. Insights are derived using linear regression.
	IncreaseInInefficientWaitInPct *int `mandatory:"true" json:"increaseInInefficientWaitInPct"`

	// Improved Percent Threshold is used to derive improving SQLs.
	ImprovedInPct *int `mandatory:"true" json:"improvedInPct"`
}

func (m SqlInsightThresholds) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlInsightThresholds) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
