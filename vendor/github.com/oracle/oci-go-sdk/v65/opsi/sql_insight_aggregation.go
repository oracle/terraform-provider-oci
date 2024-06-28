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

// SqlInsightAggregation Represents a SQL Insight.
type SqlInsightAggregation struct {

	// Insight text.
	// For example `Degrading SQLs`, `Variant SQLs`,
	//   `Inefficient SQLs`, `Improving SQLs`, `SQLs with Plan Changes`,
	//   `Degrading SQLs have increasing IO Time above 50%`,
	//   `Degrading SQLs are variant`,
	//   `2 of the 2 variant SQLs have plan changes`,
	//   `Inefficient SQLs have increasing CPU Time above 50%
	Text *string `mandatory:"true" json:"text"`

	// SQL counts for a given insight. For example insight text `2 of 10 SQLs have degrading response time` will have values as [2,10]"
	Values []int `mandatory:"true" json:"values"`

	// Insight category. It would be one of the following
	// DEGRADING,
	// VARIANT,
	// INEFFICIENT,
	// CHANGING_PLANS,
	// IMPROVING,
	// DEGRADING_VARIANT,
	// DEGRADING_INEFFICIENT,
	// DEGRADING_CHANGING_PLANS,
	// DEGRADING_INCREASING_IO,
	// DEGRADING_INCREASING_CPU,
	// DEGRADING_INCREASING_INEFFICIENT_WAIT,
	// DEGRADING_CHANGING_PLANS_AND_INCREASING_IO,
	// DEGRADING_CHANGING_PLANS_AND_INCREASING_CPU,
	// DEGRADING_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT,VARIANT_INEFFICIENT,
	// VARIANT_CHANGING_PLANS,
	// VARIANT_INCREASING_IO,
	// VARIANT_INCREASING_CPU,
	// VARIANT_INCREASING_INEFFICIENT_WAIT,
	// VARIANT_CHANGING_PLANS_AND_INCREASING_IO,
	// VARIANT_CHANGING_PLANS_AND_INCREASING_CPU,
	// VARIANT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT,
	// INEFFICIENT_CHANGING_PLANS,
	// INEFFICIENT_INCREASING_INEFFICIENT_WAIT,
	// INEFFICIENT_CHANGING_PLANS_AND_INCREASING_INEFFICIENT_WAIT
	Category *string `mandatory:"true" json:"category"`
}

func (m SqlInsightAggregation) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlInsightAggregation) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
