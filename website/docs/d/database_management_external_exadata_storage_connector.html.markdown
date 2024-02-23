---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_storage_connector"
sidebar_current: "docs-oci-datasource-database_management-external_exadata_storage_connector"
description: |-
  Provides details about a specific External Exadata Storage Connector in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_exadata_storage_connector
This data source provides details about a specific External Exadata Storage Connector resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the Exadata storage server connector specified by exadataStorageConnectorId.


## Example Usage

```hcl
data "oci_database_management_external_exadata_storage_connector" "test_external_exadata_storage_connector" {
	#Required
	external_exadata_storage_connector_id = oci_database_management_external_exadata_storage_connector.test_external_exadata_storage_connector.id
}
```

## Argument Reference

The following arguments are supported:

* `external_exadata_storage_connector_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the connector to the Exadata storage server.


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
* `time_created` - The timestamp of the creation of the Exadata resource.
* `time_updated` - The timestamp of the last update of the Exadata resource.
* `version` - The version of the Exadata resource.

