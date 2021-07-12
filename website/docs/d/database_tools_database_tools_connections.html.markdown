---
subcategory: "Database Tools"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_tools_database_tools_connections"
sidebar_current: "docs-oci-datasource-database_tools-database_tools_connections"
description: |-
  Provides the list of Database Tools Connections in Oracle Cloud Infrastructure Database Tools service
---

# Data Source: oci_database_tools_database_tools_connections
This data source provides the list of Database Tools Connections in Oracle Cloud Infrastructure Database Tools service.

Returns a list of DatabaseToolsConnections.


## Example Usage

```hcl
data "oci_database_tools_database_tools_connections" "test_database_tools_connections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.database_tools_connection_display_name
	state = var.database_tools_connection_state
	type = var.database_tools_connection_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `type` - (Optional) A filter to return only resources their endpointServiceId matches the given endpointServiceId.


## Attributes Reference

The following attributes are exported:

* `database_tools_connection_collection` - The list of database_tools_connection_collection.

### DatabaseToolsConnection Reference

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

