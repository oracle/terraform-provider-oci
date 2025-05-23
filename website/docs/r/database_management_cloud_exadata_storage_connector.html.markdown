---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_exadata_storage_connector"
sidebar_current: "docs-oci-resource-database_management-cloud_exadata_storage_connector"
description: |-
  Provides the Cloud Exadata Storage Connector resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_cloud_exadata_storage_connector
This resource provides the Cloud Exadata Storage Connector resource in Oracle Cloud Infrastructure Database Management service.

Creates the Exadata storage server connector after validating the connection information.


## Example Usage

```hcl
resource "oci_database_management_cloud_exadata_storage_connector" "test_cloud_exadata_storage_connector" {
	#Required
	agent_id = oci_cloud_bridge_agent.test_agent.id
	connection_uri = var.cloud_exadata_storage_connector_connection_uri
	credential_info {
		#Required
		password = var.cloud_exadata_storage_connector_credential_info_password
		username = var.cloud_exadata_storage_connector_credential_info_username

		#Optional
		ssl_trust_store_location = var.cloud_exadata_storage_connector_credential_info_ssl_trust_store_location
		ssl_trust_store_password = var.cloud_exadata_storage_connector_credential_info_ssl_trust_store_password
		ssl_trust_store_type = var.cloud_exadata_storage_connector_credential_info_ssl_trust_store_type
	}
	storage_server_id = oci_database_management_storage_server.test_storage_server.id

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.cloud_exadata_storage_connector_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `agent_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent for the Exadata storage server.
* `connection_uri` - (Required) (Updatable) The unique string of the connection. For example, "https://<storage-server-name>/MS/RESTService/".
* `credential_info` - (Required) (Updatable) The user credential information.
	* `password` - (Required) (Updatable) The password of the user.
	* `ssl_trust_store_location` - (Optional) (Updatable) The full path of the SSL truststore location in the agent.
	* `ssl_trust_store_password` - (Optional) (Updatable) The password of the SSL truststore location in the agent.
	* `ssl_trust_store_type` - (Optional) (Updatable) The SSL truststore type.
	* `username` - (Required) (Updatable) The name of the user.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) The name of the Exadata storage server connector.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `storage_server_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent for the Exadata storage server.
* `connection_uri` - The unique string of the connection. For example, "https://<storage-server-name>/MS/RESTService/".
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
* `internal_id` - The internal ID of the Exadata resource.
* `lifecycle_details` - The details of the lifecycle state of the Exadata resource.
* `resource_type` - The type of Exadata resource.
* `state` - The current lifecycle state of the database resource.
* `status` - The status of the Exadata resource.
* `storage_server_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata storage server.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The timestamp of the creation of the Exadata resource.
* `time_updated` - The timestamp of the last update of the Exadata resource.
* `version` - The version of the Exadata resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Exadata Storage Connector
	* `update` - (Defaults to 20 minutes), when updating the Cloud Exadata Storage Connector
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Exadata Storage Connector


## Import

CloudExadataStorageConnectors can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_cloud_exadata_storage_connector.test_cloud_exadata_storage_connector "id"
```

