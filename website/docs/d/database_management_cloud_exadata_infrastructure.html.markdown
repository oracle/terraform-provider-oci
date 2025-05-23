---
subcategory: "Database Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_database_management_cloud_exadata_infrastructure"
sidebar_current: "docs-oci-datasource-database_management-cloud_exadata_infrastructure"
description: |-
  Provides details about a specific Cloud Exadata Infrastructure in Oracle Cloud Infrastructure Database Management service
---

# Data Source: oci_database_management_cloud_exadata_infrastructure
This data source provides details about a specific Cloud Exadata Infrastructure resource in Oracle Cloud Infrastructure Database Management service.

Gets the details for the Exadata infrastructure specified by exadataInfrastructureId. It includes the VMCluster and storage grid within the
Exadata infrastructure.


## Example Usage

```hcl
data "oci_database_management_cloud_exadata_infrastructure" "test_cloud_exadata_infrastructure" {
	#Required
	cloud_exadata_infrastructure_id = oci_database_management_cloud_exadata_infrastructure.test_cloud_exadata_infrastructure.id
}
```

## Argument Reference

The following arguments are supported:

* `cloud_exadata_infrastructure_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Exadata infrastructure.


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

