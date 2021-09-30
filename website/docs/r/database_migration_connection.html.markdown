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

## Example Usage

```hcl
resource "oci_database_migration_connection" "test_connection" {
	#Required
	admin_credentials {
		#Required
		password = var.connection_admin_credentials_password
		username = var.connection_admin_credentials_username
	}
	compartment_id = var.compartment_id
	database_type = var.connection_database_type
	vault_details {
		#Required
		compartment_id = var.compartment_id
		key_id = oci_kms_key.test_key.id
		vault_id = oci_kms_vault.test_vault.id
	}

	#Optional
	certificate_tdn = var.connection_certificate_tdn
	connect_descriptor {

		#Optional
		connect_string = var.connection_connect_descriptor_connect_string
		database_service_name = oci_core_service.test_service.name
		host = var.connection_connect_descriptor_host
		port = var.connection_connect_descriptor_port
	}
	database_id = oci_database_database.test_database.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.connection_display_name
	freeform_tags = {"bar-key"= "value"}
	private_endpoint {
		#Required
		compartment_id = var.compartment_id
		subnet_id = oci_core_subnet.test_subnet.id
		vcn_id = oci_core_vcn.test_vcn.id
	}
	ssh_details {
		#Required
		host = var.connection_ssh_details_host
		sshkey = var.connection_ssh_details_sshkey
		user = var.connection_ssh_details_user

		#Optional
		sudo_location = var.connection_ssh_details_sudo_location
	}
	tls_keystore = var.connection_tls_keystore
	tls_wallet = var.connection_tls_wallet
}
```

## Argument Reference

The following arguments are supported:

* `admin_credentials` - (Required) (Updatable) Database Administrator Credentials details. 
	* `password` - (Required) (Updatable) Administrator password 
	* `username` - (Required) (Updatable) Administrator username 
* `certificate_tdn` - (Optional) (Updatable) This name is the distinguished name used while creating the certificate on target database. Requires a TLS wallet to be specified. Not required for source container database connections. 
* `compartment_id` - (Required) (Updatable) OCID of the compartment 
* `connect_descriptor` - (Optional) (Updatable) Connect Descriptor details. Required for Manual and UserManagerOci connection types. If a Private Endpoint was specified for the Connection, the host should contain a valid IP address. 
	* `connect_string` - (Optional) (Updatable) Connect String. Required if no host, port nor databaseServiceName were specified. If a Private Endpoint was specified in the Connection, the host entry should be a valid IP address. Supported formats: Easy connect: <host>:<port>/<db_service_name> Long format: (description= (address=(port=<port>)(host=<host>))(connect_data=(service_name=<db_service_name>))) 
	* `database_service_name` - (Optional) (Updatable) Database service name. Required if no connectString was specified. 
	* `host` - (Optional) (Updatable) Host or IP address of the connect descriptor. Required if no connectString was specified. 
	* `port` - (Optional) (Updatable) Port of the connect descriptor. Required if no connectString was specified. 
* `database_id` - (Optional) (Updatable) The OCID of the cloud database. Required if the database connection type is Autonomous. 
* `database_type` - (Required) Database connection type. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) Database Connection display name identifier. 
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `private_endpoint` - (Optional) (Updatable) Oracle Cloud Infrastructure Private Endpoint configuration details. Not required for source container database connections, it will default to the specified Source Database Connection Private Endpoint. 
	* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the private endpoint.  
	* `subnet_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer's subnet where the private endpoint VNIC will reside. 
	* `vcn_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN where the Private Endpoint will be bound to. 
* `ssh_details` - (Optional) (Updatable) Details of the SSH key that will be used. Required for source database Manual and UserManagerOci connection types. Not required for source container database connections. 
	* `host` - (Required) (Updatable) Name of the host the SSH key is valid for. 
	* `sshkey` - (Required) (Updatable) Private SSH key string. 
	* `sudo_location` - (Optional) (Updatable) Sudo location 
	* `user` - (Required) (Updatable) SSH user 
* `tls_keystore` - (Optional) (Updatable) keystore.jks file contents; base64 encoded String. Requires a TLS wallet to be specified. Not required for source container database connections. 
* `tls_wallet` - (Optional) (Updatable) cwallet.sso containing containing the TCPS/SSL certificate; base64 encoded String. Not required for source container database connections. 
* `vault_details` - (Required) (Updatable) Oracle Cloud Infrastructure Vault details to store migration and connection credentials secrets 
	* `compartment_id` - (Required) (Updatable) OCID of the compartment where the secret containing the credentials will be created. 
	* `key_id` - (Required) (Updatable) OCID of the vault encryption key 
	* `vault_id` - (Required) (Updatable) OCID of the vault 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `admin_credentials` - Database Administrator Credentials details. 
	* `username` - Administrator username 
* `certificate_tdn` - This name is the distinguished name used while creating the certificate on target database. 
* `compartment_id` - OCID of the compartment 
* `connect_descriptor` - Connect Descriptor details. 
	* `connect_string` - Connect string. 
	* `database_service_name` - Database service name. 
	* `host` - Host of the connect descriptor. 
	* `port` - Port of the connect descriptor. 
* `credentials_secret_id` - OCID of the Secret in the Oracle Cloud Infrastructure vault containing the Database Connection credentials. 
* `database_id` - The OCID of the cloud database. 
* `database_type` - Database connection type. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - Database Connection display name identifier. 
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the resource 
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state. 
* `private_endpoint` - Oracle Cloud Infrastructure Private Endpoint configuration details. 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to contain the private endpoint. 
	* `id` - [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of a previously created Private Endpoint. 
	* `subnet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the customer's subnet where the private endpoint VNIC will reside. 
	* `vcn_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VCN where the Private Endpoint will be bound to. 
* `ssh_details` - Details of the SSH key that will be used. 
	* `host` - Name of the host the SSH key is valid for. 
	* `sudo_location` - Sudo location 
	* `user` - SSH user 
* `state` - The current state of the Connection resource. 
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time the Connection resource was created. An RFC3339 formatted datetime string. 
* `time_updated` - The time of the last Connection resource details update. An RFC3339 formatted datetime string. 
* `vault_details` - Oracle Cloud Infrastructure Vault details to store migration and connection credentials secrets 
	* `compartment_id` - OCID of the compartment where the secret containing the credentials will be created. 
	* `key_id` - OCID of the vault encryption key 
	* `vault_id` - OCID of the vault 

## Import

Connections can be imported using the `id`, e.g.

```
$ terraform import oci_database_migration_connection.test_connection "id"
```

