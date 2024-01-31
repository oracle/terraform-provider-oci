---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_database_user_object_privileges"
sidebar_current: "docs-oci-datasource-database_management-managed_database_user_object_privileges"
description: |-
  Provides the list of Managed Database User Object Privileges in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_database_user_object_privileges
This data source provides the list of Managed Database User Object Privileges in Oracle Cloud Infrastructure Database Management service.

Gets the list of object privileges granted to a specific user.

## Example Usage

```hcl
data "oci_database_management_managed_database_user_object_privileges" "test_managed_database_user_object_privileges" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_database_user_object_privilege_name
	opc_named_credential_id = var.managed_database_user_object_privilege_opc_named_credential_id
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

* `object_privilege_collection` - The list of object_privilege_collection.

### ManagedDatabaseUserObjectPrivilege Reference

The following attributes are exported:

* `items` - An array of object privileges.
	* `common` - Indicates how the object privilege was granted. Possible values: YES if the role is granted commonly (CONTAINER=ALL is used) NO if the role is granted locally (CONTAINER=ALL is not used) 
	* `grant_option` - Indicates whether the privilege is granted with the GRANT OPTION (YES) or not (NO).
	* `grantor` - The name of the user who granted the object privilege.
	* `hierarchy` - Indicates whether the privilege is granted with the HIERARCHY OPTION (YES) or not (NO).
	* `inherited` - Indicates whether the granted privilege is inherited from another container (YES) or not (NO).
	* `name` - The name of the privilege on the object.
	* `object` - The name of the object. The object can be any object, including tables, packages, indexes, sequences, and so on.
	* `owner` - The owner of the object.
	* `schema_type` - The type of object.

