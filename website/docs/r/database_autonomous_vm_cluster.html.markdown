---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_vm_cluster"
sidebar_current: "docs-oci-resource-database-autonomous_vm_cluster"
description: |-
  Provides the Autonomous Vm Cluster resource in Oracle Cloud Infrastructure Database service
---

# oci_database_autonomous_vm_cluster
This resource provides the Autonomous Vm Cluster resource in Oracle Cloud Infrastructure Database service.

Creates an Autonomous VM cluster for Exadata Cloud@Customer. To create an Autonomous VM Cluster in the Oracle cloud, see [CreateCloudAutonomousVmCluster](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/CloudAutonomousVmCluster/CreateCloudAutonomousVmCluster).


## Example Usage

```hcl
resource "oci_database_autonomous_vm_cluster" "test_autonomous_vm_cluster" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.autonomous_vm_cluster_display_name
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	vm_cluster_network_id = oci_database_vm_cluster_network.test_vm_cluster_network.id

	#Optional
	autonomous_data_storage_size_in_tbs = var.autonomous_vm_cluster_autonomous_data_storage_size_in_tbs
	compute_model = var.autonomous_vm_cluster_compute_model
	cpu_core_count_per_node = var.autonomous_vm_cluster_cpu_core_count_per_node
	db_servers = var.autonomous_vm_cluster_db_servers
	defined_tags = var.autonomous_vm_cluster_defined_tags
	freeform_tags = {"Department"= "Finance"}
	is_local_backup_enabled = var.autonomous_vm_cluster_is_local_backup_enabled
	is_mtls_enabled = var.autonomous_vm_cluster_is_mtls_enabled
	license_model = var.autonomous_vm_cluster_license_model
	maintenance_window_details {
        #Required

		#Optional
		days_of_week {
			#Required
			name = var.autonomous_vm_cluster_maintenance_window_details_days_of_week_name
		}
		hours_of_day = var.autonomous_vm_cluster_maintenance_window_details_hours_of_day
		lead_time_in_weeks = var.autonomous_vm_cluster_maintenance_window_details_lead_time_in_weeks
		months {
			#Required
			name = var.autonomous_vm_cluster_maintenance_window_details_months_name
		}
		patching_mode = var.autonomous_vm_cluster_maintenance_window_details_patching_mode
		preference = var.autonomous_vm_cluster_maintenance_window_details_preference
		weeks_of_month = var.autonomous_vm_cluster_maintenance_window_details_weeks_of_month
	}
	memory_per_oracle_compute_unit_in_gbs = var.autonomous_vm_cluster_memory_per_oracle_compute_unit_in_gbs
	scan_listener_port_non_tls = var.autonomous_vm_cluster_scan_listener_port_non_tls
	scan_listener_port_tls = var.autonomous_vm_cluster_scan_listener_port_tls
	time_zone = var.autonomous_vm_cluster_time_zone
	total_container_databases = var.autonomous_vm_cluster_total_container_databases
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_data_storage_size_in_tbs` - (Optional) (Updatable) The data disk group size to be allocated for Autonomous Databases, in TBs.
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - (Optional) The compute model of the Autonomous VM Cluster. ECPU compute model is the recommended model and OCPU compute model is legacy. 
* `cpu_core_count_per_node` - (Optional) (Updatable) The number of CPU cores to enable per VM cluster node.
* `db_servers` - (Optional) The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db servers.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_local_backup_enabled` - (Optional) If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster. 
* `is_mtls_enabled` - (Optional) Enable mutual TLS(mTLS) authentication for database while provisioning a VMCluster. Default is TLS.
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Autonomous VM cluster. The default is BRING_YOUR_OWN_LICENSE. 
* `maintenance_window_details` - (Optional) (Updatable) The scheduling details for the quarterly maintenance window. Patching and system updates take place during the maintenance window. 
	* `days_of_week` - (Optional) (Updatable) Days during the week when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the day of the week.
	* `hours_of_day` - (Optional) (Updatable) The window of hours during the day when maintenance should be performed. The window is a 4 hour slot. Valid values are
		* 0 - represents time slot 0:00 - 3:59 UTC - 4 - represents time slot 4:00 - 7:59 UTC - 8 - represents time slot 8:00 - 11:59 UTC - 12 - represents time slot 12:00 - 15:59 UTC - 16 - represents time slot 16:00 - 19:59 UTC - 20 - represents time slot 20:00 - 23:59 UTC
	* `lead_time_in_weeks` - (Optional) (Updatable) Lead time window allows user to set a lead time to prepare for a down time. The lead time is in weeks and valid value is between 1 to 4. 
	* `months` - (Optional) (Updatable) Months during the year when maintenance should be performed.
		* `name` - (Required) (Updatable) Name of the month of the year.
	* `patching_mode` - (Optional) (Updatable) Cloud Exadata infrastructure node patching method, either "ROLLING" or "NONROLLING". Default value is ROLLING.

		*IMPORTANT*: Non-rolling infrastructure patching involves system down time. See [Oracle-Managed Infrastructure Maintenance Updates](https://docs.cloud.oracle.com/iaas/Content/Database/Concepts/examaintenance.htm#Oracle) for more information. 
	* `preference` - (Optional) (Updatable) The maintenance window scheduling preference.

	* `weeks_of_month` - (Optional) (Updatable) Weeks during the month when maintenance should be performed. Weeks start on the 1st, 8th, 15th, and 22nd days of the month, and have a duration of 7 days. Weeks start and end based on calendar dates, not days of the week. For example, to allow maintenance during the 2nd week of the month (from the 8th day to the 14th day of the month), use the value 2. Maintenance cannot be scheduled for the fifth week of months that contain more than 28 days. Note that this parameter works in conjunction with the  daysOfWeek and hoursOfDay parameters to allow you to specify specific days of the week and hours that maintenance will be performed. 
* `memory_per_oracle_compute_unit_in_gbs` - (Optional) The amount of memory (in GBs) to be enabled per OCPU or ECPU.  
* `scan_listener_port_non_tls` - (Optional) The SCAN Listener Non TLS port number. Default value is 1521.
* `scan_listener_port_tls` - (Optional) The SCAN Listener TLS port number. Default value is 2484.
* `time_zone` - (Optional) The time zone to use for the Autonomous VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `total_container_databases` - (Optional) (Updatable) The total number of Autonomous Container Databases that can be created.
* `vm_cluster_network_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `autonomous_data_storage_size_in_tbs` - The data disk group size allocated for Autonomous Databases, in TBs.
* `available_autonomous_data_storage_size_in_tbs` - The data disk group size available for Autonomous Databases, in TBs.
* `available_container_databases` - The number of Autonomous Container Databases that can be created with the currently available local storage.
* `available_cpus` - The numnber of CPU cores available.
* `available_data_storage_size_in_tbs` - **Deprecated.** Use `availableAutonomousDataStorageSizeInTBs` for Autonomous Databases' data storage availability in TBs. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - The compute model of the Autonomous VM Cluster. ECPU compute model is the recommended model and OCPU compute model is legacy. See [Compute Models in Autonomous Database on Dedicated Exadata #Infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details.
* `cpu_core_count_per_node` - The number of CPU cores enabled per VM cluster node.
* `cpus_enabled` - The number of enabled CPU cores.
* `data_storage_size_in_tbs` - The total data storage allocated in TBs
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `db_servers` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db servers.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `exadata_storage_in_tbs_lowest_scaled_value` - The lowest value to which exadataStorage(in TBs) can be scaled down.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM cluster.
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster. 
* `is_mtls_enabled` - Enable mutual TLS(mTLS) authentication for database while provisioning a VMCluster. Default is TLS.
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
* `max_acds_lowest_scaled_value` - The lowest value to which maximum number of ACDs can be scaled down.
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) to be enabled per OCPU or ECPU.  
* `memory_size_in_gbs` - The memory allocated in GBs.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `node_count` - The number of nodes in the Autonomous VM Cluster. 
* `ocpus_enabled` - The number of enabled OCPU cores.
* `provisionable_autonomous_container_databases` - **Deprecated.** Use field totalContainerDatabases. 
* `provisioned_autonomous_container_databases` - The number of provisioned Autonomous Container Databases in an Autonomous VM Cluster.
* `provisioned_cpus` - The number of CPUs provisioned in an Autonomous VM Cluster.
* `reclaimable_cpus` - CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database. 
* `reserved_cpus` - The number of CPUs reserved in an Autonomous VM Cluster.
* `scan_listener_port_non_tls` - The SCAN Listener Non TLS port number. Default value is 1521.
* `scan_listener_port_tls` - The SCAN Listener TLS port number. Default value is 2484.
* `state` - The current state of the Autonomous VM cluster.
* `time_created` - The date and time that the Autonomous VM cluster was created.
* `time_database_ssl_certificate_expires` - The date and time of Database SSL certificate expiration.
* `time_ords_certificate_expires` - The date and time of ORDS certificate expiration.
* `time_zone` - The time zone to use for the Autonomous VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `total_container_databases` - The total number of Autonomous Container Databases that can be created.
* `vm_cluster_network_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Vm Cluster
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Vm Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Vm Cluster


## Import

AutonomousVmClusters can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster "id"
```

