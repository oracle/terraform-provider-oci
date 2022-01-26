// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors.
//

package apmsynthetics

import (
	"github.com/oracle/oci-go-sdk/v56/common"
)

// MonitorResult The monitor result for a specific execution.
type MonitorResult struct {

	// Type of result content.
	// Example: Zip or Raw file.
	ResultContentType *string `mandatory:"true" json:"resultContentType"`

	// Type of result.
	// Example: HAR, Screenshot, Log or Network.
	ResultType *string `mandatory:"false" json:"resultType"`

	// Monitor result data set.
	ResultDataSet []MonitorResultData `mandatory:"false" json:"resultDataSet"`

	// The OCID (https://docs.cloud.oracle.com/Content/General/Concepts/identifiers.htm) of the monitor.
	MonitorId *string `mandatory:"false" json:"monitorId"`

	// The name of the vantage point.
	VantagePoint *string `mandatory:"false" json:"vantagePoint"`

	// The specific point of time when the result of an execution is collected.
	ExecutionTime *string `mandatory:"false" json:"executionTime"`
}

func (m MonitorResult) String() string {
	return common.PointerString(m)
}
