---
subcategory: "Data Labeling Service"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_labeling_service_dataset"
sidebar_current: "docs-oci-resource-data_labeling_service-dataset"
description: |-
  Provides the Dataset resource in Oracle Cloud Infrastructure Data Labeling Service service
---

# oci_data_labeling_service_dataset
This resource provides the Dataset resource in Oracle Cloud Infrastructure Data Labeling Service service.

Creates a new Dataset.


## Example Usage

```hcl
resource "oci_data_labeling_service_dataset" "test_dataset" {
	#Required
	annotation_format = var.dataset_annotation_format
	compartment_id = var.compartment_id
	dataset_format_details {
		#Required
		format_type = var.dataset_dataset_format_details_format_type
	}
	dataset_source_details {
		#Required
		bucket = var.dataset_dataset_source_details_bucket
		namespace = var.dataset_dataset_source_details_namespace
		source_type = var.dataset_dataset_source_details_source_type

		#Optional
		prefix = var.dataset_dataset_source_details_prefix
	}
	label_set {

		#Required
		items {

			#Required
			name = var.dataset_label_set_items_name
		}
	}

	#Optional
	defined_tags = var.dataset_defined_tags
	description = var.dataset_description
	display_name = var.dataset_display_name
	freeform_tags = var.dataset_freeform_tags
	initial_record_generation_configuration {
	}
	labeling_instructions = var.dataset_labeling_instructions
}
```

## Argument Reference

The following arguments are supported:

* `annotation_format` - (Required) The annotation format name required for labeling records.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment of the resource.
* `dataset_format_details` - (Required) Specifies how to process the data. Supported formats include IMAGE and TEXT.
	* `format_type` - (Required) Format type. DOCUMENT format is for record contents that are PDFs or TIFFs. IMAGE format is for record contents that are JPEGs or PNGs. TEXT format is for record contents that are txt files.
* `dataset_source_details` - (Required) This allows the customer to specify the source of the dataset.
	* `bucket` - (Required) The object storage bucket that contains the dataset data source
	* `namespace` - (Required) Namespace of the bucket that contains the dataset data source
	* `prefix` - (Optional) A common path prefix shared by the objects that make up the dataset. Records will not be generated for objects whose name match exactly with prefix.
	* `source_type` - (Required) Source type.  OBJECT_STORAGE allows the customer to describe where the dataset is in object storage.
* `defined_tags` - (Optional) (Updatable) The defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - (Optional) (Updatable) A user provided description of the dataset
* `display_name` - (Optional) (Updatable) A user-friendly display name for the resource.
* `freeform_tags` - (Optional) (Updatable) A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `initial_record_generation_configuration` - (Optional) Initial Generate Records configuration, generates records from the Dataset's source.
* `label_set` - (Required) An ordered collection of Labels that are unique by name. 
	* `items` - (Required) An ordered collection of Labels that are unique by name.
		* `name` - (Required) An unique name for a label within its dataset.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `annotation_format` - The annotation format name required for labeling records.
* `compartment_id` - The OCID of the compartment of the resource.
* `dataset_format_details` - Specifies how to process the data. Supported formats include DOCUMENT, IMAGE and TEXT.
	* `format_type` - Format type. DOCUMENT format is for record contents that are PDFs or TIFFs. IMAGE format is for record contents that are JPEGs or PNGs. TEXT format is for record contents that are txt files.
* `dataset_source_details` - This allows the customer to specify the source of the dataset.
	* `bucket` - The object storage bucket that contains the dataset data source
	* `namespace` - Namespace of the bucket that contains the dataset data source
	* `prefix` - A common path prefix shared by the objects that make up the dataset. Records will not be generated for objects whose name match exactly with prefix.
	* `source_type` - Source type.  OBJECT_STORAGE allows the customer to describe where the dataset is in object storage.
* `defined_tags` - The defined tags for this resource. Each key is predefined and scoped to a namespace. For example: `{"foo-namespace": {"bar-key": "value"}}` 
* `description` - A user provided description of the dataset
* `display_name` - A user-friendly display name for the resource.
* `freeform_tags` - A simple key-value pair that is applied without any predefined name, type, or scope. It exists for cross-compatibility only. For example: `{"bar-key": "value"}` 
* `id` - The OCID of the Dataset.
* `initial_record_generation_configuration` - Initial Generate Records configuration, generates records from the Dataset's source.
* `label_set` - An ordered collection of Labels that are unique by name. 
	* `items` - An ordered collection of Labels that are unique by name.
		* `name` - An unique name for a label within its dataset.
* `labeling_instructions` - The labeling instructions for human labelers in rich text format
* `lifecycle_details` - A message describing the current state in more detail. For example, it can be used to provide actionable information for a resource in FAILED or NEEDS_ATTENTION state.
* `state` - The state of a dataset. CREATING - The dataset is being created.  It will transition to ACTIVE when it is ready for labeling. ACTIVE   - The dataset is ready for labeling. UPDATING - The dataset is being updated.  It and its related resources may be unavailable for other updates until it returns to ACTIVE. NEEDS_ATTENTION - A dataset updation operation has failed due to validation or other errors and needs attention. DELETING - The dataset and its related resources are being deleted. DELETED  - The dataset has been deleted and is no longer available. FAILED   - The dataset has failed due to validation or other errors. 
* `time_created` - The date and time the resource was created, in the timestamp format defined by RFC3339.
* `time_updated` - The date and time the resource was last updated, in the timestamp format defined by RFC3339.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Dataset
	* `update` - (Defaults to 20 minutes), when updating the Dataset
	* `delete` - (Defaults to 20 minutes), when destroying the Dataset


## Import

Datasets can be imported using the `id`, e.g.

```
$ terraform import oci_data_labeling_service_dataset.test_dataset "id"
```

