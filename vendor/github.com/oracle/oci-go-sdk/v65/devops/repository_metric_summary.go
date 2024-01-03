// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// DevOps API
//
// Use the DevOps API to create DevOps projects, configure code repositories,  add artifacts to deploy, build and test software applications, configure  target deployment environments, and deploy software applications.  For more information, see DevOps (https://docs.cloud.oracle.com/Content/devops/using/home.htm).
//

package devops

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// RepositoryMetricSummary The metrics details of a repository resource.
type RepositoryMetricSummary struct {

	// Type of metric
	MetricName MetricNameEnum `mandatory:"false" json:"metricName,omitempty"`

	// The duration of the returned aggregated data in seconds.
	DurationInSeconds *int64 `mandatory:"false" json:"durationInSeconds"`

	// The qualifiers provided in the definition of the returned metric.
	Dimensions map[string]string `mandatory:"false" json:"dimensions"`

	// The start time associated with the value of the metric.
	StartTimestampInEpochSeconds *int64 `mandatory:"false" json:"startTimestampInEpochSeconds"`

	// Represents the total number of the metric being calculated.
	Count *float64 `mandatory:"false" json:"count"`

	// Represents the total duration in days calculated corresponding to the total no. of PRs.
	// This is used only for "PULL_REQUEST_REVIEW_START_DURATION_IN_DAYS" and "PULL_REQUEST_REVIEW_DURATION_IN_DAYS" metrics.
	Sum *float64 `mandatory:"false" json:"sum"`
}

func (m RepositoryMetricSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m RepositoryMetricSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingMetricNameEnum(string(m.MetricName)); !ok && m.MetricName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MetricName: %s. Supported values are: %s.", m.MetricName, strings.Join(GetMetricNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
