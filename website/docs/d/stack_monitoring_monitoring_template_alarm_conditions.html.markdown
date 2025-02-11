---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitoring_template_alarm_conditions"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitoring_template_alarm_conditions"
description: |-
  Provides the list of Monitoring Template Alarm Conditions in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitoring_template_alarm_conditions
This data source provides the list of Monitoring Template Alarm Conditions in Oracle Cloud Infrastructure Stack Monitoring service.

Returns a list of Alarm Conditions.

## Example Usage

```hcl
data "oci_stack_monitoring_monitoring_template_alarm_conditions" "test_monitoring_template_alarm_conditions" {
	#Required
	monitoring_template_id = oci_stack_monitoring_monitoring_template.test_monitoring_template.id

	#Optional
	metric_name = oci_monitoring_metric.test_metric.name
	resource_types = var.monitoring_template_alarm_condition_resource_types
	state = var.monitoring_template_alarm_condition_state
	status = var.monitoring_template_alarm_condition_status
}
```

## Argument Reference

The following arguments are supported:

* `metric_name` - (Optional) metricName filter.
* `monitoring_template_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoring template.
* `resource_types` - (Optional) Multiple resource types filter.
* `state` - (Optional) A filter to return alarm condition based on Lifecycle State.
* `status` - (Optional) A filter to return alarm condition based on input status.


## Attributes Reference

The following attributes are exported:

* `alarm_condition_collection` - The list of alarm_condition_collection.

### MonitoringTemplateAlarmCondition Reference

The following attributes are exported:

* `composite_type` - The OCID of the composite resource type like EBS/PEOPLE_SOFT.
* `condition_type` - Type of defined monitoring template.
* `conditions` - Monitoring template conditions
	* `body` - The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.
	* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm.
	* `severity` - Severity - Critical/Warning
	* `should_append_note` - Whether the note need to add into bottom of the body for mapping the alarms information with template or not.
	* `should_append_url` - Whether the URL need to add into bottom of the body for mapping the alarms information with template or not.
	* `trigger_delay` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING".
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Alarm Condition.
* `metric_name` - The metric name.
* `monitoring_template_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoring template.
* `namespace` - The stack monitoring service or application emitting the metric that is evaluated by the alarm.
* `resource_type` - The resource type OCID.
* `state` - The current lifecycle state of the monitoring template
* `status` - The current status of the monitoring template i.e. whether it is Published or Unpublished
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the alarm condition was created. Format defined by RFC3339.
* `time_updated` - The date and time the alarm condition was updated. Format defined by RFC3339.

