// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
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

// CustomExpressionQueryScalingConfiguration The scaling configuration for the custom metric expression rule.
type CustomExpressionQueryScalingConfiguration struct {

	// The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of
	// the Monitoring service interprets results for each returned time series as Boolean values,
	// where zero represents false and a non-zero value represents true. A true value means that the trigger
	// rule condition has been met. The query must specify a metric, statistic, interval, and trigger
	// rule (threshold or absence). Supported values for interval: `1m`-`60m` (also `1h`). You can optionally
	// specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.
	// Example of threshold alarm:
	//   -----
	//     CPUUtilization[1m]{resourceId = "MODEL_DEPLOYMENT_OCID"}.grouping().mean() < 25
	//     CPUUtilization[1m]{resourceId = "MODEL_DEPLOYMENT_OCID"}.grouping().mean() > 75
	//   -----
	Query *string `mandatory:"true" json:"query"`

	// The period of time that the condition defined in the alarm must persist before the alarm state
	// changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the
	// alarm must persist in breaching the condition for five minutes before the alarm updates its
	// state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five
	// minutes before the alarm updates its state to "OK."
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M.
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`

	// The value is used for adjusting the count of instances by.
	InstanceCountAdjustment *int `mandatory:"false" json:"instanceCountAdjustment"`
}

// GetPendingDuration returns PendingDuration
func (m CustomExpressionQueryScalingConfiguration) GetPendingDuration() *string {
	return m.PendingDuration
}

// GetInstanceCountAdjustment returns InstanceCountAdjustment
func (m CustomExpressionQueryScalingConfiguration) GetInstanceCountAdjustment() *int {
	return m.InstanceCountAdjustment
}

func (m CustomExpressionQueryScalingConfiguration) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomExpressionQueryScalingConfiguration) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CustomExpressionQueryScalingConfiguration) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomExpressionQueryScalingConfiguration CustomExpressionQueryScalingConfiguration
	s := struct {
		DiscriminatorParam string `json:"scalingConfigurationType"`
		MarshalTypeCustomExpressionQueryScalingConfiguration
	}{
		"QUERY",
		(MarshalTypeCustomExpressionQueryScalingConfiguration)(m),
	}

	return json.Marshal(&s)
}
