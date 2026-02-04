// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
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

// MetricCollection Metric collection specification.
type MetricCollection struct {

	// The metric name for this metric collection.
	// A valid value starts with an alphabetical character and includes only
	// alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	MetricName *string `mandatory:"true" json:"metricName"`

	// Output field in the query to be used as the metric value.
	MetricQueryFieldName *string `mandatory:"true" json:"metricQueryFieldName"`

	// Selected dimension fields for the metric collection.
	Dimensions []DimensionField `mandatory:"true" json:"dimensions"`

	// Output table in the query.
	QueryTableName *string `mandatory:"false" json:"queryTableName"`
}

func (m MetricCollection) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricCollection) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
