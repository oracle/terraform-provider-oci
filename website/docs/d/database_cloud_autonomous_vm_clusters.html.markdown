---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_autonomous_vm_clusters"
sidebar_current: "docs-oci-datasource-database-cloud_autonomous_vm_clusters"
description: |-
  Provides the list of Cloud Autonomous Vm Clusters in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_cloud_autonomous_vm_clusters
This data source provides the list of Cloud Autonomous Vm Clusters in Oracle Cloud Infrastructure Database service.

Lists Autonomous Exadata VM clusters in the Oracle cloud. For Exadata Cloud@Customer systems, see [ListAutonomousVmClusters](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/AutonomousVmCluster/ListAutonomousVmClusters).


## Example Usage

```hcl
data "oci_database_cloud_autonomous_vm_clusters" "test_cloud_autonomous_vm_clusters" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	availability_domain = var.cloud_autonomous_vm_cluster_availability_domain
	cloud_exadata_infrastructure_id = oci_database_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
	display_name = var.cloud_autonomous_vm_cluster_display_name
	state = var.cloud_autonomous_vm_cluster_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) A filter to return only resources that match the given availability domain exactly.
* `cloud_exadata_infrastructure_id` - (Optional) If provided, filters the results for the specified cloud Exadata infrastructure.
* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `cloud_autonomous_vm_clusters` - The list of cloud_autonomous_vm_clusters.

### CloudAutonomousVmCluster Reference

The following attributes are exported:

* `autonomous_data_storage_percentage` - The percentage of the data storage used for the Autonomous Databases in an Autonomous VM Cluster.
* `autonomous_data_storage_size_in_tbs` - The data disk group size allocated for Autonomous Databases, in TBs.
* `availability_domain` - The name of the availability domain that the cloud Autonomous VM cluster is located in.
* `available_autonomous_data_storage_size_in_tbs` - The data disk group size available for Autonomous Databases, in TBs.
* `available_container_databases` - The number of Autonomous Container Databases that can be created with the currently available local storage.
* `available_cpus` - CPU cores available for allocation to Autonomous Databases.
* `cloud_exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud Exadata infrastructure.
* `cluster_time_zone` - The time zone of the Cloud Autonomous VM Cluster.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compute_model` - The compute model of the Cloud Autonomous VM Cluster. ECPU compute model is the recommended model and OCPU compute model is legacy. See [Compute Models in Autonomous Database on Dedicated Exadata #Infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/dedicated/adbak) for more details.
* `cpu_core_count` - The number of CPU cores on the cloud Autonomous VM cluster.
* `cpu_core_count_per_node` - The number of CPU cores enabled per VM cluster node.
* `cpu_percentage` - The percentage of total number of CPUs used in an Autonomous VM Cluster.
* `data_storage_size_in_gb` - The total data storage allocated, in gigabytes (GB).
* `data_storage_size_in_tbs` - The total data storage allocated, in terabytes (TB).
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `db_servers` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Db servers.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - User defined description of the cloud Autonomous VM cluster.
* `display_name` - The user-friendly name for the cloud Autonomous VM cluster. The name does not need to be unique.
* `domain` - The domain name for the cloud Autonomous VM cluster.
* `exadata_storage_in_tbs_lowest_scaled_value` - The lowest value to which exadataStorage (in TBs) can be scaled down.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `hostname` - The hostname for the cloud Autonomous VM cluster.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Autonomous VM cluster.
* `is_mtls_enabled_vm_cluster` - Enable mutual TLS(mTLS) authentication for database at time of provisioning a VMCluster. This is applicable to database TLS Certificates only. Default is TLS
* `last_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance run.
* `last_update_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last maintenance update history. This value is updated when a maintenance update starts.
* `license_model` - The Oracle license model that applies to the Oracle Autonomous Database. Bring your own license (BYOL) allows you to apply your current on-premises Oracle software licenses to equivalent, highly automated Oracle services in the cloud. License Included allows you to subscribe to new Oracle Database software licenses and the Oracle Database service. Note that when provisioning an [Autonomous Database on dedicated Exadata infrastructure](https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html), this attribute must be null. It is already set at the Autonomous Exadata Infrastructure level. When provisioning an [Autonomous Database Serverless] (https://docs.oracle.com/en/cloud/paas/autonomous-database/index.html) database, if a value is not specified, the system defaults the value to `BRING_YOUR_OWN_LICENSE`. Bring your own license (BYOL) also allows you to select the DB edition using the optional parameter.

	This cannot be updated in parallel with any of the following: cpuCoreCount, computeCount, maxCpuCoreCount, dataStorageSizeInTBs, adminPassword, isMTLSConnectionRequired, dbWorkload, privateEndpointLabel, nsgIds, dbVersion, dbName, scheduledOperations, dbToolsDetails, or isFreeTier. 
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
* `max_acds_lowest_scaled_value` - The lowest value to which maximum number of ACDs can be scaled down.
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) enabled per OCPU or ECPU.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `next_maintenance_run_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the next maintenance run.
* `node_count` - The number of database servers in the cloud VM cluster. 
* `non_provisionable_autonomous_container_databases` - The number of non-provisionable Autonomous Container Databases in an Autonomous VM Cluster.
* `nsg_ids` - The list of [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) for the network security groups (NSGs) to which this resource belongs. Setting this to an empty list removes all resources from all NSGs. For more information about NSGs, see [Security Rules](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/securityrules.htm). **NsgIds restrictions:**
	* A network security group (NSG) is optional for Autonomous Databases with private access. The nsgIds list can be empty. 
