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

Onboard a tenant to Fleet Application Management Service


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
* `is_cost_tracking_tag_enabled` - (Optional) A value determining if cost tracking tag is enabled or not
* `is_fams_tag_enabled` - (Optional) A value determining FAMS tag is enabled or not


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - Tenancy OCID
* `id` - The unique id of the resource.
* `is_cost_tracking_tag_enabled` - A value determining if cost tracking tag is enabled or not
* `is_fams_tag_enabled` - A value determining FAMS tag is enabled or not
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
* `resource_region` - Associated region
* `state` - The current state of the Onboarding.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.
* `version` - Version of FAMS the tenant is onboarded to.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Onboarding
	* `update` - (Defaults to 20 minutes), when updating the Onboarding
	* `delete` - (Defaults to 20 minutes), when destroying the Onboarding


## Import

Import is not supported for this resource.

