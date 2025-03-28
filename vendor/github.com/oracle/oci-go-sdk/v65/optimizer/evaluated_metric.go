// Copyright (c) 2016, 2018, 2025, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Cloud Advisor API
//
// Use the Cloud Advisor API to find potential inefficiencies in your tenancy and address them.
// Cloud Advisor can help you save money, improve performance, strengthen system resilience, and improve security.
// For more information, see Cloud Advisor (https://docs.oracle.com/iaas/Content/CloudAdvisor/Concepts/cloudadvisoroverview.htm).
//

package optimizer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// EvaluatedMetric One of the metrics that will be evaluated by profiles using this profile level.
type EvaluatedMetric struct {

	// The name of the metric (e.g., `CpuUtilization`).
	Name *string `mandatory:"true" json:"name"`

	// The name of the statistic (e.g., `p95`).
	Statistic *string `mandatory:"true" json:"statistic"`

	// The threshold that must be crossed for the recommendation to appear.
	Threshold *float64 `mandatory:"true" json:"threshold"`

	// Optional. The metric value that the recommendation will target.
	Target *float64 `mandatory:"false" json:"target"`
}

func (m EvaluatedMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m EvaluatedMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
