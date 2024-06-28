---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_dedicated_ai_cluster"
sidebar_current: "docs-oci-resource-generative_ai-dedicated_ai_cluster"
description: |-
  Provides the Dedicated Ai Cluster resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_dedicated_ai_cluster
This resource provides the Dedicated Ai Cluster resource in Oracle Cloud Infrastructure Generative AI service.

Creates a dedicated AI cluster.

## Example Usage

```hcl
resource "oci_generative_ai_dedicated_ai_cluster" "test_dedicated_ai_cluster" {
	#Required
	compartment_id = var.compartment_id
	type = var.dedicated_ai_cluster_type
	unit_count = var.dedicated_ai_cluster_unit_count
	unit_shape = var.dedicated_ai_cluster_unit_shape

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.dedicated_ai_cluster_description
	display_name = var.dedicated_ai_cluster_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment OCID to create the dedicated AI cluster in.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) An optional description of the dedicated AI cluster.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `type` - (Required) The dedicated AI cluster type indicating whether this is a fine-tuning/training processor or hosting/inference processor.

	Allowed values are:
	* HOSTING
	* FINE_TUNING 
* `unit_count` - (Required) (Updatable) The number of dedicated units in this AI cluster.
* `unit_shape` - (Required) The shape of dedicated unit in this AI cluster. The underlying hardware configuration is hidden from customers.

	Allowed values are:
	* LARGE_COHERE
	* LARGE_COHERE_V2
	* SMALL_COHERE
	* SMALL_COHERE_V2
	* EMBED_COHERE
	* LLAMA2_70
	* LARGE_GENERIC
	* LARGE_COHERE_V2_2 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `capacity` - The total capacity for a dedicated AI cluster.
	* `capacity_type` - The type of the dedicated AI cluster capacity.
	* `total_endpoint_capacity` - The total number of endpoints that can be hosted on this dedicated AI cluster.
	* `used_endpoint_capacity` - The number of endpoints hosted on this dedicated AI cluster.
* `compartment_id` - The compartment OCID to create the dedicated AI cluster in.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the dedicated AI cluster.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the dedicated AI cluster.
* `lifecycle_details` - A message describing the current state with detail that can provide actionable information.
* `previous_state` - Dedicated AI clusters are compute resources that you can use for fine-tuning custom models or for hosting endpoints for custom models. The clusters are dedicated to your models and not shared with users in other tenancies.

	To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives Oracle Cloud Infrastructure resource access to users. See [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and [Getting Access to Generative AI Resouces](https://docs.cloud.oracle.com/iaas/Content/generative-ai/iam-policies.htm). 
* `state` - The current state of the dedicated AI cluster.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the dedicated AI cluster was created, in the format defined by RFC 3339
* `time_updated` - The date and time the dedicated AI cluster was updated, in the format defined by RFC 3339
* `type` - The dedicated AI cluster type indicating whether this is a fine-tuning/training processor or hosting/inference processor.
* `unit_count` - The number of dedicated units in this AI cluster.
* `unit_shape` - The shape of dedicated unit in this AI cluster. The underlying hardware configuration is hidden from customers.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dedicated Ai Cluster
	* `update` - (Defaults to 20 minutes), when updating the Dedicated Ai Cluster
	* `delete` - (Defaults to 20 minutes), when destroying the Dedicated Ai Cluster


## Import

DedicatedAiClusters can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster "id"
```

