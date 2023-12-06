---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_external_exadata_infrastructure"
sidebar_current: "docs-oci-resource-database_management-external_exadata_infrastructure"
description: |-
  Provides the External Exadata Infrastructure resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_external_exadata_infrastructure
This resource provides the External Exadata Infrastructure resource in Oracle Cloud Infrastructure Database Management service.

Creates an Oracle Cloud Infrastructure resource for the Exadata infrastructure and enables the Monitoring service for the Exadata infrastructure.
The following resource/subresources are created:
  Infrastructure
  Storage server connectors
  Storage servers
  Storage grids


## Example Usage

```hcl
resource "oci_database_management_external_exadata_infrastructure" "test_external_exadata_infrastructure" {
	#Required
	compartment_id = var.compartment_id
	db_system_ids = var.external_exadata_infrastructure_db_system_ids
	display_name = var.external_exadata_infrastructure_display_name

	#Optional
	discovery_key = var.external_exadata_infrastructure_discovery_key
	license_model = var.external_exadata_infrastructure_license_model
	storage_server_names = var.external_exadata_infrastructure_storage_server_names
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `db_system_ids` - (Required) (Updatable) The list of DB systems in the Exadata infrastructure.
* `discovery_key` - (Optional) (Updatable) The unique key of the discovery request.
* `display_name` - (Required) (Updatable) The name of the Exadata infrastructure.
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the database management resources. 
* `storage_server_names` - (Optional) (Updatable) The list of all the Exadata storage server names to be included for monitoring purposes. If not specified, all the Exadata storage servers associated with the DB systems are included.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_compartments` - The list of [OCIDs] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartments.
* `database_systems` - A list of DB systems.
	* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	* `internal_id` - The internal ID of the Exadata resource.
	* `license_model` - The Oracle license model that applies to the database management resources. 
	* `lifecycle_details` - The details of the lifecycle state of the Exadata resource.
	* `resource_type` - The type of Exadata resource.
	* `state` - The current lifecycle state of the database resource.
	* `status` - The status of the Exadata resource.
	* `time_created` - The timestamp of the creation of the Exadata resource.
	* `time_updated` - The timestamp of the last update of the Exadata resource.
	* `version` - The version of the Exadata resource.
* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
* `internal_id` - The internal ID of the Exadata resource.
* `license_model` - The Oracle license model that applies to the database management resources. 
* `lifecycle_details` - The details of the lifecycle state of the Exadata resource.
* `rack_size` - The rack size of the Exadata infrastructure.
* `resource_type` - The type of Exadata resource.
* `state` - The current lifecycle state of the database resource.
* `status` - The status of the Exadata resource.
* `storage_grid` - The Exadata storage server grid of the Exadata infrastructure.
	* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
	* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata resource.
	* `internal_id` - The internal ID of the Exadata resource.
	* `lifecycle_details` - The details of the lifecycle state of the Exadata resource.
	* `resource_type` - The type of Exadata resource.
	* `server_count` - The number of Exadata storage servers in the Exadata infrastructure.
	* `state` - The current lifecycle state of the database resource.
	* `status` - The status of the Exadata resource.
	* `time_created` - The timestamp of the creation of the Exadata resource.
	* `time_updated` - The timestamp of the last update of the Exadata resource.
	* `version` - The version of the Exadata resource.
* `time_created` - The timestamp of the creation of the Exadata resource.
* `time_updated` - The timestamp of the last update of the Exadata resource.
* `version` - The version of the Exadata resource.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the External Exadata Infrastructure
	* `update` - (Defaults to 20 minutes), when updating the External Exadata Infrastructure
	* `delete` - (Defaults to 20 minutes), when destroying the External Exadata Infrastructure


## Import

ExternalExadataInfrastructures can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_external_exadata_infrastructure.test_external_exadata_infrastructure "id"
```

