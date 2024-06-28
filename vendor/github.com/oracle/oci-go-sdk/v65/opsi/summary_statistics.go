// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Ops Insights API
//
// Use the Ops Insights API to perform data extraction operations to obtain database
// resource utilization, performance statistics, and reference information. For more information,
// see About Oracle Cloud Infrastructure Ops Insights (https://docs.cloud.oracle.com/en-us/iaas/operations-insights/doc/operations-insights.html).
//

package opsi

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// SummaryStatistics Contains common summary statistics.
type SummaryStatistics struct {

	// The smallest number in the data set.
	Minimum *float64 `mandatory:"true" json:"minimum"`

	// The largest number in the data set.
	Maximum *float64 `mandatory:"true" json:"maximum"`

	// The average number in the data set.
	Average *float64 `mandatory:"true" json:"average"`

	// The middle number in the data set.
	Median *float64 `mandatory:"true" json:"median"`

	// The middle number between the smallest number and the median of the data set. It's also known as the 25th quartile.
	LowerQuartile *float64 `mandatory:"true" json:"lowerQuartile"`

	// The middle number between the median and the largest number of the data set. It's also known as the 75th quartile.
	UpperQuartile *float64 `mandatory:"true" json:"upperQuartile"`
}

func (m SummaryStatistics) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SummaryStatistics) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
