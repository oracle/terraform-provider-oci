---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_routing_profile"
sidebar_current: "docs-oci-resource-generative_ai-routing_profile"
description: |-
  Provides the Routing Profile resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_routing_profile
This resource provides the Routing Profile resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/RoutingProfile

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates a routing profile.

## Example Usage

```hcl
resource "oci_generative_ai_routing_profile" "test_routing_profile" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.routing_profile_display_name

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.routing_profile_description
	freeform_tags = {"Department"= "Finance"}
	model_routing_policy {

		#Optional
		allowed_models = var.routing_profile_model_routing_policy_allowed_models
	}
	region_routing_policy {

		#Optional
		allowed_regions = var.routing_profile_region_routing_policy_allowed_regions
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The OCID of the compartment in which to create the routing profile.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - (Optional) (Updatable) An optional description of the routing profile.
* `display_name` - (Required) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `model_routing_policy` - (Optional) (Updatable) The model routing policy for the routing profile. Routing candidates are selected from the allowed models list.
	* `allowed_models` - (Optional) (Updatable) The ordered list of model names the routing candidate is selected from (for example, `meta.llama-3-70b-instruct`).

		The order of entries is preserved. Duplicate entries are not allowed.
* `region_routing_policy` - (Optional) (Updatable) The cross-region routing policy for the routing profile. If this policy is not specified, only the local region is allowed.
	* `allowed_regions` - (Optional) (Updatable) The ordered list of regions that routing is allowed to reach. Each entry is a region identifier (for example, `us-chicago-1`).

		The order of entries is preserved. Duplicate entries are not allowed.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment in which the routing profile lives.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
* `description` - An optional description of the routing profile.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `id` - An OCID that uniquely identifies this routing profile resource.
* `lifecycle_details` - A message describing the current state of the routing profile in more detail that can provide actionable information.
* `model_routing_policy` - The model routing policy for the routing profile. Routing candidates are selected from the allowed models list.
	* `allowed_models` - The ordered list of model names the routing candidate is selected from (for example, `meta.llama-3-70b-instruct`).

		The order of entries is preserved. Duplicate entries are not allowed.
* `previous_state` - A bounded snapshot of the previous state of a routing profile. This object intentionally omits `previousState` to avoid recursive nesting.
	* `compartment_id` - The OCID of the compartment in which the routing profile lives.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}`
	* `description` - An optional description of the routing profile.
	* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
	* `id` - An OCID that uniquely identifies this routing profile resource.
	* `lifecycle_details` - A message describing the current state of the routing profile in more detail that can provide actionable information.
	* `model_routing_policy` - The model routing policy for the routing profile. Routing candidates are selected from the allowed models list.
		* `allowed_models` - The ordered list of model names the routing candidate is selected from (for example, `meta.llama-3-70b-instruct`).

			The order of entries is preserved. Duplicate entries are not allowed.
	* `region_routing_policy` - The cross-region routing policy for the routing profile. If this policy is not specified, only the local region is allowed.
		* `allowed_regions` - The ordered list of regions that routing is allowed to reach. Each entry is a region identifier (for example, `us-chicago-1`).

			The order of entries is preserved. Duplicate entries are not allowed.
	* `state` - The current state of the routing profile.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
	* `time_created` - The date and time that the routing profile was created in the format of an RFC3339 datetime string.
	* `time_updated` - The date and time that the routing profile was updated in the format of an RFC3339 datetime string.
* `region_routing_policy` - The cross-region routing policy for the routing profile. If this policy is not specified, only the local region is allowed.
	* `allowed_regions` - The ordered list of regions that routing is allowed to reach. Each entry is a region identifier (for example, `us-chicago-1`).

		The order of entries is preserved. Duplicate entries are not allowed.
* `state` - The current state of the routing profile.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - The date and time that the routing profile was created in the format of an RFC3339 datetime string.
* `time_updated` - The date and time that the routing profile was updated in the format of an RFC3339 datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Routing Profile
	* `update` - (Defaults to 20 minutes), when updating the Routing Profile
	* `delete` - (Defaults to 20 minutes), when destroying the Routing Profile


## Import

RoutingProfiles can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_routing_profile.test_routing_profile "id"
```
