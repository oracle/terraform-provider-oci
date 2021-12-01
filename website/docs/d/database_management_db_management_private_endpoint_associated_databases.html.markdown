---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_db_management_private_endpoint_associated_databases"
sidebar_current: "docs-oci-datasource-database_management-db_management_private_endpoint_associated_databases"
description: |-
  Provides the list of Db Management Private Endpoint Associated Databases in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_db_management_private_endpoint_associated_databases
This data source provides the list of Db Management Private Endpoint Associated Databases in Oracle Cloud Infrastructure Database Management service.

Gets the list of databases using a specific Database Management private endpoint.

## Example Usage

```hcl
data "oci_database_management_db_management_private_endpoint_associated_databases" "test_db_management_private_endpoint_associated_databases" {
	#Required
	compartment_id = var.compartment_id
	db_management_private_endpoint_id = oci_database_management_db_management_private_endpoint.test_db_management_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_management_private_endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint.


## Attributes Reference

The following attributes are exported:

* `associated_database_collection` - The list of associated_database_collection.

### DbManagementPrivateEndpointAssociatedDatabase Reference

The following attributes are exported:

* `items` - A list of databases using a Database Management private endpoint.
	* `compartment_id` - The compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
	* `name` - The name of the database.
	* `time_registered` - The time when Database Management was enabled for the database.

