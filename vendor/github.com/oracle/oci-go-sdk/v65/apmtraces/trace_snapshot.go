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

// TraceSnapshot Definition of a trace snapshot object.
type TraceSnapshot struct {

	// Unique identifier (traceId) for the trace that represents the span set.  Note that this field is
	// defined as traceKey in the API and it maps to the traceId in the trace data in Application Performance
	// Monitoring.
	Key *string `mandatory:"true" json:"key"`

	// List of spans.
	SpanSnapshots []SpanSnapshot `mandatory:"true" json:"spanSnapshots"`

	// Start time of the trace.
	TimeStarted *common.SDKTime `mandatory:"false" json:"timeStarted"`

	// End time of the trace.
	TimeEnded *common.SDKTime `mandatory:"false" json:"timeEnded"`

	// Trace snapshots properties.
	TraceSnapshotDetails []SnapshotDetail `mandatory:"false" json:"traceSnapshotDetails"`
}

func (m TraceSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TraceSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
