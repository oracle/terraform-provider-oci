---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_volume_backup_policy_assignments"
sidebar_current: "docs-oci-datasource-core-volume_backup_policy_assignments"
description: |-
  Provides the list of Volume Backup Policy Assignments in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_volume_backup_policy_assignments
This data source provides the list of Volume Backup Policy Assignments in Oracle Cloud Infrastructure Core service.

Gets the volume backup policy assignment for the specified volume. The
`assetId` query parameter is required, and the returned list will contain at most
one item, since volume can only have one volume backup policy assigned at a time.


## Example Usage

```hcl
data "oci_core_volume_backup_policy_assignments" "test_volume_backup_policy_assignments" {
	#Required
	asset_id = oci_core_volume.test_volume.id
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

* `asset_id` - The OCID of the volume the policy has been assigned to.
* `id` - The OCID of the volume backup policy assignment.
* `policy_id` - The OCID of the volume backup policy that has been assigned to the volume. 
* `time_created` - The date and time the volume backup policy was assigned to the volume. The format is defined by [RFC3339](https://tools.ietf.org/html/rfc3339). 
* `xrc_kms_key_id` - The OCID of the Vault service key which is the master encryption key for the block / boot volume cross region backups, which will be used in the destination region to encrypt the backup's encryption keys. For more information about the Vault service and encryption keys, see [Overview of Vault service](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm) and [Using Keys](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Tasks/usingkeys.htm). 

