// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Stack Monitoring API
//
// Stack Monitoring API.
//

package stackmonitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AnomalyDataPoint anomaly evaluation result fo the data point
type AnomalyDataPoint struct {

	// timestamp of when the metric was collected
	Timestamp *common.SDKTime `mandatory:"true" json:"timestamp"`

	// value for the metric data point
	Value *float64 `mandatory:"true" json:"value"`

	// if the value is anomaly or not 0 indicates not an anomaly -1 indicates value is below the threshold +1 indicates value is above the threshold
	Anomaly *float64 `mandatory:"true" json:"anomaly"`

	// lower threshold for the metric value
	Low *float64 `mandatory:"false" json:"low"`

	// upper threshold for the metric value
	High *float64 `mandatory:"false" json:"high"`
}

func (m AnomalyDataPoint) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AnomalyDataPoint) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
