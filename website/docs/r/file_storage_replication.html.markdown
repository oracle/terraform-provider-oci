---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_replication"
sidebar_current: "docs-oci-resource-file_storage-replication"
description: |-
  Provides the Replication resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_replication
This resource provides the Replication resource in Oracle Cloud Infrastructure File Storage service.

Creates a new replication in the specified compartment.
Replications are the primary resource that governs the policy of cross-region replication between source
and target file systems. Replications are associated with a secondary resource called a [`ReplicationTarget`](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ReplicationTarget)
located in another availability domain.
The associated replication target resource is automatically created along with the replication resource.
The replication retrieves the delta of data between two snapshots of a source file system
and sends it to the associated `ReplicationTarget`, which retrieves the delta and applies it to the target
file system.
Only unexported file systems can be used as target file systems.
For more information, see [Using Replication](https://docs.cloud.oracle.com/iaas/Content/File/Tasks/FSreplication.htm).

For information about access control and compartments, see
[Overview of the IAM
Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

For information about availability domains, see [Regions and
Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the
`ListAvailabilityDomains` operation in the Identity and Access
Management Service API.

All Oracle Cloud Infrastructure Services resources, including
replications, get an Oracle-assigned, unique ID called an
Oracle Cloud Identifier ([OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm)).
When you create a resource, you can find its OCID in the response.
You can also retrieve a resource's OCID by using a List API operation on that resource
type, or by viewing the resource in the Console.


## Example Usage

```hcl
resource "oci_file_storage_replication" "test_replication" {
	#Required
	compartment_id = var.compartment_id
	source_id = oci_file_storage_file_system.test_source.id
	target_id = oci_file_storage_file_system.test_target.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.replication_display_name
	freeform_tags = {"Department"= "Finance"}
	replication_interval = var.replication_replication_interval
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the replication.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information. An associated replication target will also created with the same `displayName`. Example: `My replication` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `replication_interval` - (Optional) (Updatable) Duration in minutes between replication snapshots.
* `source_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source file system. 
* `target_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target file system. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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
* `recovery_point_time` - The [`snapshotTime`](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/Snapshot/snapshotTime) of the most recent recoverable replication snapshot in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. Example: `2021-04-04T20:01:29.100Z` 
* `replication_interval` - Duration in minutes between replication snapshots.
* `replication_target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [`ReplicationTarget`](https://docs.cloud.oracle.com/iaas/api/#/en/iaas/latest/ReplicationTarget). 
* `source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the source file system. 
* `state` - The current state of this replication. This resource can be in a `FAILED` state if replication target is deleted instead of the replication resource. 
* `target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the target file system. 
* `time_created` - The date and time the replication was created in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2021-01-04T20:01:29.100Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Replication
	* `update` - (Defaults to 20 minutes), when updating the Replication
	* `delete` - (Defaults to 20 minutes), when destroying the Replication


## Import

Replications can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_replication.test_replication "id"
```

