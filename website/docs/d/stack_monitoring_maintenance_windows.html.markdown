---
subcategory: "Stack Monitoring"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_stack_monitoring_maintenance_windows"
sidebar_current: "docs-oci-datasource-stack_monitoring-maintenance_windows"
description: |-
  Provides the list of Maintenance Windows in Oracle Cloud Infrastructure Stack Monitoring service
---

# Data Source: oci_stack_monitoring_maintenance_windows
This data source provides the list of Maintenance Windows in Oracle Cloud Infrastructure Stack Monitoring service.

Returns a list of maintenance windows.

## Example Usage

```hcl
data "oci_stack_monitoring_maintenance_windows" "test_maintenance_windows" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	lifecycle_details = var.maintenance_window_lifecycle_details
	name = var.maintenance_window_name
	status = var.maintenance_window_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which data is listed.
* `lifecycle_details` - (Optional) A filter to return maintenance windows with matching lifecycleDetails.
* `name` - (Optional) A filter to return maintenance windows that match exact resource name.
* `status` - (Optional) A filter to return only maintenance windows with matching lifecycleState.


## Attributes Reference

The following attributes are exported:

* `maintenance_window_collection` - The list of maintenance_window_collection.

### MaintenanceWindow Reference

The following attributes are exported:

* `compartment_id` - Compartment Identifier [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). 
* `description` - Maintenance Window description.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of maintenance window. 
* `lifecycle_details` - Lifecycle Details of the Maintenance Window.
* `name` - Maintenance Window name.
* `resources` - List of resource Ids which are part of the Maintenance Window 
	* `are_members_included` - Flag to indicate if the members of the resource has to be include in the Maintenance Window. 
	* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource part of the Maintenance window. 
* `resources_details` - List of resource details that are part of the Maintenance Window. 
	* `name` - Name of the monitored resource 
	* `number_of_members` - Number of members of the resource 
	* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of monitored resource part of the Maintenance window. 
	* `type` - Type of the monitored resource 
* `schedule` - Schedule information of the Maintenance Window 
	* `maintenance_window_duration` - Duration time of each recurrence of each Maintenance Window. It must be specified as a string in ISO 8601 extended format. 
	* `maintenance_window_recurrences` - A RFC5545 formatted recurrence string which represents the Maintenance Window Recurrence. Please refer this for details:https://datatracker.ietf.org/doc/html/rfc5545#section-3.3.10 FREQ: Frequency of the Maintenance Window. The supported values are: DAILY and WEEKLY. BYDAY: Comma separated days for Weekly Maintenance Window. BYHOUR: Specifies the start hour of each recurrence after `timeMaintenanceWindowStart` value. BYMINUTE: Specifies the start minute of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00 BYSECOND: Specifies the start second of each reccurrence after `timeMaintenanceWindowStart` value. The default value is 00 Other Rules are not supported. 
	* `schedule_type` - Property to identify the type of the Maintenance Window. 
	* `time_maintenance_window_end` - Start time of Maintenance window. A RFC3339 formatted datetime string 
	* `time_maintenance_window_start` - Start time of Maintenance window. A RFC3339 formatted datetime string 
* `state` - Lifecycle state of the monitored resource.
* `time_created` - The time the the maintenance window was created. An RFC3339 formatted datetime string 
* `time_updated` - The time the the mainteance window was updated. An RFC3339 formatted datetime string 

