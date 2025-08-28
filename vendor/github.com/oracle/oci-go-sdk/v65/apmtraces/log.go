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

// Log Definition of a log object.
type Log struct {

	// Unique identifier (logId) for the logKey.  Note that this field is
	// defined as logKey in the API and it maps to the logId in Application Performance Monitoring.
	LogKey *string `mandatory:"true" json:"logKey"`

	// Time used by the time picker (RecordedTime).  Either the timeCreated if present or the timeObserved.
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// Name of the event.
	EventName *string `mandatory:"false" json:"eventName"`

	// Unique identifier for the trace (traceId) associated with this log.
	TraceKey *string `mandatory:"false" json:"traceKey"`

	// Unique identifier for the span (spanId) associated with this log.
	SpanKey *string `mandatory:"false" json:"spanKey"`

	// Trace flags.
	TraceFlags *int `mandatory:"false" json:"traceFlags"`

	// Time that the log event occurred (CreatedTime).
	TimeCreated *common.SDKTime `mandatory:"false" json:"timeCreated"`

	// Time that the log was received by apm (ObservedTime).
	TimeObserved *common.SDKTime `mandatory:"false" json:"timeObserved"`

	// Log Severity text (SeverityText).  Also known as Log level.
	SeverityText *string `mandatory:"false" json:"severityText"`

	// Log Severity number (SeverityNumber).
	SeverityNumber *int `mandatory:"false" json:"severityNumber"`

	// Log body (Body).
	Body *string `mandatory:"false" json:"body"`

	// Full values for attributes that are too long to be stored as a log attribute (Overflow).
	OverflowAttributes *string `mandatory:"false" json:"overflowAttributes"`

	// List of attributes associated with the logs.
	Attributes []Attribute `mandatory:"false" json:"attributes"`

	// Metadata about the attributes in the logs.
	AttributeMetadata map[string]AttributeMetadata `mandatory:"false" json:"attributeMetadata"`
}

func (m Log) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Log) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}
