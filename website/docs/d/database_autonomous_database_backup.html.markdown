---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_autonomous_database_backup"
sidebar_current: "docs-oci-datasource-database-autonomous_database_backup"
description: |-
  Provides details about a specific Autonomous Database Backup in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_autonomous_database_backup
This data source provides details about a specific Autonomous Database Backup resource in Oracle Cloud Infrastructure Database service.

Gets information about the specified Autonomous Database backup.

## Example Usage

```hcl
data "oci_database_autonomous_database_backup" "test_autonomous_database_backup" {
	#Required
	autonomous_database_backup_id = "${oci_database_autonomous_database_backup.test_autonomous_database_backup.id}"
}
```

## Argument Reference

The following arguments are supported:

* `autonomous_database_backup_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.


## Attributes Reference

The following attributes are exported:

* `autonomous_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_size_in_tbs` - The size of the database in terabytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Autonomous Database backup.
* `is_automatic` - Indicates whether the backup is user-initiated or automatic.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.

