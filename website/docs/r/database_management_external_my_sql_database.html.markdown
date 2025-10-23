---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_my_sql_database"
sidebar_current: "docs-oci-resource-database_management-external_my_sql_database"
description: |-
  Provides the External My Sql Database resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_my_sql_database
This resource provides the External My Sql Database resource in Oracle Cloud Infrastructure Database Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-management/latest/ExternalMySqlDatabase

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemanagement

Creates an external MySQL database.


## Example Usage

```hcl
resource "oci_database_management_external_my_sql_database" "test_external_my_sql_database" {
	#Required
	compartment_id = var.compartment_id
	db_name = var.external_my_sql_database_db_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) OCID of compartment for the External MySQL Database.
* `db_name` - (Required) (Updatable) Name of the External MySQL Database.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of compartment for the External MySQL Database.
* `db_name` - Display Name of the External MySQL Database.
* `external_database_id` - OCID of External MySQL Database.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External My Sql Database
	* `update` - (Defaults to 20 minutes), when updating the External My Sql Database
	* `delete` - (Defaults to 20 minutes), when destroying the External My Sql Database


## Import

ExternalMySqlDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_my_sql_database.test_external_my_sql_database "id"
```

