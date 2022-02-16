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
	"fmt"
	"github.com/oracle/oci-go-sdk/v58/common"
	"strings"
)

// SqlPlanInsights Represents collection of SQL Plan Insights.
type SqlPlanInsights struct {

	// SQL Plan Insight text.
	// For example `Number of Plans Used`, `Most Executed Plan`,
	//   `Best Performing Plan`, `Worst Performing Plan`,
	//   `Plan With Most IO`,
	//   `Plan with Most CPU`
	Text *string `mandatory:"true" json:"text"`

	// SQL execution plan hash value for a given insight. For example `Most Executed Plan` insight will have value as "3975467901"
	Value *int64 `mandatory:"true" json:"value"`

	// SQL Insight category. For example PLANS_USED, MOST_EXECUTED, BEST_PERFORMER, WORST_PERFORMER, MOST_CPU or MOST_IO.
	Category *string `mandatory:"true" json:"category"`
}

func (m SqlPlanInsights) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m SqlPlanInsights) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
