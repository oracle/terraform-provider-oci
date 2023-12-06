---
subcategory: "Ai Language"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_language_model_type"
sidebar_current: "docs-oci-datasource-ai_language-model_type"
description: |-
  Provides details about a specific Model Type in Oracle Cloud Infrastructure Ai Language service
---

# Data Source: oci_ai_language_model_type
This data source provides details about a specific Model Type resource in Oracle Cloud Infrastructure Ai Language service.

Gets model capabilities

## Example Usage

```hcl
data "oci_ai_language_model_type" "test_model_type" {
	#Required
	model_type = var.model_type_model_type
}
```

## Argument Reference

The following arguments are supported:

* `model_type` - (Required) Results like version and model supported info by specifying model type


## Attributes Reference

The following attributes are exported:

* `capabilities` - Model information capabilities related to version
	* `capability` - Model information capabilities related to version
		* `details` - values
* `versions` - Model versions available for this model type

