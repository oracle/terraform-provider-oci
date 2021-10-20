---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster"
sidebar_current: "docs-oci-resource-database-vm_cluster"
description: |-
  Provides the Vm Cluster resource in Oracle Cloud Infrastructure Database service
---

# oci_database_vm_cluster
This resource provides the Vm Cluster resource in Oracle Cloud Infrastructure Database service.

Creates an Exadata Cloud@Customer VM cluster.


## Example Usage

```hcl
resource "oci_database_vm_cluster" "test_vm_cluster" {
	#Required
	compartment_id = var.compartment_id
	cpu_core_count = var.vm_cluster_cpu_core_count
	display_name = var.vm_cluster_display_name
	exadata_infrastructure_id = oci_database_exadata_infrastructure.test_exadata_infrastructure.id
	gi_version = var.vm_cluster_gi_version
	ssh_public_keys = var.vm_cluster_ssh_public_keys
	vm_cluster_network_id = oci_database_vm_cluster_network.test_vm_cluster_network.id

	#Optional
	data_storage_size_in_tbs = var.vm_cluster_data_storage_size_in_tbs
	db_node_storage_size_in_gbs = var.vm_cluster_db_node_storage_size_in_gbs
	db_servers = var.vm_cluster_db_servers
	defined_tags = var.vm_cluster_defined_tags
	freeform_tags = {"Department"= "Finance"}
	is_local_backup_enabled = var.vm_cluster_is_local_backup_enabled
	is_sparse_diskgroup_enabled = var.vm_cluster_is_sparse_diskgroup_enabled
	license_model = var.vm_cluster_license_model
	memory_size_in_gbs = var.vm_cluster_memory_size_in_gbs
	time_zone = var.vm_cluster_time_zone
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpu_core_count` - (Required) (Updatable) The number of CPU cores to enable for the VM cluster. *Note:* If `cpu_core_count` is modified in `DISCONNECTED` state, the provider could experience a drift in Terraform state. To remediate this, refresh your Terraform state and update the configuration file when the Oracle Cloud Infrastructure connection is established.
* `data_storage_size_in_tbs` - (Optional) (Updatable) The data disk group size to be allocated in TBs.
* `db_node_storage_size_in_gbs` - (Optional) (Updatable) The local node storage to be allocated in GBs.
* `db_servers` - (Optional) The list of Db server.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) The user-friendly name for the VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - (Required) The Oracle Grid Infrastructure software version for the VM cluster.
* `is_local_backup_enabled` - (Optional) If true, database backup on local Exadata storage is configured for the VM cluster. If false, database backup on local Exadata storage is not available in the VM cluster. 
* `is_sparse_diskgroup_enabled` - (Optional) If true, the sparse disk group is configured for the VM cluster. If false, the sparse disk group is not created. 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the VM cluster. The default is BRING_YOUR_OWN_LICENSE. 
* `memory_size_in_gbs` - (Optional) (Updatable) The memory to be allocated in GBs.
* `ssh_public_keys` - (Required) (Updatable) The public key portion of one or more key pairs used for SSH access to the VM cluster.
* `time_zone` - (Optional) The time zone to use for the VM cluster. For details, see [DB System Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vm_cluster_network_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `cpus_enabled` - The number of enabled CPU cores.
* `data_storage_size_in_tbs` - Size, in terabytes, of the DATA disk group.
* `db_node_storage_size_in_gbs` - The local node storage allocated in GBs.
* `db_servers` - The list of Db server.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the Exadata Cloud@Customer VM cluster. The name does not need to be unique.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `gi_version` - The Oracle Grid Infrastructure software version for the VM cluster.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
* `is_local_backup_enabled` - If true, database backup on local Exadata storage is configured for the VM cluster. If false, database backup on local Exadata storage is not available in the VM cluster. 
* `is_sparse_diskgroup_enabled` - If true, sparse disk group is configured for the VM cluster. If false, sparse disk group is not created. 
* `last_patch_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation starts.
* `license_model` - The Oracle license model that applies to the VM cluster. The default is LICENSE_INCLUDED. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `memory_size_in_gbs` - The memory allocated in GBs.
* `shape` - The shape of the Exadata infrastructure. The shape determines the amount of CPU, storage, and memory resources allocated to the instance. 
* `ssh_public_keys` - The public key portion of one or more key pairs used for SSH access to the VM cluster.
* `state` - The current state of the VM cluster.
* `system_version` - Operating system version of the image.
* `time_created` - The date and time that the VM cluster was created.
* `time_zone` - The time zone of the Exadata infrastructure. For details, see [Exadata Infrastructure Time Zones](https://docs.cloud.oracle.com/iaas/Content/Database/References/timezones.htm).
* `vm_cluster_network_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster network.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Vm Cluster
	* `update` - (Defaults to 20 minutes), when updating the Vm Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Vm Cluster


## Import

VmClusters can be imported using the `id`, e.g.

```
$ terraform import oci_database_vm_cluster.test_vm_cluster "id"
```

