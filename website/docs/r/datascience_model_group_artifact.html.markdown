---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_group_artifact"
sidebar_current: "docs-oci-resource-datascience-model_group_artifact"
description: |-
  Provides the Model Group Artifact resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_model_group_artifact
This resource provides the Model Group Artifact resource in Oracle Cloud Infrastructure Data Science service.
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/data-science/latest/ModelGroupArtifact

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/datascience

Creates artifact for the Model Group.

## Example Usage

```hcl
resource "oci_datascience_model_group_artifact" "test_model_group_artifact" {
	#Required
	model_group_artifact = var.model_group_artifact_model_group_artifact
	content_length = var.model_group_artifact_content_length
	model_group_id = oci_datascience_model_group.test_model_group.id

	#Optional
	content_disposition = var.model_group_artifact_content_disposition
}
```

## Argument Reference

The following arguments are supported:

* `model_group_artifact` - (Required) The model group artifact to upload.
* `content_disposition` - (Optional) This header allows you to specify a filename during upload. This file name is used to dispose of the file contents while downloading the file. If this optional field is not populated in the request, then the OCID of the model is used for the file name when downloading. Example: `{"Content-Disposition": "attachment" "filename"="model.tar.gz" "Content-Length": "2347" "Content-Type": "application/gzip"}` 
* `content_length` - (Required) The content length of the body.
* `model_group_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the modelGroup.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Model Group Artifact
	* `update` - (Defaults to 20 minutes), when updating the Model Group Artifact
	* `delete` - (Defaults to 20 minutes), when destroying the Model Group Artifact


## Import

ModelGroupArtifacts can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model_group_artifact.test_model_group_artifact "id"
```

