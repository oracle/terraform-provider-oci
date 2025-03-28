// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
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

// SpanSnapshot Definition of a span snapshot object.
type SpanSnapshot struct {

	// Unique identifier (spanId) for the trace span.
	Key *string `mandatory:"true" json:"key"`

	// Start time of the span.
	TimeStarted *common.SDKTime `mandatory:"true" json:"timeStarted"`

	// End time of the span.
	TimeEnded *common.SDKTime `mandatory:"true" json:"timeEnded"`

	// Span name associated with the trace.
	SpanName *string `mandatory:"false" json:"spanName"`

	// Span snapshots properties.
	SpanSnapshotDetails []SnapshotDetail `mandatory:"false" json:"spanSnapshotDetails"`

	// Thread snapshots.
	ThreadSnapshots []ThreadSnapshot `mandatory:"false" json:"threadSnapshots"`

	// An array of child span snapshots.
	Children []SpanSnapshot `mandatory:"false" json:"children"`
}

func (m SpanSnapshot) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SpanSnapshot) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
