---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_block_volume_replica"
sidebar_current: "docs-oci-datasource-core-block_volume_replica"
description: |-
  Provides details about a specific Block Volume Replica in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_block_volume_replica
This data source provides details about a specific Block Volume Replica resource in Oracle Cloud Infrastructure Core service.

Gets information for the specified block volume replica.

## Example Usage

```hcl
data "oci_core_block_volume_replica" "test_block_volume_replica" {
	#Required
	block_volume_replica_id = oci_core_block_volume_replica.test_block_volume_replica.id
}
```

## Argument Reference

The following arguments are supported:

* `block_volume_replica_id` - (Required) The OCID of the block volume replica.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the block volume replica.  Example: `Uocm:PHX-AD-1` 
* `block_volume_id` - The OCID of the source block volume.
* `compartment_id` - The OCID of the compartment that contains the block volume replica.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The block volume replica's Oracle ID (OCID).
* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the block volume replica, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 
* `size_in_gbs` - The size of the source block volume, in GBs. 
* `state` - The current state of a block volume replica.
* `time_created` - The date and time the block volume replica was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_last_synced` - The date and time the block volume replica was last synced from the source block volume. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

