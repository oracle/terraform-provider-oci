---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_vm_cluster_add_virtual_machine"
sidebar_current: "docs-oci-resource-database-vm_cluster_add_virtual_machine"
description: |-
  Provides the Vm Cluster Add Virtual Machine resource in Oracle Cloud Infrastructure Database service
---

# oci_database_vm_cluster_add_virtual_machine
This resource provides the Vm Cluster Add Virtual Machine resource in Oracle Cloud Infrastructure Database service.

Add Virtual Machines to the VM cluster. Applies to Exadata Cloud@Customer instances only.


## Example Usage

```hcl
resource "oci_database_vm_cluster_add_virtual_machine" "test_vm_cluster_add_virtual_machine" {
	#Required
	db_servers {
		#Required
		db_server_id = oci_database_db_server.test_db_server.id
	}
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
}
```

##### Note: You may also need to add `db_servers` and `cpu_core_count` to the ignore_changes for the resource `oci_database_vm_cluster` list if you see a diff on a subsequent apply

## Argument Reference

The following arguments are supported:

* `db_servers` - (Required) The list of Exacc DB servers for the cluster to be added.
	* `db_server_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exacc Db server.
* `vm_cluster_id` - (Required) The VM cluster [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).


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
	* `create` - (Defaults to 20 minutes), when creating the Vm Cluster Add Virtual Machine
	* `update` - (Defaults to 20 minutes), when updating the Vm Cluster Add Virtual Machine
	* `delete` - (Defaults to 20 minutes), when destroying the Vm Cluster Add Virtual Machine


## Import

VmClusterAddVirtualMachine can be imported using the `id`, e.g.

```
$ terraform import oci_database_vm_cluster_add_virtual_machine.test_vm_cluster_add_virtual_machine "id"
```

