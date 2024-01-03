// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LogAnalyticsSourceMetric LogAnalyticsSourceMetric
type LogAnalyticsSourceMetric struct {

	// A flag specifying whether or not the metric source is enabled.
	IsMetricSourceEnabled *bool `mandatory:"false" json:"isMetricSourceEnabled"`

	// The metric name.
	MetricName *string `mandatory:"false" json:"metricName"`

	// The source internal name.
	SourceName *string `mandatory:"false" json:"sourceName"`

	// The entity type.
	EntityType *string `mandatory:"false" json:"entityType"`
}

func (m LogAnalyticsSourceMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LogAnalyticsSourceMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
