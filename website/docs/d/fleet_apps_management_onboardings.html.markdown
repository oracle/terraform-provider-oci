---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_onboardings"
sidebar_current: "docs-oci-datasource-fleet_apps_management-onboardings"
description: |-
  Provides the list of Onboardings in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_onboardings
This data source provides the list of Onboardings in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of onboarding information for the Tenancy.


## Example Usage

```hcl
data "oci_fleet_apps_management_onboardings" "test_onboardings" {

	#Optional
	compartment_id = var.compartment_id
	id = var.onboarding_id
	state = var.onboarding_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `id` - (Optional) unique onboarding identifier
* `state` - (Optional) A filter to return only resources whose lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `onboarding_collection` - The list of onboarding_collection.

### Onboarding Reference

The following attributes are exported:

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

