---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_onboarding"
sidebar_current: "docs-oci-resource-fleet_apps_management-onboarding"
description: |-
  Provides the Onboarding resource in Oracle Cloud Infrastructure Fleet Apps Management service
---

# oci_fleet_apps_management_onboarding
This resource provides the Onboarding resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Onboard a tenant to Fleet Application Management.
The onboarding process lets Fleet Application Management create a few required policies that you need to start using it and its features.


## Example Usage

```hcl
resource "oci_fleet_apps_management_onboarding" "test_onboarding" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	is_cost_tracking_tag_enabled = var.onboarding_is_cost_tracking_tag_enabled
	is_fams_tag_enabled = var.onboarding_is_fams_tag_enabled
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) Tenancy OCID
* `is_cost_tracking_tag_enabled` - (Optional) A value determining if the cost tracking tag is enabled or not. Allow Fleet Application Management to tag resources with cost tracking tag using "Oracle$FAMS-Tags.FAMSManaged" tag. 
* `is_fams_tag_enabled` - (Optional) A value determining if the Fleet Application Management tagging is enabled or not. Allow Fleet Application Management to tag resources with fleet name using "Oracle$FAMS-Tags.FleetName" tag. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `applied_policies` - Summary of the Fleet Application Management Onboard Policy.
	* `id` - The unique id of the resource.
	* `statements` - Policy statements.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `compartment_id` - Tenancy OCID
* `discovery_frequency` - Provide discovery frequency.
* `id` - The unique id of the resource.
* `is_cost_tracking_tag_enabled` - A value determining if the cost tracking tag is enabled or not. Allow Fleet Application Management to tag resources with cost tracking tag using "Oracle$FAMS-Tags.FAMSManaged" tag. 
* `is_fams_tag_enabled` - A value determining if the Fleet Application Management tagging is enabled or not. Allow Fleet Application Management to tag resources with fleet name using "Oracle$FAMS-Tags.FleetName" tag. 
* `items` - List of Fleet Application Management Onboardings.
	* `applied_policies` - Summary of the Fleet Application Management Onboard Policy.
		* `id` - The unique id of the resource.
		* `statements` - Policy statements.
		* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
		* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
		* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
	* `compartment_id` - Tenancy OCID
	* `discovery_frequency` - Provide discovery frequency.
	* `id` - The unique id of the resource.
	* `is_cost_tracking_tag_enabled` - A value determining if the cost tracking tag is enabled or not. Allow Fleet Application Management to tag resources with cost tracking tag using "Oracle$FAMS-Tags.FAMSManaged" tag. 
	* `is_fams_tag_enabled` - A value determining if the Fleet Application Management tagging is enabled or not. Allow Fleet Application Management to tag resources with fleet name using "Oracle$FAMS-Tags.FleetName" tag. 
	* `resource_region` - Associated region
	* `state` - The current state of the Onboarding.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
	* `version` - The version of Fleet Application Management that the tenant is onboarded to.
* `resource_region` - Associated region
* `state` - The current state of the Onboarding.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `version` - The version of Fleet Application Management that the tenant is onboarded to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Onboarding
	* `update` - (Defaults to 20 minutes), when updating the Onboarding
	* `delete` - (Defaults to 20 minutes), when destroying the Onboarding


## Import

Import is not supported for this resource.

