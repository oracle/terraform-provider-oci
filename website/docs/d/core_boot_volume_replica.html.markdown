---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume_replica"
sidebar_current: "docs-oci-datasource-core-boot_volume_replica"
description: |-
  Provides details about a specific Boot Volume Replica in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_boot_volume_replica
This data source provides details about a specific Boot Volume Replica resource in Oracle Cloud Infrastructure Core service.

Gets information for the specified boot volume replica.

## Example Usage

```hcl
data "oci_core_boot_volume_replica" "test_boot_volume_replica" {
	#Required
	boot_volume_replica_id = oci_core_boot_volume_replica.test_boot_volume_replica.id
}
```

## Argument Reference

The following arguments are supported:

* `boot_volume_replica_id` - (Required) The OCID of the boot volume replica.


## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain of the boot volume replica.  Example: `Uocm:PHX-AD-1` 
* `boot_volume_id` - The OCID of the source boot volume.
* `compartment_id` - The OCID of the compartment that contains the boot volume replica.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The boot volume replica's Oracle ID (OCID).
* `image_id` - The image OCID used to create the boot volume the replica is replicated from. 
* `kms_key_id` - The OCID of the Vault service key to assign as the master encryption key for the boot volume replica, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm).
* `image_id` - The image OCID used to create the boot volume the replica is replicated from.
* `size_in_gbs` - The size of the source boot volume, in GBs. 
* `state` - The current state of a boot volume replica.
* `time_created` - The date and time the boot volume replica was created. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `time_last_synced` - The date and time the boot volume replica was last synced from the source boot volume. Format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 

