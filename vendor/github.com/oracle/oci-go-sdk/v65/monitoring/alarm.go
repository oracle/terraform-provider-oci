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

// Alarm The properties that define an alarm.
// For information about alarms, see
// Alarms Overview (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#AlarmsOverview).
// To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized,
// talk to an administrator. If you're an administrator who needs to write policies to give users access, see
// Getting Started with Policies (https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/policygetstarted.htm).
// For information about endpoints and signing API requests, see
// About the API (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/usingapi.htm).
// For information about available SDKs and tools, see
// SDKS and Other Tools (https://docs.cloud.oracle.com/iaas/Content/API/Concepts/sdks.htm).
type Alarm struct {

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the alarm.
	Id *string `mandatory:"true" json:"id"`

	// A user-friendly name for the alarm. It does not have to be unique, and it's changeable.
	// This value determines the title of each alarm notification.
	// Example: `High CPU Utilization`
	DisplayName *string `mandatory:"true" json:"displayName"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the alarm.
	CompartmentId *string `mandatory:"true" json:"compartmentId"`

	// The OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the metric
	// being evaluated by the alarm.
	MetricCompartmentId *string `mandatory:"true" json:"metricCompartmentId"`

	// The source service or application emitting the metric that is evaluated by the alarm.
	// Example: `oci_computeagent`
	Namespace *string `mandatory:"true" json:"namespace"`

	// The Monitoring Query Language (MQL) expression to evaluate for the alarm. The Alarms feature of
	// the Monitoring service interprets results for each returned time series as Boolean values,
	// where zero represents false and a non-zero value represents true. A true value means that the trigger
	// rule condition has been met. The query must specify a metric, statistic, interval, and trigger
	// rule (threshold or absence). Supported values for interval depend on the specified time range. More
	// interval values are supported for smaller time ranges. You can optionally
	// specify dimensions and grouping functions.
	// Also, you can customize the
	// absence detection period (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/create-edit-alarm-query-absence-detection-period.htm).
	// Supported grouping functions: `grouping()`, `groupBy()`.
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
	// Example of absence alarm with custom absence detection period of 20 hours:
	//   -----
	//
	//     CpuUtilization[1m]{availabilityDomain="cumS:PHX-AD-1"}.absent(20h)
	//
	//   -----
	Query *string `mandatory:"true" json:"query"`

	// The perceived type of response required when the alarm is in the "FIRING" state.
	// Example: `CRITICAL`
	Severity AlarmSeverityEnum `mandatory:"true" json:"severity"`

	// A list of destinations for alarm notifications.
	// Each destination is represented by the OCID (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
	// of a related resource, such as a NotificationTopic.
	// Supported destination services: Notifications, Streaming.
	// Limit: One destination per supported destination service.
	Destinations []string `mandatory:"true" json:"destinations"`

	// Whether the alarm is enabled.
	// Example: `true`
	IsEnabled *bool `mandatory:"true" json:"isEnabled"`

	// The current lifecycle state of the alarm.
	// Example: `DELETED`
	LifecycleState AlarmLifecycleStateEnum `mandatory:"true" json:"lifecycleState"`

	// The date and time the alarm was created. Format defined by RFC3339.
	// Example: `2023-02-01T01:02:29.600Z`
	TimeCreated *common.SDKTime `mandatory:"true" json:"timeCreated"`

	// The date and time the alarm was last updated. Format defined by RFC3339.
	// Example: `2023-02-03T01:02:29.600Z`
	TimeUpdated *common.SDKTime `mandatory:"true" json:"timeUpdated"`

	// When true, the alarm evaluates metrics from all compartments and subcompartments. The parameter can
	// only be set to true when metricCompartmentId is the tenancy OCID (the tenancy is the root compartment).
	// A true value requires the user to have tenancy-level permissions. If this requirement is not met,
	// then the call is rejected. When false, the alarm evaluates metrics from only the compartment specified
	// in metricCompartmentId. Default is false.
	// Example: `true`
	MetricCompartmentIdInSubtree *bool `mandatory:"false" json:"metricCompartmentIdInSubtree"`

	// Resource group to match for metric data retrieved by the alarm. A resource group is a custom string that you can match when retrieving custom metrics. Only one resource group can be applied per metric.
	// A valid resourceGroup value starts with an alphabetical character and includes only alphanumeric characters, periods (.), underscores (_), hyphens (-), and dollar signs ($).
	// Example: `frontend-fleet`
	ResourceGroup *string `mandatory:"false" json:"resourceGroup"`

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

	// The human-readable content of the delivered alarm notification.
	// Optionally include dynamic variables (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/update-alarm-dynamic-variables.htm).
	// Oracle recommends providing guidance
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
	MessageFormat AlarmMessageFormatEnum `mandatory:"false" json:"messageFormat,omitempty"`

	// The frequency for re-submitting alarm notifications, if the alarm keeps firing without
	// interruption. Format defined by ISO 8601. For example, `PT4H` indicates four hours.
	// Minimum: PT1M. Maximum: P30D.
	// Default value: null (notifications are not re-submitted).
	// Example: `PT2H`
	RepeatNotificationDuration *string `mandatory:"false" json:"repeatNotificationDuration"`

	// The configuration details for suppressing an alarm.
	Suppression *Suppression `mandatory:"false" json:"suppression"`

	// Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.
	// Example: `{"Department": "Finance"}`
	FreeformTags map[string]string `mandatory:"false" json:"freeformTags"`

	// Usage of predefined tag keys. These predefined keys are scoped to namespaces.
	// Example: `{"Operations": {"CostCenter": "42"}}`
	DefinedTags map[string]map[string]interface{} `mandatory:"false" json:"definedTags"`

	// A set of overrides that control evaluations of the alarm.
	// Each override can specify values for query, severity, body, and pending duration.
	// When an alarm contains overrides, the Monitoring service evaluates each override in order, beginning with the first override in the array (index position `0`),
	// and then evaluates the alarm's base values (`ruleName` value of `BASE`).
	Overrides []AlarmOverride `mandatory:"false" json:"overrides"`

	// Identifier of the alarm's base values for alarm evaluation, for use when the alarm contains overrides.
	// Default value is `BASE`. For information about alarm overrides, see AlarmOverride.
	RuleName *string `mandatory:"false" json:"ruleName"`

	// The version of the alarm notification to be delivered. Allowed value: `1.X`
	// The value must start with a number (up to four digits), followed by a period and an uppercase X.
	NotificationVersion *string `mandatory:"false" json:"notificationVersion"`

	// Customizable notification title (`title` alarm message parameter (https://docs.cloud.oracle.com/iaas/Content/Monitoring/alarm-message-format.htm)).
	// Optionally include dynamic variables (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/update-alarm-dynamic-variables.htm).
	// The notification title appears as the subject line in a formatted email message and as the title in a Slack message.
	NotificationTitle *string `mandatory:"false" json:"notificationTitle"`

	// Customizable slack period to wait for metric ingestion before evaluating the alarm.
	// Specify a string in ISO 8601 format (`PT10M` for ten minutes or `PT1H`
	// for one hour). Minimum: PT3M. Maximum: PT2H. Default: PT3M.
	// For more information about the slack period, see
	// About the Internal Reset Period (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Concepts/monitoringoverview.htm#reset).
	EvaluationSlackDuration *string `mandatory:"false" json:"evaluationSlackDuration"`

	// Customizable alarm summary (`alarmSummary` alarm message parameter (https://docs.cloud.oracle.com/iaas/Content/Monitoring/alarm-message-format.htm)).
	// Optionally include dynamic variables (https://docs.cloud.oracle.com/iaas/Content/Monitoring/Tasks/update-alarm-dynamic-variables.htm).
	// The alarm summary appears within the body of the alarm message and in responses to
	// ListAlarmsStatus
	// GetAlarmHistory and
	// RetrieveDimensionStates.
	AlarmSummary *string `mandatory:"false" json:"alarmSummary"`
}

func (m Alarm) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m Alarm) ValidateEnumValue() (bool, error) {
	errMessage := []string{}
	if _, ok := GetMappingAlarmSeverityEnum(string(m.Severity)); !ok && m.Severity != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for Severity: %s. Supported values are: %s.", m.Severity, strings.Join(GetAlarmSeverityEnumStringValues(), ",")))
	}
	if _, ok := GetMappingAlarmLifecycleStateEnum(string(m.LifecycleState)); !ok && m.LifecycleState != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for LifecycleState: %s. Supported values are: %s.", m.LifecycleState, strings.Join(GetAlarmLifecycleStateEnumStringValues(), ",")))
	}

	if _, ok := GetMappingAlarmMessageFormatEnum(string(m.MessageFormat)); !ok && m.MessageFormat != "" {
		errMessage = append(errMessage, fmt.Sprintf("unsupported enum value for MessageFormat: %s. Supported values are: %s.", m.MessageFormat, strings.Join(GetAlarmMessageFormatEnumStringValues(), ",")))
	}
	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}

// AlarmSeverityEnum Enum with underlying type: string
type AlarmSeverityEnum string

// Set of constants representing the allowable values for AlarmSeverityEnum
const (
	AlarmSeverityCritical AlarmSeverityEnum = "CRITICAL"
	AlarmSeverityError    AlarmSeverityEnum = "ERROR"
	AlarmSeverityWarning  AlarmSeverityEnum = "WARNING"
	AlarmSeverityInfo     AlarmSeverityEnum = "INFO"
)

var mappingAlarmSeverityEnum = map[string]AlarmSeverityEnum{
	"CRITICAL": AlarmSeverityCritical,
	"ERROR":    AlarmSeverityError,
	"WARNING":  AlarmSeverityWarning,
	"INFO":     AlarmSeverityInfo,
}

var mappingAlarmSeverityEnumLowerCase = map[string]AlarmSeverityEnum{
	"critical": AlarmSeverityCritical,
	"error":    AlarmSeverityError,
	"warning":  AlarmSeverityWarning,
	"info":     AlarmSeverityInfo,
}

// GetAlarmSeverityEnumValues Enumerates the set of values for AlarmSeverityEnum
func GetAlarmSeverityEnumValues() []AlarmSeverityEnum {
	values := make([]AlarmSeverityEnum, 0)
	for _, v := range mappingAlarmSeverityEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmSeverityEnumStringValues Enumerates the set of values in String for AlarmSeverityEnum
func GetAlarmSeverityEnumStringValues() []string {
	return []string{
		"CRITICAL",
		"ERROR",
		"WARNING",
		"INFO",
	}
}

// GetMappingAlarmSeverityEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmSeverityEnum(val string) (AlarmSeverityEnum, bool) {
	enum, ok := mappingAlarmSeverityEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AlarmMessageFormatEnum Enum with underlying type: string
type AlarmMessageFormatEnum string

// Set of constants representing the allowable values for AlarmMessageFormatEnum
const (
	AlarmMessageFormatRaw          AlarmMessageFormatEnum = "RAW"
	AlarmMessageFormatPrettyJson   AlarmMessageFormatEnum = "PRETTY_JSON"
	AlarmMessageFormatOnsOptimized AlarmMessageFormatEnum = "ONS_OPTIMIZED"
)

var mappingAlarmMessageFormatEnum = map[string]AlarmMessageFormatEnum{
	"RAW":           AlarmMessageFormatRaw,
	"PRETTY_JSON":   AlarmMessageFormatPrettyJson,
	"ONS_OPTIMIZED": AlarmMessageFormatOnsOptimized,
}

var mappingAlarmMessageFormatEnumLowerCase = map[string]AlarmMessageFormatEnum{
	"raw":           AlarmMessageFormatRaw,
	"pretty_json":   AlarmMessageFormatPrettyJson,
	"ons_optimized": AlarmMessageFormatOnsOptimized,
}

// GetAlarmMessageFormatEnumValues Enumerates the set of values for AlarmMessageFormatEnum
func GetAlarmMessageFormatEnumValues() []AlarmMessageFormatEnum {
	values := make([]AlarmMessageFormatEnum, 0)
	for _, v := range mappingAlarmMessageFormatEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmMessageFormatEnumStringValues Enumerates the set of values in String for AlarmMessageFormatEnum
func GetAlarmMessageFormatEnumStringValues() []string {
	return []string{
		"RAW",
		"PRETTY_JSON",
		"ONS_OPTIMIZED",
	}
}

// GetMappingAlarmMessageFormatEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmMessageFormatEnum(val string) (AlarmMessageFormatEnum, bool) {
	enum, ok := mappingAlarmMessageFormatEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}

// AlarmLifecycleStateEnum Enum with underlying type: string
type AlarmLifecycleStateEnum string

// Set of constants representing the allowable values for AlarmLifecycleStateEnum
const (
	AlarmLifecycleStateActive   AlarmLifecycleStateEnum = "ACTIVE"
	AlarmLifecycleStateDeleting AlarmLifecycleStateEnum = "DELETING"
	AlarmLifecycleStateDeleted  AlarmLifecycleStateEnum = "DELETED"
)

var mappingAlarmLifecycleStateEnum = map[string]AlarmLifecycleStateEnum{
	"ACTIVE":   AlarmLifecycleStateActive,
	"DELETING": AlarmLifecycleStateDeleting,
	"DELETED":  AlarmLifecycleStateDeleted,
}

var mappingAlarmLifecycleStateEnumLowerCase = map[string]AlarmLifecycleStateEnum{
	"active":   AlarmLifecycleStateActive,
	"deleting": AlarmLifecycleStateDeleting,
	"deleted":  AlarmLifecycleStateDeleted,
}

// GetAlarmLifecycleStateEnumValues Enumerates the set of values for AlarmLifecycleStateEnum
func GetAlarmLifecycleStateEnumValues() []AlarmLifecycleStateEnum {
	values := make([]AlarmLifecycleStateEnum, 0)
	for _, v := range mappingAlarmLifecycleStateEnum {
		values = append(values, v)
	}
	return values
}

// GetAlarmLifecycleStateEnumStringValues Enumerates the set of values in String for AlarmLifecycleStateEnum
func GetAlarmLifecycleStateEnumStringValues() []string {
	return []string{
		"ACTIVE",
		"DELETING",
		"DELETED",
	}
}

// GetMappingAlarmLifecycleStateEnum performs case Insensitive comparison on enum value and return the desired enum
func GetMappingAlarmLifecycleStateEnum(val string) (AlarmLifecycleStateEnum, bool) {
	enum, ok := mappingAlarmLifecycleStateEnumLowerCase[strings.ToLower(val)]
	return enum, ok
}
