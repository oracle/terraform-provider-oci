---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup_policy_assignment"
sidebar_current: "docs-oci-resource-core-volume_backup_policy_assignment"
description: |-
  Provides the Volume Backup Policy Assignment resource in Oracle Cloud Infrastructure Core service
---

# oci_core_volume_backup_policy_assignment
This resource provides the Volume Backup Policy Assignment resource in Oracle Cloud Infrastructure Core service.

Assigns a volume backup policy to the specified volume or volume group. Note that a given volume or volume group can
only have one backup policy assigned to it. If this operation is used for a volume or volume group that already
has a different backup policy assigned, the prior backup policy will be silently unassigned.


## Example Usage

```hcl
resource "oci_core_volume_backup_policy_assignment" "test_volume_backup_policy_assignment" {
	#Required
	asset_id = oci_core_volume.test_volume.id
	policy_id = oci_core_volume_backup_policy.test_volume_backup_policy.id

	#Optional
	xrc_kms_key_id = oci_kms_key.test_key.id
}
```

## Argument Reference

The following arguments are supported:

* `asset_id` - (Required) The OCID of the volume or volume group to assign the policy to.
* `policy_id` - (Required) The OCID of the volume backup policy to assign to the volume.
* `xrc_kms_key_id` - (Optional) The OCID of the Vault service key which is the master encryption key for the block / boot volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `asset_id` - The OCID of the volume or volume group the policy has been assigned to.
* `id` - The OCID of the volume backup policy assignment.
* `policy_id` - The OCID of the volume backup policy that has been assigned to the volume or volume group. For a volume group, only the **user defined** policy is allowed to use. For more information, see [Policy-Based Backups](https://docs.oracle.com/en-us/iaas/Content/Block/Tasks/schedulingvolumebackups.htm).
* `time_created` - The date and time the volume backup policy was assigned to the volume. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `xrc_kms_key_id` - The OCID of the Vault service key which is the master encryption key for the block / boot volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Volume Backup Policy Assignment
	* `update` - (Defaults to 20 minutes), when updating the Volume Backup Policy Assignment
	* `delete` - (Defaults to 20 minutes), when destroying the Volume Backup Policy Assignment


## Import

VolumeBackupPolicyAssignments can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment "id"
```

