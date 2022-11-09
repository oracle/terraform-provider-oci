---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_vm_clusters"
sidebar_current: "docs-oci-datasource-database-autonomous_vm_clusters"
description: |-
  Provides the list of Autonomous Vm Clusters in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_vm_clusters
This data source provides the list of Autonomous Vm Clusters in Oracle Cloud Infrastructure Database service.

Gets a list of Exadata Cloud@Customer Autonomous VM clusters in the specified compartment. To list Autonomous VM Clusters in the Oracle Cloud, see [ListCloudAutonomousVmClusters](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudAutonomousVmCluster/ListCloudAutonomousVmClusters).


## Example Usage

```hcl
data "oci_database_autonomous_vm_clusters" "test_autonomous_vm_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.autonomous_vm_cluster_display_name
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	state = var.autonomous_vm_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `exadata_infrastructure_id` - (Optional) If provided, filters the results for the given Exadata Infrastructure.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `autonomous_vm_clusters` - The list of autonomous_vm_clusters.

### AutonomousVmCluster Reference

The following attributes are exported:

* `autonomous_data_storage_size_in_tbs` - The data disk group size allocated for Autonomous Databases, in TBs.
* `available_autonomous_data_storage_size_in_tbs` - The data disk group size available for Autonomous Databases, in TBs.
* `available_container_databases` - The number of Autonomous Container Databases that can be created with the currently available local storage.
* `available_cpus` - The numnber of CPU cores available.
* `available_data_storage_size_in_tbs` - **Deprecated.** Use `availableAutonomousDataStorageSizeInTBs` for Autonomous Databases' data storage availability in TBs. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count_per_node` - The number of CPU cores enabled per VM cluster node.
* `cpus_enabled` - The number of enabled CPU cores.
* `data_storage_size_in_gb` - The total data storage allocated in GBs.
* `data_storage_size_in_tbs` - The total data storage allocated in TBs
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM cluster.
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster. 
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `license_model` - The Oracle license model that applies to the Autonomous VM cluster. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `maintenance_window` - The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `days_of_week` - Days during the week when maintenance should be performed.
		* `name` - Name of the day of the week.
	* `hours_of_day` - The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - Months during the year when maintenance should be performed.
		* `name` - Name of the month of the year.
	* `preference` - The maintenance window scheduling preference.
	* `weeks_of_month` - Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) enabled per each OCPU core.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `ocpus_enabled` - The number of enabled OCPU cores.
* `reclaimable_cpus` - CPU cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database.
* `state` - The current state of the Autonomous VM cluster.
* `time_created` - The date and time that the Autonomous VM cluster was created.
* `time_zone` - The time zone to use for the Autonomous VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `total_container_databases` - The total number of Autonomous Container Databases that can be created.
* `vm_cluster_network_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.

