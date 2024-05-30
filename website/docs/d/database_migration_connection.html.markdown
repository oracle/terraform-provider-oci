---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_connection"
sidebar_current: "docs-oci-datasource-database_migration-connection"
description: |-
Provides details about a specific Connection in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_connection
This data source provides details about a specific Connection resource in Oracle Cloud Infrastructure Database Migration service.

Display Database Connection details.

Note: If you wish to use the DMS deprecated API version /20210929 it is necessary to pin the Terraform Provider version to v5.47.0. Newer Terraform provider versions will not support the DMS deprecated API version /20210929

## Example Usage

```hcl
data "oci_database_migration_connection" "test_connection" {
	#Required
	connection_id = oci_database_migration_connection.test_connection.id
}
```

## Argument Reference

The following arguments are supported:

* `connection_id` - (Required) The OCID of the database connection.


## Attributes Reference

The following attributes are exported:

* `additional_attributes` - An array of name-value pair attribute entries.
  * `name` - The name of the property entry.
  * `value` - The value of the property entry.
* `compartment_id` - The OCID of the compartment.
* `connection_string` - Connect descriptor or Easy Connect Naming method used to connect to a database.
* `connection_type` - Defines the type of connection. For example, ORACLE.
* `database_id` - The OCID of the database being referenced.
* `database_name` - The name of the database being referenced.
* `db_system_id` - The OCID of the database system being referenced.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `description` - A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"}
* `host` - The IP Address of the host.
* `id` - The OCID of the connection being referenced.
* `ingress_ips` - List of ingress IP addresses from where to connect to this connection's privateIp.
  * `ingress_ip` - A Private Endpoint IPv4 or IPv6 Address created in the customer's subnet.
* `key_id` - The OCID of the key used in cryptographic operations.
* `lifecycle_details` - The message describing the current state of the connection's lifecycle in detail. For example, can be used to provide actionable information for a connection in a Failed state.
* `nsg_ids` - An array of Network Security Group OCIDs used to define network access for Connections.
* `password` - The password (credential) used when creating or updating this resource.
* `port` - The port to be used for the connection.
* `private_endpoint_id` - The OCID of the resource being referenced.
* `replication_password` - The password (credential) used when creating or updating this resource.
* `replication_username` - The username (credential) used when creating or updating this resource.
* `secret_id` - The OCID of the resource being referenced.
* `security_protocol` - Security Protocol to be used for the connection.
* `ssh_host` - Name of the host the SSH key is valid for.
* `ssh_key` - Private SSH key string.
* `ssh_sudo_location` - Sudo location
* `ssh_user` - The username (credential) used when creating or updating this resource.
* `ssl_mode` - SSL mode to be used for the connection.
* `state` - The Connection's current lifecycle state.
* `subnet_id` - Oracle Cloud Infrastructure resource ID.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `technology_type` - The type of MySQL source or target connection. Example: OCI_MYSQL represents Oracle Cloud Infrastructure MySQL HeatWave Database Service
* `time_created` - The time when this resource was created. An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
* `time_updated` - The time when this resource was updated. An RFC3339 formatted datetime string such as `2016-08-25T21:10:29.600Z`.
* `username` - The username (credential) used when creating or updating this resource.
* `vault_id` - Oracle Cloud Infrastructure resource ID.
