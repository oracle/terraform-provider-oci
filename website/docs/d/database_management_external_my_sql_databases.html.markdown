---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_my_sql_databases"
sidebar_current: "docs-oci-datasource-database_management-external_my_sql_databases"
description: |-
  Provides the list of External My Sql Databases in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_my_sql_databases
This data source provides the list of External My Sql Databases in Oracle Cloud Infrastructure Database Management service.

Gets the list of External MySQL Databases. 


## Example Usage

```hcl
data "oci_database_management_external_my_sql_databases" "test_external_my_sql_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.external_my_sql_database_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) The parameter to filter by MySQL Database System type.


## Attributes Reference

The following attributes are exported:

* `external_my_sql_database_collection` - The list of external_my_sql_database_collection.

### ExternalMySqlDatabase Reference

The following attributes are exported:

* `compartment_id` - OCID of compartment for the External MySQL Database.
* `db_name` - Display Name of the External MySQL Database.
* `external_database_id` - OCID of External MySQL Database.

