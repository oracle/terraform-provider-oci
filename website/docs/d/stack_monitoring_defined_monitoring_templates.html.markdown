---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_defined_monitoring_templates"
sidebar_current: "docs-oci-datasource-stack_monitoring-defined_monitoring_templates"
description: |-
  Provides the list of Defined Monitoring Templates in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_defined_monitoring_templates
This data source provides the list of Defined Monitoring Templates in Oracle Cloud Infrastructure Stack Monitoring service.

List Defined Monitoring Templates.

## Example Usage

```hcl
data "oci_stack_monitoring_defined_monitoring_templates" "test_defined_monitoring_templates" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.defined_monitoring_template_display_name
	resource_types = var.defined_monitoring_template_resource_types
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the tenancy(root) for which  defined monitored templates should be listed. 
* `display_name` - (Optional) A filter to return monitoring template based on name.
* `resource_types` - (Optional) Multiple resource types filter.


## Attributes Reference

The following attributes are exported:

* `defined_monitoring_template_collection` - The list of defined_monitoring_template_collection.

### DefinedMonitoringTemplate Reference

The following attributes are exported:

* `items` - List of defined Monitoring Template.
	* `composite_type` - Type of composite resource type OCID like EBS/PEOPLE_SOFT.
	* `defined_alarm_conditions` - Defined Monitoring template alarm conditions
		* `condition_type` - Type of defined monitoring template.
		* `conditions` - Monitoring template conditions.
			* `body` - The human-readable content of the delivered alarm notification. Oracle recommends providing guidance to operators for resolving the alarm condition. Consider adding links to standard runbook practices. Avoid entering confidential information.
			* `query` - The Monitoring Query Language (MQL) expression to evaluate for the alarm.
			* `severity` - Severity - Critical/Warning
			* `should_append_note` - Whether the note need to add into bottom of the body for mapping the alarms information with template or not.
			* `should_append_url` - Whether the URL need to add into bottom of the body for mapping the alarms information with template or not.
			* `trigger_delay` - The period of time that the condition defined in the alarm must persist before the alarm state changes from "OK" to "FIRING".
		* `metric_name` - The metric name.
	* `display_name` - The name of the definedMonitoringTemplate.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the definedMonitoringTemplate.
	* `namespace` - The stack monitoring service or application emitting the metric that is evaluated by the alarm.
	* `resource_type` - The resource types OCID.
	* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The date and time the monitoringTemplate was created. Format defined by RFC3339.
	* `time_updated` - The date and time the monitoringTemplate was updated. Format defined by RFC3339.

