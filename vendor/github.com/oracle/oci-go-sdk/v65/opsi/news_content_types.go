// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// NewsContentTypes Content types that the news report can handle.
type NewsContentTypes struct {

	// Supported resources for capacity planning content type.
	CapacityPlanningResources []NewsContentTypesResourceEnum `mandatory:"false" json:"capacityPlanningResources"`

	// Supported resources for SQL insights - fleet analysis content type.
	SqlInsightsFleetAnalysisResources []NewsSqlInsightsContentTypesResourceEnum `mandatory:"false" json:"sqlInsightsFleetAnalysisResources"`

	// Supported resources for SQL insights - plan changes content type.
	SqlInsightsPlanChangesResources []NewsSqlInsightsContentTypesResourceEnum `mandatory:"false" json:"sqlInsightsPlanChangesResources"`

	// Supported resources for SQL insights - top databases content type.
	SqlInsightsTopDatabasesResources []NewsSqlInsightsContentTypesResourceEnum `mandatory:"false" json:"sqlInsightsTopDatabasesResources"`

	// Supported resources for SQL insights - top SQL by insights content type.
	SqlInsightsTopSqlByInsightsResources []NewsSqlInsightsContentTypesResourceEnum `mandatory:"false" json:"sqlInsightsTopSqlByInsightsResources"`

	// Supported resources for SQL insights - top SQL content type.
	SqlInsightsTopSqlResources []NewsSqlInsightsContentTypesResourceEnum `mandatory:"false" json:"sqlInsightsTopSqlResources"`

	// Supported resources for SQL insights - performance degradation content type.
	SqlInsightsPerformanceDegradationResources []NewsSqlInsightsContentTypesResourceEnum `mandatory:"false" json:"sqlInsightsPerformanceDegradationResources"`

	// Supported resources for actionable insights content type.
	ActionableInsightsResources []ActionableInsightsContentTypesResourceEnum `mandatory:"false" json:"actionableInsightsResources"`
}

func (m NewsContentTypes) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m NewsContentTypes) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
