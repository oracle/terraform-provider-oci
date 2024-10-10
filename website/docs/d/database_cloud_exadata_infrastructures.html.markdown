---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_exadata_infrastructures"
sidebar_current: "docs-oci-datasource-database-cloud_exadata_infrastructures"
description: |-
  Provides the list of Cloud Exadata Infrastructures in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_cloud_exadata_infrastructures
This data source provides the list of Cloud Exadata Infrastructures in Oracle Cloud Infrastructure Database service.

Gets a list of the cloud Exadata infrastructure resources in the specified compartment. Applies to Exadata Cloud Service instances and Autonomous Database on dedicated Exadata infrastructure only.


## Example Usage

```hcl
data "oci_database_cloud_exadata_infrastructures" "test_cloud_exadata_infrastructures" {
	#Required
	compartment_id = var.compartment_id

	#Optional
    cluster_placement_group_id = var.cloud_exadata_infrastructure_cluster_placement_group_id
	display_name = var.cloud_exadata_infrastructure_display_name
	state = var.cloud_exadata_infrastructure_state
}
```

## Argument Reference

The following arguments are supported:

* `cluster_placement_group_id` - (Optional) A filter to return only resources that match the given cluster placement group ID exactly.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `cloud_exadata_infrastructures` - The list of cloud_exadata_infrastructures.

### CloudExadataInfrastructure Reference

The following attributes are exported:

* `activated_storage_count` - The requested number of additional storage servers activated for the Exadata infrastructure.
* `additional_storage_count` - The requested number of additional storage servers for the Exadata infrastructure.
* `availability_domain` - The name of the availability domain that the cloud Exadata infrastructure resource is located in.
* `available_storage_size_in_gbs` - The available storage can be allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).
* `cluster_placement_group_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cluster placement group of the Exadata Infrastructure.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_count` - The number of compute servers for the cloud Exadata infrastructure.
* `cpu_count` - The total number of CPU cores allocated.
* `customer_contacts` - The list of customer email addresses that receive information from Oracle about the specified Oracle Cloud Infrastructure Database service resource. Oracle uses these email addresses to send notifications about planned and unplanned software maintenance updates, information about system hardware, and other information needed by administrators. Up to 10 email addresses can be added to the customer contacts for a cloud Exadata infrastructure instance. 
	* `email` - The email address used by Oracle to send notifications regarding databases and infrastructure.
* `data_storage_size_in_tbs` - Size, in terabytes, of the DATA disk group. 
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `db_server_version` - The software version of the database servers (dom0) in the cloud Exadata infrastructure. Example: 20.1.15 
* `defined_file_system_configurations` - Details of the file system configuration of the Exadata infrastructure.
	* `is_backup_partition` - If true, the file system is used to create a backup prior to Exadata VM OS update.
	* `is_resizable` - If true, the file system resize is allowed for the Exadata Infrastructure cluster. If false, the file system resize is not allowed.
	* `min_size_gb` - The minimum size of file system.
	* `mount_point` - The mount point of file system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the cloud Exadata infrastructure resource. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure resource.
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `custom_action_timeout_in_mins` - Determines the amount of time the system will wait before the start of each database server patching operation. Custom action timeout is in minutes and valid value is between 15 to 120 (inclusive). 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `is_custom_action_timeout_enabled` - If true, enables the configuration of a custom action timeout (waiting period) between database server patching operations.
	* `is_monthly_patching_enabled` - If true, enables the monthly patching option.
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `patching_mode` - Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - The maintenance window scheduling preference.
	* `skip_ru` - If true, skips the release update (RU) for the quarter. You cannot skip two consecutive quarters. An RU skip request will only be honoured if the current version of the Autonomous Container Database is supported for current quarter. 
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `max_cpu_count` - The total number of CPU cores available.
* `max_data_storage_in_tbs` - The total available DATA disk group size.
* `max_db_node_storage_in_gbs` - The total local node storage available in GBs.
* `max_memory_in_gbs` - The total memory available in GBs.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `monthly_db_server_version` - The monthly software version of the database servers (dom0) in the cloud Exadata infrastructure. Example: 20.1.15 
* `monthly_storage_server_version` - The monthly software version of the storage servers (cells) in the cloud Exadata infrastructure. Example: 20.1.15 
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `shape` - The model name of the cloud Exadata infrastructure resource. 
* `state` - The current lifecycle state of the cloud Exadata infrastructure resource.
* `storage_count` - The number of storage servers for the cloud Exadata infrastructure.
* `storage_server_version` - The software version of the storage servers (cells) in the cloud Exadata infrastructure. Example: 20.1.15 
* `subscription_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subscription with which resource needs to be associated with.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the cloud Exadata infrastructure resource was created.
* `total_storage_size_in_gbs` - The total storage allocated to the cloud Exadata infrastructure resource, in gigabytes (GB).

