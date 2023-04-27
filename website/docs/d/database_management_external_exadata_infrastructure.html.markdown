---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_infrastructure"
sidebar_current: "docs-oci-datasource-database_management-external_exadata_infrastructure"
description: |-
  Provides details about a specific External Exadata Infrastructure in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_external_exadata_infrastructure
This data source provides details about a specific External Exadata Infrastructure resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the the Exadata infrastructure specified by externalExadataInfrastructureId. It includes the database systems and storage grid within the
Exadata infrastructure.


## Example Usage

```hcl
data "oci_database_management_external_exadata_infrastructure" "test_external_exadata_infrastructure" {
	#Required
	external_exadata_infrastructure_id = oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure.id
}
```

## Argument Reference

The following arguments are supported:

* `external_exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.


## Attributes Reference

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

