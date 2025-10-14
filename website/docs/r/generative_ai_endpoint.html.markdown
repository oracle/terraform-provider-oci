---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_endpoint"
sidebar_current: "docs-oci-resource-generative_ai-endpoint"
description: |-
  Provides the Endpoint resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_endpoint
This resource provides the Endpoint resource in Oracle Cloud Infrastructure Generative AI service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/generative-ai/latest/Endpoint

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/generative_ai

Creates an endpoint.

The header contains an opc-work-request-id, which is the id for the WorkRequest that tracks the endpoint creation progress.


## Example Usage

```hcl
resource "oci_generative_ai_endpoint" "test_endpoint" {
	#Required
	compartment_id = var.compartment_id
	dedicated_ai_cluster_id = oci_generative_ai_dedicated_ai_cluster.test_dedicated_ai_cluster.id
	model_id = oci_generative_ai_model.test_model.id

	#Optional
	content_moderation_config {
		#Required
		is_enabled = var.endpoint_content_moderation_config_is_enabled

		#Optional
		mode = var.endpoint_content_moderation_config_mode
		model_id = oci_generative_ai_model.test_model.id
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.endpoint_description
	display_name = var.endpoint_display_name
	freeform_tags = {"Department"= "Finance"}
	generative_ai_private_endpoint_id = oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The compartment OCID to create the endpoint in.
* `content_moderation_config` - (Optional) (Updatable) The configuration details, whether to add the content moderation feature to the model. Content moderation removes toxic and biased content from responses.
	* `is_enabled` - (Required) (Updatable) Whether to enable the content moderation feature.
	* `mode` - (Optional) (Updatable) Enum for the modes of operation for inference protection.
	* `model_id` - (Optional) (Updatable) The OCID of the model used for the feature.
* `dedicated_ai_cluster_id` - (Required) The OCID of the dedicated AI cluster on which a model will be deployed to.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) An optional description of the endpoint.
* `display_name` - (Optional) (Updatable) A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}`
* `generative_ai_private_endpoint_id` - (Optional) (Updatable) The OCID of the Generative AI private endpoint to which this endpoint is attached to.
* `model_id` - (Required) The OCID of the model that's used to create this endpoint.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID to create the endpoint in.
* `content_moderation_config` - The configuration details, whether to add the content moderation feature to the model. Content moderation removes toxic and biased content from responses.
	* `is_enabled` - Whether to enable the content moderation feature.
	* `mode` - Enum for the modes of operation for inference protection.
	* `model_id` - The OCID of the model used for the feature.
* `dedicated_ai_cluster_id` - The OCID of the dedicated AI cluster on which the model will be deployed to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the endpoint.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `generative_ai_private_endpoint_id` - The OCID of the Generative AI private endpoint to which this endpoint is attached to.
* `id` - An OCID that uniquely identifies this endpoint resource.
* `lifecycle_details` - A message describing the current state of the endpoint in more detail that can provide actionable information.
* `model_id` - The OCID of the model that's used to create this endpoint.
* `previous_state` - To host a custom model for inference, create an endpoint for that model on a dedicated AI cluster of type HOSTING. 

	To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives Oracle Cloud Infrastructure resource access to users. See [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and [Getting Access to Generative AI Resouces](https://docs.cloud.oracle.com/iaas/Content/generative-ai/iam-policies.htm). 
* `state` - The current state of the endpoint.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the endpoint was created in the format of an RFC3339 datetime string.
* `time_updated` - The date and time that the endpoint was updated in the format of an RFC3339 datetime string.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Endpoint
	* `update` - (Defaults to 20 minutes), when updating the Endpoint
	* `delete` - (Defaults to 20 minutes), when destroying the Endpoint


## Import

Endpoints can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_endpoint.test_endpoint "id"
```

