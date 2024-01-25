---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_databases_user_system_privileges"
sidebar_current: "docs-oci-datasource-database_management-managed_databases_user_system_privileges"
description: |-
  Provides the list of Managed Databases User System Privileges in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_managed_databases_user_system_privileges
This data source provides the list of Managed Databases User System Privileges in Oracle Cloud Infrastructure Database Management service.

Gets the list of system privileges granted to a specific user.

## Example Usage

```hcl
data "oci_database_management_managed_databases_user_system_privileges" "test_managed_databases_user_system_privileges" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	user_name = oci_identity_user.test_user.name

	#Optional
	name = var.managed_databases_user_system_privilege_name
	opc_named_credential_id = var.managed_databases_user_system_privilege_opc_named_credential_id
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

* `system_privilege_collection` - The list of system_privilege_collection.

### ManagedDatabasesUserSystemPrivilege Reference

The following attributes are exported:

* `items` - An array of system privileges.
	* `admin_option` - Indicates whether the system privilege is granted with the ADMIN option (YES) or not (NO).
	* `common` - Indicates how the system privilege was granted. Possible values: YES if the system privilege is granted commonly (CONTAINER=ALL is used) NO if the system privilege is granted locally (CONTAINER=ALL is not used) 
	* `inherited` - Indicates whether the granted system privilege is inherited from another container (YES) or not (NO).
	* `name` - The name of a system privilege.

