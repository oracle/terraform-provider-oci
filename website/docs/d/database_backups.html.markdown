---
layout: "oci"
page_title: "OCI: oci_database_backups"
sidebar_current: "docs-oci-datasource-database-backups"
description: |-
  Provides a list of Backups
---

# Data Source: oci_database_backups
The Backups data source allows access to the list of OCI backups

Gets a list of backups based on the databaseId or compartmentId specified. Either one of the query parameters must be provided.


## Example Usage

```hcl
data "oci_database_backups" "test_backups" {

	#Optional
	database_id = "${oci_database_database.test_database.id}"
	// or
	// compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The compartment OCID.
* `database_id` - (Optional) The OCID of the database.


## Attributes Reference

The following attributes are exported:

* `backups` - The list of backups.

### Backup Reference

The following attributes are exported:

* `availability_domain` - The name of the Availability Domain that the backup is located in.
* `compartment_id` - The OCID of the compartment.
* `database_edition` - The Oracle Database Edition of the DbSystem on which the backup was taken. 
* `database_id` - The OCID of the database.
* `db_data_size_in_mbs` - Size of the database in mega-bytes at the time the backup was taken. 
* `display_name` - The user-friendly name for the backup. It does not have to be unique.
* `id` - The OCID of the backup.
* `lifecycle_details` - Additional information about the current lifecycleState.
* `state` - The current state of the backup.
* `time_ended` - The date and time the backup was completed.
* `time_started` - The date and time the backup starts.
* `type` - The type of backup.

