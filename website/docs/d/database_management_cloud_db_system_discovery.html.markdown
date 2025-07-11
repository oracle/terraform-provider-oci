---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_db_system_discovery"
sidebar_current: "docs-oci-datasource-database_management-cloud_db_system_discovery"
description: |-
  Provides details about a specific Cloud Db System Discovery in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_db_system_discovery
This data source provides details about a specific Cloud Db System Discovery resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the cloud DB system discovery resource specified by `cloudDbSystemDiscoveryId`.


## Example Usage

```hcl
data "oci_database_management_cloud_db_system_discovery" "test_cloud_db_system_discovery" {
	#Required
	cloud_db_system_discovery_id = oci_database_management_cloud_db_system_discovery.test_cloud_db_system_discovery.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_db_system_discovery_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system discovery.


## Attributes Reference

The following attributes are exported:

* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the cloud DB system discovery. 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `dbaas_parent_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the parent cloud DB Infrastructure. For VM Dbsystems , it will be the DBSystem Id. For ExaCS and ExaCC,  it will be the cloudVmClusterId and vmClusterId respectively. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `deployment_type` - The deployment type of cloud dbsystem.
* `discovered_components` - The list of DB system components that were found in the DB system discovery.
	* `adr_home_directory` - The directory that stores tracing and logging incidents when Automatic Diagnostic Repository (ADR) is enabled.
	* `asm_instances` - The list of asm instances for the cloud Asm.
		* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the ASM instance.
		* `host_name` - The name of the host on which the ASM instance is running.
		* `instance_name` - The name of the ASM instance.
	* `associated_components` - The list of associated components.
		* `association_type` - The association type.
		* `component_id` - The identifier of the associated component.
		* `component_type` - The type of associated component.
	* `can_enable_all_current_pdbs` - Indicates whether Diagnostics & Management should be enabled for all the current pluggable databases in the container database.
	* `cluster_id` - The unique identifier of the Oracle cluster.
	* `cluster_instances` - The list of cluster instances for the cloud cluster.
		* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the cluster instance.
		* `cluster_id` - The unique identifier of the Oracle cluster.
		* `connector` - The connector details used to connect to the cloud DB system component.
			* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the cloud DB system connector. 
			* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
			* `connection_info` - The connection details required to connect to a cloud DB system component.
				* `component_type` - The component type.
				* `connection_credentials` - The credentials used to connect to the Cloud ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
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
			* `connection_status` - The status of connectivity to the cloud DB system component.
			* `connector_type` - The type of connector.
			* `display_name` - The user-friendly name for the cloud connector. The name does not have to be unique.
			* `time_connection_status_last_updated` - The date and time the connectionStatus of the cloud DB system connector was last updated.
		* `crs_base_directory` - The Oracle base location of Cluster Ready Services (CRS).
		* `host_name` - The name of the host on which the cluster instance is running.
		* `node_role` - The role of the cluster node.
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `component_id` - The identifier of the discovered DB system component.
	* `component_name` - The name of the discovered DB system component.
	* `component_type` - The cloud DB system component type.
	* `connector` - The connector details used to connect to the cloud DB system component.
		* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the cloud DB system connector. 
		* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
		* `connection_info` - The connection details required to connect to a cloud DB system component.
			* `component_type` - The component type.
			* `connection_credentials` - The credentials used to connect to the Cloud ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
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
		* `connection_status` - The status of connectivity to the cloud DB system component.
		* `connector_type` - The type of connector.
		* `display_name` - The user-friendly name for the cloud connector. The name does not have to be unique.
		* `time_connection_status_last_updated` - The date and time the connectionStatus of the cloud DB system connector was last updated.
	* `container_database_id` - The unique identifier of the parent Container Database (CDB).
	* `cpu_core_count` - The number of CPU cores available on the DB node.
	* `crs_base_directory` - The Oracle base location of Cluster Ready Services (CRS).
	* `db_edition` - The Oracle Database edition.
	* `db_id` - The Oracle Database ID.
	* `db_instances` - The list of database instances.
		* `adr_home_directory` - The Automatic Diagnostic Repository (ADR) home directory for the DB instance.
		* `host_name` - The name of the host on which the DB instance is running.
		* `instance_name` - The name of the DB instance.
		* `node_name` - The name of the DB instance node.
		* `oracle_home` - The Oracle home location of the DB instance.
	* `db_node_name` - The name of the DB node.
	* `db_packs` - The database packs licensed for the cloud Oracle Database.
	* `db_role` - The role of the Oracle Database in Oracle Data Guard configuration.
	* `db_type` - The type of Oracle Database. Indicates whether the database is a Container Database, Pluggable Database, or a Non-container Database. 
	* `db_unique_name` - The `DB_UNIQUE_NAME` of the cloud database.
	* `db_version` - The Oracle Database version.
	* `dbaas_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Dbaas Oracle Cloud Infrastructure resource matching the discovered DB system component.
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
	* `host_name` - The name of the host on which the cluster instance is running.
	* `instance_name` - The name of the ASM instance.
	* `is_auto_enable_pluggable_database` - Indicates whether Diagnostics & Management should be enabled automatically for all the pluggable databases in the container database.
	* `is_cluster` - Indicates whether the Oracle Database is part of a cluster.
	* `is_flex_cluster` - Indicates whether the cluster is an Oracle Flex Cluster or not.
	* `is_flex_enabled` - Indicates whether Oracle Flex ASM is enabled or not.
	* `is_selected_for_monitoring` - Indicates whether the DB system component should be provisioned as an Oracle Cloud Infrastructure resource or not.
	* `listener_alias` - The listener alias.
	* `listener_type` - The type of listener.
	* `log_directory` - The destination directory of the listener log file.
	* `memory_size_in_gbs` - The total memory in gigabytes (GB) on the DB node.
	* `network_configurations` - The list of network address configurations of the cloud cluster.
		* `network_number` - The network number.
		* `network_type` - The network type.
		* `subnet` - The subnet for the network.
	* `node_name` - The name of the DB instance node.
	* `node_role` - The role of the cluster node.
	* `ocr_file_location` - The location of the Oracle Cluster Registry (OCR) file.
	* `oracle_home` - The Oracle home location of the DB instance.
	* `pluggable_databases` - The list of Pluggable Databases.
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
		* `connector` - The connector details used to connect to the cloud DB system component.
			* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the management agent used for the cloud DB system connector. 
			* `connection_failure_message` - The error message indicating the reason for connection failure or `null` if the connection was successful. 
			* `connection_info` - The connection details required to connect to a cloud DB system component.
				* `component_type` - The component type.
				* `connection_credentials` - The credentials used to connect to the Cloud ASM instance. Currently only the `DETAILS` type is supported for creating MACS connector credentials. 
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
			* `connection_status` - The status of connectivity to the cloud DB system component.
			* `connector_type` - The type of connector.
			* `display_name` - The user-friendly name for the cloud connector. The name does not have to be unique.
			* `time_connection_status_last_updated` - The date and time the connectionStatus of the cloud DB system connector was last updated.
		* `container_database_id` - The unique identifier of the parent Container Database (CDB).
		* `guid` - The unique identifier of the PDB.
	* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Oracle Cloud Infrastructure resource matching the discovered DB system component.
	* `scan_configurations` - The list of Single Client Access Name (SCAN) configurations of the cloud cluster.
		* `network_number` - The network number from which SCAN VIPs are obtained.
		* `scan_name` - The name of the SCAN listener.
		* `scan_port` - The port number of the SCAN listener.
		* `scan_protocol` - The protocol of the SCAN listener.
	* `status` - The state of the discovered DB system component.
	* `trace_directory` - The destination directory of the listener trace file.
	* `version` - The version of Oracle Clusterware running in the cluster.
	* `vip_configurations` - The list of Virtual IP (VIP) configurations of the cloud cluster.
		* `address` - The VIP name or IP address.
		* `network_number` - The network number from which VIPs are obtained.
		* `node_name` - The name of the node with the VIP.
* `display_name` - The user-friendly name for the DB system. The name does not have to be unique.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `grid_home` - The directory in which Oracle Grid Infrastructure is installed.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the cloud DB system discovery.
* `lifecycle_details` - Additional information about the current lifecycle state.
* `resource_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the existing Oracle Cloud Infrastructure resource matching the discovered DB system.
* `state` - The current lifecycle state of the cloud DB system discovery resource.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the cloud DB system discovery was created.
* `time_updated` - The date and time the cloud DB system discovery was last updated.

