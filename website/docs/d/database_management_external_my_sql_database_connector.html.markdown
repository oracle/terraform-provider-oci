---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_my_sql_database_connector"
sidebar_current: "docs-oci-datasource-database_management-external_my_sql_database_connector"
description: |-
  Provides details about a specific External My Sql Database Connector in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_my_sql_database_connector
This data source provides details about a specific External My Sql Database Connector resource in Oracle Cloud Infrastructure Database Management service.

Retrieves the MySQL database connector.


## Example Usage

```hcl
data "oci_database_management_external_my_sql_database_connector" "test_external_my_sql_database_connector" {
	#Required
	external_my_sql_database_connector_id = oci_database_management_external_my_sql_database_connector.test_external_my_sql_database_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `external_my_sql_database_connector_id` - (Required) The OCID of the External MySQL Database Connector.


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

