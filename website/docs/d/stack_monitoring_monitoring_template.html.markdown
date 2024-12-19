---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_monitoring_template"
sidebar_current: "docs-oci-datasource-stack_monitoring-monitoring_template"
description: |-
  Provides details about a specific Monitoring Template in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_monitoring_template
This data source provides details about a specific Monitoring Template resource in Oracle Cloud Infrastructure Stack Monitoring service.

Gets a Monitoring Template by identifier

## Example Usage

```hcl
data "oci_stack_monitoring_monitoring_template" "test_monitoring_template" {
	#Required
	monitoring_template_id = oci_stack_monitoring_monitoring_template.test_monitoring_template.id
}
```

## Argument Reference

The following arguments are supported:

* `monitoring_template_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the monitoring template.


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

