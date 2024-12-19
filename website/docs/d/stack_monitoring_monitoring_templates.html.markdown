---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitoring_templates"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitoring_templates"
description: |-
  Provides the list of Monitoring Templates in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitoring_templates
This data source provides the list of Monitoring Templates in Oracle Cloud Infrastructure Stack Monitoring service.

Returns a list of Monitoring Templates.

## Example Usage

```hcl
data "oci_stack_monitoring_monitoring_templates" "test_monitoring_templates" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.monitoring_template_display_name
	metric_name = oci_monitoring_metric.test_metric.name
	monitoring_template_id = oci_stack_monitoring_monitoring_template.test_monitoring_template.id
	namespace = var.monitoring_template_namespace
	resource_types = var.monitoring_template_resource_types
	state = var.monitoring_template_state
	status = var.monitoring_template_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which data is listed.
* `display_name` - (Optional) A filter to return monitoring template based on name.
* `metric_name` - (Optional) metricName filter.
* `monitoring_template_id` - (Optional) A filter to return monitoring template based on input monitoringTemplateId
* `namespace` - (Optional) namespace filter.
* `resource_types` - (Optional) Multiple resource types filter.
* `state` - (Optional) A filter to return monitoring template based on Lifecycle State
* `status` - (Optional) A filter to return monitoring template based on input status


## Attributes Reference

The following attributes are exported:

* `monitoring_template_collection` - The list of monitoring_template_collection.

### MonitoringTemplate Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment containing the monitoringTemplate.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A user-friendly description for the monitoring template. It does not have to be unique, and it's changeable. Avoid entering confidential information.
* `destinations` - A list of destinations for alarm notifications. Each destination is represented by the OCID of a related resource.
* `display_name` - A user-friendly name for the monitoring template. It should be unique, and it's mutable in nature. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoringTemplate
* `is_alarms_enabled` - Whether the alarm is enabled or disabled. Default value is enabled.
* `is_split_notification_enabled` - Whether the alarm notification is enabled or disabled, it will be Enabled by default.
* `members` - List of members of this monitoring template.
	* `composite_type` - The OCID of the composite resource type like EBS or Peoplesoft.
	* `id` - The OCID of the resourceInstance/resourceType/resourceGroup
	* `type` - Type of the member reference RESOURCE_INSTANCE, RESOURCE_TYPE, RESOURCE_GROUP
* `message_format` - The format to use for alarm notifications.
* `repeat_notification_duration` - The frequency for re-submitting alarm notifications, if the alarm keeps firing without interruption. Format defined by ISO 8601. For example, PT4H indicates four hours. Minimum- PT1M. Maximum - P30D.
* `state` - The current lifecycle state of the monitoring template.
* `status` - The current status of the monitoring template i.e. whether it is Applied or NotApplied.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tenant_id` - Tenant Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)
* `time_created` - The date and time the monitoringTemplate was created. Format defined by RFC3339.
* `time_updated` - The date and time the monitoringTemplate was last updated. Format defined by RFC3339.
* `total_alarm_conditions` - Total Alarm Conditions
* `total_applied_alarm_conditions` - Total Applied Alarm Conditions

