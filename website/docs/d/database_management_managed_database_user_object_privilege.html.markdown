---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_user_object_privilege"
sidebar_current: "docs-oci-datasource-database_management-managed_database_user_object_privilege"
description: |-
  Provides details about a specific Managed Database User Object Privilege in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_user_object_privilege
This data source provides details about a specific Managed Database User Object Privilege resource in Oracle Cloud Infrastructure Database Management service.

Gets the list of Object Privileges granted for the specified user.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_object_privilege" "test_managed_database_user_object_privilege" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_object_privilege_name
}
```

## Argument Reference

The following arguments are supported:

* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `name` - (Optional) A filter to return only resources that match the entire name.
* `user_name` - (Required) The name of the user whose details are to be viewed.


## Attributes Reference

The following attributes are exported:

* `items` - An array of User resources.
	* `common` - Indicates how the grant was made. Possible values: YES if the role was granted commonly (CONTAINER=ALL was used) NO if the role was granted locally (CONTAINER=ALL was not used) 
	* `grant_option` - Indicates whether the privilege was granted with the GRANT OPTION (YES) or not (NO)
	* `grantor` - The name of the user who performed the grant
	* `hierarchy` - Indicates whether the privilege was granted with the HIERARCHY OPTION (YES) or not (NO)
	* `inherited` - Indicates whether the role grant was inherited from another container (YES) or not (NO)
	* `name` - The name of the privilege on the object.
	* `object` - The name of the object. The object can be any object, including tables, packages, indexes, sequences, and so on.
	* `owner` - The owner of the object.
	* `schema_type` - The type of the object.

