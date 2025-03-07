---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_defined_metadata_artifact"
sidebar_current: "docs-oci-resource-datascience-model_defined_metadata_artifact"
description: |-
  Provides the Model Defined Metadata Artifact resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_model_defined_metadata_artifact
This resource provides the Model Defined Metadata Artifact resource in Oracle Cloud Infrastructure Data Science service.

Creates model defined metadata artifact for specified model.

## Example Usage

```hcl
resource "oci_datascience_model_defined_metadata_artifact" "test_model_defined_metadata_artifact" {
	#Required
	model_defined_metadatum_artifact = var.model_defined_metadata_artifact_model_defined_metadatum_artifact
	content_length = var.model_defined_metadata_artifact_content_length
	metadatum_key_name = oci_kms_key.test_key.name
	model_id = oci_datascience_model.test_model.id

	#Optional
	content_disposition = var.model_defined_metadata_artifact_content_disposition
}
```

## Argument Reference

The following arguments are supported:

* `model_defined_metadatum_artifact` - (Required) (Updatable) The model defined metadata artifact to upload.
* `content_disposition` - (Optional) (Updatable) This header allows you to specify a filename during upload. This file name is used to dispose of the file contents while downloading the file. If this optional field is not populated in the request, then the OCID of the model is used for the file name when downloading. Example: `{"Content-Disposition": "attachment" "filename"="model.tar.gz" "Content-Length": "2347" "Content-Type": "application/gzip"}` 
* `content_length` - (Required) (Updatable) The content length of the body.
* `metadatum_key_name` - (Required) The name of the model metadatum in the metadata.
* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the model.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model Defined Metadata Artifact
	* `update` - (Defaults to 20 minutes), when updating the Model Defined Metadata Artifact
	* `delete` - (Defaults to 20 minutes), when destroying the Model Defined Metadata Artifact


## Import

ModelDefinedMetadataArtifacts can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model_defined_metadata_artifact.test_model_defined_metadata_artifact "id"
```

