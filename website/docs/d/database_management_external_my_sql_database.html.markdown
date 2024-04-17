---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_my_sql_database"
sidebar_current: "docs-oci-datasource-database_management-external_my_sql_database"
description: |-
  Provides details about a specific External My Sql Database in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_my_sql_database
This data source provides details about a specific External My Sql Database resource in Oracle Cloud Infrastructure Database Management service.

Retrieves the external MySQL database information.


## Example Usage

```hcl
data "oci_database_management_external_my_sql_database" "test_external_my_sql_database" {
	#Required
	external_my_sql_database_id = oci_database_management_external_my_sql_database.test_external_my_sql_database.id
}
```

## Argument Reference

The following arguments are supported:

* `external_my_sql_database_id` - (Required) The OCID of the External MySQL Database.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of compartment for the External MySQL Database.
* `db_name` - Display Name of the External MySQL Database.
* `external_database_id` - OCID of External MySQL Database.

