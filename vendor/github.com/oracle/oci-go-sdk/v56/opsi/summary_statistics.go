// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
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
	"github.com/oracle/oci-go-sdk/v56/common"
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
