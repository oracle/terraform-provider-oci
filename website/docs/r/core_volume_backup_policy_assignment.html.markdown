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

Assigns a policy to the specified asset, such as a volume. Note that a given asset can
only have one policy assigned to it; if this method is called for an asset that previously
has a different policy assigned, the prior assignment will be silently deleted.


## Example Usage

```hcl
resource "oci_core_volume_backup_policy_assignment" "test_volume_backup_policy_assignment" {
	#Required
	asset_id = "${oci_core_volume.test_volume.id}"
	policy_id = "${oci_core_volume_backup_policy.test_volume_backup_policy.id}"
}
```

## Argument Reference

The following arguments are supported:

* `asset_id` - (Required) The OCID of the asset (e.g. a volume) to which to assign the policy.
* `policy_id` - (Required) The OCID of the volume backup policy to assign to an asset.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `asset_id` - The OCID of the asset (e.g. a volume) to which the policy has been assigned.
* `id` - The OCID of the volume backup policy assignment.
* `policy_id` - The OCID of the volume backup policy that has been assigned to an asset.
* `time_created` - The date and time the volume backup policy assignment was created. Format defined by RFC3339. 

## Import

VolumeBackupPolicyAssignments can be imported using the `id`, e.g.

```
$ terraform import oci_core_volume_backup_policy_assignment.test_volume_backup_policy_assignment "id"
```

