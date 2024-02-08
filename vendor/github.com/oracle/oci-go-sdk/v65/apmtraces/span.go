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

// Span Definition of a span object.
type Span struct {

	// Unique identifier (spanId) for the span.  Note that this field is
	// defined as spanKey in the API and it maps to the spanId in the trace data
	// in Application Performance Monitoring.
	Key *string `mandatory:"true" json:"key"`

	// Unique identifier for the trace.
	TraceKey *string `mandatory:"true" json:"traceKey"`

	// Span start time.  Timestamp when the span was started.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// Span end time.  Timestamp when the span was completed.
	TimeEnded *common.SDKTime `mandatory:"true" json:"timeEnded"`

	// Total span duration in milliseconds.
	DurationInMs *int64 `mandatory:"true" json:"durationInMs"`

	// Span name associated with the trace.  This is usually the method or URI of the request.
	OperationName *string `mandatory:"true" json:"operationName"`

	// Indicates if the span has an error.
	IsError *bool `mandatory:"true" json:"isError"`

	// Unique parent identifier for the span if one exists. For root spans this will be null.
	ParentSpanKey *string `mandatory:"false" json:"parentSpanKey"`

	// Service name associated with the span.
	ServiceName *string `mandatory:"false" json:"serviceName"`

	// Kind associated with the span.
	Kind *string `mandatory:"false" json:"kind"`

	// List of tags associated with the span.
	Tags []Tag `mandatory:"false" json:"tags"`

	// List of logs associated with the span.
	Logs []SpanLogCollection `mandatory:"false" json:"logs"`
}

func (m Span) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Span) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
