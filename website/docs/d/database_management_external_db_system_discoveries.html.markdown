---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_db_system_discoveries"
sidebar_current: "docs-oci-datasource-database_management-external_db_system_discoveries"
description: |-
  Provides the list of External Db System Discoveries in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_db_system_discoveries
This data source provides the list of External Db System Discoveries in Oracle Cloud Infrastructure Database Management service.

Lists the external DB system discovery resources in the specified compartment.

## Example Usage

```hcl
data "oci_database_management_external_db_system_discoveries" "test_external_db_system_discoveries" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.external_db_system_discovery_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to only return the resources that match the entire display name.


## Attributes Reference

The following attributes are exported:

* `external_db_system_discovery_collection` - The list of external_db_system_discovery_collection.

### ExternalDbSystemDiscovery Reference

The following attributes are exported:

* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the external DB system discovery. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `discovered_components` - The list of DB system components that were found in the DB system discovery.
	* `adr_home_directory` - The directory that stores tracing and logging incidents when Automatic Diagnostic Repository (ADR) is enabled.
	* `asm_instances` - 
		* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the ASM instance.
		* `host_name` - The name of the host on which the ASM instance is running.
		* `instance_name` - The name of the ASM instance.
	* `associated_components` - The list of associated components.
		* `association_type` - The association type.
		* `component_id` - The identifier of the associated component.
		* `component_type` - The type of associated component.
	* `cluster_id` - The unique identifier of the Oracle cluster.
	* `cluster_instances` - 
		* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
		* `cluster_id` - The unique identifier of the Oracle cluster.
		* `connector` - The connector details used to connect to the external DB system component.
			* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the external DB system connector. 
			* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
			* `connection_info` - The connection details required to connect to an external DB system component.
				* `component_type` - The component type.
				* `connection_credentials` - The credentials used to connect to the ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
					* `credential_name` - The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

						For example: inventorydb.abc112233445566778899 
					* `credential_type` - The type of credential used to connect to the ASM instance.
					* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Named Credential where the database password metadata is stored. 
					* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
					* `role` - The role of the user connecting to the ASM instance.
					* `ssl_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
					* `user_name` - The user name used to connect to the ASM instance.
				* `connection_string` - The Oracle Database connection string. 
					* `host_name` - The host name of the database or the SCAN name in case of a RAC database.
					* `hosts` - The list of host names of the ASM instances.
					* `port` - The port used to connect to the ASM instance.
					* `protocol` - The protocol used to connect to the ASM instance.
					* `service` - The service name of the ASM instance.
				* `database_credential` - The credential to connect to the database to perform tablespace administration tasks.
					* `credential_type` - The type of the credential for tablespace administration tasks.
					* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential where the database password metadata is stored. 
					* `password` - The database user's password encoded using BASE64 scheme.
					* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database password is stored. 
					* `role` - The role of the database user.
					* `username` - The user to connect to the database.
			* `connection_status` - The status of connectivity to the external DB system component.
			* `connector_type` - The type of connector.
			* `display_name` - The user-friendly name for the external connector. The name does not have to be unique.
			* `time_connection_status_last_updated` - The date and time the connectionStatus of the external DB system connector was last updated.
		* `crs_base_directory` - The Oracle base location of Cluster Ready Services (CRS).
		* `host_name` - The name of the host on which the cluster instance is running.
		* `node_role` - The role of the cluster node.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `component_id` - The identifier of the discovered DB system component.
	* `component_name` - The name of the discovered DB system component.
	* `component_type` - The external DB system component type.
	* `connector` - The connector details used to connect to the external DB system component.
		* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the external DB system connector. 
		* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
		* `connection_info` - The connection details required to connect to an external DB system component.
			* `component_type` - The component type.
			* `connection_credentials` - The credentials used to connect to the ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
				* `credential_name` - The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

					For example: inventorydb.abc112233445566778899 
				* `credential_type` - The type of credential used to connect to the ASM instance.
				* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Named Credential where the database password metadata is stored. 
				* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
				* `role` - The role of the user connecting to the ASM instance.
				* `ssl_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
				* `user_name` - The user name used to connect to the ASM instance.
			* `connection_string` - The Oracle Database connection string. 
				* `host_name` - The host name of the database or the SCAN name in case of a RAC database.
				* `hosts` - The list of host names of the ASM instances.
				* `port` - The port used to connect to the ASM instance.
				* `protocol` - The protocol used to connect to the ASM instance.
				* `service` - The service name of the ASM instance.
			* `database_credential` - The credential to connect to the database to perform tablespace administration tasks.
				* `credential_type` - The type of the credential for tablespace administration tasks.
				* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential where the database password metadata is stored. 
				* `password` - The database user's password encoded using BASE64 scheme.
				* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database password is stored. 
				* `role` - The role of the database user.
				* `username` - The user to connect to the database.
		* `connection_status` - The status of connectivity to the external DB system component.
		* `connector_type` - The type of connector.
		* `display_name` - The user-friendly name for the external connector. The name does not have to be unique.
		* `time_connection_status_last_updated` - The date and time the connectionStatus of the external DB system connector was last updated.
	* `container_database_id` - The unique identifier of the parent Container Database (CDB).
	* `cpu_core_count` - The number of CPU cores available on the DB node.
	* `crs_base_directory` - The Oracle base location of Cluster Ready Services (CRS).
	* `db_edition` - The Oracle Database edition.
	* `db_id` - The Oracle Database ID.
	* `db_node_name` - The name of the DB node.
	* `db_packs` - The database packs licensed for the external Oracle Database.
	* `db_role` - The role of the Oracle Database in Oracle Data Guard configuration.
	* `db_type` - The type of Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or a Non-container Database. 
	* `db_unique_name` - The `DB_UNIQUE_NAME` of the external database.
	* `db_version` - The Oracle Database version.
	* `display_name` - The user-friendly name for the discovered DB system component. The name does not have to be unique.
	* `endpoints` - The list of protocol addresses the listener is configured to listen on.
		* `host` - The host name or IP address.
		* `key` - The unique name of the service.
		* `port` - The port number.
		* `protocol` - The listener protocol.
		* `services` - The list of services registered with the listener.
	* `grid_home` - The directory in which ASM is installed. This is the same directory in which Oracle Grid Infrastructure is installed.
	* `guid` - The unique identifier of the PDB.
	* `home_directory` - The location of the DB home.
	* `host_name` - The name of the host on which the external listener is running.
	* `instance_name` - The name of the ASM instance.
	* `is_cluster` - Indicates whether the Oracle Database is part of a cluster.
	* `is_flex_cluster` - Indicates whether the cluster is an Oracle Flex Cluster or not.
	* `is_flex_enabled` - Indicates whether Oracle Flex ASM is enabled or not.
	* `is_selected_for_monitoring` - Indicates whether the DB system component should be provisioned as an Oracle Cloud Infrastructure resource or not.
	* `listener_alias` - The listener alias.
	* `listener_type` - The type of listener.
	* `log_directory` - The destination directory of the listener log file.
	* `memory_size_in_gbs` - The total memory in gigabytes (GB) on the DB node.
	* `network_configurations` - The list of network address configurations of the external cluster.
		* `network_number` - The network number.
		* `network_type` - The network type.
		* `subnet` - The subnet for the network.
	* `node_role` - The role of the cluster node.
	* `ocr_file_location` - The location of the Oracle Cluster Registry (OCR) file.
	* `oracle_home` - The Oracle home location of the listener.
	* `pluggable_databases` - The list of Pluggable Databases.
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
		* `connector` - The connector details used to connect to the external DB system component.
			* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the external DB system connector. 
			* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
			* `connection_info` - The connection details required to connect to an external DB system component.
				* `component_type` - The component type.
				* `connection_credentials` - The credentials used to connect to the ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
					* `credential_name` - The name of the credential information that used to connect to the DB system resource. The name should be in "x.y" format, where the length of "x" has a maximum of 64 characters, and length of "y" has a maximum of 199 characters. The name strings can contain letters, numbers and the underscore character only. Other characters are not valid, except for the "." character that separates the "x" and "y" portions of the name. *IMPORTANT* - The name must be unique within the Oracle Cloud Infrastructure region the credential is being created in. If you specify a name that duplicates the name of another credential within the same Oracle Cloud Infrastructure region, you may overwrite or corrupt the credential that is already using the name.

						For example: inventorydb.abc112233445566778899 
					* `credential_type` - The type of credential used to connect to the ASM instance.
					* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Named Credential where the database password metadata is stored. 
					* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the user password.
					* `role` - The role of the user connecting to the ASM instance.
					* `ssl_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the secret containing the SSL keystore and truststore details.
					* `user_name` - The user name used to connect to the ASM instance.
				* `connection_string` - The Oracle Database connection string. 
					* `host_name` - The host name of the database or the SCAN name in case of a RAC database.
					* `hosts` - The list of host names of the ASM instances.
					* `port` - The port used to connect to the ASM instance.
					* `protocol` - The protocol used to connect to the ASM instance.
					* `service` - The service name of the ASM instance.
				* `database_credential` - The credential to connect to the database to perform tablespace administration tasks.
					* `credential_type` - The type of the credential for tablespace administration tasks.
					* `named_credential_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the named credential where the database password metadata is stored. 
					* `password` - The database user's password encoded using BASE64 scheme.
					* `password_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database password is stored. 
					* `role` - The role of the database user.
					* `username` - The user to connect to the database.
			* `connection_status` - The status of connectivity to the external DB system component.
			* `connector_type` - The type of connector.
			* `display_name` - The user-friendly name for the external connector. The name does not have to be unique.
			* `time_connection_status_last_updated` - The date and time the connectionStatus of the external DB system connector was last updated.
		* `container_database_id` - The unique identifier of the parent Container Database (CDB).
		* `guid` - The unique identifier of the PDB.
	* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Oracle Cloud Infrastructure resource matching the discovered DB system component.
	* `scan_configurations` - The list of Single Client Access Name (SCAN) configurations of the external cluster.
		* `network_number` - The network number from which SCAN VIPs are obtained.
		* `scan_name` - The name of the SCAN listener.
		* `scan_port` - The port number of the SCAN listener.
		* `scan_protocol` - The protocol of the SCAN listener.
	* `status` - The state of the discovered DB system component.
	* `trace_directory` - The destination directory of the listener trace file.
	* `version` - The version of Oracle Clusterware running in the cluster.
	* `vip_configurations` - The list of Virtual IP (VIP) configurations of the external cluster.
		* `address` - The VIP name or IP address.
		* `network_number` - The network number from which VIPs are obtained.
		* `node_name` - The name of the node with the VIP.
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `grid_home` - The directory in which Oracle Grid Infrastructure is installed.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the external DB system discovery.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Oracle Cloud Infrastructure resource matching the discovered DB system.
* `state` - The current lifecycle state of the external DB system discovery resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the external DB system discovery was created.
* `time_updated` - The date and time the external DB system discovery was last updated.

