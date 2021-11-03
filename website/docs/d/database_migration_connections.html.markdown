---
subcategory: "Database Migration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_migration_connections"
sidebar_current: "docs-oci-datasource-database_migration-connections"
description: |-
  Provides the list of Connections in Oracle Cloud Infrastructure Database Migration service
---

# Data Source: oci_database_migration_connections
This data source provides the list of Connections in Oracle Cloud Infrastructure Database Migration service.

List all Database Connections.

## Example Usage

```hcl
data "oci_database_migration_connections" "test_connections" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.connection_display_name
	state = var.connection_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given. 
* `state` - (Optional) The current state of the Database Migration Deployment. 


## Attributes Reference

The following attributes are exported:

* `connection_collection` - The list of connection_collection.

### Connection Reference

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

