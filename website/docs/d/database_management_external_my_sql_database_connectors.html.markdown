---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_my_sql_database_connectors"
sidebar_current: "docs-oci-datasource-database_management-external_my_sql_database_connectors"
description: |-
  Provides the list of External My Sql Database Connectors in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_my_sql_database_connectors
This data source provides the list of External My Sql Database Connectors in Oracle Cloud Infrastructure Database Management service.

Gets the list of External MySQL Database connectors. 


## Example Usage

```hcl
data "oci_database_management_external_my_sql_database_connectors" "test_external_my_sql_database_connectors" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	name = var.external_my_sql_database_connector_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `name` - (Optional) The parameter to filter by MySQL Database System type.


## Attributes Reference

The following attributes are exported:

* `my_sql_connector_collection` - The list of my_sql_connector_collection.

### ExternalMySqlDatabaseConnector Reference

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

