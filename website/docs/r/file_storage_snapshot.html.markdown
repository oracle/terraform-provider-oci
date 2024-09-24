---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_snapshot"
sidebar_current: "docs-oci-resource-file_storage-snapshot"
description: |-
  Provides the Snapshot resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_snapshot
This resource provides the Snapshot resource in Oracle Cloud Infrastructure File Storage service.

Creates a new snapshot of the specified file system. You
can access the snapshot at `.snapshot/<name>`.


## Example Usage

```hcl
resource "oci_file_storage_snapshot" "test_snapshot" {
	#Required
	file_system_id = oci_file_storage_file_system.test_file_system.id
	name = var.snapshot_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	expiration_time = var.snapshot_expiration_time
	freeform_tags = {"Department"= "Finance"}
	locks {
		#Required
		type = var.snapshot_locks_type

		#Optional
		message = var.snapshot_locks_message
		related_resource_id = oci_cloud_guard_resource.test_resource.id
		time_created = var.snapshot_locks_time_created
	}
}
```

## Argument Reference

The following arguments are supported:

* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `expiration_time` - (Optional) (Updatable) The time when this snapshot will be deleted.
* `file_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system to take a snapshot of.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `name` - (Required) Name of the snapshot. This value is immutable. It must also be unique with respect to all other non-DELETED snapshots on the associated file system.

	Avoid entering confidential information.

	Example: `Sunday` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `expiration_time` - The time when this snapshot will be deleted.
* `file_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system from which the snapshot was created. 
* `filesystem_snapshot_policy_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system snapshot policy that created this snapshot. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot.
* `is_clone_source` - Specifies whether the snapshot has been cloned. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `lifecycle_details` - Additional information about the current `lifecycleState`.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `name` - Name of the snapshot. This value is immutable.

	Avoid entering confidential information.

	Example: `Sunday` 
* `provenance_id` - An [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) identifying the parent from which this snapshot was cloned. If this snapshot was not cloned, then the `provenanceId` is the same as the snapshot `id` value. If this snapshot was cloned, then the `provenanceId` value is the parent's `provenanceId`. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `snapshot_time` - The date and time the snapshot was taken, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. This value might be the same or different from `timeCreated` depending on the following factors:
	* If the snapshot is created in the original file system directory.
	* If the snapshot is cloned from a file system.
	* If the snapshot is replicated from a file system.

	Example: `2020-08-25T21:10:29.600Z` 
* `snapshot_type` - Specifies the generation type of the snapshot. 
* `state` - The current state of the snapshot.
* `time_created` - The date and time the snapshot was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Snapshot
	* `update` - (Defaults to 20 minutes), when updating the Snapshot
	* `delete` - (Defaults to 20 minutes), when destroying the Snapshot


## Import

Snapshots can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_snapshot.test_snapshot "id"
```

