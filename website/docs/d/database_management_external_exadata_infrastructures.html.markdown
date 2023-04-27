---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_infrastructures"
sidebar_current: "docs-oci-datasource-database_management-external_exadata_infrastructures"
description: |-
  Provides the list of External Exadata Infrastructures in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_exadata_infrastructures
This data source provides the list of External Exadata Infrastructures in Oracle Cloud Infrastructure Database Management service.

Lists the Exadata infrastructures for a specific compartment.


## Example Usage

```hcl
data "oci_database_management_external_exadata_infrastructures" "test_external_exadata_infrastructures" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.external_exadata_infrastructure_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) The optional single value query filter parameter on the entity display name.


## Attributes Reference

The following attributes are exported:

* `external_exadata_infrastructure_collection` - The list of external_exadata_infrastructure_collection.

### ExternalExadataInfrastructure Reference

The following attributes are exported:

* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment.
* `database_compartments` - The list of [OCIDs] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartments
* `database_systems` - A list of database systems.
	* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of compartment.
	* `display_name` - The name of the resource. English letters, numbers, "-", "_" and "." only.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	* `internal_id` - The internal ID.
	* `license_model` - The Oracle license model that applies to the database management resources. 
	* `lifecycle_details` - The details of the lifecycle state.
	* `resource_type` - The type of resource.
	* `state` - The current lifecycle state of the database resource.
	* `status` - The status of the entity.
	* `time_created` - The timestamp of the creation.
	* `time_updated` - The timestamp of the last update.
	* `version` - The version of the resource.
* `display_name` - The name of the resource. English letters, numbers, "-", "_" and "." only.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
* `internal_id` - The internal ID.
* `license_model` - The Oracle license model that applies to the database management resources. 
* `lifecycle_details` - The details of the lifecycle state.
* `rack_size` - The rack size of the Exadata infrastructure.
* `resource_type` - The type of resource.
* `state` - The current lifecycle state of the database resource.
* `status` - The status of the entity.
* `storage_grid` - The storage server grid of the Exadata infrastructure.
	* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
	* `display_name` - The name of the resource. English letters, numbers, "-", "_" and "." only.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	* `internal_id` - The internal ID.
	* `lifecycle_details` - The details of the lifecycle state.
	* `resource_type` - The type of resource.
	* `server_count` - The number of the storage servers in the Exadata infrastructure.
	* `state` - The current lifecycle state of the database resource.
	* `status` - The status of the entity.
	* `time_created` - The timestamp of the creation.
	* `time_updated` - The timestamp of the last update.
	* `version` - The version of the resource.
* `time_created` - The timestamp of the creation.
* `time_updated` - The timestamp of the last update.
* `version` - The version of the resource.

