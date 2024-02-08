// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Operations Insights API
//
// Use the Operations Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Operations Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// DataObjectQueryTimeFilters Time filters to be applied in the data object query.
type DataObjectQueryTimeFilters struct {

	// Specify time period in ISO 8601 format with respect to current time.
	// Default is last 30 days represented by P30D.
	// If timePeriod is specified, then timeStart and timeEnd will be ignored.
	// Examples: P90D (last 90 days), P4W (last 4 weeks), P2M (last 2 months), P1Y (last 12 months).
	TimePeriod *string `mandatory:"false" json:"timePeriod"`

	// Start time in UTC in RFC3339 formatted datetime string. Example: 2021-10-30T00:00:00.000Z.
	// timeStart and timeEnd are used together. If timePeriod is specified, this parameter is ignored.
	TimeStart *common.SDKTime `mandatory:"false" json:"timeStart"`

	// End time in UTC in RFC3339 formatted datetime string. Example: 2021-10-30T00:00:00.000Z.
	// timeStart and timeEnd are used together. If timePeriod is specified, this parameter is ignored.
	// If timeEnd is not specified, current time is used as timeEnd.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`
}

func (m DataObjectQueryTimeFilters) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m DataObjectQueryTimeFilters) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
