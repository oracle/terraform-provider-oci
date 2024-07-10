---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_connection"
sidebar_current: "docs-oci-resource-database_migration-connection"
description: |-
Provides the Connection resource in Oracle Cloud Infrastructure Database Migration service
---

# oci_database_migration_connection
This resource provides the Connection resource in Oracle Cloud Infrastructure Database Migration service.

Create a Database Connection resource that contains the details to connect to either a Source or Target Database
in the migration.

Note: If you wish to use the DMS deprecated API version /20210929 it is necessary to pin the Terraform Provider version to v5.47.0. Newer Terraform provider versions will not support the DMS deprecated API version /20210929

## Example Usage

```hcl
resource "oci_database_migration_connection" "test_connection" {
	#Required
	compartment_id = var.compartment_id
	connection_type = var.connection_connection_type
	display_name = var.connection_display_name
	key_id = oci_kms_key.test_key.id
	password = var.connection_password
	technology_type = var.connection_technology_type
	username = var.connection_username
	vault_id = oci_kms_vault.test_vault.id

	#Optional
	additional_attributes {

		#Optional
		name = var.connection_additional_attributes_name
		value = var.connection_additional_attributes_value
	}
	connection_string = var.connection_connection_string
	database_id = oci_database_database.test_database.id
	database_name = oci_database_database.test_database.name
	db_system_id = oci_database_db_system.test_db_system.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.connection_description
	freeform_tags = var.connection_freeform_tags
	host = var.connection_host
	nsg_ids = var.connection_nsg_ids
	port = var.connection_port
	replication_password = var.connection_replication_password
	replication_username = var.connection_replication_username
	security_protocol = var.connection_security_protocol
	ssh_host = var.connection_ssh_host
	ssh_key = var.connection_ssh_key
	ssh_sudo_location = var.connection_ssh_sudo_location
	ssh_user = var.connection_ssh_user
	ssl_ca = var.connection_ssl_ca
	ssl_cert = var.connection_ssl_cert
	ssl_crl = var.connection_ssl_crl
	ssl_key = var.connection_ssl_key
	ssl_mode = var.connection_ssl_mode
	subnet_id = oci_core_subnet.test_subnet.id
	wallet = var.connection_wallet
}
```

## Argument Reference

The following arguments are supported:

* `additional_attributes` - (Applicable when connection_type=MYSQL) (Updatable) An array of name-value pair attribute entries.
  * `name` - (Required when connection_type=MYSQL) (Updatable) The name of the property entry.
  * `value` - (Required when connection_type=MYSQL) (Updatable) The value of the property entry.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment.
* `connection_string` - (Applicable when connection_type=ORACLE) (Updatable) Connect descriptor or Easy Connect Naming method used to connect to a database.
* `connection_type` - (Required) (Updatable) Defines the type of connection. For example, ORACLE.
* `database_id` - (Applicable when connection_type=ORACLE) (Updatable) The OCID of the database being referenced.
* `database_name` - (Required when connection_type=MYSQL) (Updatable) The name of the database being referenced.
* `db_system_id` - (Applicable when connection_type=MYSQL) (Updatable) The OCID of the database system being referenced.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `description` - (Optional) (Updatable) A user-friendly description. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable.  Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace.  For more information, see Resource Tags. Example: {"Department": "Finance"}
* `host` - (Applicable when connection_type=MYSQL) (Updatable) The IP Address of the host.
* `key_id` - (Required) (Updatable) The OCID of the key used in cryptographic operations.
* `nsg_ids` - (Optional) (Updatable) An array of Network Security Group OCIDs used to define network access for Connections.
* `password` - (Required) (Updatable) The password (credential) used when creating or updating this resource.
* `port` - (Applicable when connection_type=MYSQL) (Updatable) The port to be used for the connection.
* `replication_password` - (Optional) (Updatable) The password (credential) used when creating or updating this resource.
* `replication_username` - (Optional) (Updatable) The username (credential) used when creating or updating this resource.
* `security_protocol` - (Required when connection_type=MYSQL) (Updatable) Security Type for MySQL.
* `ssh_host` - (Applicable when connection_type=ORACLE) (Updatable) Name of the host the SSH key is valid for.
* `ssh_key` - (Applicable when connection_type=ORACLE) (Updatable) Private SSH key string.
* `ssh_sudo_location` - (Applicable when connection_type=ORACLE) (Updatable) Sudo location
* `ssh_user` - (Applicable when connection_type=ORACLE) (Updatable) The username (credential) used when creating or updating this resource.
* `ssl_ca` - (Applicable when connection_type=MYSQL) (Updatable) Database Certificate - The base64 encoded content of mysql.pem file containing the server public key (for 1 and 2-way SSL).
* `ssl_cert` - (Applicable when connection_type=MYSQL) (Updatable) Client Certificate - The base64 encoded content of client-cert.pem file  containing the client public key (for 2-way SSL).
* `ssl_crl` - (Applicable when connection_type=MYSQL) (Updatable) Certificates revoked by certificate authorities (CA). Server certificate must not be on this list (for 1 and 2-way SSL). Note: This is an optional and that too only applicable if TLS/MTLS option is selected.
* `ssl_key` - (Applicable when connection_type=MYSQL) (Updatable) Client Key - The client-key.pem containing the client private key (for 2-way SSL).
* `ssl_mode` - (Applicable when connection_type=MYSQL) (Updatable) SSL modes for MySQL.
* `subnet_id` - (Optional) (Updatable) Oracle Cloud Infrastructure resource ID.
* `technology_type` - (Required) The type of MySQL source or target connection. Example: OCI_MYSQL represents Oracle Cloud Infrastructure MySQL HeatWave Database Service
* `username` - (Required) (Updatable) The username (credential) used when creating or updating this resource.
* `vault_id` - (Required) (Updatable) Oracle Cloud Infrastructure resource ID.
* `wallet` - (Applicable when connection_type=ORACLE) (Updatable) The wallet contents used to make connections to a database.  This attribute is expected to be base64 encoded.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
* `create` - (Defaults to 20 minutes), when creating the Connection
* `update` - (Defaults to 20 minutes), when updating the Connection
* `delete` - (Defaults to 20 minutes), when destroying the Connection


## Import

Connections can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_connection.test_connection "id"
```
