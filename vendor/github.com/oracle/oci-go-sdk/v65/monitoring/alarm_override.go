// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetricData, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// AlarmOverride Values to use for an independent evaluation of the alarm.
// You can specify values for query, severity, body, and pending duration.
// When an alarm contains overrides, the Monitoring service evaluates each override in order,
// beginning with the first override in the array (index position `0`),
// and then evaluates the alarm's base values (`ruleName` value of `BASE`)
type AlarmOverride struct {

	// The period of time that the condition defined in the alarm must persist before the alarm state
	// changes from "OK" to "FIRING". For example, a value of 5 minutes means that the
	// alarm must persist in breaching the condition for five minutes before the alarm updates its
	// state to "FIRING".
	// The duration is specified as a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT1M. Maximum: PT1H. Default: PT1M.
	// Under the default value of PT1M, the first evaluation that breaches the alarm updates the
	// state to "FIRING".
	// The alarm updates its status to "OK" when the breaching condition has been clear for
	// the most recent minute.
	// Example: `PT5M`
	PendingDuration *string `mandatory:"false" json:"pendingDuration"`

	// The perceived severity of the alarm with regard to the affected system.
	// Example: `CRITICAL`
	Severity AlarmSeverityEnum `mandatory:"false" json:"severity,omitempty"`

	// The human-readable content of the delivered alarm notification. Oracle recommends providing guidance
	// to operators for resolving the alarm condition. Consider adding links to standard runbook
	// practices. Avoid entering confidential information.
	// Example: `High CPU usage alert. Follow runbook instructions for resolution.`
	Body *string `mandatory:"false" json:"body"`

	// A user-friendly description for this alarm override. Must be unique across all `ruleName` values for the alarm.
	RuleName *string `mandatory:"false" json:"ruleName"`

	// The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of
	// the Monitoring service interprets results for each returned time series as Boolean values,
	// where zero represents false and a non-zero value represents true. A true value means that the trigger
	// rule condition has been met. The query must specify a metric, statistic, interval, and trigger
	// rule (threshold or absence). Supported values for interval depend on the specified time range. More
	// interval values are supported for smaller time ranges. You can optionally
	// specify dimensions and grouping functions. Supported grouping functions: `grouping()`, `groupBy()`.
	// For information about writing MQL expressions, see
	// Editing the MQL Expression for a Query (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/query-metric-mql.htm).
	// For details about MQL, see
	// Monitoring Query Language (MQL) Reference (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Reference/mql.htm).
	// For available dimensions, review the metric definition for the supported service. See
	// Supported Services (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#SupportedServices).
	// Example of threshold alarm:
	//   -----
	//     CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.groupBy(availabilityDomain).percentile(0.9) > 85
	//   -----
	// Example of absence alarm:
	//   -----
	//     CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent()
	//   -----
	Query *string `mandatory:"false" json:"query"`
}

func (m AlarmOverride) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m AlarmOverride) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAlarmSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlarmSeverityEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
