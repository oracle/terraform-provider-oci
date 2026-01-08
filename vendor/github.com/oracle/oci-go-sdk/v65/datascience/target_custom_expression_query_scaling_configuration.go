// Copyright (c) 2016, 2018, 2026, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Data Science API
//
// Use the Data Science API to organize your data science work, access data and computing resources, and build, train, deploy and manage models and model deployments. For more information, see Data Science (https://docs.oracle.com/iaas/data-science/using/data-science.htm).
//

package datascience

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// TargetCustomExpressionQueryScalingConfiguration The scaling configuration for the custom expression query for the workload.
type TargetCustomExpressionQueryScalingConfiguration struct {

	// The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of the Monitoring service
	// interprets results for each returned time series as Boolean values, where zero represents false and a non-zero value
	// represents true. A true value means that the trigger rule condition has been met. The query must specify a metric,
	// statistic, interval, and trigger rule (threshold or absence). Supported values for interval: 1m-60m (also 1h).
	// You can optionally specify dimensions and grouping functions. Supported grouping functions: grouping(), groupBy().
	// Example of threshold alarm:
	//   ```
	//   CPUUtilization[1m]{resourceId = "Model_Deployment_OCID"}.grouping().mean() < 25
	//   CPUUtilization[1m]{resourceId = "Model_Deployment_OCID"}.grouping().mean() > 75
	//   ```
	Query *string `mandatory:"true" json:"query"`

	// A metric value at which the scaling operation will be triggered.
	Threshold *float32 `mandatory:"true" json:"threshold"`

	// Namespace to read the metrics from. Default value will be service metric namespace.
	MetricNamespace *string `mandatory:"false" json:"metricNamespace"`
}

func (m TargetCustomExpressionQueryScalingConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m TargetCustomExpressionQueryScalingConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf("%s", strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m TargetCustomExpressionQueryScalingConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeTargetCustomExpressionQueryScalingConfiguration TargetCustomExpressionQueryScalingConfiguration
	s := struct {
		DiscriminatorParam string `json:"targetScalingConfigurationType"`
		MarshalTypeTargetCustomExpressionQueryScalingConfiguration
	}{
		"QUERY",
		(MarshalTypeTargetCustomExpressionQueryScalingConfiguration)(m),
	}

	return json.Marshal(&s)
}
