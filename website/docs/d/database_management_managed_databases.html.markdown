---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_databases"
sidebar_current: "docs-oci-datasource-database_management-managed_databases"
description: |-
  Provides the list of Managed Databases in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_databases
This data source provides the list of Managed Databases in Oracle Cloud Infrastructure Database Management service.

Gets the Managed Database for a specific ID or the list of Managed Databases in a specific compartment.
Managed Databases can also be filtered based on the name parameter. Only one of the parameters, ID or name
should be provided. If none of these parameters is provided, all the Managed Databases in the compartment are listed.


## Example Usage

```hcl
data "oci_database_management_managed_databases" "test_managed_databases" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	id = var.managed_database_id
	name = var.managed_database_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `id` - (Optional) The identifier of the resource.
* `name` - (Optional) A filter to return only resources that match the entire name.


## Attributes Reference

The following attributes are exported:

* `managed_database_collection` - The list of managed_database_collection.

### ManagedDatabase Reference

The following attributes are exported:

* `additional_details` - The additional details specific to a type of database defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_status` - The status of the Oracle Database. Indicates whether the status of the database is UP, DOWN, or UNKNOWN at the current time. 
* `database_sub_type` - The subtype of the Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or a Non-container Database.
* `database_type` - The type of Oracle Database installation.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `is_cluster` - Indicates whether the Oracle Database is part of a cluster.
* `managed_database_groups` - A list of Managed Database Groups that the Managed Database belongs to.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which the Managed Database Group resides.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database Group.
	* `name` - The name of the Managed Database Group.
* `name` - The name of the Managed Database.
* `parent_container_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent Container Database if Managed Database is a Pluggable Database. 
* `time_created` - The date and time the Managed Database was created.

