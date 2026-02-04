---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_advanced_cluster_file_system"
sidebar_current: "docs-oci-resource-database-advanced_cluster_file_system"
description: |-
  Provides the Advanced Cluster File System resource in Oracle Cloud Infrastructure Database service
---

# oci_database_advanced_cluster_file_system
This resource provides the Advanced Cluster File System resource in Oracle Cloud Infrastructure Database service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database/latest/AdvancedClusterFileSystem

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database

Creates an advanced cluster file system resource.


## Example Usage

```hcl
resource "oci_database_advanced_cluster_file_system" "test_advanced_cluster_file_system" {
	#Required
	name = var.advanced_cluster_file_system_name
	storage_in_gbs = var.advanced_cluster_file_system_storage_in_gbs
	vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id

	#Optional
	compartment_id = var.compartment_id
	defined_tags = var.advanced_cluster_file_system_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `name` - (Required) The user-friendly name for the Advanced cluster file system. The name has to be unique.
* `storage_in_gbs` - (Required) (Updatable) The total storage needed for advanced cluster file system in GBs.
* `vm_cluster_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.
* `state` - (Optional) (Updatable) The target state for the Advanced Cluster File System. Could be set to `ACTIVE` or `INACTIVE`. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `description` - Description of the advanced cluster file system.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the advanced cluster file system.
* `is_mounted` - True if the file system is mounted on all VMs within VM Cluster.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `mount_point` - The mount point of file system.
* `name` - The user-friendly name for the Advanced cluster file system. The file system name is unique for a cluster.
* `state` - The current state of the advanced cluster file system. Valid states are CREATING, AVAILABLE, UPDATING, FAILED, DELETED. 
* `storage_in_gbs` - The total storage allocated for advanced cluster file system in GBs.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `time_created` - The date and time the advanced cluster file system was created.
* `time_updated` - The last date and time that the advanced cluster file system was updated.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts). This parameter and `secretId` are required for Customer Managed Keys.
* `vm_cluster_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM cluster.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Advanced Cluster File System
	* `update` - (Defaults to 20 minutes), when updating the Advanced Cluster File System
	* `delete` - (Defaults to 20 minutes), when destroying the Advanced Cluster File System


## Import

AdvancedClusterFileSystems can be imported using the `id`, e.g.

```
$ terraform import oci_database_advanced_cluster_file_system.test_advanced_cluster_file_system "id"
```

