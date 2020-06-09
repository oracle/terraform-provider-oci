---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_backups"
sidebar_current: "docs-oci-datasource-database-backups"
description: |-
  Provides the list of Backups in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_backups
This data source provides the list of Backups in Oracle Cloud Infrastructure Database service.

Gets a list of backups based on the databaseId or compartmentId specified. Either one of the query parameters must be provided.


## Example Usage

```hcl
data "oci_database_backups" "test_backups" {

	#Optional
	compartment_id = "${var.compartment_id}"
	database_id = "${oci_database_database.test_database.id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `database_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.


## Attributes Reference

The following attributes are exported:

* `backups` - The list of backups.

### Backup Reference

The following attributes are exported:

* `availability_domain` - The name of the availability domain where the database backup is stored.
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_edition` - The Oracle Database edition of the DB system from which the database backup was taken. 
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_size_in_gbs` - The size of the database in gigabytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. The name does not have to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the backup.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `shape` - Shape of the backup's source database.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup was completed.
* `time_started` - The date and time the backup started.
* `type` - The type of backup.
* `version` - Version of the backup's source database

