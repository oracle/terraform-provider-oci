---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database"
sidebar_current: "docs-oci-resource-data_safe-target_database"
description: |-
  Provides the Target Database resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_target_database
This resource provides the Target Database resource in Oracle Cloud Infrastructure Data Safe service.

Registers the specified database with Data Safe and creates a Data Safe target database in the Data Safe Console.


## Example Usage

```hcl
resource "oci_data_safe_target_database" "test_target_database" {
	#Required
	compartment_id = var.compartment_id
	database_details {
		#Required
		database_type = var.target_database_database_details_database_type
		infrastructure_type = var.target_database_database_details_infrastructure_type

		#Optional
		autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
		db_system_id = oci_database_db_system.test_db_system.id
		instance_id = oci_core_instance.test_instance.id
		ip_addresses = var.target_database_database_details_ip_addresses
		listener_port = var.target_database_database_details_listener_port
		service_name = oci_core_service.test_service.name
		vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
	}

	#Optional
	connection_option {
		#Required
		connection_type = var.target_database_connection_option_connection_type

		#Optional
		datasafe_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
		on_prem_connector_id = oci_data_safe_on_prem_connector.test_on_prem_connector.id
	}
	credentials {
		#Required
		password = var.target_database_credentials_password
		user_name = oci_identity_user.test_user.name
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.target_database_description
	display_name = var.target_database_display_name
	freeform_tags = {"Department"= "Finance"}
	peer_target_database_details {
		#Required
		database_details {
			#Required
			database_type = var.target_database_peer_target_database_details_database_details_database_type
			infrastructure_type = var.target_database_peer_target_database_details_database_details_infrastructure_type

			#Optional
			autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
			db_system_id = oci_database_db_system.test_db_system.id
			instance_id = oci_core_instance.test_instance.id
			ip_addresses = var.target_database_peer_target_database_details_database_details_ip_addresses
			listener_port = var.target_database_peer_target_database_details_database_details_listener_port
			service_name = oci_core_service.test_service.name
			vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
		}

		#Optional
		dataguard_association_id = oci_certificates_management_association.test_association.id
		description = var.target_database_peer_target_database_details_description
		display_name = var.target_database_peer_target_database_details_display_name
		tls_config {
			#Required
			status = var.target_database_peer_target_database_details_tls_config_status

			#Optional
			certificate_store_type = var.target_database_peer_target_database_details_tls_config_certificate_store_type
			key_store_content = var.target_database_peer_target_database_details_tls_config_key_store_content
			store_password = var.target_database_peer_target_database_details_tls_config_store_password
			trust_store_content = var.target_database_peer_target_database_details_tls_config_trust_store_content
		}
	}
	tls_config {
		#Required
		status = var.target_database_tls_config_status

		#Optional
		certificate_store_type = var.target_database_tls_config_certificate_store_type
		key_store_content = var.target_database_tls_config_key_store_content
		store_password = var.target_database_tls_config_store_password
		trust_store_content = var.target_database_tls_config_trust_store_content
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment in which to create the Data Safe target database.
* `connection_option` - (Optional) (Updatable) Types of connection supported by Data Safe.
	* `connection_type` - (Required) (Updatable) The connection type used to connect to the database. Allowed values:
		* PRIVATE_ENDPOINT - Represents connection through private endpoint in Data Safe.
		* ONPREM_CONNECTOR - Represents connection through on-premises connector in Data Safe. 
	* `datasafe_private_endpoint_id` - (Required when connection_type=PRIVATE_ENDPOINT) (Updatable) The OCID of the Data Safe private endpoint.
	* `on_prem_connector_id` - (Required when connection_type=ONPREM_CONNECTOR) (Updatable) The OCID of the on-premises connector.
* `credentials` - (Optional) (Updatable) The database credentials required for Data Safe to connect to the database.
	* `password` - (Required) (Updatable) The password of the database user.
	* `user_name` - (Required) (Updatable) The database user name.
* `database_details` - (Required) (Updatable) Details of the database for the registration in Data Safe. 
	* `autonomous_database_id` - (Required when database_type=AUTONOMOUS_DATABASE) (Updatable) The OCID of the Autonomous Database registered as a target database in Data Safe.
	* `database_type` - (Required) (Updatable) The database type.
	* `db_system_id` - (Applicable when database_type=DATABASE_CLOUD_SERVICE) (Updatable) The OCID of the cloud database registered as a target database in Data Safe.
	* `infrastructure_type` - (Required) (Updatable) The infrastructure type the database is running on.
	* `instance_id` - (Applicable when database_type=INSTALLED_DATABASE) (Updatable) The OCID of the compute instance on which the database is running.
	* `ip_addresses` - (Applicable when database_type=INSTALLED_DATABASE) (Updatable) The list of database host IP Addresses. Fully qualified domain names can be used if connectionType is 'ONPREM_CONNECTOR'. 
	* `listener_port` - (Required when database_type=DATABASE_CLOUD_SERVICE | INSTALLED_DATABASE) (Updatable) The port number of the database listener.
	* `service_name` - (Required when database_type=DATABASE_CLOUD_SERVICE | INSTALLED_DATABASE) (Updatable) The service name of the database registered as target database.
	* `vm_cluster_id` - (Applicable when database_type=DATABASE_CLOUD_SERVICE) (Updatable) The OCID of the VM cluster in which the database is running.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) The description of the target database in Data Safe.
* `display_name` - (Optional) (Updatable) The display name of the target database in Data Safe. The name is modifiable and does not need to be unique.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `peer_target_database_details` - (Optional) The details of the database to be registered as a peer target database.
	* `database_details` - (Required) Details of the database for the registration in Data Safe. 
		* `autonomous_database_id` - (Required when database_type=AUTONOMOUS_DATABASE) The OCID of the Autonomous Database registered as a target database in Data Safe.
		* `database_type` - (Required) The database type.
		* `db_system_id` - (Applicable when database_type=DATABASE_CLOUD_SERVICE) The OCID of the cloud database registered as a target database in Data Safe.
		* `infrastructure_type` - (Required) The infrastructure type the database is running on.
		* `instance_id` - (Applicable when database_type=INSTALLED_DATABASE) The OCID of the compute instance on which the database is running.
		* `ip_addresses` - (Applicable when database_type=INSTALLED_DATABASE) The list of database host IP Addresses. Fully qualified domain names can be used if connectionType is 'ONPREM_CONNECTOR'. 
		* `listener_port` - (Required when database_type=DATABASE_CLOUD_SERVICE | INSTALLED_DATABASE) The port number of the database listener.
		* `service_name` - (Required when database_type=DATABASE_CLOUD_SERVICE | INSTALLED_DATABASE) The service name of the database registered as target database.
		* `vm_cluster_id` - (Applicable when database_type=DATABASE_CLOUD_SERVICE) The OCID of the VM cluster in which the database is running.
	* `dataguard_association_id` - (Optional) The OCID of the Data Guard Association resource in which the database being registered is considered as peer database to the primary database.
	* `description` - (Optional) The description of the peer target database in Data Safe.
	* `display_name` - (Optional) The display name of the peer target database in Data Safe. The name is modifiable and does not need to be unique.
	* `tls_config` - (Optional) The details required to establish a TLS enabled connection.
		* `certificate_store_type` - (Optional) The format of the certificate store.
		* `key_store_content` - (Optional) Base64 encoded string of key store file content.
		* `status` - (Required) Status to represent whether the database connection is TLS enabled or not.
		* `store_password` - (Optional) The password to read the trust store and key store files, if they are password protected.
		* `trust_store_content` - (Optional) Base64 encoded string of trust store file content.
* `tls_config` - (Optional) (Updatable) The details required to establish a TLS enabled connection.
	* `certificate_store_type` - (Optional) (Updatable) The format of the certificate store.
	* `key_store_content` - (Optional) (Updatable) Base64 encoded string of key store file content.
	* `status` - (Required) (Updatable) Status to represent whether the database connection is TLS enabled or not.
	* `store_password` - (Optional) (Updatable) The password to read the trust store and key store files, if they are password protected.
	* `trust_store_content` - (Optional) (Updatable) Base64 encoded string of trust store file content.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Target Database
	* `update` - (Defaults to 20 minutes), when updating the Target Database
	* `delete` - (Defaults to 20 minutes), when destroying the Target Database


## Import

TargetDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_target_database.test_target_database "id"
```

