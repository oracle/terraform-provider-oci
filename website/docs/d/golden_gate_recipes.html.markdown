---
subcategory: "Golden Gate"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_golden_gate_recipes"
sidebar_current: "docs-oci-datasource-golden_gate-recipes"
description: |-
  Provides the list of Recipes in Oracle Cloud Infrastructure Golden Gate service
---

# Data Source: oci_golden_gate_recipes
This data source provides the list of Recipes in Oracle Cloud Infrastructure Golden Gate service.

Returns an array of Recipe Summary.


## Example Usage

```hcl
data "oci_golden_gate_recipes" "test_recipes" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.recipe_display_name
	recipe_type = var.recipe_recipe_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment that contains the work request. Work requests should be scoped  to the same compartment as the resource the work request affects. If the work request concerns  multiple resources, and those resources are not in the same compartment, it is up to the service team  to pick the primary resource whose compartment should be used. 
* `display_name` - (Optional) A filter to return only the resources that match the entire 'displayName' given. 
* `recipe_type` - (Optional) The pipeline's recipe type. The default value is ZERO_ETL. 


## Attributes Reference

The following attributes are exported:

* `recipe_summary_collection` - The list of recipe_summary_collection.

### Recipe Reference

The following attributes are exported:

* `items` - Array of Recipe Summary 
	* `description` - Metadata about this specific object. 
	* `display_name` - An object's Display Name. 
	* `name` - An object's Display Name. 
	* `recipe_type` - The type of the recipe 
	* `supported_source_technology_types` - Array of supported technology types for this recipe. 
	* `supported_target_technology_types` - Array of supported technology types for this recipe. 

