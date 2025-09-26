---
subcategory: "Generative AI"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generative_ai_imported_model"
sidebar_current: "docs-oci-resource-generative_ai-imported_model"
description: |-
  Provides the Imported Model resource in Oracle Cloud Infrastructure Generative AI service
---

# oci_generative_ai_imported_model
This resource provides the Imported Model resource in Oracle Cloud Infrastructure Generative AI service.

Import a model from ModelDataSource.

The header contains an opc-work-request-id, which is the id for the WorkRequest that tracks the importedModel creation progress.


## Example Usage

```hcl
resource "oci_generative_ai_imported_model" "test_imported_model" {
	#Required
	compartment_id = var.compartment_id
	data_source {

		#Optional
		access_token = var.imported_model_data_source_access_token
		branch = var.imported_model_data_source_branch
		bucket = var.imported_model_data_source_bucket
		commit = var.imported_model_data_source_commit
		model_id = oci_generative_ai_model.test_model.id
		namespace = var.imported_model_data_source_namespace
		prefix_name = var.imported_model_data_source_prefix_name
		region = var.imported_model_data_source_region
		source_type = var.imported_model_data_source_source_type
	}

	#Optional
	capabilities = var.imported_model_capabilities
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.imported_model_description
	display_name = var.imported_model_display_name
	freeform_tags = {"Department"= "Finance"}
	vendor = var.imported_model_vendor
	version = var.imported_model_version
}
```

## Argument Reference

The following arguments are supported:

* `capabilities` - (Optional) Specifies the intended use or supported capabilities of the imported model.
* `compartment_id` - (Required) (Updatable) The compartment OCID from which the model is imported.
* `data_source` - (Required) Defines the source location and method used to import the model. Supports importing from Hugging Face,  an Object Storage location, or by referencing an already imported model. 
	* `access_token` - (Applicable when source_type=HUGGING_FACE_MODEL) Hugging Face access token to authenticate requests for restricted models.  This token will be securely stored in Oracle Cloud Infrastructure Vault. 
	* `branch` - (Applicable when source_type=HUGGING_FACE_MODEL) The name of the branch in the Hugging Face repository to import the model from.  If not specified, "main" will be used by default.  If you provide both a branch and a commit hash, the model will be imported from the specified commit. 
	* `bucket` - (Required when source_type=OBJECT_STORAGE_OBJECT) The name of the Object Storage bucket.
	* `commit` - (Applicable when source_type=HUGGING_FACE_MODEL) The commit hash in the Hugging Face repository to import the model from.  If both a branch and a commit are provided, the commit hash will be used. 
	* `model_id` - (Required when source_type=HUGGING_FACE_MODEL) The full model OCID from Hugging Face, typically in the format "org/model-name" (e.g., "meta-llama/Llama-2-7b"). 
	* `namespace` - (Required when source_type=OBJECT_STORAGE_OBJECT) The namespace of the Object Storage where the files are stored.
	* `prefix_name` - (Required when source_type=OBJECT_STORAGE_OBJECT) The prefix path (or folder) within the bucket where files are located.
	* `region` - (Applicable when source_type=OBJECT_STORAGE_OBJECT) The full canonical Oracle Cloud region identifier (e.g., "us-ashburn-1") where the object storage bucket  containing the files resides. 
	* `source_type` - (Optional) Specifies the source type for model import.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) An optional description of the imported model.
* `display_name` - (Optional) (Updatable) A user-friendly name for the imported model.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `vendor` - (Optional) (Updatable) The provider of the imported model.
* `version` - (Optional) (Updatable) The version of the imported model.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `capabilities` - Specifies the intended use or supported capabilities of the imported model.
* `compartment_id` - The compartment OCID from which the model is imported.
* `data_source` - Defines the source location and method used to import the model. Supports importing from Hugging Face,  an Object Storage location, or by referencing an already imported model. 
	* `access_token` - Hugging Face access token to authenticate requests for restricted models.  This token will be securely stored in Oracle Cloud Infrastructure Vault. 
	* `branch` - The name of the branch in the Hugging Face repository to import the model from.  If not specified, "main" will be used by default.  If you provide both a branch and a commit hash, the model will be imported from the specified commit. 
	* `bucket` - The name of the Object Storage bucket.
	* `commit` - The commit hash in the Hugging Face repository to import the model from.  If both a branch and a commit are provided, the commit hash will be used. 
	* `model_id` - The full model OCID from Hugging Face, typically in the format "org/model-name" (e.g., "meta-llama/Llama-2-7b"). 
	* `namespace` - The namespace of the Object Storage where the files are stored.
	* `prefix_name` - The prefix path (or folder) within the bucket where files are located.
	* `region` - The full canonical Oracle Cloud region identifier (e.g., "us-ashburn-1") where the object storage bucket  containing the files resides. 
	* `source_type` - Specifies the source type for model import.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - An optional description of the imported model.
* `display_name` - A user-friendly name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - An OCID that uniquely identifies an imported model.
* `lifecycle_details` - Additional information about the current state of the imported model, providing more detailed and actionable context.
* `previous_state` - Represents a model imported into the system based on an external data source, such as Hugging Face or Object Storage.

	To use any of the API operations, you must be authorized in an IAM policy. If you're not authorized, talk to an administrator who gives Oracle Cloud Infrastructure resource access to users. See [Getting Started with Policies](https://docs.cloud.oracle.com/iaas/Content/Identity/policiesgs/get-started-with-policies.htm) and [Getting Access to Generative AI Resources](https://docs.cloud.oracle.com/iaas/Content/generative-ai/iam-policies.htm). 
* `state` - The lifecycle state of the imported model.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time that the imported model was created in the format of an RFC3339 datetime string.
* `time_updated` - The date and time that the imported model was updated in the format of an RFC3339 datetime string.
* `vendor` - The provider of the imported model.
* `version` - The version of the imported model.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Imported Model
	* `update` - (Defaults to 20 minutes), when updating the Imported Model
	* `delete` - (Defaults to 20 minutes), when destroying the Imported Model


## Import

ImportedModels can be imported using the `id`, e.g.

```
$ terraform import oci_generative_ai_imported_model.test_imported_model "id"
```

