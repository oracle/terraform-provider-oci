---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_file_systems"
sidebar_current: "docs-oci-datasource-file_storage-file_systems"
description: |-
  Provides the list of File Systems in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_file_systems
This data source provides the list of File Systems in Oracle Cloud Infrastructure File Storage service.

Lists the file system resources in the specified compartment, or by the specified compartment and
file system snapshot policy.


## Example Usage

```hcl
data "oci_file_storage_file_systems" "test_file_systems" {
	#Required
	availability_domain = var.file_system_availability_domain
	compartment_id = var.compartment_id

	#Optional
	display_name = var.file_system_display_name
	filesystem_snapshot_policy_id = oci_file_storage_filesystem_snapshot_policy.test_filesystem_snapshot_policy.id
	id = var.file_system_id
	parent_file_system_id = oci_file_storage_file_system.test_file_system.id
	source_snapshot_id = oci_file_storage_snapshot.test_snapshot.id
	state = var.file_system_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `filesystem_snapshot_policy_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system snapshot policy that is associated with the file systems. 
* `id` - (Optional) Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type. 
* `parent_file_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system that contains the source snapshot of a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
* `source_snapshot_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm).
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `file_systems` - The list of file_systems.

### FileSystem Reference

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
* `metered_bytes` - The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system. For more information, see [File System Usage and Metering](https://docs.cloud.oracle.com/iaas/Content/File/Concepts/FSutilization.htm). 
* `replication_target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the replication target associated with the file system. Empty if the file system is not being used as target in a replication. 
* `source_details` - Source information for the file system. 
	* `parent_file_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the file system that contains the source snapshot of a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
	* `source_snapshot_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source snapshot used to create a cloned file system. See [Cloning a File System](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/cloningFS.htm). 
* `state` - The current state of the file system.
* `time_created` - The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

