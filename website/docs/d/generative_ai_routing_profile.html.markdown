---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_routing_profile"
sidebar_current: "docs-oci-datasource-generative_ai-routing_profile"
description: |-
  Provides details about a specific Routing Profile in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_routing_profile
This data source provides details about a specific Routing Profile resource in Oracle Cloud Infrastructure Generative AI service.

Gets information about a routing profile.

## Example Usage

```hcl
data "oci_generative_ai_routing_profile" "test_routing_profile" {
	#Required
	routing_profile_id = oci_generative_ai_routing_profile.test_routing_profile.id
}
```

## Argument Reference

The following arguments are supported:

* `routing_profile_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the routing profile.


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
