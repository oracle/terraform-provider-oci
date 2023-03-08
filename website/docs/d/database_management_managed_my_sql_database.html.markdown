---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_database"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_database"
description: |-
  Provides details about a specific Managed My Sql Database in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_database
This data source provides details about a specific Managed My Sql Database resource in Oracle Cloud Infrastructure Database Management service.

Retrieves General Information for given MySQL Instance.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_database" "test_managed_my_sql_database" {
	#Required
	managed_my_sql_database_id = oci_database_management_managed_my_sql_database.test_managed_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `managed_my_sql_database_id` - (Required) The OCID of ManagedMySqlDatabase.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_name` - MySQL Database Name
* `db_version` - MySQL Database Version
* `id` - The OCID of the Managed MySql Database.
* `name` - The name of the Managed MySQL Database.
* `time_created` - The date and time the Managed Database was created.

