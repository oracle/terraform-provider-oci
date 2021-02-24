---
subcategory: "Data Science"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_datascience_model_provenance"
sidebar_current: "docs-oci-resource-datascience-model_provenance"
description: |-
  Provides the Model Provenance resource in Oracle Cloud Infrastructure Data Science service
---

# oci_datascience_model_provenance
This resource provides the Model Provenance resource in Oracle Cloud Infrastructure Data Science service.

Creates provenance information for the specified model.

## Example Usage

```hcl
resource "oci_datascience_model_provenance" "test_model_provenance" {
	#Required
	model_id = oci_datascience_model.test_model.id

	#Optional
	git_branch = var.model_provenance_git_branch
	git_commit = var.model_provenance_git_commit
	repository_url = var.model_provenance_repository_url
	script_dir = var.model_provenance_script_dir
	training_script = var.model_provenance_training_script
}
```

## Argument Reference

The following arguments are supported:

* `git_branch` - (Optional) (Updatable) For model reproducibility purposes. Branch of the git repository associated with model training.
* `git_commit` - (Optional) (Updatable) For model reproducibility purposes. Commit ID of the git repository associated with model training.
* `model_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/API/Concepts/identifiers.htm) of the model.
* `repository_url` - (Optional) (Updatable) For model reproducibility purposes. URL of the git repository associated with model training.
* `script_dir` - (Optional) (Updatable) For model reproducibility purposes. Path to model artifacts.
* `training_script` - (Optional) (Updatable) For model reproducibility purposes. Path to the python script or notebook in which the model was trained." 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `git_branch` - For model reproducibility purposes. Branch of the git repository associated with model training.
* `git_commit` - For model reproducibility purposes. Commit ID of the git repository associated with model training.
* `repository_url` - For model reproducibility purposes. URL of the git repository associated with model training.
* `script_dir` - For model reproducibility purposes. Path to model artifacts.
* `training_script` - For model reproducibility purposes. Path to the python script or notebook in which the model was trained." 

## Import

ModelProvenances can be imported using the `id`, e.g.

```
$ terraform import oci_datascience_model_provenance.test_model_provenance "models/{modelId}/provenance" 
```

