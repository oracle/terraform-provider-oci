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

// Trace Definition of a trace object.
type Trace struct {

	// Unique identifier (traceId) for the trace that represents the span set.  Note that this field is
	// defined as traceKey in the API and it maps to the traceId in the trace data in Application Performance
	// Monitoring.
	Key *string `mandatory:"true" json:"key"`

	// An array of spans in the trace.
	Spans []Span `mandatory:"true" json:"spans"`

	// Root span name associated with the trace. This is the flow start operation name.
	// Null is displayed if the root span is not yet completed.
	RootSpanOperationName *string `mandatory:"false" json:"rootSpanOperationName"`

	// Start time of the earliest span in the span collection.
	TimeEarliestSpanStarted *common.SDKTime `mandatory:"false" json:"timeEarliestSpanStarted"`

	// End time of the span that most recently ended in the span collection.
	TimeLatestSpanEnded *common.SDKTime `mandatory:"false" json:"timeLatestSpanEnded"`

	// The number of spans that have been processed by the system for the trace.  Note that there
	// could be additional spans that have not been processed or reported yet if the trace is still
	// in progress.
	SpanCount *int `mandatory:"false" json:"spanCount"`

	// The number of spans with errors that have been processed by the system for the trace.
	// Note that the number of spans with errors will be less than or equal to the total number of spans in the trace.
	ErrorSpanCount *int `mandatory:"false" json:"errorSpanCount"`

	// Service associated with the trace.
	RootSpanServiceName *string `mandatory:"false" json:"rootSpanServiceName"`

	// Start time of the root span for the span collection.
	TimeRootSpanStarted *common.SDKTime `mandatory:"false" json:"timeRootSpanStarted"`

	// End time of the root span for the span collection.
	TimeRootSpanEnded *common.SDKTime `mandatory:"false" json:"timeRootSpanEnded"`

	// Time taken for the root span operation to complete in milliseconds.
	RootSpanDurationInMs *int `mandatory:"false" json:"rootSpanDurationInMs"`

	// Time between the start of the earliest span and the end of the most recent span in milliseconds.
	TraceDurationInMs *int `mandatory:"false" json:"traceDurationInMs"`

	// Boolean flag that indicates whether the trace has an error.
	IsFault *bool `mandatory:"false" json:"isFault"`

	// The status of the trace.
	// The trace statuses are defined as follows:
	// complete - a root span has been recorded, but there is no information on the errors.
	// success - a complete root span is recorded there is a successful error type and error code - HTTP 200.
	// incomplete - the root span has not yet been received.
	// error - the root span returned with an error. There may or may not be an associated error code or error type.
	TraceStatus *string `mandatory:"false" json:"traceStatus"`

	// Error type of the trace.
	TraceErrorType *string `mandatory:"false" json:"traceErrorType"`

	// Error code of the trace.
	TraceErrorCode *string `mandatory:"false" json:"traceErrorCode"`

	// A summary of the spans by service.
	ServiceSummaries []TraceServiceSummary `mandatory:"false" json:"serviceSummaries"`

	// Source of trace (traces, syn_traces).
	SourceName TraceSourceNameEnum `mandatory:"false" json:"sourceName,omitempty"`

	SpanSummary *TraceSpanSummary `mandatory:"false" json:"spanSummary"`
}

func (m Trace) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Trace) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingTraceSourceNameEnum(string(m.SourceName)); !ok && m.SourceName != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for SourceName: %s. Supported values are: %s.", m.SourceName, strings.Join(GetTraceSourceNameEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// TraceSourceNameEnum Enum with underlying type: string
type TraceSourceNameEnum string

// Set of constants representing the allowable values for TraceSourceNameEnum
const (
	TraceSourceNameTraces    TraceSourceNameEnum = "TRACES"
	TraceSourceNameSynTraces TraceSourceNameEnum = "SYN_TRACES"
)

var mappingTraceSourceNameEnum = map[string]TraceSourceNameEnum{
	"TRACES":     TraceSourceNameTraces,
	"SYN_TRACES": TraceSourceNameSynTraces,
}

var mappingTraceSourceNameEnumLowerCase = map[string]TraceSourceNameEnum{
	"traces":     TraceSourceNameTraces,
	"syn_traces": TraceSourceNameSynTraces,
}

// GetTraceSourceNameEnumValues Enumerates the set of values for TraceSourceNameEnum
func GetTraceSourceNameEnumValues() []TraceSourceNameEnum {
	values := make([]TraceSourceNameEnum, 0)
	for _, v := range mappingTraceSourceNameEnum {
		values = append(values, v)
	}
	return values
}

// GetTraceSourceNameEnumStringValues Enumerates the set of values in String for TraceSourceNameEnum
func GetTraceSourceNameEnumStringValues() []string {
	return []string{
		"TRACES",
		"SYN_TRACES",
	}
}

// GetMappingTraceSourceNameEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingTraceSourceNameEnum(val string) (TraceSourceNameEnum, bool) {
	enum, ok := mappingTraceSourceNameEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
