---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_file_system"
sidebar_current: "docs-oci-resource-file_storage-file_system"
description: |-
  Provides the File System resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_file_system
This resource provides the File System resource in Oracle Cloud Infrastructure File Storage service.

Creates a new file system in the specified compartment and
availability domain. Instances can mount file systems in
another availability domain, but doing so might increase
latency when compared to mounting instances in the same
availability domain.

After you create a file system, you can associate it with a mount
target. Instances can then mount the file system by connecting to the
mount target's IP address. You can associate a file system with
more than one mount target at a time.

For information about access control and compartments, see
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

For information about Network Security Groups access control, see
[Network Security Groups](https://docs.cloud.oracle.com/iaas/Content/Network/Concepts/networksecuritygroups.htm).

For information about availability domains, see [Regions and
Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the
`ListAvailabilityDomains` operation in the Identity and Access
Management Service API.

All Oracle Cloud Infrastructure resources, including
file systems, get an Oracle-assigned, unique ID called an Oracle
Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
When you create a resource, you can find its OCID in the response.
You can also retrieve a resource's OCID by using a List API operation on that resource
type or by viewing the resource in the Console.


## Example Usage

```hcl
resource "oci_file_storage_file_system" "test_file_system" {
	#Required
	availability_domain = var.file_system_availability_domain
	compartment_id = var.compartment_id

	#Optional
	clone_attach_status = var.file_system_clone_attach_status
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.file_system_display_name
	filesystem_snapshot_policy_id = oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id
	freeform_tags = {"Department"= "Finance"}
	kms_key_id = oci_kms_key.test_key.id
	locks {
		#Required
		type = var.file_system_locks_type

		#Optional
		message = var.file_system_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.file_system_locks_time_created
	}
	source_snapshot_id = oci_file_storage_snapshot.test_snapshot.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1` 
* `clone_attach_status` - (Optional) Specifies whether the clone file system is attached to its parent file system. If the value is set to 'DETACH', then the file system will be created, which is deep copied from the snapshot specified by sourceSnapshotId, else will remain attached to its parent. 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the file system in.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system` 
* `filesystem_snapshot_policy_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated file system snapshot policy, which controls the frequency of snapshot creation and retention period of the taken snapshots.

	May be unset as a blank value. 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `kms_key_id` - (Optional) (Updatable) The OCID of KMS key used to encrypt the encryption keys associated with this file system. May be unset as a blank or deleted from the configuration to remove the KMS key.
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `source_snapshot_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `detach_clone_trigger` - (Optional) (Updatable) An optional property when incremented triggers Detach Clone. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the file system is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `clone_attach_status` - Specifies whether the file system is attached to its parent file system.
* `clone_count` - Specifies the total number of children of a file system.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the file system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system` 
* `filesystem_snapshot_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated file system snapshot policy, which controls the frequency of snapshot creation and retention period of the taken snapshots. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system.
* `is_clone_parent` - Specifies whether the file system has been cloned. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `is_hydrated` - Specifies whether the data has finished copying from the source to the clone. Hydration can take up to several hours to complete depending on the size of the source. The source and clone remain available during hydration, but there may be some performance impact. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm#hydration). 
* `is_targetable` - Specifies whether the file system can be used as a target file system for replication. The system sets this value to `true` if the file system is unexported, hasn't yet been specified as a target file system in any replication resource, and has no user snapshots. After the file system has been specified as a target in a replication, or if the file system contains user snapshots, the system sets this value to `false`. For more information, see [Using Replication](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/using-replication.htm). 
* `kms_key_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the KMS key used to encrypt the encryption keys associated with this file system. 
* `lifecycle_details` - Additional information about the current 'lifecycleState'.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `metered_bytes` - The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system. For more information, see [File System Usage and Metering](https://docs.cloud.oracle.com/iaas/Content/File/Concepts/FSutilization.htm). 
* `replication_target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the replication target associated with the file system. Empty if the file system is not being used as target in a replication. 
* `source_details` - Source information for the file system. 
	* `parent_file_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system that contains the source snapshot of a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
	* `source_snapshot_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `state` - The current state of the file system.
* `time_created` - The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the File System
	* `update` - (Defaults to 20 minutes), when updating the File System
	* `delete` - (Defaults to 20 minutes), when destroying the File System


## Import

FileSystems can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_file_system.test_file_system "id"
```

