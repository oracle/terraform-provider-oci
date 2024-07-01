---
subcategory: "Opsi"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_opsi_database_insight"
sidebar_current: "docs-oci-resource-opsi-database_insight"
description: |-
  Provides the Database Insight resource in Oracle Cloud Infrastructure Opsi service
---

# oci_opsi_database_insight
This resource provides the Database Insight resource in Oracle Cloud Infrastructure Opsi service.

Create a Database Insight resource for a Enterprise Manager(EM) managed database in Operations Insights. The database will be enabled in Operations Insights. Database metric collection and analysis will be started. The Database Insight resource for Autonomous Database and Management Agent managed external Database needs to be created by Database service terraform provider. 


## Example Usage

```hcl
resource "oci_opsi_database_insight" "test_database_insight" {
	#Required
	compartment_id = var.compartment_id
	entity_source = var.database_insight_entity_source

	#Optional
	connection_details {

		#Optional
		hosts {

			#Optional
			host_ip = var.database_insight_connection_details_hosts_host_ip
			port = var.database_insight_connection_details_hosts_port
		}
		protocol = var.database_insight_connection_details_protocol
		service_name = oci_core_service.test_service.name
	}
	credential_details {
		#Required
		credential_source_name = var.database_insight_credential_details_credential_source_name
		credential_type = var.database_insight_credential_details_credential_type

		#Optional
		password_secret_id = oci_vault_secret.test_secret.id
		role = var.database_insight_credential_details_role
		user_name = oci_identity_user.test_user.name
		wallet_secret_id = oci_vault_secret.test_secret.id
	}
	database_id = oci_database_database.test_database.id
	database_resource_type = var.database_insight_database_resource_type
	dbm_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	defined_tags = {"foo-namespace.bar-key"= "value"}
	deployment_type = var.database_insight_deployment_type
	enterprise_manager_bridge_id = oci_opsi_enterprise_manager_bridge.test_enterprise_manager_bridge.id
	enterprise_manager_entity_identifier = var.database_insight_enterprise_manager_entity_identifier
	enterprise_manager_identifier = var.database_insight_enterprise_manager_identifier
	exadata_insight_id = oci_opsi_exadata_insight.test_exadata_insight.id
	freeform_tags = {"bar-key"= "value"}
	opsi_private_endpoint_id = oci_dataflow_private_endpoint.test_private_endpoint.id
	service_name = oci_core_service.test_service.name
	system_tags = var.database_insight_system_tags
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) Compartment Identifier of database
* `connection_details` - (Optional) Connection details of the private endpoints.
	* `hosts` - (Required when entity_source=PE_COMANAGED_DATABASE) List of hosts and port for private endpoint accessed database resource.
		* `host_ip` - (Applicable when entity_source=PE_COMANAGED_DATABASE) Host IP used for connection requests for Cloud DB resource.
		* `port` - (Applicable when entity_source=PE_COMANAGED_DATABASE) Listener port number used for connection requests for rivate endpoint accessed db resource.
* `protocol` - (Optional) Protocol used for connection requests for private endpoint accssed database resource.
* `service_name` - (Optional) Database service name used for connection requests.
* `credential_details` - (Required when entity_source=AUTONOMOUS_DATABASE | PE_COMANAGED_DATABASE) User credential details to connect to the database. 
* `credential_source_name` - (Required) Credential source name that had been added in Management Agent wallet. This is supplied in the External Database Service.
* `credential_type` - (Required) Credential type. 
* `password_secret_id` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) The secret [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) mapping to the database credentials.
* `role` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) database user role.
* `user_name` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) database user name.
* `wallet_secret_id` - (Applicable when credential_type=CREDENTIALS_BY_VAULT) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Secret where the database keystore contents are stored. This is used for TCPS support in BM/VM/ExaCS cases.
* `database_id` - (Required when entity_source=AUTONOMOUS_DATABASE | MACS_MANAGED_EXTERNAL_DATABASE | MDS_MYSQL_DATABASE_SYSTEM | PE_COMANAGED_DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the database.
* `database_resource_type` - (Required when entity_source=AUTONOMOUS_DATABASE | MACS_MANAGED_EXTERNAL_DATABASE | MDS_MYSQL_DATABASE_SYSTEM | PE_COMANAGED_DATABASE) Oracle Cloud Infrastructure database resource type
* `dbm_private_endpoint_id` - (Applicable when entity_source=PE_COMANAGED_DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}`
* `deployment_type` - (Required when entity_source=PE_COMANAGED_DATABASE) Database Deployment Type
* `enterprise_manager_bridge_id` - (Required when entity_source=EM_MANAGED_EXTERNAL_DATABASE) OPSI Enterprise Manager Bridge OCID
* `enterprise_manager_entity_identifier` - (Required when entity_source=EM_MANAGED_EXTERNAL_DATABASE) Enterprise Manager Entity Unique Identifier
* `enterprise_manager_identifier` - (Required when entity_source=EM_MANAGED_EXTERNAL_DATABASE) Enterprise Manager Unique Identifier
* `entity_source` - (Required) (Updatable) Source of the database entity.
* `exadata_insight_id` - (Applicable when entity_source=EM_MANAGED_EXTERNAL_DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata insight.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `opsi_private_endpoint_id` - (Applicable when entity_source=AUTONOMOUS_DATABASE | PE_COMANAGED_DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the OPSI private endpoint
* `dbm_private_endpoint_id` - (Applicable when entity_source=PE_COMANAGED_DATABASE) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Database Management private endpoint. This field and opsi_private_endpoint_id are mutually exclusive. If DBM private endpoint ID is provided, a new OPSI private endpoint ID will be created.
* `service_name` - (Required when entity_source=PE_COMANAGED_DATABASE) Database service name used for connection requests.
* `system_tags` - (Applicable when entity_source=PE_COMANAGED_DATABASE) System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `status` - (Optional) (Updatable) Status of the resource. Example: "ENABLED", "DISABLED". Resource can be either enabled or disabled by updating the value of status field to either "ENABLED" or "DISABLED"

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values. The resource destruction here is basically a soft delete. User cannot create resource using the same EM managed bridge OCID. If resource is in enabled state during destruction, the resource will be disabled automatically before performing delete operation.

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Compartment identifier of the database
* `connection_credential_details` - User credential details to connect to the database. This is supplied via the External Database Service.
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
* `time_created` - The time the database insight was first enabled. An RFC3339 formatted datetime string
* `time_updated` - The time the database insight was updated. An RFC3339 formatted datetime string

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Database Insight
	* `update` - (Defaults to 20 minutes), when updating the Database Insight
	* `delete` - (Defaults to 20 minutes), when destroying the Database Insight


## Import

DatabaseInsights can be imported using the `id`, e.g.

```
$ terraform import oci_opsi_database_insight.test_database_insight "id"
```
