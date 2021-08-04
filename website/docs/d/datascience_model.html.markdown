---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model"
sidebar_current: "docs-oci-datasource-datascience-model"
description: |-
  Provides details about a specific Model in Oracle Cloud Infrastructure Data Science service
---

# Data Source: oci_datascience_model
This data source provides details about a specific Model resource in Oracle Cloud Infrastructure Data Science service.

Gets the specified model's information.

## Example Usage

```hcl
data "oci_datascience_model" "test_model" {
	#Required
	model_id = oci_datascience_model.test_model.id
}
```

## Argument Reference

The following arguments are supported:

* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model's compartment.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user who created the model.
* `custom_metadata_list` - An array of custom metadata details for the model.
	* `category` - Category of model metadata which should be null for defined metadata.For custom metadata is should be one of the following values "Performance,Training Profile,Training and Validation Datasets,Training Environment,other".
	* `description` - Description of model metadata
	* `key` - key of the model Metadata. This can be custom key(user defined) as well as Oracle Cloud Infrastructure defined. Example of Oracle defined keys - useCaseType, libraryName, libraryVersion, estimatorClass, hyperParameters. Example of user defined keys - BaseModel
	* `value` - Value of model metadata
* `defined_metadata_list` - An array of defined metadata details for the model.
	* `category` - Category of model metadata which should be null for defined metadata.For custom metadata is should be one of the following values "Performance,Training Profile,Training and Validation Datasets,Training Environment,other".
	* `description` - Description of model metadata
	* `key` - key of the model Metadata. This can be custom key(user defined) as well as Oracle Cloud Infrastructure defined. Example of Oracle defined keys - useCaseType, libraryName, libraryVersion, estimatorClass, hyperParameters. Example of user defined keys - BaseModel
	* `value` - Value of model metadata
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A short description of the model.
* `display_name` - A user-friendly display name for the resource. It does not have to be unique and can be modified. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.
* `input_schema` - Input schema file content in String format
* `output_schema` - Output schema file content in String format
* `project_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the project associated with the model.
* `state` - The state of the model.
* `time_created` - The date and time the resource was created in the timestamp format defined by [RFC3339](https://tools.ietf.org/html/rfc3339). Example: 2019-08-25T21:10:29.41Z 

