---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database_peer_target_database"
sidebar_current: "docs-oci-resource-data_safe-target_database_peer_target_database"
description: |-
  Provides the Target Database Peer Target Database resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_target_database_peer_target_database
This resource provides the Target Database Peer Target Database resource in Oracle Cloud Infrastructure Data Safe service.

Creates the peer target database under the primary target database in Data Safe.

## Example Usage

```hcl
resource "oci_data_safe_target_database_peer_target_database" "test_target_database_peer_target_database" {
	#Required
	database_details {
		#Required
		database_type = var.target_database_peer_target_database_database_details_database_type
		infrastructure_type = var.target_database_peer_target_database_database_details_infrastructure_type

		#Optional
		autonomous_database_id = oci_database_autonomous_database.test_autonomous_database.id
		db_system_id = oci_database_db_system.test_db_system.id
		instance_id = oci_core_instance.test_instance.id
		ip_addresses = var.target_database_peer_target_database_database_details_ip_addresses
		listener_port = var.target_database_peer_target_database_database_details_listener_port
		service_name = oci_core_service.test_service.name
		vm_cluster_id = oci_database_vm_cluster.test_vm_cluster.id
	}
	target_database_id = oci_data_safe_target_database.test_target_database.id

	#Optional
	dataguard_association_id = oci_certificates_management_association.test_association.id
	description = var.target_database_peer_target_database_description
	display_name = var.target_database_peer_target_database_display_name
	tls_config {
		#Required
		status = var.target_database_peer_target_database_tls_config_status

		#Optional
		certificate_store_type = var.target_database_peer_target_database_tls_config_certificate_store_type
		key_store_content = var.target_database_peer_target_database_tls_config_key_store_content
		store_password = var.target_database_peer_target_database_tls_config_store_password
		trust_store_content = var.target_database_peer_target_database_tls_config_trust_store_content
	}
}
```

## Argument Reference

The following arguments are supported:

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
* `dataguard_association_id` - (Optional) The OCID of the Data Guard Association resource in which the database being registered is considered as peer database to the primary database.
* `description` - (Optional) (Updatable) The description of the peer target database in Data Safe.
* `display_name` - (Optional) (Updatable) The display name of the peer target database in Data Safe. The name is modifiable and does not need to be unique.
* `target_database_id` - (Required) The OCID of the Data Safe target database.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Target Database Peer Target Database
	* `update` - (Defaults to 20 minutes), when updating the Target Database Peer Target Database
	* `delete` - (Defaults to 20 minutes), when destroying the Target Database Peer Target Database


## Import

TargetDatabasePeerTargetDatabases can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_target_database_peer_target_database.test_target_database_peer_target_database "targetDatabases/{targetDatabaseId}/peerTargetDatabases/{peerTargetDatabaseId}" 
```

