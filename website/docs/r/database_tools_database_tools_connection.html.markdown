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

Creates a new Database Tools connection.


## Example Usage

```hcl
resource "oci_database_tools_database_tools_connection" "test_database_tools_connection" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.database_tools_connection_display_name
	type = var.database_tools_connection_type
	user_name = oci_identity_user.test_user.name
	user_password {
		#Required
		secret_id = oci_vault_secret.test_secret.id
		value_type = var.database_tools_connection_user_password_value_type
	}

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
	locks {
		#Required
		type = var.database_tools_connection_locks_type

		#Optional
		message = var.database_tools_connection_locks_message
		related_resource_id = oci_usage_proxy_resource.test_resource.id
		time_created = var.database_tools_connection_locks_time_created
	}
	private_endpoint_id = oci_database_tools_database_tools_private_endpoint.test_database_tools_private_endpoint.id
	proxy_client {
		proxy_authentication_type = var.database_tools_connection_proxy_client_proxy_authentication_type
		roles = var.database_tools_connection_proxy_client_roles
		user_name = oci_identity_user.test_user.name
		user_password {
			#Required
			secret_id = oci_vault_secret.test_secret.id
			value_type = var.database_tools_connection_proxy_client_user_password_value_type
		}		
	}
	private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	related_resource {
		entity_type = var.database_tools_connection_related_resource_entity_type
		identifier = var.database_tools_connection_related_resource_identifier
	}
	runtime_support = var.database_tools_connection_runtime_support
	url = var.database_tools_connection_url
}
```

## Argument Reference

The following arguments are supported:

* `advanced_properties` - (Optional) (Updatable) The advanced connection properties key-value pair (e.g., `oracle.net.ssl_server_dn_match`).
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
* `connection_string` - (Required when type=MYSQL | ORACLE_DATABASE | POSTGRESQL) (Updatable) The connect descriptor or Easy Connect Naming method use to connect to the database.
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
* `locks` - (Optional) Locks associated with this resource.
	* `message` - (Optional) A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - (Optional) The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - (Optional) When the lock was created.
	* `type` - (Required) Type of the lock.
* `private_endpoint_id` - (Applicable when type=MYSQL | ORACLE_DATABASE | POSTGRESQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools private endpoint used to access the database in the customer VCN.
* `proxy_client` - (Applicable when type=ORACLE_DATABASE) (Updatable) The proxy client information.
	* `proxy_authentication_type` - (Required) (Updatable) The proxy authentication type.
	* `roles` - (Applicable when proxy_authentication_type=USER_NAME) (Updatable) A list of database roles for the client. These roles are enabled if the proxy is authorized to use the roles on behalf of the client.
	* `user_name` - (Required when proxy_authentication_type=USER_NAME) (Updatable) The user name.
	* `user_password` - (Applicable when proxy_authentication_type=USER_NAME) (Updatable) The user password.
		* `secret_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
		* `value_type` - (Required) (Updatable) The value type of the user password.
* `related_resource` - (Applicable when type=MYSQL | ORACLE_DATABASE | POSTGRESQL) (Updatable) The related resource
	* `entity_type` - (Required when type=MYSQL | ORACLE_DATABASE | POSTGRESQL) (Updatable) The resource entity type.
	* `identifier` - (Required when type=MYSQL | ORACLE_DATABASE | POSTGRESQL) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
* `runtime_support` - (Optional) Specifies whether this connection is supported by the Database Tools Runtime.
* `type` - (Required) (Updatable) The DatabaseToolsConnection type.
* `url` - (Required when type=GENERIC_JDBC) (Updatable) The JDBC URL used to connect to the Generic JDBC database system.
* `user_name` - (Required) (Updatable) The database user name.
* `user_password` - (Required) (Updatable) The user password.
	* `secret_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	* `value_type` - (Required) (Updatable) The value type of the user password.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `advanced_properties` - The advanced connection properties key-value pair (for example, `oracle.net.ssl_server_dn_match`).
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment containing the Database Tools connection.
* `connection_string` - The connect descriptor or Easy Connect Naming method used to connect to the database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools connection.
* `key_stores` - The Oracle wallet or Java Keystores containing trusted certificates for authenticating the server's public certificate and the client private key and associated certificates required for client authentication. 
	* `key_store_content` - The key store content.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the key store.
		* `value_type` - The value type of the key store content.
	* `key_store_password` - The key store password.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the key store password.
		* `value_type` - The value type of the key store password.
	* `key_store_type` - The key store type.
* `lifecycle_details` - A message describing the current state in more detail. For example, this message can be used to provide actionable information for a resource in the Failed state.
* `locks` - Locks associated with this resource.
	* `message` - A message added by the creator of the lock. This is typically used to give an indication of why the resource is locked. 
	* `related_resource_id` - The id of the resource that is locking this resource. Indicates that deleting this resource will remove the lock. 
	* `time_created` - When the lock was created.
	* `type` - Type of the lock.
* `private_endpoint_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Tools private endpoint used to access the database in the customer VCN.
* `proxy_client` - The proxy client information.
	* `proxy_authentication_type` - The proxy authentication type.
	* `roles` - A list of database roles for the client. These roles are enabled if the proxy is authorized to use the roles on behalf of the client.
	* `user_name` - The user name.
	* `user_password` - The user password.
		* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
		* `value_type` - The value type of the user password.
* `related_resource` - A related resource
	* `entity_type` - The resource entity type.
	* `identifier` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the related resource.
* `runtime_support` - Specifies whether this connection is supported by the Database Tools Runtime.
* `state` - The current state of the Database Tools connection.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Database Tools connection was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the DatabaseToolsConnection was updated. An RFC3339 formatted datetime string.
* `type` - The Database Tools connection type.
* `url` - The JDBC URL used to connect to the Generic JDBC database system.
* `user_name` - The database user name.
* `user_password` - The user password.
	* `secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
	* `value_type` - The value type of the user password.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Tools Connection
	* `update` - (Defaults to 20 minutes), when updating the Database Tools Connection
	* `delete` - (Defaults to 20 minutes), when destroying the Database Tools Connection


## Import

DatabaseToolsConnections can be imported using the `id`, e.g.

```
$ terraform import oci_database_tools_database_tools_connection.test_database_tools_connection "id"
```

