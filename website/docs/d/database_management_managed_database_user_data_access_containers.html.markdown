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

Gets the list of containers for a specific user. This is only applicable if ALL_CONTAINERS !='Y'.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_data_access_containers" "test_managed_database_user_data_access_containers" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_data_access_container_name
	opc_named_credential_id = var.managed_database_user_data_access_container_opc_named_credential_id
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `opc_named_credential_id` - (Optional) The OCID of the Named Credential.
* `user_name` - (Required) The name of the user whose details are to be viewed.


## Attributes Reference

The following attributes are exported:

* `data_access_container_collection` - The list of data_access_container_collection.

### ManagedDatabaseUserDataAccessContainer Reference

The following attributes are exported:

* `items` - An array of container resources.
	* `name` - The name of the container included in the attribute.

