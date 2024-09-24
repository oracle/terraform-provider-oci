---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_onboarding_policies"
sidebar_current: "docs-oci-datasource-fleet_apps_management-onboarding_policies"
description: |-
  Provides the list of Onboarding Policies in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_onboarding_policies
This data source provides the list of Onboarding Policies in Oracle Cloud Infrastructure Fleet Apps Management service.

Returns a list of onboarding policy information for FAMS.


## Example Usage

```hcl
data "oci_fleet_apps_management_onboarding_policies" "test_onboarding_policies" {
}
```

## Argument Reference

The following arguments are supported:



## Attributes Reference

The following attributes are exported:

* `onboarding_policy_collection` - The list of onboarding_policy_collection.

### OnboardingPolicy Reference

The following attributes are exported:

* `items` - List of FleetAppManagementService Onboard policies.
	* `id` - The unique id of the resource.
	* `statements` - Policy statements.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `time_created` - The time this resource was created. An RFC3339 formatted datetime string.
	* `time_updated` - The time this resource was last updated. An RFC3339 formatted datetime string.

