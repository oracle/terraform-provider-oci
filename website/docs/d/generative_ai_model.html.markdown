---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_model"
sidebar_current: "docs-oci-datasource-generative_ai-model"
description: |-
  Provides details about a specific Model in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_model
This data source provides details about a specific Model resource in Oracle Cloud Infrastructure Generative AI service.

Gets information about a custom model.

## Example Usage

```hcl
data "oci_generative_ai_model" "test_model" {
	#Required
	model_id = oci_generative_ai_model.test_model.id
}
```

## Argument Reference

The following arguments are supported:

* `model_id` - (Required) The model OCID


## Attributes Reference

The following attributes are exported:

* `capabilities` - Describes what this model can be used for.
* `compartment_id` - The compartment OCID for fine-tuned models. For pretrained models, this value is null.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `id` - An ID that uniquely identifies a pretrained or fine-tuned model.
* `state` - The lifecycle state of the model.
* `time_deprecated` - Corresponds to the time when the custom model and its associated foundation model will be deprecated.
* `type` - The model type indicating whether this is a pretrained/base model or a custom/fine-tuned model.
* `version` - The version of the model.

