---
layout: "oci"
page_title: "OCI: oci_database_db_home"
sidebar_current: "docs-oci-datasource-database-db_home"
description: |-
  Provides details about a specific DbHome
---

# Data Source: oci_database_db_home
The `oci_database_db_home` data source provides details about a specific DbHome

Gets information about the specified database home.

## Example Usage

```hcl
data "oci_database_db_home" "test_db_home" {
	#Required
	db_home_id = "${oci_database_db_system.test_db_system.id}"
}
```

## Argument Reference

The following arguments are supported:

* `db_home_id` - (Required) The database home [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm).


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_system_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the DB system.
* `db_version` - The Oracle Database version.
* `display_name` - The user-provided name for the database home. The name does not need to be unique.
* `id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the database home.
* `last_patch_history_entry_id` - The [OCID](https://docs.us-phoenix-1.oraclecloud.com/Content/General/Concepts/identifiers.htm) of the last patch history. This value is updated as soon as a patch operation is started.
* `state` - The current state of the database home.
* `time_created` - The date and time the database home was created.

