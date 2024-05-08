---
subcategory: "Cloud Guard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_cloud_guard_security_recipe"
sidebar_current: "docs-oci-datasource-cloud_guard-security_recipe"
description: |-
  Provides details about a specific Security Recipe in Oracle Cloud Infrastructure Cloud Guard service
---

# Data Source: oci_cloud_guard_security_recipe
This data source provides details about a specific Security Recipe resource in Oracle Cloud Infrastructure Cloud Guard service.

Returns a security zone recipe (SecurityRecipe resource) identified by securityRecipeId.

## Example Usage

```hcl
data "oci_cloud_guard_security_recipe" "test_security_recipe" {
	#Required
	security_recipe_id = oci_cloud_guard_security_recipe.test_security_recipe.id
}
```

## Argument Reference

The following arguments are supported:

* `security_recipe_id` - (Required) The unique identifier of the security zone recipe. (`SecurityRecipe`)


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the recipe
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - The recipe's description
* `display_name` - The recipe's display name
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}`

	Avoid entering confidential information. 
* `id` - Unique identifier that can’t be changed after creation
* `lifecycle_details` - A message describing the current state in more detail. For example, this can be used to provide actionable information for a recipe in the `Failed` state.
* `owner` - The owner of the recipe
* `security_policies` - The list of security policy IDs that are included in the recipe
* `state` - The current lifecycle state of the recipe
* `time_created` - The time the recipe was created. An RFC3339 formatted datetime string.
* `time_updated` - The time the recipe was last updated. An RFC3339 formatted datetime string.

