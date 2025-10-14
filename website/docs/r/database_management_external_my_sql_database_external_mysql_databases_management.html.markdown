---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_my_sql_database_external_mysql_databases_management"
sidebar_current: "docs-oci-resource-database_management-external_my_sql_database_external_mysql_databases_management"
description: |-
  Provides the External My Sql Database External Mysql Databases Management resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_my_sql_database_external_mysql_databases_management
This resource provides the External My Sql Database External Mysql Databases Management resource in Oracle Cloud Infrastructure Database Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-management/latest/ExternalMySqlDatabaseExternalMysqlDatabasesManagement

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemanagement
Enables Database Management for an external MySQL Database.


## Example Usage

```hcl
resource "oci_database_management_external_my_sql_database_external_mysql_databases_management" "test_external_my_sql_database_external_mysql_databases_management" {
	#Required
	external_my_sql_database_id = oci_database_management_external_my_sql_database.test_external_my_sql_database.id
	enable_external_mysql_database = var.enable_external_mysql_database

	#Optional
	connector_id = oci_database_management_connector.test_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `connector_id` - (Optional) OCID of External MySQL Database connector.
* `external_my_sql_database_id` - (Required) The OCID of the External MySQL Database.
* `enable_external_mysql_database` - (Required) (Updatable) A required field when set to `true` calls enable action and when set to `false` calls disable action.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External My Sql Database External Mysql Databases Management
	* `update` - (Defaults to 20 minutes), when updating the External My Sql Database External Mysql Databases Management
	* `delete` - (Defaults to 20 minutes), when destroying the External My Sql Database External Mysql Databases Management
