---
subcategory: "Database Tools Runtime"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_runtime_database_tools_connection_credential"
sidebar_current: "docs-oci-resource-database_tools_runtime-database_tools_connection_credential"
description: |-
  Provides the Database Tools Connection Credential resource in Oracle Cloud Infrastructure Database Tools Runtime service
---

# oci_database_tools_runtime_database_tools_connection_credential
This resource provides the Database Tools Connection Credential resource in Oracle Cloud Infrastructure Database Tools Runtime service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/database_tools_runtime

Creates a credential for the user specified by the key.

## Example Usage

```hcl
resource "oci_database_tools_runtime_database_tools_connection_credential" "test_database_tools_connection_credential" {
	#Required
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
	key = var.database_tools_connection_credential_key
	password = var.database_tools_connection_credential_password
	type = var.database_tools_connection_credential_type
	user_name = oci_identity_user.test_user.name
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.
* `key` - (Required) The credential_name to be created
* `password` - (Required) (Updatable) The password for the new credential.
* `type` - (Required) (Updatable) The type of credential.
* `user_name` - (Required) (Updatable) The username for the new credential.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `enabled` - Indicates whether this credential is enabled (TRUE) or not (FALSE)
* `key` - Name of the credential
* `key_type` - Indicates whether this refers to a public synonym or not.
* `owner` - Owner of the credential
* `related_resource` - A related resource for a credential.
	* `identifier` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
	* `type` - The related resource type.
* `user_name` - Name of the user that will be used to log in to the remote database or the remote or local operating system
* `windows_domain` - For a Windows target, the Windows domain to use when logging in

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Connection Credential
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Connection Credential
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Connection Credential


## Import

DatabaseToolsConnectionCredentials can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_runtime_database_tools_connection_credential.test_database_tools_connection_credential "databaseToolsConnections/{databaseToolsConnectionId}/credentials/{credentialKey}" 
```

