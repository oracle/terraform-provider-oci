---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_boot_volume_replicas"
sidebar_current: "docs-oci-datasource-core-boot_volume_replicas"
description: |-
  Provides the list of Boot Volume Replicas in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_boot_volume_replicas
This data source provides the list of Boot Volume Replicas in Oracle Cloud Infrastructure Core service.

Lists the boot volume replicas in the specified compartment and availability domain.


## Example Usage

```hcl
data "oci_core_boot_volume_replicas" "test_boot_volume_replicas" {

	#Optional
	availability_domain = var.boot_volume_replica_availability_domain
	compartment_id = var.compartment_id
	display_name = var.boot_volume_replica_display_name
	state = var.boot_volume_replica_state
	volume_group_replica_id = oci_core_volume_group_replica.test_volume_group_replica.id
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `state` - (Optional) A filter to only return resources that match the given lifecycle state. The state value is case-insensitive. 
* `volume_group_replica_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the volume group replica.
* `state` - (Optional) A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `boot_volume_replicas` - The list of boot_volume_replicas.

### BootVolumeReplica Reference

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

