---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitoring_template"
sidebar_current: "docs-oci-resource-stack_monitoring-monitoring_template"
description: |-
  Provides the Monitoring Template resource in Oracle Cloud Infrastructure Stack Monitoring service
---

# oci_stack_monitoring_monitoring_template
This resource provides the Monitoring Template resource in Oracle Cloud Infrastructure Stack Monitoring service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/stack-monitoring/latest/MonitoringTemplate

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/stack_monitoring

Creates a new monitoring template for a given compartment.

## Example Usage

```hcl
resource "oci_stack_monitoring_monitoring_template" "test_monitoring_template" {
	#Required
	compartment_id = var.compartment_id
	destinations = var.monitoring_template_destinations
	display_name = var.monitoring_template_display_name
	members {
		#Required
		id = var.monitoring_template_members_id
		type = var.monitoring_template_members_type

		#Optional
		composite_type = var.monitoring_template_members_composite_type
	}

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.monitoring_template_description
	freeform_tags = {"bar-key"= "value"}
	is_alarms_enabled = var.monitoring_template_is_alarms_enabled
	is_split_notification_enabled = var.monitoring_template_is_split_notification_enabled
	message_format = var.monitoring_template_message_format
	repeat_notification_duration = var.monitoring_template_repeat_notification_duration
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment containing the monitoringTemplate.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A user-friendly description for the monitoring template. It does not have to be unique, and it's changeable. Avoid entering confidential information.
* `destinations` - (Required) (Updatable) A list of destinations for alarm notifications. Each destination is represented by the OCID of a related resource, such as a topic.
* `display_name` - (Required) (Updatable) A user-friendly name for the monitoring template. It is unique and mutable in nature. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `is_alarms_enabled` - (Optional) (Updatable) Whether the alarm is enabled or disabled, it will be Enabled by default.
* `is_split_notification_enabled` - (Optional) (Updatable) Whether the alarm notification is enabled or disabled, it will be Enabled by default.
* `members` - (Required) (Updatable) List of members of this monitoring template
	* `composite_type` - (Optional) (Updatable) The OCID of the composite resource type like EBS or Peoplesoft.
	* `id` - (Required) (Updatable) The OCID of the resourceInstance/resourceType/resourceGroup
	* `type` - (Required) (Updatable) Type of the member reference RESOURCE_INSTANCE, RESOURCE_TYPE, RESOURCE_GROUP
* `message_format` - (Optional) (Updatable) The format to use for alarm notifications.
* `repeat_notification_duration` - (Optional) (Updatable) The frequency for re-submitting alarm notifications, if the alarm keeps firing without interruption. Format defined by ISO 8601. For example, PT4H indicates four hours. Minimum- PT1M. Maximum - P30D.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Monitoring Template
	* `update` - (Defaults to 20 minutes), when updating the Monitoring Template
	* `delete` - (Defaults to 20 minutes), when destroying the Monitoring Template


## Import

MonitoringTemplates can be imported using the `id`, e.g.

```
$ terraform import oci_stack_monitoring_monitoring_template.test_monitoring_template "id"
```

