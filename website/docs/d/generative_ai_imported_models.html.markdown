---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_imported_models"
sidebar_current: "docs-oci-datasource-generative_ai-imported_models"
description: |-
  Provides the list of Imported Models in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_imported_models
This data source provides the list of Imported Models in Oracle Cloud Infrastructure Generative AI service.

Lists imported models in a specific compartment.

## Example Usage

```hcl
data "oci_generative_ai_imported_models" "test_imported_models" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	capability = var.imported_model_capability
	display_name = var.imported_model_display_name
	id = var.imported_model_id
	state = var.imported_model_state
	vendor = var.imported_model_vendor
}
```

## Argument Reference

The following arguments are supported:

* `capability` - (Optional) A filter to return only resources their capability matches the given capability.
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The ID of the importedModel.
* `state` - (Optional) A filter to return only resources their lifecycleState matches the given lifecycleState.
* `vendor` - (Optional) A filter to return only resources that match the entire vendor given.


## Attributes Reference

The following attributes are exported:

* `imported_model_collection` - The list of imported_model_collection.

### ImportedModel Reference

The following attributes are exported:


