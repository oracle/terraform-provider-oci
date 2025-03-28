// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.oracle.com/iaas/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// ProjectedDataItem The timestamp of the projected event and their corresponding resource value.
// `highValue` and `lowValue` are the uncertainty bounds of the corresponding value.
type ProjectedDataItem struct {

	// The timestamp in which the current sampling period ends in RFC 3339 format.
	EndTimestamp *common.SDKTime `mandatory:"true" json:"endTimestamp"`

	// Total amount used of the resource metric type (CPU, STORAGE).
	Usage *float64 `mandatory:"true" json:"usage"`

	// Upper uncertainty bound of the current usage value.
	HighValue *float64 `mandatory:"true" json:"highValue"`

	// Lower uncertainty bound of the current usage value.
	LowValue *float64 `mandatory:"true" json:"lowValue"`
}

func (m ProjectedDataItem) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m ProjectedDataItem) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
