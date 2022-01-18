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
	defined_tags = var.autonomous_vm_cluster_defined_tags
	freeform_tags = {"Department"= "Finance"}
	is_local_backup_enabled = var.autonomous_vm_cluster_is_local_backup_enabled
	license_model = var.autonomous_vm_cluster_license_model
	time_zone = var.autonomous_vm_cluster_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_local_backup_enabled` - (Optional) If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster. 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the Autonomous VM cluster. The default is BRING_YOUR_OWN_LICENSE. 
* `time_zone` - (Optional) The time zone to use for the Autonomous VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vm_cluster_network_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `available_cpus` - The numnber of CPU cores available.
* `available_data_storage_size_in_tbs` - The data storage available in TBs
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpus_enabled` - The number of enabled CPU cores.
* `data_storage_size_in_tbs` - The total data storage allocated in TBs
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Autonomous VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous VM cluster.
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the Autonomous VM cluster. If false, database backup on local Exadata storage is not available in the Autonomous VM cluster. 
* `license_model` - The Oracle license model that applies to the Autonomous VM cluster. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `state` - The current state of the Autonomous VM cluster.
* `time_created` - The date and time that the Autonomous VM cluster was created.
* `time_zone` - The time zone to use for the Autonomous VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vm_cluster_network_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Autonomous Vm Cluster
	* `update` - (Defaults to 20 minutes), when updating the Autonomous Vm Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Autonomous Vm Cluster


## Import

AutonomousVmClusters can be imported using the `id`, e.g.

```
$ terraform import oci_database_autonomous_vm_cluster.test_autonomous_vm_cluster "id"
```

