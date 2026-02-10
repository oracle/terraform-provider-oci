---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_imported_model"
sidebar_current: "docs-oci-datasource-generative_ai-imported_model"
description: |-
  Provides details about a specific Imported Model in Oracle Cloud Infrastructure Generative AI service
---

# Data Source: oci_generative_ai_imported_model
This data source provides details about a specific Imported Model resource in Oracle Cloud Infrastructure Generative AI service.

Gets information about an imported model.

## Example Usage

```hcl
data "oci_generative_ai_imported_model" "test_imported_model" {
	#Required
	imported_model_id = oci_generative_ai_imported_model.test_imported_model.id
}
```

## Argument Reference

The following arguments are supported:

* `imported_model_id` - (Required) The importedModel OCID


## Attributes Reference

The following attributes are exported:

* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `lifecycle_details` - Additional information about the current state of the imported model, providing more detailed and actionable context.
* `time_created` - The date and time that the imported model was created in the format of an RFC3339 datetime string.

