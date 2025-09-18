// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Database Service API
//
// The API for the Database Service. Use this API to manage resources such as databases and DB Systems. For more information, see Overview of the Database Service (https://docs.oracle.com/iaas/Content/Database/Concepts/databaseoverview.htm).
//

package database

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EstimateCostSavingSummary Autonomous AI Database cost savings.
type EstimateCostSavingSummary struct {

	// The epoch time at which cost aggregation starts.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// The epoch time at which cost aggregation ends.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Indicates if CPU autoscaling is applied.
	IsCpuAutoscale *bool `mandatory:"false" json:"isCpuAutoscale"`

	// CPU cost for a given time period under regular billing plan, in ECPU hours.
	EstimatedUsageWithoutElasticPool *int64 `mandatory:"false" json:"estimatedUsageWithoutElasticPool"`

	// CPU cost for a given time period under elastic pool billing plan, in ECPU hours.
	UsageWithElasticPool *int64 `mandatory:"false" json:"usageWithElasticPool"`

	// Estimated cost savings in percentage with elastic pool utilization.
	CostSavingsWithElasticPool *float64 `mandatory:"false" json:"costSavingsWithElasticPool"`
}

func (m EstimateCostSavingSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EstimateCostSavingSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
