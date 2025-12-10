---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_models"
sidebar_current: "docs-oci-datasource-generative_ai-models"
description: |-
  Provides the list of Models in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_models
This data source provides the list of Models in Oracle Cloud Infrastructure Generative AI service.

Lists the models in a specific compartment. Includes pretrained base models and fine-tuned custom models.

## Example Usage

```hcl
data "oci_generative_ai_models" "test_models" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	capability = var.model_capability
	display_name = var.model_display_name
	id = var.model_id
	state = var.model_state
	vendor = var.model_vendor
}
```

## Argument Reference

The following arguments are supported:

* `capability` - (Optional) A filter to return only resources their capability matches the given capability.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The ID of the model.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `vendor` - (Optional) A filter to return only resources that match the entire vendor given.


## Attributes Reference

The following attributes are exported:

* `model_collection` - The list of model_collection.

### Model Reference

The following attributes are exported:


