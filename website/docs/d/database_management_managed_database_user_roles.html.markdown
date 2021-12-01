---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_user_roles"
sidebar_current: "docs-oci-datasource-database_management-managed_database_user_roles"
description: |-
  Provides the list of Managed Database User Roles in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_user_roles
This data source provides the list of Managed Database User Roles in Oracle Cloud Infrastructure Database Management service.

Gets the list of roles granted for the specified user.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_roles" "test_managed_database_user_roles" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_role_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `user_name` - (Required) The name of the user whose details are to be viewed.


## Attributes Reference

The following attributes are exported:

* `role_collection` - The list of role_collection.

### ManagedDatabaseUserRole Reference

The following attributes are exported:

* `items` - An array of User resources.
	* `admin_option` - Indicates whether the grant was with the ADMIN OPTION (YES) or not (NO)
	* `common` - Indicates how the grant was made. Possible values: YES if the role was granted commonly (CONTAINER=ALL was used) NO if the role was granted locally (CONTAINER=ALL was not used) 
	* `default_role` - Indicates whether the role is designated as a DEFAULT ROLE for the user (YES) or not (NO)
	* `delegate_option` - Indicates whether the grant was with the DELEGATE OPTION (YES) or not (NO)
	* `inherited` - Indicates whether the role grant was inherited from another container (YES) or not (NO)
	* `name` - The name of a granted role

