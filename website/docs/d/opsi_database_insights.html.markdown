---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_database_insights"
sidebar_current: "docs-oci-datasource-opsi-database_insights"
description: |-
 Provides the list of Database Insights in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_database_insights
This data source provides the list of Database Insights in Oracle Cloud Infrastructure Opsi service.

Gets a list of database insights based on the query parameters specified. Either compartmentId or id query parameter must be specified.
When both compartmentId and compartmentIdInSubtree are specified, a list of database insights in that compartment and in all sub-compartments will be returned.


## Example Usage

```hcl
data "oci_opsi_database_insights" "test_database_insights" {

	#Optional
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.database_insight_compartment_id_in_subtree
	database_id = oci_database_database.test_database.id
	database_type = var.database_insight_database_type
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
	exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
	fields = var.database_insight_fields
	id = var.database_insight_id
	opsi_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	state = var.database_insight_state
	status = var.database_insight_status
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) A flag to search all resources within a given compartment and all sub-compartments.
* `database_id` - (Applicable when entity_source=AUTONOMOUS_DATABASE | MACS_MANAGED_CLOUD_DATABASE | MACS_MANAGED_EXTERNAL_DATABASE | MDS_MYSQL_DATABASE_SYSTEM | PE_COMANAGED_DATABASE) Optional list of database [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated DBaaS entity.
* `database_type` - (Optional) Filter by one or more database type. Possible values are ADW-S, ATP-S, ADW-D, ATP-D, EXTERNAL-PDB, EXTERNAL-NONCDB.
* `enterprise_manager_bridge_id` - (Applicable when entity_source=EM_MANAGED_EXTERNAL_DATABASE) Unique Enterprise Manager bridge identifier
* `exadata_insight_id` - (Applicable when entity_source=EM_MANAGED_EXTERNAL_DATABASE) [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of exadata insight resource.
* `fields` - (Optional) Specifies the fields to return in a database summary response. By default all fields are returned if omitted.
* `id` - (Optional) Optional list of database insight resource [OCIDs](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).
* `opsi_private_endpoint_id` - (Applicable when entity_source=AUTONOMOUS_DATABASE | PE_COMANAGED_DATABASE) Unique Operations Insights PrivateEndpoint identifier
* `state` - (Optional) Lifecycle states
* `status` - (Optional) Resource Status


## Attributes Reference

The following attributes are exported:

* `database_insights_collection` - The list of database_insights_collection.

### DatabaseInsight Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the database
* `connection_credential_details` - User credential details to connect to the database.
	* `credential_source_name` - Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
	* `credential_type` - CREDENTIALS_BY_SOURCE is supplied via the External Database Service. CREDENTIALS_BY_VAULT is supplied by secret service to connection PE_COMANAGED_DATABASE and ADB as well. CREDENTIALS_BY_IAM is used db-token to connect only for Autonomous Database. 
	* `password_secret_id` - The secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) mapping to the database credentials.
	* `role` - database user role.
	* `user_name` - database user name.
	* `wallet_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database keystore contents are stored.
* `connection_details` - Connection details to connect to the database. HostName, protocol, and port should be specified.
	* `host_name` - Name of the listener host that will be used to create the connect string to the database.
	* `hosts` - List of hosts and port for private endpoint accessed database resource.
		* `host_ip` - Host IP used for connection requests for Cloud DB resource.
		* `port` - Listener port number used for connection requests for rivate endpoint accessed db resource.
	* `port` - Listener port number used for connection requests.
	* `protocol` - Protocol used for connection requests for private endpoint accssed database resource.
	* `service_name` - Database service name used for connection requests.
	* `connector_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of External Database Connector
* `credential_details` - User credential details to connect to the database.
		* `credential_source_name` - Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
		* `credential_type` - CREDENTIALS_BY_SOURCE is supplied via the External Database Service. CREDENTIALS_BY_VAULT is supplied by secret service to connection PE_COMANAGED_DATABASE and ADB as well. CREDENTIALS_BY_IAM is used db-token to connect only for Autonomous Database. 
	* `password_secret_id` - The secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) mapping to the database credentials.
	* `role` - database user role.
	* `user_name` - database user name.
	* `wallet_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database keystore contents are stored.
* `database_connection_status_details` - A message describing the status of the database connection of this resource. For example, it can be used to provide actionable information about the permission and content validity of the database connection.
* `database_display_name` - Display name of database
* `database_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_name` - Name of database
* `database_resource_type` - Oracle Cloud Infrastructure database resource type
* `database_type` - Ops Insights internal representation of the database type.
* `database_version` - The version of the database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `enterprise_manager_bridge_id` - OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_display_name` - Enterprise Manager Entity Display Name
* `enterprise_manager_entity_identifier` - Enterprise Manager Entity Unique Identifier
* `enterprise_manager_entity_name` - Enterprise Manager Entity Name
* `enterprise_manager_entity_type` - Enterprise Manager Entity Type
* `enterprise_manager_identifier` - Enterprise Manager Unqiue Identifier
* `entity_source` - Source of the database entity.
* `exadata_insight_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`
* `id` - Database insight identifier
* `is_advanced_features_enabled` - Flag is to identify if advanced features for autonomous database is enabled or not
* `is_heat_wave_cluster_attached` - Specifies if MYSQL DB System has heatwave cluster attached.
* `is_highly_available` - Specifies if MYSQL DB System is highly available.
* `lifecycle_details` - A message describing the current state in more detail. For example, can be used to provide actionable information for a resource in Failed state.
* `opsi_private_endpoint_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
* `parent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the VM Cluster or DB System ID, depending on which configuration the resource belongs to.
* `processor_count` - Processor count. This is the OCPU count for Autonomous Database and CPU core count for other database types.
* `root_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata Infrastructure.
* `state` - The current state of the database.
* `status` - Indicates the status of a database insight in Operations Insights
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The time the the database insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the database insight was updated. An RFC3339 formatted datetime string