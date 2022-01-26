// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// LogAnalytics API
//
// The LogAnalytics API for the LogAnalytics service.
//

package loganalytics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
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
