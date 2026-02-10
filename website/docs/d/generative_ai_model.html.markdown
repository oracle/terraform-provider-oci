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
* `description` - An optional description of the model.
* `state` - The lifecycle state of the model.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `vendor` - The provider of the base model.
* `version` - The version of the model.


