// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Trace Explorer API
//
// Use the Application Performance Monitoring Trace Explorer API to query traces and associated spans in Trace Explorer. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmtraces

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// QueryResultMetadataSummary Summary containing the metadata about the query result set.
type QueryResultMetadataSummary struct {

	// A collection of QueryResultRowTypeSummary objects that describe the type and properties of the individual row elements of the query rows
	// being returned.  The i-th element in this list contains the QueryResultRowTypeSummary of the i-th key-value pair in the QueryResultRowData map.
	QueryResultRowTypeSummaries []QueryResultRowTypeSummary `mandatory:"false" json:"queryResultRowTypeSummaries"`

	// Source of the query result set (traces, spans, and so on).
	SourceName *string `mandatory:"false" json:"sourceName"`

	// Columns or attributes of the query rows  which are group by values.  This is a list of ResultsGroupedBy summary objects,
	// and the list will contain as many elements as the attributes and aggregate functions in the group by clause in the select query.
	QueryResultsGroupedBy []QueryResultsGroupedBySummary `mandatory:"false" json:"queryResultsGroupedBy"`

	// Order by which the query results are organized.  This is a list of queryResultsOrderedBy summary objects, and the list
	// will contain more than one OrderedBy summary object, if the sort was multidimensional.
	QueryResultsOrderedBy []QueryResultsOrderedBySummary `mandatory:"false" json:"queryResultsOrderedBy"`

	QueryResultsTopologyInfo *QueryResultsTopologyInfo `mandatory:"false" json:"queryResultsTopologyInfo"`

	// Interval for the time series function in minutes.
	TimeSeriesIntervalInMins *int `mandatory:"false" json:"timeSeriesIntervalInMins"`
}

func (m QueryResultMetadataSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryResultMetadataSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
