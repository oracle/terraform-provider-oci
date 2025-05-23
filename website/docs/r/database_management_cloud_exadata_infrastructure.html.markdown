---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_exadata_infrastructure"
sidebar_current: "docs-oci-resource-database_management-cloud_exadata_infrastructure"
description: |-
  Provides the Cloud Exadata Infrastructure resource in Oracle Cloud Infrastructure Database Management service
---

# oci_database_management_cloud_exadata_infrastructure
This resource provides the Cloud Exadata Infrastructure resource in Oracle Cloud Infrastructure Database Management service.

Creates an Oracle Cloud Infrastructure resource for the Exadata infrastructure and enables the Monitoring service for the Exadata infrastructure.
The following resource/subresources are created:
  Infrastructure
  Storage server connectors
  Storage servers
  Storage grids


## Example Usage

```hcl
resource "oci_database_management_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
	#Required
	compartment_id = var.compartment_id
	vm_cluster_ids = var.cloud_exadata_infrastructure_vm_cluster_ids

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	discovery_key = var.cloud_exadata_infrastructure_discovery_key
	display_name = var.cloud_exadata_infrastructure_display_name
	freeform_tags = {"Department"= "Finance"}
	license_model = var.cloud_exadata_infrastructure_license_model
	storage_server_names = var.cloud_exadata_infrastructure_storage_server_names
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `discovery_key` - (Optional) (Updatable) The unique key of the discovery request.
* `display_name` - (Optional) (Updatable) The name of the Exadata infrastructure.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `license_model` - (Optional) (Updatable) The Oracle license model that applies to the database management resources. 
* `storage_server_names` - (Optional) (Updatable) The list of all the Exadata storage server names to be included for monitoring purposes. If not specified, all the Exadata storage servers associated with the VM Clusters are included.
* `vm_cluster_ids` - (Required) (Updatable) The list of VM Clusters in the Exadata infrastructure.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `database_compartments` - The list of [OCIDs] (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartments.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `deployment_type` - The infrastructure deployment type.
* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
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
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). System tags can be viewed by users, but can only be created by the system.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The timestamp of the creation of the Exadata resource.
* `time_updated` - The timestamp of the last update of the Exadata resource.
* `version` - The version of the Exadata resource.
* `vm_clusters` - The list of VM Clusters in the Exadata infrastructure.
	* `additional_details` - The additional details of the resource defined in `{"key": "value"}` format. Example: `{"bar-key": "value"}` 
	* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
	* `deployment_type` - The infrastructure deployment type.
	* `display_name` - The name of the Exadata resource. English letters, numbers, "-", "_" and "." only.
	* `home_directory` - The Oracle home directory.
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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Cloud Exadata Infrastructure
	* `update` - (Defaults to 20 minutes), when updating the Cloud Exadata Infrastructure
	* `delete` - (Defaults to 20 minutes), when destroying the Cloud Exadata Infrastructure


## Import

CloudExadataInfrastructures can be imported using the `id`, e.g.

```
$ terraform import oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure "id"
```

