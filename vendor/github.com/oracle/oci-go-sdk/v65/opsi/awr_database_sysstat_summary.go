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

// AwrDatabaseSysstatSummary The summary of the AWR SYSSTAT data.
type AwrDatabaseSysstatSummary struct {

	// The name of the SYSSTAT.
	Name *string `mandatory:"true" json:"name"`

	// The name of the SYSSTAT category.
	Category *string `mandatory:"false" json:"category"`

	// The start time of the SYSSTAT.
	TimeBegin *common.SDKTime `mandatory:"false" json:"timeBegin"`

	// The end time of the SYSSTAT.
	TimeEnd *common.SDKTime `mandatory:"false" json:"timeEnd"`

	// The average value of the SYSSTAT. The units are stats name/val per the time period {timeBegin - timeEnd}.
	AvgValue *float64 `mandatory:"false" json:"avgValue"`

	// The last value of the SYSSTAT. The units are stats name/val per the time period {timeBegin - timeEnd}.
	CurrentValue *float64 `mandatory:"false" json:"currentValue"`
}

func (m AwrDatabaseSysstatSummary) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AwrDatabaseSysstatSummary) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
