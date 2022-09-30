---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_replication_target"
sidebar_current: "docs-oci-datasource-file_storage-replication_target"
description: |-
  Provides details about a specific Replication Target in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_replication_target
This data source provides details about a specific Replication Target resource in Oracle Cloud Infrastructure File Storage service.

Gets the specified replication target's information.

## Example Usage

```hcl
data "oci_file_storage_replication_target" "test_replication_target" {
	#Required
	replication_target_id = oci_file_storage_replication_target.test_replication_target.id
}
```

## Argument Reference

The following arguments are supported:

* `replication_target_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the replication target.


## Attributes Reference

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

