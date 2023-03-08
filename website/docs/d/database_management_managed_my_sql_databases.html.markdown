---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_my_sql_databases"
sidebar_current: "docs-oci-datasource-database_management-managed_my_sql_databases"
description: |-
  Provides the list of Managed My Sql Databases in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_my_sql_databases
This data source provides the list of Managed My Sql Databases in Oracle Cloud Infrastructure Database Management service.

Gets the list of Managed MySQL Databases in a specific compartment.


## Example Usage

```hcl
data "oci_database_management_managed_my_sql_databases" "test_managed_my_sql_databases" {
	#Required
	compartment_id = var.compartment_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.


## Attributes Reference

The following attributes are exported:

* `managed_my_sql_database_collection` - The list of managed_my_sql_database_collection.

### ManagedMySqlDatabase Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment.
* `db_name` - MySQL Database Name
* `db_version` - MySQL Database Version
* `id` - The OCID of the Managed MySql Database.
* `name` - The name of the Managed MySQL Database.
* `time_created` - The date and time the Managed Database was created.

