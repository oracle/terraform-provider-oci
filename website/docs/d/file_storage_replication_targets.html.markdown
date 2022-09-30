---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_replication_targets"
sidebar_current: "docs-oci-datasource-file_storage-replication_targets"
description: |-
  Provides the list of Replication Targets in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_replication_targets
This data source provides the list of Replication Targets in Oracle Cloud Infrastructure File Storage service.

Lists the replication target resources in the specified compartment.


## Example Usage

```hcl
data "oci_file_storage_replication_targets" "test_replication_targets" {
	#Required
	availability_domain = var.replication_target_availability_domain
	compartment_id = var.compartment_id

	#Optional
	display_name = var.replication_target_display_name
	id = var.replication_target_id
	state = var.replication_target_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `id` - (Optional) Filter results by [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm). Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `replication_targets` - The list of replication_targets.

### ReplicationTarget Reference

The following attributes are exported:

* `availability_domain` - The availability domain the replication target is in. Must be in the same availability domain as the target file system. Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the replication.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `delta_progress` - Percentage progress of the current replication cycle. 
* `delta_status` - The current state of the snapshot during replication operations.
* `display_name` - A user-friendly name. This name is same as the replication display name for the associated resource. Example: `My Replication` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the replication target.
* `last_snapshot_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last snapshot snapshot which was completely applied to the target file system. Empty while the initial snapshot is being applied. 
* `lifecycle_details` - Additional information about the current `lifecycleState`.
* `recovery_point_time` - The snapshotTime of the most recent recoverable replication snapshot in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. Example: `2021-04-04T20:01:29.100Z` 
* `replication_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of replication.
* `source_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of source filesystem.
* `state` - The current state of this replication.
* `target_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of target filesystem.
* `time_created` - The date and time the replication target was created in target region. in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format. Example: `2021-01-04T20:01:29.100Z` 

