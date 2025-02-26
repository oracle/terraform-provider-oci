---
subcategory: "Ai Document"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_ai_document_model_type"
sidebar_current: "docs-oci-datasource-ai_document-model_type"
description: |-
  Provides details about a specific Model Type in Oracle Cloud Infrastructure Ai Document service
---

# Data Source: oci_ai_document_model_type
This data source provides details about a specific Model Type resource in Oracle Cloud Infrastructure Ai Document service.

Gets model capabilities

## Example Usage

```hcl
data "oci_ai_document_model_type" "test_model_type" {
	#Required
	model_type = var.model_type_model_type

	#Optional
	compartment_id = var.compartment_id
	model_sub_type = var.model_type_model_sub_type
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `model_sub_type` - (Optional) The sub type based upon model selected.
* `model_type` - (Required) The type of the Document model.


## Attributes Reference

The following attributes are exported:

* `capabilities` - Model information capabilities related to version
	* `capability` - Model information capabilities related to version
		* `details` - values
* `versions` - Model versions available for this model and submodel type

