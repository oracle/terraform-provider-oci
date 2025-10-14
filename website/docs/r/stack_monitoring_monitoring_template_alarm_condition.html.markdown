---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitoring_template_alarm_condition"
sidebar_current: "docs-oci-resource-stack_monitoring-monitoring_template_alarm_condition"
description: |-
  Provides the Monitoring Template Alarm Condition resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitoring_template_alarm_condition
This resource provides the Monitoring Template Alarm Condition resource in Oracle Cloud Infrastructure Stack Monitoring service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/stack-monitoring/latest/MonitoringTemplateAlarmCondition

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/stack_monitoring

Create a new alarm condition in same monitoringTemplate compartment.

## Example Usage

```hcl
resource "oci_stack_monitoring_monitoring_template_alarm_condition" "test_monitoring_template_alarm_condition" {
	#Required
	condition_type = var.monitoring_template_alarm_condition_condition_type
	conditions {
		#Required
		query = var.monitoring_template_alarm_condition_conditions_query
		severity = var.monitoring_template_alarm_condition_conditions_severity

		#Optional
		body = var.monitoring_template_alarm_condition_conditions_body
		should_append_note = var.monitoring_template_alarm_condition_conditions_should_append_note
		should_append_url = var.monitoring_template_alarm_condition_conditions_should_append_url
		trigger_delay = var.monitoring_template_alarm_condition_conditions_trigger_delay
	}
	metric_name = oci_monitoring_metric.test_metric.name
	monitoring_template_id = oci_stack_monitoring_monitoring_template.test_monitoring_template.id
	namespace = var.monitoring_template_alarm_condition_namespace
	resource_type = var.monitoring_template_alarm_condition_resource_type

	#Optional
	composite_type = var.monitoring_template_alarm_condition_composite_type
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `composite_type` - (Optional) (Updatable) The OCID of the composite resource type like EBS/PEOPLE_SOFT.
* `condition_type` - (Required) (Updatable) Type of defined monitoring template.
* `conditions` - (Required) (Updatable) Monitoring template conditions.
	* `body` - (Optional) (Updatable) The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.
	* `query` - (Required) (Updatable) The Monitoring Query Language (MQL) expression to evaluate for the alarm.
	* `severity` - (Required) (Updatable) Severity - Critical/Warning
	* `should_append_note` - (Optional) (Updatable) Whether the note need to add into bottom of the body for mapping the alarms information with template or not.
	* `should_append_url` - (Optional) (Updatable) Whether the URL need to add into bottom of the body for mapping the alarms information with template or not.
	* `trigger_delay` - (Optional) (Updatable) The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING".
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `metric_name` - (Required) (Updatable) The metric name.
* `monitoring_template_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoring template.
* `namespace` - (Required) (Updatable) The stack monitoring service or application emitting the metric that is evaluated by the alarm.
* `resource_type` - (Required) (Updatable) The resource group OCID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitoring Template Alarm Condition
	* `update` - (Defaults to 20 minutes), when updating the Monitoring Template Alarm Condition
	* `delete` - (Defaults to 20 minutes), when destroying the Monitoring Template Alarm Condition


## Import

MonitoringTemplateAlarmConditions can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitoring_template_alarm_condition.test_monitoring_template_alarm_condition "monitoringTemplates/{monitoringTemplateId}/alarmConditions/{alarmConditionId}" 
```

