---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_db_homes"
sidebar_current: "docs-oci-datasource-database-db_homes"
description: |-
  Provides the list of Db Homes in Oracle Cloud Infrastructure Database service
---

# Data Source: oci_database_db_homes
This data source provides the list of Db Homes in Oracle Cloud Infrastructure Database service.

Gets a list of database homes in the specified DB system and compartment. A database home is a directory where Oracle database software is installed.


## Example Usage

```hcl
data "oci_database_db_homes" "test_db_homes" {
	#Required
	compartment_id = "${var.compartment_id}"
	db_system_id = "${oci_database_db_system.test_db_system.id}"

	#Optional
	display_name = "${var.db_home_display_name}"
	state = "${var.db_home_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `db_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. The match is not case sensitive.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state exactly.


## Attributes Reference

The following attributes are exported:

* `db_homes` - The list of db_homes.

### DbHome Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_version` - The Oracle Database version.
* `display_name` - The user-provided name for the database home. The name does not need to be unique.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database home.
* `last_patch_history_entry_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
* `state` - The current state of the database home.
* `time_created` - The date and time the database home was created.

