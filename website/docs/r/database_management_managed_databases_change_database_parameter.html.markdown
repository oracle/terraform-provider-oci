---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_managed_databases_change_database_parameter"
sidebar_current: "docs-oci-resource-database_management-managed_databases_change_database_parameter"
description: |-
  Provides the Managed Databases Change Database Parameter resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_managed_databases_change_database_parameter
This resource provides the Managed Databases Change Database Parameter resource in Oracle Cloud Infrastructure Database Management service.

Changes database parameter values. There are two kinds of database
parameters:

- Dynamic parameters: They can be changed for the current Oracle
Database instance. The changes take effect immediately.
- Static parameters: They cannot be changed for the current instance.
You must change these parameters and then restart the database before
changes take effect.

**Note:** If the instance is started using a text initialization
parameter file, the parameter changes are applicable only for the
current instance. You must update them manually to be passed to
a future instance.


## Example Usage

```hcl
resource "oci_database_management_managed_databases_change_database_parameter" "test_managed_databases_change_database_parameter" {
	#Required
	managed_database_id = oci_database_management_managed_database.test_managed_database.id
	parameters {
		#Required
		name = var.managed_databases_change_database_parameter_parameters_name
		value = var.managed_databases_change_database_parameter_parameters_value

		#Optional
		update_comment = var.managed_databases_change_database_parameter_parameters_update_comment
	}
	scope = var.managed_databases_change_database_parameter_scope

	#Optional
	credentials {

		#Optional
		password = var.managed_databases_change_database_parameter_credentials_password
		role = var.managed_databases_change_database_parameter_credentials_role
		secret_id = oci_vault_secret.test_secret.id
		user_name = oci_identity_user.test_user.name
	}
	database_credential {
		#Required
		credential_type = var.managed_databases_change_database_parameter_database_credential_credential_type

		#Optional
		named_credential_id = oci_database_management_named_credential.test_named_credential.id
		password = var.managed_databases_change_database_parameter_database_credential_password
		password_secret_id = oci_vault_secret.test_secret.id
		role = var.managed_databases_change_database_parameter_database_credential_role
		username = var.managed_databases_change_database_parameter_database_credential_username
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
* `parameters` - (Required) A list of database parameters and their values.
	* `name` - (Required) The parameter name.
	* `update_comment` - (Optional) A comment string to associate with the change in parameter value. It cannot contain control characters or a line break. 
	* `value` - (Required) The parameter value.
* `scope` - (Required) The clause used to specify when the parameter change takes effect.

	Use `MEMORY` to make the change in memory and affect it immediately. Use `SPFILE` to make the change in the server parameter file. The change takes effect when the database is next shut down and started up again. Use `BOTH` to make the change in memory and in the server parameter file. The change takes effect immediately and persists after the database is shut down and started up again. 


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
	* `create` - (Defaults to 20 minutes), when creating the Managed Databases Change Database Parameter
	* `update` - (Defaults to 20 minutes), when updating the Managed Databases Change Database Parameter
	* `delete` - (Defaults to 20 minutes), when destroying the Managed Databases Change Database Parameter


## Import

Import is not supported for this resource.

