---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_endpoint"
sidebar_current: "docs-oci-datasource-generative_ai-endpoint"
description: |-
  Provides details about a specific Endpoint in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_endpoint
This data source provides details about a specific Endpoint resource in Oracle Cloud Infrastructure Generative AI service.

Gets information about an endpoint.

## Example Usage

```hcl
data "oci_generative_ai_endpoint" "test_endpoint" {
	#Required
	endpoint_id = oci_generative_ai_endpoint.test_endpoint.id
}
```

## Argument Reference

The following arguments are supported:

* `endpoint_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint.


## Attributes Reference

The following attributes are exported:

* `content_moderation_config` - The configuration details, whether to add the content moderation feature to the model. Content moderation removes toxic and biased content from responses.
	* `is_enabled` - Whether to enable the content moderation feature.
	* `mode` - Enum for the modes of operation for inference protection.
	* `model_id` - The OCID of the model used for the feature.
* `dedicated_ai_cluster_id` - The OCID of the dedicated AI cluster on which the model will be deployed to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the endpoint.
* `display_name` - A user-friendly name. Does not have to be unique, and it's changeable.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - An OCID that uniquely identifies this endpoint resource.
* `lifecycle_details` - A message describing the current state of the endpoint in more detail that can provide actionable information.
* `model_id` - The OCID of the model that's used to create this endpoint.
* `state` - The current state of the endpoint.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the endpoint was created in the format of an RFC3339 datetime string.
* `time_updated` - The date and time that the endpoint was updated in the format of an RFC3339 datetime string.

