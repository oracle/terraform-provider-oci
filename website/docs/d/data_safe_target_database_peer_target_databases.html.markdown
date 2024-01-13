---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_target_database_peer_target_databases"
sidebar_current: "docs-oci-datasource-data_safe-target_database_peer_target_databases"
description: |-
  Provides the list of Target Database Peer Target Databases in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_target_database_peer_target_databases
This data source provides the list of Target Database Peer Target Databases in Oracle Cloud Infrastructure Data Safe service.

Lists all the peer target databases under the primary target database identified by the OCID passed as path parameter.

## Example Usage

```hcl
data "oci_data_safe_target_database_peer_target_databases" "test_target_database_peer_target_databases" {
	#Required
	target_database_id = oci_data_safe_target_database.test_target_database.id
}
```

## Argument Reference

The following arguments are supported:

* `target_database_id` - (Required) The OCID of the Data Safe target database.


## Attributes Reference

The following attributes are exported:

* `peer_target_database_collection` - The list of peer_target_database_collection.

### TargetDatabasePeerTargetDatabase Reference

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

