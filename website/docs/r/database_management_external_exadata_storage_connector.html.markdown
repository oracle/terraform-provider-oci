---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_storage_connector"
sidebar_current: "docs-oci-resource-database_management-external_exadata_storage_connector"
description: |-
  Provides the External Exadata Storage Connector resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_exadata_storage_connector
This resource provides the External Exadata Storage Connector resource in Oracle Cloud Infrastructure Database Management service.

Create the storage server connector after validating the connection information.
Or only validates the connection information for creating the connection to the storage server.
The connector for one storage server is associated with the Exadata infrastructure discovery or existing Exadata infrastructure.


## Example Usage

```hcl
resource "oci_database_management_external_exadata_storage_connector" "test_external_exadata_storage_connector" {
	#Required
	agent_id = oci_cloud_bridge_agent.test_agent.id
	connection_uri = var.external_exadata_storage_connector_connection_uri
	connector_name = var.external_exadata_storage_connector_connector_name
	credential_info {
		#Required
		password = var.external_exadata_storage_connector_credential_info_password
		username = var.external_exadata_storage_connector_credential_info_username

		#Optional
		ssl_trust_store_location = var.external_exadata_storage_connector_credential_info_ssl_trust_store_location
		ssl_trust_store_password = var.external_exadata_storage_connector_credential_info_ssl_trust_store_password
		ssl_trust_store_type = var.external_exadata_storage_connector_credential_info_ssl_trust_store_type
	}
	storage_server_id = oci_database_management_storage_server.test_storage_server.id
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent for the Exadata storage server.
* `connection_uri` - (Required) (Updatable) The unique connection string of the connection. For example, "https://slcm21celadm02.us.oracle.com:443/MS/RESTService/".
* `connector_name` - (Required) (Updatable) The connector name if Oracle Cloud Infrastructure connector is created.
* `credential_info` - (Required) (Updatable) The user credential information.
	* `password` - (Required) (Updatable) The password of the user.
	* `ssl_trust_store_location` - (Optional) (Updatable) The full path of the SSL trust store Location in the agent.
	* `ssl_trust_store_password` - (Optional) (Updatable) The password of the SSL trust store Location in the agent.
	* `ssl_trust_store_type` - (Optional) (Updatable) The SSL trust store type.
	* `username` - (Required) (Updatable) The name of the user.
* `storage_server_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent for the Exadata storage server.
* `connection_uri` - The unique connection string of the connection. For example, "https://slcm21celadm02.us.oracle.com:443/MS/RESTService/".
* `display_name` - The name of the resource. English letters, numbers, "-", "_" and "." only.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata infrastructure system.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
* `internal_id` - The internal ID.
* `lifecycle_details` - The details of the lifecycle state.
* `resource_type` - The type of resource.
* `state` - The current lifecycle state of the database resource.
* `status` - The status of the entity.
* `storage_server_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server.
* `time_created` - The timestamp of the creation.
* `time_updated` - The timestamp of the last update.
* `version` - The version of the resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Exadata Storage Connector
	* `update` - (Defaults to 20 minutes), when updating the External Exadata Storage Connector
	* `delete` - (Defaults to 20 minutes), when destroying the External Exadata Storage Connector


## Import

ExternalExadataStorageConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_exadata_storage_connector.test_external_exadata_storage_connector "id"
```

