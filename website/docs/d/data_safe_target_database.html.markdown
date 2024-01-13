---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database"
sidebar_current: "docs-oci-datasource-data_safe-target_database"
description: |-
  Provides details about a specific Target Database in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_database
This data source provides details about a specific Target Database resource in Oracle Cloud Infrastructure Data Safe service.

Returns the details of the specified Data Safe target database.


## Example Usage

```hcl
data "oci_data_safe_target_database" "test_target_database" {
	#Required
	target_database_id = oci_data_safe_target_database.test_target_database.id
}
```

## Argument Reference

The following arguments are supported:

* `target_database_id` - (Required) The OCID of the Data Safe target database.


## Attributes Reference

The following attributes are exported:

* `associated_resource_ids` - The OCIDs of associated resources like database, Data Safe private endpoint etc.
* `compartment_id` - The OCID of the compartment which contains the Data Safe target database.
* `connection_option` - Types of connection supported by Data Safe.
	* `connection_type` - The connection type used to connect to the database. Allowed values:
		* PRIVATE_ENDPOINT - Represents connection through private endpoint in Data Safe.
		* ONPREM_CONNECTOR - Represents connection through on-premises connector in Data Safe. 
	* `datasafe_private_endpoint_id` - The OCID of the Data Safe private endpoint.
	* `on_prem_connector_id` - The OCID of the on-premises connector.
* `credentials` - The database credentials required for Data Safe to connect to the database.
	* `password` - The password of the database user.
	* `user_name` - The database user name.
* `database_details` - Details of the database for the registration in Data Safe. 
	* `autonomous_database_id` - The OCID of the Autonomous Database registered as a target database in Data Safe.
	* `database_type` - The database type.
	* `db_system_id` - The OCID of the cloud database registered as a target database in Data Safe.
	* `infrastructure_type` - The infrastructure type the database is running on.
	* `instance_id` - The OCID of the compute instance on which the database is running.
	* `ip_addresses` - The list of database host IP Addresses. Fully qualified domain names can be used if connectionType is 'ONPREM_CONNECTOR'. 
	* `listener_port` - The port number of the database listener.
	* `service_name` - The service name of the database registered as target database.
	* `vm_cluster_id` - The OCID of the VM cluster in which the database is running.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the target database in Data Safe.
* `display_name` - The display name of the target database in Data Safe.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the Data Safe target database.
* `lifecycle_details` - Details about the current state of the target database in Data Safe.
* `peer_target_databases` - The OCIDs of associated resources like Database, Data Safe private endpoint etc.
	* `database_details` - Details of the database for the registration in Data Safe. 
		* `autonomous_database_id` - The OCID of the Autonomous Database registered as a target database in Data Safe.
		* `database_type` - The database type.
		* `db_system_id` - The OCID of the cloud database registered as a target database in Data Safe.
		* `infrastructure_type` - The infrastructure type the database is running on.
		* `instance_id` - The OCID of the compute instance on which the database is running.
		* `ip_addresses` - The list of database host IP Addresses. Fully qualified domain names can be used if connectionType is 'ONPREM_CONNECTOR'. 
		* `listener_port` - The port number of the database listener.
		* `service_name` - The service name of the database registered as target database.
		* `vm_cluster_id` - The OCID of the VM cluster in which the database is running.
	* `database_unique_name` - Unique name of the database associated to the peer target database.
	* `dataguard_association_id` - The OCID of the Data Guard Association resource in which the database associated to the peer target database is considered as peer database to the primary database.
	* `description` - The description of the peer target database in Data Safe.
	* `display_name` - The display name of the peer target database in Data Safe.
	* `key` - The secondary key assigned for the peer target database in Data Safe.
	* `lifecycle_details` - Details about the current state of the peer target database in Data Safe.
	* `role` - Role of the database associated to the peer target database.
	* `state` - The current state of the peer target database in Data Safe.
	* `time_created` - The date and time of the peer target database registration in Data Safe.
	* `tls_config` - The details required to establish a TLS enabled connection.
		* `certificate_store_type` - The format of the certificate store.
		* `key_store_content` - Base64 encoded string of key store file content.
		* `status` - Status to represent whether the database connection is TLS enabled or not.
		* `store_password` - The password to read the trust store and key store files, if they are password protected.
		* `trust_store_content` - Base64 encoded string of trust store file content.
* `state` - The current state of the target database in Data Safe.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the database was registered in Data Safe and created as a target database in Data Safe.
* `time_updated` - The date and time of the target database update in Data Safe.
* `tls_config` - The details required to establish a TLS enabled connection.
	* `certificate_store_type` - The format of the certificate store.
	* `key_store_content` - Base64 encoded string of key store file content.
	* `status` - Status to represent whether the database connection is TLS enabled or not.
	* `store_password` - The password to read the trust store and key store files, if they are password protected.
	* `trust_store_content` - Base64 encoded string of trust store file content.

