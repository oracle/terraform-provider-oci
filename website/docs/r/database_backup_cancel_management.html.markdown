---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_backup_cancel_management"
sidebar_current: "docs-oci-resource-database-backup_cancel_management"
description: |-
  Provides the Backup Cancel Management resource in Oracle Cloud Infrastructure Database service
---

# oci_database_backup_cancel_management
This resource provides the Backup Cancel Management resource in Oracle Cloud Infrastructure Database service.

Cancel automatic full/incremental create backup workrequests specified by the backup Id. This cannot be used on manual backups.

## Example Usage

```hcl
resource "oci_database_backup_cancel_management" "test_backup_cancel_management" {
	#Required
	backup_id = oci_database_backup.test_backup.id
    cancel_backup_trigger = 1
}
```

## Argument Reference

The following arguments are supported:

* `backup_id` - (Required) The backup [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `cancel_backup_trigger` - (Optional) When changed to a different integer, re-triggers cancel backup on the backup specified by the backup_id


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Backup Cancel Management
	* `update` - (Defaults to 20 minutes), when updating the Backup Cancel Management
	* `delete` - (Defaults to 20 minutes), when destroying the Backup Cancel Management


## Import

Import is not supported for this resource.

