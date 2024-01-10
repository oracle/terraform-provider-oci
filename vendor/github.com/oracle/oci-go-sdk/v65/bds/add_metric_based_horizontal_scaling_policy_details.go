// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Big Data Service API
//
// REST API for Oracle Big Data Service. Use this API to build, deploy, and manage fully elastic Big Data Service clusters. Build on Hadoop, Spark and Data Science distributions, which can be fully integrated with existing enterprise data in Oracle Database and Oracle applications.
//

package bds

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AddMetricBasedHorizontalScalingPolicyDetails Details of a metric based horizontal autoscaling policy.
// In a metric-based autoscaling policy, an autoscaling action is triggered when a performance metric exceeds a threshold.
type AddMetricBasedHorizontalScalingPolicyDetails struct {
	ScaleOutConfig *MetricBasedHorizontalScaleOutConfig `mandatory:"false" json:"scaleOutConfig"`

	ScaleInConfig *MetricBasedHorizontalScaleInConfig `mandatory:"false" json:"scaleInConfig"`
}

func (m AddMetricBasedHorizontalScalingPolicyDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AddMetricBasedHorizontalScalingPolicyDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m AddMetricBasedHorizontalScalingPolicyDetails) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeAddMetricBasedHorizontalScalingPolicyDetails AddMetricBasedHorizontalScalingPolicyDetails
	s := struct {
		DiscriminatorParam string `json:"policyType"`
		MarshalTypeAddMetricBasedHorizontalScalingPolicyDetails
	}{
		"METRIC_BASED_HORIZONTAL_SCALING_POLICY",
		(MarshalTypeAddMetricBasedHorizontalScalingPolicyDetails)(m),
	}

	return json.Marshal(&s)
}
