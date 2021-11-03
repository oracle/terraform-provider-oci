---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_connection"
sidebar_current: "docs-oci-resource-database_tools-database_tools_connection"
description: |-
  Provides the Database Tools Connection resource in Oracle Cloud Infrastructure Database Tools service
---

# oci_database_tools_database_tools_connection
This resource provides the Database Tools Connection resource in Oracle Cloud Infrastructure Database Tools service.

Creates a new DatabaseToolsConnection.


## Example Usage

```hcl
resource "oci_database_tools_database_tools_connection" "test_database_tools_connection" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.database_tools_connection_display_name
	type = var.database_tools_connection_type

	#Optional
	advanced_properties = var.database_tools_connection_advanced_properties
	connection_string = var.database_tools_connection_connection_string
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	key_stores {

		#Optional
		key_store_content {
			#Required
			value_type = var.database_tools_connection_key_stores_key_store_content_value_type

			#Optional
			secret_id = oci_vault_secret.test_secret.id
		}
		key_store_password {
			#Required
			value_type = var.database_tools_connection_key_stores_key_store_password_value_type

			#Optional
			secret_id = oci_vault_secret.test_secret.id
		}
		key_store_type = var.database_tools_connection_key_stores_key_store_type
	}
	private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	related_resource {
		#Required
		entity_type = var.database_tools_connection_related_resource_entity_type
		identifier = var.database_tools_connection_related_resource_identifier
	}
	user_name = oci_identity_user.test_user.name
	user_password {
		#Required
		value_type = var.database_tools_connection_user_password_value_type

		#Optional
		secret_id = oci_vault_secret.test_secret.id
	}
}
```

## Argument Reference

The following arguments are supported:

* `advanced_properties` - (Optional) (Updatable) Advanced connection properties key-value pair (e.g., oracle.net.ssl_server_dn_match).
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the containing Compartment.
* `connection_string` - (Optional) (Updatable) Connect descriptor or Easy Connect Naming method to connect to the database.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `key_stores` - (Optional) (Updatable) Oracle wallet or Java Keystores containing trusted certificates for authenticating the server's public certificate and the client private key and associated certificates required for client authentication. 
	* `key_store_content` - (Optional) (Updatable) The key store content.
		* `secret_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the key store.
		* `value_type` - (Required) (Updatable) The value type of the key store content.
	* `key_store_password` - (Optional) (Updatable) The key store password.
		* `secret_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the key store password.
		* `value_type` - (Required) (Updatable) The value type of the key store password.
	* `key_store_type` - (Optional) (Updatable) The key store type.
* `private_endpoint_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DatabaseToolsPrivateEndpoint used to access the database in the Customer VCN.
* `related_resource` - (Optional) (Updatable) The related resource
	* `entity_type` - (Required) (Updatable) The resource entity type.
	* `identifier` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
* `type` - (Required) (Updatable) The DatabaseToolsConnection type.
* `user_name` - (Optional) (Updatable) Database user name.
* `user_password` - (Optional) (Updatable) The user password.
	* `secret_id` - (Optional) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	* `value_type` - (Required) (Updatable) The value type of the user password.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `advanced_properties` - Advanced connection properties key-value pair (e.g., oracle.net.ssl_server_dn_match).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the containing Compartment.
* `connection_string` - Connect descriptor or Easy Connect Naming method to connect to the database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DatabaseToolsConnection.
* `key_stores` - Oracle wallet or Java Keystores containing trusted certificates for authenticating the server's public certificate and the client private key and associated certificates required for client authentication. 
	* `key_store_content` - The key store content.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the key store.
		* `value_type` - The value type of the key store content.
	* `key_store_password` - The key store password.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the key store password.
		* `value_type` - The value type of the key store password.
	* `key_store_type` - The key store type.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `private_endpoint_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the DatabaseToolsPrivateEndpoint used to access the database in the Customer VCN.
* `related_resource` - A related resource
	* `entity_type` - The resource entity type.
	* `identifier` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
* `state` - The current state of the DatabaseToolsConnection.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the DatabaseToolsConnection was created. An RFC3339 formatted datetime string
* `time_updated` - The time the DatabaseToolsConnection was updated. An RFC3339 formatted datetime string
* `type` - The DatabaseToolsConnection type.
* `user_name` - Database user name.
* `user_password` - The user password.
	* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	* `value_type` - The value type of the user password.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Connection
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Connection


## Import

DatabaseToolsConnections can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_database_tools_connection.test_database_tools_connection "id"
```

