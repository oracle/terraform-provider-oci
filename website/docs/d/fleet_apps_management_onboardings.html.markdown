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
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `onboarding_collection` - The list of onboarding_collection.

### Onboarding Reference

The following attributes are exported:

* `items` - List of FleetAppManagementService Onboardings.
	* `compartment_id` - Tenancy OCID
	* `id` - The unique id of the resource.
	* `is_cost_tracking_tag_enabled` - A value determining if cost tracking tag is enabled or not
	* `is_fams_tag_enabled` - A value determining FAMS tag is enabled or not
	* `resource_region` - Associated region
	* `state` - The current state of the Onboarding.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
	* `version` - Version of FAMS the tenant is onboarded to.

