---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_replications"
sidebar_current: "docs-oci-datasource-file_storage-replications"
description: |-
  Provides the list of Replications in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_replications
This data source provides the list of Replications in Oracle Cloud Infrastructure File Storage service.

Lists the replication resources in the specified compartment.


## Example Usage

```hcl
data "oci_file_storage_replications" "test_replications" {
	#Required
	availability_domain = var.replication_availability_domain
	compartment_id = var.compartment_id

	#Optional
	display_name = var.replication_display_name
	file_system_id = oci_file_storage_file_system.test_file_system.id
	id = var.replication_id
	state = var.replication_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `file_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source file system.
* `id` - (Optional) Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `replications` - The list of replications.

### Replication Reference

The following attributes are exported:

* `availability_domain` - The availability domain the replication is in. The replication must be in the same availability domain as the source file system. Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the replication.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `delta_progress` - Percentage progress of the current replication cycle. 
* `delta_status` - The current state of the snapshot during replication operations.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My replication` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the replication.
* `last_snapshot_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last snapshot that has been replicated completely. Empty if the copy of the initial snapshot is not complete. 
* `lifecycle_details` - Additional information about the current 'lifecycleState'.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The ID of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `recovery_point_time` - The [`snapshotTime`](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Snapshot/snapshotTime) of the most recent recoverable replication snapshot in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. Example: `2021-04-04T20:01:29.100Z` 
* `replication_interval` - Duration in minutes between replication snapshots.
* `replication_target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [`ReplicationTarget`](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ReplicationTarget). 
* `source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source file system. 
* `state` - The current state of this replication. This resource can be in a `FAILED` state if replication target is deleted instead of the replication resource. 
* `target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target file system. 
* `time_created` - The date and time the replication was created in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2021-01-04T20:01:29.100Z` 

