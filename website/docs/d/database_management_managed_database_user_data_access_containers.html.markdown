---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_user_data_access_containers"
sidebar_current: "docs-oci-datasource-database_management-managed_database_user_data_access_containers"
description: |-
  Provides the list of Managed Database User Data Access Containers in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_user_data_access_containers
This data source provides the list of Managed Database User Data Access Containers in Oracle Cloud Infrastructure Database Management service.

Gets the list of Containers if it does not apply to all containers for the specified user.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_data_access_containers" "test_managed_database_user_data_access_containers" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_data_access_container_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `user_name` - (Required) The name of the user whose details are to be viewed.


## Attributes Reference

The following attributes are exported:

* `data_access_container_collection` - The list of data_access_container_collection.

### ManagedDatabaseUserDataAccessContainer Reference

The following attributes are exported:

* `items` - An array of Container resources.
	* `name` - The name of a container included in this attribute if it does not apply to all containers.

