---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup_policy_assignments"
sidebar_current: "docs-oci-datasource-core-volume_backup_policy_assignments"
description: |-
  Provides the list of Volume Backup Policy Assignments in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_backup_policy_assignments
This data source provides the list of Volume Backup Policy Assignments in Oracle Cloud Infrastructure Core service.

Gets the volume backup policy assignment for the specified asset. Note that the
assetId query parameter is required, and that the returned list will contain at most
one item (since any given asset can only have one policy assigned to it).


## Example Usage

```hcl
data "oci_core_volume_backup_policy_assignments" "test_volume_backup_policy_assignments" {
	#Required
	asset_id = "${oci_core_volume.test_volume.id}"
}
```

## Argument Reference

The following arguments are supported:

* `asset_id` - (Required) The OCID of an asset (e.g. a volume).


## Attributes Reference

The following attributes are exported:

* `volume_backup_policy_assignments` - The list of volume_backup_policy_assignments.

### VolumeBackupPolicyAssignment Reference

The following attributes are exported:

* `asset_id` - The OCID of the asset (e.g. a volume) to which the policy has been assigned.
* `id` - The OCID of the volume backup policy assignment.
* `policy_id` - The OCID of the volume backup policy that has been assigned to an asset.
* `time_created` - The date and time the volume backup policy assignment was created. Format defined by RFC3339. 

