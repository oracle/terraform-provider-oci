// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// MetricBasedHorizontalScaleInConfig Configration for a metric based horizontal scale-in policy.
type MetricBasedHorizontalScaleInConfig struct {
	Metric *AutoScalePolicyMetricRule `mandatory:"false" json:"metric"`

	// This value is the minimum number of nodes the cluster can be scaled-in to.
	MinNodeCount *int `mandatory:"false" json:"minNodeCount"`

	// This value is the number of nodes to remove during a scale-in event.
	StepSize *int `mandatory:"false" json:"stepSize"`
}

func (m MetricBasedHorizontalScaleInConfig) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m MetricBasedHorizontalScaleInConfig) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
