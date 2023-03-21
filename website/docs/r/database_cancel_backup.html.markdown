---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_cancel_backup"
sidebar_current: "docs-oci-resource-database-cancel-backup"
description: |-
  Provides the Cancel Backup resource in the Oracle Cloud Infrastructure Database service.
---

# oci_database_cancel_backup
This resource provides the Cancel Backup resource in the Oracle Cloud Infrastructure Database service.

Cancels an automatic backup, which is in the provisioning workflow.


## Example Usage

```hcl
resource "oci_database_cancel_backup" "test_cancel_backup" {
	#Required
	backup_id = oci_database_backup.test_backup.id

    #Optional
	cancel_backup_trigger = 1
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the automatic backup.
* `cancel_backup_trigger` - (Optional) (Updatable) An optional integer, when updated triggers activation of cancel backup config.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain where the database backup is stored.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_edition` - The Oracle Database edition of the DB system from which the database backup was taken. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_size_in_gbs` - The size of the database in gigabytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
* `kms_key_id` - The OCID of the key container that is used as the master encryption key in database transparent data encryption (TDE) operations.
* `kms_key_version_id` - The OCID of the key container version that is used in database transparent data encryption (TDE) operations KMS Key can have multiple key versions. If none is specified, the current key version (latest) of the Key Id is used for the operation. 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `shape` - Shape of the backup's source database.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup was completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.
* `vault_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure [vault](https://docs.cloud.oracle.com/iaas/Content/KeyManagement/Concepts/keyoverview.htm#concepts).
* `version` - Version of the backup's source database

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
    * `create` - (Defaults to 20 minutes), when creating the Cancel Backup
	* `update` - (Defaults to 20 minutes), when updating the Cancel Backup
	* `delete` - (Defaults to 20 minutes), when destroying the Cancel Backup

