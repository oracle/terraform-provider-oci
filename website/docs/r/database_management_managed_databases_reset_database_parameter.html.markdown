---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_databases_reset_database_parameter"
sidebar_current: "docs-oci-resource-database_management-managed_databases_reset_database_parameter"
description: |-
  Provides the Managed Databases Reset Database Parameter resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_managed_databases_reset_database_parameter
This resource provides the Managed Databases Reset Database Parameter resource in Oracle Cloud Infrastructure Database Management service.

Resets database parameter values to their default or startup values.


## Example Usage

```hcl
resource "oci_database_management_managed_databases_reset_database_parameter" "test_managed_databases_reset_database_parameter" {
	#Required
	credentials {

		#Optional
		password = var.managed_databases_reset_database_parameter_credentials_password
		role = var.managed_databases_reset_database_parameter_credentials_role
		secret_id = oci_vault_secret.test_secret.id
		user_name = oci_identity_user.test_user.name
	}
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	parameters = var.managed_databases_reset_database_parameter_parameters
	scope = var.managed_databases_reset_database_parameter_scope
}
```

## Argument Reference

The following arguments are supported:

* `credentials` - (Required) The database credentials used to perform management activity.
	* `password` - (Optional) The password for the database user name. 
	* `role` - (Optional) The role of the database user. Indicates whether the database user is a normal user or sysdba.
	* `secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	* `user_name` - (Optional) The database user name used to perform management activity. 
* `managed_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Managed Database.
* `parameters` - (Required) A list of database parameter names.
* `scope` - (Required) The clause used to specify when the parameter change takes effect.

	Use `MEMORY` to make the change in memory and ensure that it takes effect immediately. Use `SPFILE` to make the change in the server parameter file. The change takes effect when the database is next shut down and started up again. Use `BOTH` to make the change in memory and in the server parameter file. The change takes effect immediately and persists after the database is shut down and started up again. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `status` - A map with the parameter name as key and its update status as value. 
	* `error_code` - An error code that defines the failure or `null` if the parameter was updated successfully. 
	* `error_message` - The error message indicating the reason for failure or `null` if the parameter was updated successfully. 
	* `status` - The status of the parameter update.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Databases Reset Database Parameter
	* `update` - (Defaults to 20 minutes), when updating the Managed Databases Reset Database Parameter
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Databases Reset Database Parameter


## Import

Import is not supported for this resource.

