---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_zone"
sidebar_current: "docs-oci-resource-cloud_guard-security_zone"
description: |-
  Provides the Security Zone resource in Oracle Cloud Infrastructure Cloud Guard service
---

# oci_cloud_guard_security_zone
This resource provides the Security Zone resource in Oracle Cloud Infrastructure Cloud Guard service.

Creates a security zone for a compartment. A security zone enforces all security zone policies in a given security zone recipe. Any actions that violate a policy are denied. By default, any subcompartments are also in the same security zone.


## Example Usage

```hcl
resource "oci_cloud_guard_security_zone" "test_security_zone" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.security_zone_display_name
	security_zone_recipe_id = oci_cloud_guard_security_zone_recipe.test_security_zone_recipe.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.security_zone_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment for the security zone
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) The security zone's description
* `display_name` - (Required) (Updatable) The security zone's name
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `security_zone_recipe_id` - (Required) (Updatable) The OCID of the recipe (`SecurityRecipe`) for the security zone


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Security Zone
	* `update` - (Defaults to 20 minutes), when updating the Security Zone
	* `delete` - (Defaults to 20 minutes), when destroying the Security Zone


## Import

SecurityZones can be imported using the `id`, e.g.

```
$ terraform import oci_cloud_guard_security_zone.test_security_zone "id"
```

