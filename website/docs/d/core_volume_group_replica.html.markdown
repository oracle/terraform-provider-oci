---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_group_replica"
sidebar_current: "docs-oci-datasource-core-volume_group_replica"
description: |-
  Provides details about a specific Volume Group Replica in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_group_replica
This data source provides details about a specific Volume Group Replica resource in Oracle Cloud Infrastructure Core service.

Gets information for the specified volume group replica.

## Example Usage

```hcl
data "oci_core_volume_group_replica" "test_volume_group_replica" {
	#Required
	volume_group_replica_id = oci_core_volume_group_replica.test_volume_group_replica.id
}
```

## Argument Reference

The following arguments are supported:

* `volume_group_replica_id` - (Required) The OCID of the volume replica group.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the volume group replica.
* `compartment_id` - The OCID of the compartment that contains the volume group replica.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The OCID for the volume group replica.
* `member_replicas` - Volume replicas within this volume group replica.
	* `volume_replica_id` - The volume replica ID.
* `size_in_gbs` - The aggregate size of the volume group replica in GBs.
* `state` - The current state of a volume group.
* `time_created` - The date and time the volume group replica was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_last_synced` - The date and time the volume group replica was last synced from the source volume group. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `volume_group_id` - The OCID of the source volume group.

