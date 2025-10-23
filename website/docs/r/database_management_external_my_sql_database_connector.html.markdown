---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_my_sql_database_connector"
sidebar_current: "docs-oci-resource-database_management-external_my_sql_database_connector"
description: |-
  Provides the External My Sql Database Connector resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_my_sql_database_connector
This resource provides the External My Sql Database Connector resource in Oracle Cloud Infrastructure Database Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/database-management/latest/ExternalMySqlDatabaseConnector

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/databasemanagement

Creates an external MySQL connector resource.


## Example Usage

```hcl
resource "oci_database_management_external_my_sql_database_connector" "test_external_my_sql_database_connector" {
	#Required
	compartment_id = var.compartment_id
	connector_details {
		#Required
		credential_type = var.external_my_sql_database_connector_connector_details_credential_type
		display_name = var.external_my_sql_database_connector_connector_details_display_name
		external_database_id = oci_database_management_external_database.test_external_database.id
		host_name = var.external_my_sql_database_connector_connector_details_host_name
		macs_agent_id = oci_cloud_bridge_agent.test_agent.id
		network_protocol = var.external_my_sql_database_connector_connector_details_network_protocol
		port = var.external_my_sql_database_connector_connector_details_port
		ssl_secret_id = oci_vault_secret.test_secret.id
	}
	is_test_connection_param = var.external_my_sql_database_connector_is_test_connection_param
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) OCID of compartment for the External MySQL Database.
* `connector_details` - (Required) (Updatable) Create Details of external database connector.
	* `credential_type` - (Required) (Updatable) Type of the credential.
	* `display_name` - (Required) (Updatable) External MySQL Database Connector Name.
	* `external_database_id` - (Required) (Updatable) OCID of MySQL Database resource.
	* `host_name` - (Required) (Updatable) Host name for Connector.
	* `macs_agent_id` - (Required) (Updatable) Agent Id of the MACS agent.
	* `network_protocol` - (Required) (Updatable) Protocol to be used to connect to External MySQL Database; TCP, TCP with SSL or Socket.
	* `port` - (Required) (Updatable) Port number to connect to External MySQL Database.
	* `ssl_secret_id` - (Required) (Updatable) If using existing SSL secret to connect, OCID for the secret resource.
* `is_test_connection_param` - (Required) Parameter indicating whether database connection needs to be tested.
* `check_connection_status_trigger` - (Optional) (Updatable) An optional property when incremented triggers Check Connection Status. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `associated_services` - Oracle Cloud Infrastructure Services associated with this connector.
* `compartment_id` - OCID of compartment for the External MySQL connector.
* `connection_status` - Connection Status
* `connector_type` - Connector Type.
* `credential_type` - Credential type used to connect to database.
* `external_database_id` - OCID of MySQL Database resource
* `host_name` - Host name for Connector.
* `id` - OCID of MySQL Database Connector.
* `macs_agent_id` - Agent Id of the MACS agent.
* `name` - External MySQL Database Connector Name.
* `network_protocol` - Network Protocol.
* `port` - Connector port.
* `source_database` - Name of MySQL Database.
* `source_database_type` - Type of MySQL Database.
* `ssl_secret_id` - OCID of the SSL secret, if TCPS with SSL is used to connect to database.
* `ssl_secret_name` - Name of the SSL secret, if TCPS with SSL is used to connect to database.
* `state` - Indicates lifecycle  state of the resource.
* `time_connection_status_updated` - Time when connection status was last updated.
* `time_created` - Connector creation time.
* `time_updated` - Connector update time.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External My Sql Database Connector
	* `update` - (Defaults to 20 minutes), when updating the External My Sql Database Connector
	* `delete` - (Defaults to 20 minutes), when destroying the External My Sql Database Connector


## Import

ExternalMySqlDatabaseConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_my_sql_database_connector.test_external_my_sql_database_connector "id"
```

