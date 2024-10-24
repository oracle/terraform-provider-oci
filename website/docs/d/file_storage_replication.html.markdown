---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_replication"
sidebar_current: "docs-oci-datasource-file_storage-replication"
description: |-
  Provides details about a specific Replication in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_replication
This data source provides details about a specific Replication resource in Oracle Cloud Infrastructure File Storage service.

Gets the specified replication's information.

## Example Usage

```hcl
data "oci_file_storage_replication" "test_replication" {
	#Required
	replication_id = oci_file_storage_replication.test_replication.id
}
```

## Argument Reference

The following arguments are supported:

* `replication_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the replication.


## Attributes Reference

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

