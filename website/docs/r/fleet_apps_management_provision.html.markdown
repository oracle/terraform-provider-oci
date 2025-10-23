---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_provision"
sidebar_current: "docs-oci-resource-fleet_apps_management-provision"
description: |-
  Provides the Provision resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_provision
This resource provides the Provision resource in Oracle Cloud Infrastructure Fleet Apps Management service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/fleet-management/latest/Provision

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/fleet_apps_management

Creates a Provision.


## Example Usage

```hcl
resource "oci_fleet_apps_management_provision" "test_provision" {
	#Required
	compartment_id = var.compartment_id
	config_catalog_item_id = oci_fleet_apps_management_catalog_item.test_catalog_item.id
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
	package_catalog_item_id = oci_fleet_apps_management_catalog_item.test_catalog_item.id
	tf_variable_region_id = oci_identity_region.test_region.id
	tf_variable_tenancy_id = oci_identity_tenancy.test_tenancy.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	display_name = var.provision_display_name
	freeform_tags = {"bar-key"= "value"}
	provision_description = var.provision_provision_description
	tf_variable_compartment_id = oci_identity_compartment.test_compartment.id
	tf_variable_current_user_id = oci_identity_user.test_user.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment to create the FamProvision in. 
* `config_catalog_item_id` - (Required) A [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item to a file with key/value pairs to set up variables for createStack API.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable. Avoid entering confidential information.
* `fleet_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Fleet.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `package_catalog_item_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Catalog Item.
* `provision_description` - (Optional) (Updatable) A description of the provision.
* `tf_variable_compartment_id` - (Optional) An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
* `tf_variable_current_user_id` - (Optional) An optional variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
* `tf_variable_region_id` - (Required) A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.
* `tf_variable_tenancy_id` - (Required) A mandatory variable added to a list of RMS variables for createStack API. Overrides the one supplied in configuration file.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Provision
	* `update` - (Defaults to 20 minutes), when updating the Provision
	* `delete` - (Defaults to 20 minutes), when destroying the Provision


## Import

Provisions can be imported using the `id`, e.g.

```
$ terraform import oci_fleet_apps_management_provision.test_provision "id"
```

