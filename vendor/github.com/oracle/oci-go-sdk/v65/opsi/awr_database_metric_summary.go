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

// AwrDatabaseMetricSummary The summary of the AWR metric data for a particular metric at a specific time.
type AwrDatabaseMetricSummary struct {

	// The name of the metric.
	Name *string `mandatory:"true" json:"name"`

	// The time of the sampling.
	Timestamp *common.SDKTime `mandatory:"false" json:"timestamp"`

	// The average value of the sampling period.
	AvgValue *float64 `mandatory:"false" json:"avgValue"`

	// The minimum value of the sampling period.
	MinValue *float64 `mandatory:"false" json:"minValue"`

	// The maximum value of the sampling period.
	MaxValue *float64 `mandatory:"false" json:"maxValue"`
}

func (m AwrDatabaseMetricSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseMetricSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
