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

// QueryDetails Input arguments for running a log anlaytics query. If the request is set to run in asynchronous mode
// then shouldIncludeColumns and shouldIncludeFields can be overwritten when retrieving the results.
type QueryDetails struct {

	// Compartment Identifier OCID  (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// Query to perform. Must conform to logging analytic querylanguage syntax. Syntax errors will be returned if present.
	QueryString *string `mandatory:"true" json:"queryString"`

	// Default subsystem to qualify fields with in the queryString if not specified.
	SubSystem SubSystemNameEnum `mandatory:"true" json:"subSystem"`

	// Flag to search all child compartments of the compartment Id specified in the compartmentId query parameter.
	CompartmentIdInSubtree *bool `mandatory:"false" json:"compartmentIdInSubtree"`

	// Saved search OCID for this query if known.
	SavedSearchId *string `mandatory:"false" json:"savedSearchId"`

	// Maximum number of results to count.  Note a maximum of 2001 will be enforced; that is, actualMaxTotalCountUsed = Math.min(maxTotalCount, 2001).
	MaxTotalCount *int `mandatory:"false" json:"maxTotalCount"`

	TimeFilter *TimeRange `mandatory:"false" json:"timeFilter"`

	// List of filters to be applied when the query executes. More than one filter per field is not permitted.
	ScopeFilters []ScopeFilter `mandatory:"false" json:"scopeFilters"`

	// Amount of time, in seconds, allowed for a query to execute. If this time expires before the query is complete, any partial results will be returned.
	QueryTimeoutInSeconds *int `mandatory:"false" json:"queryTimeoutInSeconds"`

	// Option to run the query asynchronously. This will lead to a LogAnalyticsQueryJobWorkRequest being submitted and the {workRequestId} will be returned to use for fetching the results.
	ShouldRunAsync *bool `mandatory:"false" json:"shouldRunAsync"`

	// Execution mode for the query if running asynchronously i.e (shouldRunAsync is set to true).
	AsyncMode JobModeEnum `mandatory:"false" json:"asyncMode,omitempty"`

	// Include the total number of results from the query. Note, this value will always be equal to or less than maxTotalCount.
	ShouldIncludeTotalCount *bool `mandatory:"false" json:"shouldIncludeTotalCount"`

	// Include columns in response
	ShouldIncludeColumns *bool `mandatory:"false" json:"shouldIncludeColumns"`

	// Include fields in response
	ShouldIncludeFields *bool `mandatory:"false" json:"shouldIncludeFields"`

	// Controls if query should ignore pre-calculated results if available and only use raw data. If set and no acceleration data is found it will fallback to raw data.
	ShouldUseAcceleration *bool `mandatory:"false" json:"shouldUseAcceleration"`
}

func (m QueryDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m QueryDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingSubSystemNameEnum(string(m.SubSystem)); !ok && m.SubSystem != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SubSystem: %s. Supported values are: %s.", m.SubSystem, strings.Join(GetSubSystemNameEnumStringValues(), ",")))
	}

	if _, ok := GetMappingJobModeEnum(string(m.AsyncMode)); !ok && m.AsyncMode != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for AsyncMode: %s. Supported values are: %s.", m.AsyncMode, strings.Join(GetJobModeEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
