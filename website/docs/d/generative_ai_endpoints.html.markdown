---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_endpoints"
sidebar_current: "docs-oci-datasource-generative_ai-endpoints"
description: |-
  Provides the list of Endpoints in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_endpoints
This data source provides the list of Endpoints in Oracle Cloud Infrastructure Generative AI service.

Lists the endpoints of a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_endpoints" "test_endpoints" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.endpoint_display_name
	generative_ai_private_endpoint_id = oci_generative_ai_generative_ai_private_endpoint.test_generative_ai_private_endpoint.id
	id = var.endpoint_id
	state = var.endpoint_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `generative_ai_private_endpoint_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the private endpoint.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the endpoint.
* `state` - (Optional) A filter to return only resources that their lifecycle state matches the given lifecycle state.


## Attributes Reference

The following attributes are exported:

* `endpoint_collection` - The list of endpoint_collection.

### Endpoint Reference

The following attributes are exported:

* `content_moderation_config` - The configuration details, whether to add the content moderation feature to the model. Content moderation removes toxic and biased content from responses.
	* `is_enabled` - Whether to enable the content moderation feature.
	* `mode` - Enum for the modes of operation for inference protection.
	* `model_id` - The OCID of the model used for the feature.
* `dedicated_ai_cluster_id` - The OCID of the dedicated AI cluster on which the model will be deployed to.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `lifecycle_details` - A message describing the current state of the endpoint in more detail that can provide actionable information.
* `state` - The current state of the endpoint.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 

