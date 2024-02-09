// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Autoscaling API
//
// Use the Autoscaling API to dynamically scale compute resources to meet application requirements. For more information about
// autoscaling, see Autoscaling (https://docs.cloud.oracle.com/Content/Compute/Tasks/autoscalinginstancepools.htm). For information about the
// Compute service, see Compute (https://docs.cloud.oracle.com/Content/Compute/home.htm).
//

package autoscaling

import (
	"encoding/json"
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// CustomMetric Metric details for triggering an autoscaling action based on a custom MQL query.
type CustomMetric struct {

	// The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of
	// the Monitoring service interprets results for each returned time series as Boolean values,
	// where zero represents false and a non-zero value represents true. A true value means that the trigger
	// rule condition has been met. The query must specify a metric, statistic, interval, and trigger
	// rule (threshold or absence). Supported values for interval: `1m`-`60m` (also `1h`). You can optionally
	// specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.
	// Example of threshold alarm:
	//   -----
	//     CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85
	//   -----
	Query *string `mandatory:"true" json:"query"`

	// The namespace for the query.
	Namespace *string `mandatory:"true" json:"namespace"`

	// The OCID of the compartment containing the metrics.
	MetricCompartmentId *string `mandatory:"true" json:"metricCompartmentId"`

	// The period of time that the condition defined in the alarm must persist before the alarm state
	// changes from "OK" to "FIRING" or vice versa. For example, a value of 5 minutes means that the
	// alarm must persist in breaching the condition for five minutes before the alarm updates its
	// state to "FIRING"; likewise, the alarm must persist in not breaching the condition for five
	// minutes before the alarm updates its state to "OK."
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT3M. Maximum: PT1H. Default: PT3M.
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`

	// The resource group for the query.
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`
}

// GetPendingDuration returns PendingDuration
func (m CustomMetric) GetPendingDuration() *string {
	return m.PendingDuration
}

func (m CustomMetric) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m CustomMetric) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// MarshalJSON marshals to json representation
func (m CustomMetric) MarshalJSON() (buff []byte, e error) {
	type MarshalTypeCustomMetric CustomMetric
	s := struct {
		DiscriminatorParam string `json:"metricSource"`
		MarshalTypeCustomMetric
	}{
		"CUSTOM_QUERY",
		(MarshalTypeCustomMetric)(m),
	}

	return json.Marshal(&s)
}
