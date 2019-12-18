---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_container_database"
sidebar_current: "docs-oci-resource-database-autonomous_container_database"
description: |-
  Provides the Autonomous Container Database resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_container_database
This resource provides the Autonomous Container Database resource in Oracle Cloud Infrastructure Database service.

Create a new Autonomous Container Database in the specified Autonomous Exadata Infrastructure.


## Example Usage

```hcl
resource "oci_database_autonomous_container_database" "test_autonomous_container_database" {
	#Required
	autonomous_exadata_infrastructure_id = "${oci_database_autonomous_exadata_infrastructure.test_autonomous_exadata_infrastructure.id}"
	display_name = "${var.autonomous_container_database_display_name}"
	patch_model = "${var.autonomous_container_database_patch_model}"

	#Optional
	backup_config {

		#Optional
		recovery_window_in_days = "${var.autonomous_container_database_backup_config_recovery_window_in_days}"
	}
	compartment_id = "${var.compartment_id}"
	defined_tags = {"Operations.CostCenter"= "42"}
	freeform_tags = {"Department"= "Finance"}
	maintenance_window_details {
		#Required
		preference = "${var.autonomous_container_database_maintenance_window_details_preference}"

		#Optional
		days_of_week {
			#Required
			name = "${var.autonomous_container_database_maintenance_window_details_days_of_week_name}"
		}
		hours_of_day = "${var.autonomous_container_database_maintenance_window_details_hours_of_day}"
		lead_time_in_weeks = "${var.autonomous_container_database_maintenance_window_details_lead_time_in_weeks}"
		months {
			#Required
			name = "${var.autonomous_container_database_maintenance_window_details_months_name}"
		}
		weeks_of_month = "${var.autonomous_container_database_maintenance_window_details_weeks_of_month}"
	}
	service_level_agreement_type = "${var.autonomous_container_database_service_level_agreement_type}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_exadata_infrastructure_id` - (Required) The OCID of the Autonomous Exadata Infrastructure.
* `backup_config` - (Optional) (Updatable) 
	* `recovery_window_in_days` - (Optional) (Updatable) Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Autonomous Container Database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The display name for the Autonomous Container Database.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `maintenance_window_details` - (Optional) (Updatable) 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `preference` - (Required) (Updatable) The maintenance window scheduling preference.
	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `patch_model` - (Required) (Updatable) Database Patch model preference.
* `service_level_agreement_type` - (Optional) The service level agreement type of the Autonomous Container Database. The default is STANDARD. For a mission critical Autonomous Container Database, the specified Autonomous Exadata Infrastructure must be associated with a remote Autonomous Exadata Infrastructure.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_exadata_infrastructure_id` - The OCID of the Autonomous Exadata Infrastructure.
* `availability_domain` - The availability domain of the Autonomous Container Database.
* `backup_config` - 
	* `recovery_window_in_days` - Number of days between the current and the earliest point of recoverability covered by automatic backups. This value applies to automatic backups. After a new automatic backup has been created, Oracle removes old automatic backups that are created before the window. When the value is updated, it is applied to all existing automatic backups. 
* `compartment_id` - The OCID of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-provided name for the Autonomous Container Database.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Autonomous Container Database.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `maintenance_window` - 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `patch_model` - Database patch model preference.
* `service_level_agreement_type` - The service level agreement type of the container database. The default is STANDARD.
* `state` - The current state of the Autonomous Container Database.
* `time_created` - The date and time the Autonomous Container Database was created.

## Import

AutonomousContainerDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_container_database.test_autonomous_container_database "id"
```

