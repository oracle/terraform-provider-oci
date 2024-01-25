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
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	parameters = var.managed_databases_reset_database_parameter_parameters
	scope = var.managed_databases_reset_database_parameter_scope

	#Optional
	credentials {

		#Optional
		password = var.managed_databases_reset_database_parameter_credentials_password
		role = var.managed_databases_reset_database_parameter_credentials_role
		secret_id = oci_vault_secret.test_secret.id
		user_name = oci_identity_user.test_user.name
	}
	database_credential {
		#Required
		credential_type = var.managed_databases_reset_database_parameter_database_credential_credential_type

		#Optional
		named_credential_id = oci_database_management_named_credential.test_named_credential.id
		password = var.managed_databases_reset_database_parameter_database_credential_password
		password_secret_id = oci_vault_secret.test_secret.id
		role = var.managed_databases_reset_database_parameter_database_credential_role
		username = var.managed_databases_reset_database_parameter_database_credential_username
	}
}
```

## Argument Reference

The following arguments are supported:

* `credentials` - (Optional) The database credentials used to perform management activity. Provide one of the following attribute set. (userName, password, role) OR (userName, secretId, role) OR (namedCredentialId) 
	* `password` - (Optional) The password for the database user name. 
	* `role` - (Optional) The role of the database user. Indicates whether the database user is a normal user or sysdba.
	* `secret_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	* `user_name` - (Optional) The database user name used to perform management activity. 
* `database_credential` - (Optional) The credential to connect to the database to perform tablespace administration tasks.
	* `credential_type` - (Required) The type of the credential for tablespace administration tasks.
	* `named_credential_id` - (Required when credential_type=NAMED_CREDENTIAL) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential where the database password metadata is stored. 
	* `password` - (Required when credential_type=PASSWORD) The database user's password encoded using BASE64 scheme.
	* `password_secret_id` - (Required when credential_type=SECRET) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database password is stored. 
	* `role` - (Applicable when credential_type=PASSWORD | SECRET) The role of the database user.
	* `username` - (Applicable when credential_type=PASSWORD | SECRET) The user to connect to the database.
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

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Managed Databases Reset Database Parameter
	* `update` - (Defaults to 20 minutes), when updating the Managed Databases Reset Database Parameter
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Databases Reset Database Parameter


## Import

Import is not supported for this resource.

