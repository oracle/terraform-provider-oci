// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Application Performance Monitoring Synthetic Monitoring API
//
// Use the Application Performance Monitoring Synthetic Monitoring API to query synthetic scripts and monitors. For more information, see Application Performance Monitoring (https://docs.oracle.com/iaas/application-performance-monitoring/index.html).
//

package apmsynthetics

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
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

	// The name of the public or dedicated vantage point.
	VantagePoint *string `mandatory:"false" json:"vantagePoint"`

	// The specific point of time when the result of an execution is collected.
	ExecutionTime *string `mandatory:"false" json:"executionTime"`
}

func (m MonitorResult) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MonitorResult) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
