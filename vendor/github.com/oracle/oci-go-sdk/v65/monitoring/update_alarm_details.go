// Copyright (c) 2016, 2018, 2023, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Monitoring API
//
// Use the Monitoring API to manage metric queries and alarms for assessing the health, capacity, and performance of your cloud resources.
// Endpoints vary by operation. For PostMetric, use the `telemetry-ingestion` endpoints; for all other operations, use the `telemetry` endpoints.
// For more information, see
// the Monitoring documentation (https://docs.cloud.oracle.com/iaas/Content/Monitoring/home.htm).
//

package monitoring

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// UpdateAlarmDetails The configuration details for updating an alarm.
type UpdateAlarmDetails struct {

	// A user-friendly name for the alarm. It does not have to be unique, and it's changeable.
	// Avoid entering confidential information.
	// This value determines the title of each alarm notification.
	// Example: `High CPU Utilization`
	DisplayName *string `mandatory:"false" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm.
	CompartmentId *string `mandatory:"false" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric
	// being evaluated by the alarm.
	MetricCompartmentId *string `mandatory:"false" json:"metricCompartmentId"`

	// When true, the alarm evaluates metrics from all compartments and subcompartments. The parameter can
	// only be set to true when metricCompartmentId is the tenancy OCID (the tenancy is the root compartment).
	// A true value requires the user to have tenancy-level permissions. If this requirement is not met,
	// then the call is rejected. When false, the alarm evaluates metrics from only the compartment specified
	// in metricCompartmentId. Default is false.
	// Example: `true`
	MetricCompartmentIdInSubtree *bool `mandatory:"false" json:"metricCompartmentIdInSubtree"`

	// The source service or application emitting the metric that is evaluated by the alarm.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"false" json:"namespace"`

	// Resource group that you want to match. A null value returns only metric data that has no resource groups. The alarm retrieves metric data associated with the specified resource group only. Only one resource group can be applied per metric.
	// A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	// Avoid entering confidential information.
	// Example: `frontend-fleet`
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

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

	// The time between calculated aggregation windows for the alarm. Supported value: `1m`
	Resolution *string `mandatory:"false" json:"resolution"`

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

	// When set to `true`, splits alarm notifications per metric stream.
	// When set to `false`, groups alarm notifications across metric streams.
	IsNotificationsPerMetricDimensionEnabled *bool `mandatory:"false" json:"isNotificationsPerMetricDimensionEnabled"`

	// The format to use for alarm notifications. The formats are:
	// * `RAW` - Raw JSON blob. Default value. When the `destinations` attribute specifies `Streaming`, all alarm notifications use this format.
	// * `PRETTY_JSON`: JSON with new lines and indents. Available when the `destinations` attribute specifies `Notifications` only.
	// * `ONS_OPTIMIZED`: Simplified, user-friendly layout. Available when the `destinations` attribute specifies `Notifications` only. Applies to Email subscription types only.
	MessageFormat UpdateAlarmDetailsMessageFormatEnum `mandatory:"false" json:"messageFormat,omitempty"`

	// A list of destinations for alarm notifications.
	// Each destination is represented by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	// of a related resource, such as a NotificationTopic.
	// Supported destination services: Notifications
	// , Streaming.
	// Limit: One destination per supported destination service.
	Destinations []string `mandatory:"false" json:"destinations"`

	// The frequency for re-submitting alarm notifications, if the alarm keeps firing without
	// interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours.
	// Minimum: PT1M. Maximum: P30D.
	// Default value: null (notifications are not re-submitted).
	// Example: `PT2H`
	RepeatNotificationDuration *string `mandatory:"false" json:"repeatNotificationDuration"`

	// The configuration details for suppressing an alarm.
	Suppression *Suppression `mandatory:"false" json:"suppression"`

	// Whether the alarm is enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"false" json:"isEnabled"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`
}

func (m UpdateAlarmDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m UpdateAlarmDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if _, ok := GetMappingAlarmSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlarmSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingUpdateAlarmDetailsMessageFormatEnum(string(m.MessageFormat)); !ok && m.MessageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageFormat: %s. Supported values are: %s.", m.MessageFormat, strings.Join(GetUpdateAlarmDetailsMessageFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// UpdateAlarmDetailsMessageFormatEnum Enum with underlying type: string
type UpdateAlarmDetailsMessageFormatEnum string

// Set of constants representing the allowable values for UpdateAlarmDetailsMessageFormatEnum
const (
	UpdateAlarmDetailsMessageFormatRaw          UpdateAlarmDetailsMessageFormatEnum = "RAW"
	UpdateAlarmDetailsMessageFormatPrettyJson   UpdateAlarmDetailsMessageFormatEnum = "PRETTY_JSON"
	UpdateAlarmDetailsMessageFormatOnsOptimized UpdateAlarmDetailsMessageFormatEnum = "ONS_OPTIMIZED"
)

var mappingUpdateAlarmDetailsMessageFormatEnum = map[string]UpdateAlarmDetailsMessageFormatEnum{
	"RAW":           UpdateAlarmDetailsMessageFormatRaw,
	"PRETTY_JSON":   UpdateAlarmDetailsMessageFormatPrettyJson,
	"ONS_OPTIMIZED": UpdateAlarmDetailsMessageFormatOnsOptimized,
}

var mappingUpdateAlarmDetailsMessageFormatEnumLowerCase = map[string]UpdateAlarmDetailsMessageFormatEnum{
	"raw":           UpdateAlarmDetailsMessageFormatRaw,
	"pretty_json":   UpdateAlarmDetailsMessageFormatPrettyJson,
	"ons_optimized": UpdateAlarmDetailsMessageFormatOnsOptimized,
}

// GetUpdateAlarmDetailsMessageFormatEnumValues Enumerates the set of values for UpdateAlarmDetailsMessageFormatEnum
func GetUpdateAlarmDetailsMessageFormatEnumValues() []UpdateAlarmDetailsMessageFormatEnum {
	values := make([]UpdateAlarmDetailsMessageFormatEnum, 0)
	for _, v := range mappingUpdateAlarmDetailsMessageFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetUpdateAlarmDetailsMessageFormatEnumStringValues Enumerates the set of values in String for UpdateAlarmDetailsMessageFormatEnum
func GetUpdateAlarmDetailsMessageFormatEnumStringValues() []string {
	return []string{
		"RAW",
		"PRETTY_JSON",
		"ONS_OPTIMIZED",
	}
}

// GetMappingUpdateAlarmDetailsMessageFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingUpdateAlarmDetailsMessageFormatEnum(val string) (UpdateAlarmDetailsMessageFormatEnum, bool) {
	enum, ok := mappingUpdateAlarmDetailsMessageFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
