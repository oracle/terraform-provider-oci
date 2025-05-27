---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_provisions"
sidebar_current: "docs-oci-datasource-fleet_apps_management-provisions"
description: |-
  Provides the list of Provisions in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_provisions
This data source provides the list of Provisions in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of all the Provisions in the specified compartment.
The query parameter `compartmentId` is required unless the query parameter `id` or `fleetId` is specified.


## Example Usage

```hcl
data "oci_fleet_apps_management_provisions" "test_provisions" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.provision_display_name
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	id = var.provision_id
	state = var.provision_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) (Updatable) The ID of the compartment in which to list resources. Empty only if the resource OCID query param is not specified. 
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `fleet_id` - (Optional) unique Fleet identifier
* `id` - (Optional) Unique identifier or OCID for listing a single provision by id. Either compartmentId or id must be provided. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `provision_collection` - The list of provision_collection.

### Provision Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `config_catalog_item_display_name` - A display Name of the Catalog Item in the Catalog.
* `config_catalog_item_id` - A [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item to a file with key/value pairs to set up variables for createStack API.
* `config_catalog_item_listing_id` - A listing ID of the Catalog Item in the Catalog.
* `config_catalog_item_listing_version` - A listing version of the Catalog Item in the Catalog.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `deployed_resources` - The deployed resources and their summary
	* `mode` - The mode of the resource. Example: "managed"
	* `resource_instance_list` - Collection of InstanceSummary
		* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which instance is deployed.
		* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.
		* `state` - Instance lifecycle state
	* `resource_name` - The name of the resource
	* `resource_provider` - The name of the Provider
	* `resource_type` - The provider resource type. Must be supported by the [Oracle Cloud Infrastructure provider](https://registry.terraform.io/providers/oracle/oci/latest/docs). Example: oci_core_instance 
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `fleet_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the FamProvision.
* `lifecycle_details` - A message that describes the current state of the FamProvision in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `package_catalog_item_display_name` - A display Name of the Catalog Item in the Catalog.
* `package_catalog_item_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item.
* `package_catalog_item_listing_id` - A listing ID of the Catalog Item in the Catalog.
* `package_catalog_item_listing_version` - A listing version of the Catalog Item in the Catalog.
* `provision_description` - A description of the provision.
* `rms_apply_job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the RMS APPLY Job.
* `stack_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the RMS Stack.
* `state` - The current state of the FamProvision.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `tf_outputs` - Outputs from the Terraform Apply job
	* `is_sensitive` - The indicator if the data for this parameter is sensitive (e.g. should the data be hidden in UI, encrypted if stored, etc.)
	* `output_description` - The output description
	* `output_name` - The output name
	* `output_type` - The output type
	* `output_value` - The output value
* `tf_variable_compartment_id` - An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
* `tf_variable_current_user_id` - An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
* `tf_variable_region_id` - A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
* `tf_variable_tenancy_id` - A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
* `time_created` - The date and time the FamProvision was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
* `time_updated` - The date and time the FamProvision was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

