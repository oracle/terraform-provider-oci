---
subcategory: "Database"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_external_database_connector"
sidebar_current: "docs-oci-resource-database-external_database_connector"
description: |-
  Provides the External Database Connector resource in Oracle Cloud Infrastructure Database service
---

# oci_database_external_database_connector
This resource provides the External Database Connector resource in Oracle Cloud Infrastructure Database service.

Creates a new external database connector.

## Example Usage

```hcl
resource "oci_database_external_database_connector" "test_external_database_connector" {
	#Required
	connection_credentials {
		#Required
		credential_name = var.external_database_connector_connection_credentials_credential_name

		#Optional
		credential_type = var.external_database_connector_connection_credentials_credential_type
		password = var.external_database_connector_connection_credentials_password
		role = var.external_database_connector_connection_credentials_role
		username = var.external_database_connector_connection_credentials_username
	}
	connection_string {
		#Required
		hostname = var.external_database_connector_connection_string_hostname
		port = var.external_database_connector_connection_string_port
		protocol = var.external_database_connector_connection_string_protocol
		service = var.external_database_connector_connection_string_service
	}
	connector_agent_id = oci_database_connector_agent.test_connector_agent.id
	display_name = var.external_database_connector_display_name
	external_database_id = oci_database_database.test_database.id

	#Optional
	connector_type = var.external_database_connector_connector_type
	defined_tags = var.external_database_connector_defined_tags
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `connection_credentials` - (Required) (Updatable) Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials. 
	* `credential_name` - (Required) (Updatable) The name of the credential information that used to connect to the database. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

		For example: inventorydb.abc112233445566778899 
	* `credential_type` - (Optional) (Updatable) The type of credential used to connect to the database.
	* `password` - (Required when credential_type=DETAILS) (Updatable) The password that will be used to connect to the database.
	* `role` - (Required when credential_type=DETAILS) (Updatable) The role of the user that will be connecting to the database.
	* `username` - (Required when credential_type=DETAILS) (Updatable) The username that will be used to connect to the database.
* `connection_string` - (Required) (Updatable) The Oracle Database connection string. 
	* `hostname` - (Required) (Updatable) The host name of the database.
	* `port` - (Required) (Updatable) The port used to connect to the database.
	* `protocol` - (Required) (Updatable) The protocol used to connect to the database.
	* `service` - (Required) (Updatable) The name of the service alias used to connect to the database.
* `connector_agent_id` - (Required) The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `connector_type` - (Optional) (Updatable) The type of connector used by the external database resource.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - (Required) (Updatable) The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique. 
* `external_database_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `connection_credentials` - Credentials used to connect to the database. Currently only the `DETAILS` type is supported for creating MACS connector crendentials. 
	* `credential_name` - The name of the credential information that used to connect to the database. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

		For example: inventorydb.abc112233445566778899 
	* `credential_type` - The type of credential used to connect to the database.
	* `password` - The password that will be used to connect to the database.
	* `role` - The role of the user that will be connecting to the database.
	* `username` - The username that will be used to connect to the database.
* `connection_status` - The status of connectivity to the external database.
* `connection_string` - The Oracle Database connection string. 
	* `hostname` - The host name of the database.
	* `port` - The port used to connect to the database.
	* `protocol` - The protocol used to connect to the database.
	* `service` - The name of the service alias used to connect to the database.
* `connector_agent_id` - The ID of the agent used for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `connector_type` - The type of connector used by the external database resource.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). 
* `display_name` - The user-friendly name for the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). The name does not have to be unique. 
* `external_database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external database resource.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the [external database connector](https://docs.cloud.oracle.com/iaas/api/#/en/database/latest/datatypes/CreateExternalDatabaseConnectorDetails). 
* `lifecycle_details` - Additional information about the current lifecycle state.
* `state` - The current lifecycle state of the external database connector resource.
* `time_connection_status_last_updated` - The date and time the `connectionStatus` of this external connector was last updated.
* `time_created` - The date and time the external connector was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Database Connector
	* `update` - (Defaults to 20 minutes), when updating the External Database Connector
	* `delete` - (Defaults to 20 minutes), when destroying the External Database Connector


## Import

ExternalDatabaseConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_database_external_database_connector.test_external_database_connector "id"
```

