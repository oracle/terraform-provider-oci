---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_vm_cluster_acd_resource_usages"
sidebar_current: "docs-oci-datasource-database-autonomous_vm_cluster_acd_resource_usages"
description: |-
  Provides the list of Autonomous Vm Cluster Acd Resource Usages in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_vm_cluster_acd_resource_usages
This data source provides the list of Autonomous Vm Cluster Acd Resource Usages in Oracle Cloud Infrastructure Database service.

Gets the list of resource usage details for all the Autonomous Container Database in the specified Autonomous Exadata VM cluster.


## Example Usage

```hcl
data "oci_database_autonomous_vm_cluster_acd_resource_usages" "test_autonomous_vm_cluster_acd_resource_usages" {
	#Required
	autonomous_vm_cluster_id = oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster.id

	#Optional
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_vm_cluster_id` - (Required) The autonomous VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `autonomous_container_database_resource_usages` - The list of autonomous_container_database_resource_usages.

### AutonomousVmClusterAcdResourceUsage Reference

The following attributes are exported:

* `autonomous_container_database_vm_usage` - list of autonomous container database resource usage per autonomous virtual machine.
	* `display_name` - The user-friendly name for the Autonomous VM. The name does not need to be unique.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM.
	* `provisioned_cpus` - CPUs/cores assigned to Autonomous Databases for the ACD instance in given Autonomus VM.
	* `reclaimable_cpus` - CPUs/cores that continue to be included in the count of OCPUs available to the Autonomous Container Database in given Autonomous VM, even after one of its Autonomous Database is terminated or scaled down. You can release them to the available OCPUs at its parent AVMC level by restarting the Autonomous Container Database. 
	* `reserved_cpus` - CPUs/cores reserved for scalability, resilliency and other overheads. This includes failover, autoscaling and idle instance overhead.
	* `used_cpus` - CPUs/cores assigned to the ACD instance in given Autonomous VM. Sum of provisioned, reserved and reclaimable CPUs/ cores to the ACD instance. 
* `available_cpus` - CPUs available for provisioning or scaling an Autonomous Database in the Autonomous Container Database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous Container Database. The name does not need to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Container Database.
* `largest_provisionable_autonomous_database_in_cpus` - Largest provisionable ADB in the Autonomous Container Database.
* `provisionable_cpus` - Valid list of provisionable CPUs for Autonomous Database.
* `provisioned_cpus` - CPUs / cores assigned to ADBs in the Autonomous Container Database.
* `reclaimable_cpus` - Number of CPUs that are reclaimable or released to the AVMC on Autonomous Container Database restart.
* `reserved_cpus` - CPUs / cores reserved for scalability, resilliency and other overheads. This includes failover, autoscaling and idle instance overhead. 
* `used_cpus` - CPUs / cores assigned to the Autonomous Container Database. Sum of provisioned, reserved and reclaimable CPUs/ cores.

