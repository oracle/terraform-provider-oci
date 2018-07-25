---
layout: "oci"
page_title: "OCI: oci_core_volume_backup_policies"
sidebar_current: "docs-oci-datasource-core-volume_backup_policies"
description: |-
  Provides a list of VolumeBackupPolicies
---

# Data Source: oci_core_volume_backup_policies
The VolumeBackupPolicies data source allows access to the list of OCI volume_backup_policies

Lists all volume backup policies available to the caller.

## Example Usage

```hcl
data "oci_core_volume_backup_policies" "test_volume_backup_policies" {
}
```

## Attributes Reference

The following attributes are exported:

* `volume_backup_policies` - The list of volume_backup_policies.

### VolumeBackupPolicy Reference

The following attributes are exported:

* `display_name` - A user-friendly name for the volume backup policy. Does not have to be unique and it's changeable. Avoid entering confidential information. 
* `id` - The OCID of the volume backup policy.
* `schedules` - The collection of schedules that this policy will apply.
	* `backup_type` - The type of backup to create.
	* `offset_seconds` - The number of seconds (positive or negative) that the backup time should be shifted from the default interval boundaries specified by the period.
	* `period` - How often the backup should occur.
	* `retention_seconds` - How long, in seconds, backups created by this schedule should be kept until being automatically deleted.
* `time_created` - The date and time the volume backup policy was created. Format defined by RFC3339. 

