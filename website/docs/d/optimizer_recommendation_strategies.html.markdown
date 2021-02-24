---
subcategory: "Optimizer"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_optimizer_recommendation_strategies"
sidebar_current: "docs-oci-datasource-optimizer-recommendation_strategies"
description: |-
  Provides the list of Recommendation Strategies in Oracle Cloud Infrastructure Optimizer service
---

# Data Source: oci_optimizer_recommendation_strategies
This data source provides the list of Recommendation Strategies in Oracle Cloud Infrastructure Optimizer service.

Lists the existing strategies.


## Example Usage

```hcl
data "oci_optimizer_recommendation_strategies" "test_recommendation_strategies" {
	#Required
	compartment_id = var.compartment_id
	compartment_id_in_subtree = var.recommendation_strategy_compartment_id_in_subtree

	#Optional
	name = var.recommendation_strategy_name
	recommendation_name = oci_optimizer_recommendation.test_recommendation.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The OCID of the compartment.
* `compartment_id_in_subtree` - (Required) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned depending on the the setting of `accessLevel`.

	Can only be set to true when performing ListCompartments on the tenancy (root compartment). 
* `name` - (Optional) Optional. A filter that returns results that match the name specified.
* `recommendation_name` - (Optional) Optional. A filter that returns results that match the recommendation name specified.


## Attributes Reference

The following attributes are exported:

* `recommendation_strategy_collection` - The list of recommendation_strategy_collection.

### RecommendationStrategy Reference

The following attributes are exported:

* `items` - A collection of recommendation strategy summaries.
	* `name` - The display name of the recommendation.
	* `strategies` - The list of strategies used.
		* `is_default` - Whether this is the default recommendation strategy.
		* `parameters_definition` - The list of strategies for the parameters.
			* `default_value` - A default value used for the strategy parameter.
			* `description` - Text describing the strategy parameter.
			* `is_required` - Whether this parameter is required.
			* `name` - The name of the strategy parameter.
			* `possible_values` - The list of possible values used for these strategy parameters.
			* `type` - The type of strategy parameter.
		* `strategy_name` - The name of the strategy.

