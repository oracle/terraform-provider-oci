---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_zone"
sidebar_current: "docs-oci-datasource-cloud_guard-security_zone"
description: |-
  Provides details about a specific Security Zone in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_security_zone
This data source provides details about a specific Security Zone resource in Oracle Cloud Infrastructure Cloud Guard service.

Gets a security zone by its identifier. A security zone is associated with a security zone recipe and enforces all security zone policies in the recipe. Any actions in the zone's compartments that violate a policy are denied.

## Example Usage

```hcl
data "oci_cloud_guard_security_zone" "test_security_zone" {
	#Required
	security_zone_id = oci_cloud_guard_security_zone.test_security_zone.id
}
```

## Argument Reference

The following arguments are supported:

* `security_zone_id` - (Required) The unique identifier of the security zone (`SecurityZone`)


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment for the security zone
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The security zone's description
* `display_name` - The security zone's name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier that is immutable on creation
* `inherited_by_compartments` - List of inherited compartments
* `lifecycle_details` - A message describing the current state in more detail. For example, this can be used to provide actionable information for a zone in the `Failed` state.
* `security_zone_recipe_id` - The OCID of the recipe (`SecurityRecipe`) for the security zone
* `security_zone_target_id` - The OCID of the target associated with the security zone
* `state` - The current state of the security zone
* `time_created` - The time the security zone was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the security zone was last updated. An RFC3339 formatted datetime string.

