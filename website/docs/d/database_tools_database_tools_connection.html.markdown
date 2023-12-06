---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_connection"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_connection"
description: |-
  Provides details about a specific Database Tools Connection in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_connection
This data source provides details about a specific Database Tools Connection resource in Oracle Cloud Infrastructure Database Tools service.

Gets details of the specified Database Tools connection.

## Example Usage

```hcl
data "oci_database_tools_database_tools_connection" "test_database_tools_connection" {
	#Required
	database_tools_connection_id = oci_database_tools_database_tools_connection.test_database_tools_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `database_tools_connection_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a Database Tools connection.


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

