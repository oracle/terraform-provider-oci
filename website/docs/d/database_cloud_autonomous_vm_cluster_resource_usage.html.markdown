---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cloud_autonomous_vm_cluster_resource_usage"
sidebar_current: "docs-oci-datasource-database-cloud_autonomous_vm_cluster_resource_usage"
description: |-
  Provides details about a specific Cloud Autonomous Vm Cluster Resource Usage in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_cloud_autonomous_vm_cluster_resource_usage
This data source provides details about a specific Cloud Autonomous Vm Cluster Resource Usage resource in Oracle Cloud Infrastructure Database service.

Get the resource usage details for the specified Cloud Autonomous Exadata VM cluster.


## Example Usage

```hcl
data "oci_database_cloud_autonomous_vm_cluster_resource_usage" "test_cloud_autonomous_vm_cluster_resource_usage" {
	#Required
	cloud_autonomous_vm_cluster_id = oci_database_cloud_autonomous_vm_cluster.test_cloud_autonomous_vm_cluster.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_autonomous_vm_cluster_id` - (Required) The Cloud VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_data_storage_size_in_tbs` - The data disk group size allocated for Autonomous Databases, in TBs.
* `autonomous_vm_resource_usage` - List of Autonomous VM resource usages.
	* `autonomous_container_database_usage` - Associated Autonomous Container Database Usages.
		* `available_cpus` - The number of CPU cores available.
		* `display_name` - The user-friendly name for the Autonomous Container Database. The name does not need to be unique.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
		* `provisioned_cpus` - CPUs/cores assigned to Autonomous Databases in the ACD instances.
		* `reclaimable_cpus` - CPUs/cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database. 
		* `reserved_cpus` - CPUs/cores reserved for scalability, resilliency and other overheads. This includes failover, autoscaling and idle instance overhead. 
		* `used_cpus` - CPUs/cores assigned to the ACD instance. Sum of provisioned, reserved and reclaimable CPUs/ cores to the ACD instance. 
	* `available_cpus` - The number of CPU cores available.
	* `display_name` - The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM Cluster.
	* `provisioned_cpus` - The number of CPUs provisioned in an Autonomous VM Cluster.
	* `reclaimable_cpus` - CPU cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database. 
	* `reserved_cpus` - The number of CPUs reserved in an Autonomous VM Cluster.
	* `used_cpus` - The number of CPU cores alloted to the Autonomous Container Databases in an Cloud Autonomous VM cluster. 
* `available_autonomous_data_storage_size_in_tbs` - The data disk group size available for Autonomous Databases, in TBs.
* `available_cpus` - The number of CPU cores available.
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `display_name` - The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
* `exadata_storage_in_tbs` - Total exadata storage allocated for the Autonomous VM Cluster. DATA + RECOVERY + SPARSE + any overhead in TBs.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Cloud Autonomous VM cluster.
* `memory_per_oracle_compute_unit_in_gbs` - The amount of memory (in GBs) to be enabled per each CPU core.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `non_provisionable_autonomous_container_databases` - The number of non-provisionable Autonomous Container Databases in an Autonomous VM Cluster.
* `provisionable_autonomous_container_databases` - The number of provisionable Autonomous Container Databases in an Autonomous VM Cluster.
* `provisioned_autonomous_container_databases` - The number of provisioned Autonomous Container Databases in an Autonomous VM Cluster.
* `provisioned_cpus` - The number of CPUs provisioned in an Autonomous VM Cluster.
* `reclaimable_cpus` - CPU cores that continue to be included in the count of OCPUs available to the Autonomous Container Database even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database. 
* `reserved_cpus` - The number of CPUs reserved in an Autonomous VM Cluster.
* `total_container_databases` - The total number of Autonomous Container Databases that can be created.
* `total_cpus` - The number of CPU cores enabled on the Cloud Autonomous VM cluster.
* `used_autonomous_data_storage_size_in_tbs` - The data disk group size used for Autonomous Databases, in TBs.
* `used_cpus` - The number of CPU cores alloted to the Autonomous Container Databases in an Cloud Autonomous VM cluster. 

