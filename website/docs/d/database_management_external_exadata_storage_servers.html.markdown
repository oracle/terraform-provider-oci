---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_storage_servers"
sidebar_current: "docs-oci-datasource-database_management-external_exadata_storage_servers"
description: |-
  Provides the list of External Exadata Storage Servers in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_exadata_storage_servers
This data source provides the list of External Exadata Storage Servers in Oracle Cloud Infrastructure Database Management service.

Lists all the storage servers for the exadata infrastructure or storage grid.


## Example Usage

```hcl
data "oci_database_management_external_exadata_storage_servers" "test_external_exadata_storage_servers" {
	#Required
	compartment_id = var.compartment_id
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id

	#Optional
	display_name = var.external_exadata_storage_server_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) The optional single value query filter parameter on the entity display name.
* `external_exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.


## Attributes Reference

The following attributes are exported:

* `external_exadata_storage_server_collection` - The list of external_exadata_storage_server_collection.

### ExternalExadataStorageServer Reference

The following attributes are exported:

* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `connector` - The connector of the storage server.
	* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
	* `agent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the agent for the Exadata storage server.
	* `connection_uri` - The unique connection string of the connection. For example, "https://slcm21celadm02.us.oracle.com:443/MS/RESTService/".
	* `display_name` - The name of the resource. English letters, numbers, "-", "_" and "." only.
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
* `cpu_count` - CPU count of the storage server
* `display_name` - The name of the resource. English letters, numbers, "-", "_" and "." only.
* `exadata_infrastructure_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata infrastructure system.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
* `internal_id` - The internal ID.
* `ip_address` - The IP address of the storage server.
* `lifecycle_details` - The details of the lifecycle state.
* `make_model` - The make model of the storage server.
* `max_flash_disk_iops` - Maximum flash disk IO operations per second of the storage server
* `max_flash_disk_throughput` - Maximum flash disk IO throughput in MB/s of the storage server
* `max_hard_disk_iops` - Maximum hard disk IO operations per second of the storage server
* `max_hard_disk_throughput` - Maximum hard disk IO throughput in MB/s of the storage server
* `memory_gb` - Storage server memory size in GB
* `resource_type` - The type of resource.
* `state` - The current lifecycle state of the database resource.
* `status` - The status of the entity.
* `storage_grid_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of Exadata storage grid.
* `time_created` - The timestamp of the creation.
* `time_updated` - The timestamp of the last update.
* `version` - The version of the resource.

