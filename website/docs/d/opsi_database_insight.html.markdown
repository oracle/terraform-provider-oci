---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_database_insight"
sidebar_current: "docs-oci-datasource-opsi-database_insight"
description: |-
  Provides details about a specific Database Insight in Oracle Cloud Infrastructure Opsi service
---

# Data Source: oci_opsi_database_insight
This data source provides details about a specific Database Insight resource in Oracle Cloud Infrastructure Opsi service.

Gets details of a database insight.

## Example Usage

```hcl
data "oci_opsi_database_insight" "test_database_insight" {
	#Required
	database_insight_id = oci_opsi_database_insight.test_database_insight.id
}
```

## Argument Reference

The following arguments are supported:

* `database_insight_id` - (Required) Unique database insight identifier


## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the database
* `connection_credential_details` - User credential details to connect to the database. 
	* `credential_source_name` - Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
	* `credential_type` - Credential type. 
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
	* `credential_type` - Credential type. 
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