* `ocpu_count` - The number of CPU cores on the cloud Autonomous VM cluster. Only 1 decimal place is allowed for the fractional part.
* `ocpus_lowest_scaled_value` - The lowest value to which ocpus can be scaled down.
* `provisionable_autonomous_container_databases` - The number of provisionable Autonomous Container Databases in an Autonomous VM Cluster.
* `provisioned_autonomous_container_databases` - The number of provisioned Autonomous Container Databases in an Autonomous VM Cluster.
* `provisioned_cpus` - The number of CPUs provisioned in an Autonomous VM Cluster.
* `reclaimable_cpus` - For Autonomous Databases on Dedicated Exadata Infrastructure:
	* These are the CPUs that continue to be included in the count of CPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available CPUs at its parent Autonomous VM Cluster level by restarting the Autonomous Container Database.
	* The CPU type (OCPUs or ECPUs) is determined by the parent Autonomous Exadata VM Cluster's compute model.  
* `reserved_cpus` - The number of CPUs reserved in an Autonomous VM Cluster.
* `scan_listener_port_non_tls` - The SCAN Listener Non TLS port. Default is 1521.
* `scan_listener_port_tls` - The SCAN Listenenr TLS port. Default is 2484.
* `security_attributes` - Security Attributes for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Oracle-ZPR": {"MaxEgressCount": {"value": "42", "mode": "audit"}}}` 
* `shape` - The model name of the Exadata hardware running the cloud Autonomous VM cluster. 
* `state` - The current state of the cloud Autonomous VM cluster.
* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the subnet the cloud Autonomous VM Cluster is associated with.

	**Subnet Restrictions:**
	* For Exadata and virtual machine 2-node RAC DB systems, do not use a subnet that overlaps with 192.168.128.0/20.

	These subnets are used by the Oracle Clusterware private interconnect on the database instance. Specifying an overlapping subnet will cause the private interconnect to malfunction. This restriction applies to both the client subnet and backup subnet. 
* `time_created` - The date and time that the cloud Autonomous VM cluster was created.
* `time_database_ssl_certificate_expires` - The date and time of Database SSL certificate expiration.
* `time_ords_certificate_expires` - The date and time of ORDS certificate expiration.
* `time_updated` - The last date and time that the cloud Autonomous VM cluster was updated.
* `total_autonomous_data_storage_in_tbs` - The total data disk group size for Autonomous Databases, in TBs.
* `total_container_databases` - The total number of Autonomous Container Databases that can be created with the allocated local storage.
* `total_cpus` - The total number of CPUs in an Autonomous VM Cluster.

