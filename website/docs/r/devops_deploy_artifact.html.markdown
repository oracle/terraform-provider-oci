---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_artifact"
sidebar_current: "docs-oci-resource-devops-deploy_artifact"
description: |-
  Provides the Deploy Artifact resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_deploy_artifact
This resource provides the Deploy Artifact resource in Oracle Cloud Infrastructure Devops service.

Creates a new deployment artifact.

## Example Usage

```hcl
resource "oci_devops_deploy_artifact" "test_deploy_artifact" {
	#Required
	argument_substitution_mode = var.deploy_artifact_argument_substitution_mode
	deploy_artifact_source {
		#Required
		deploy_artifact_source_type = var.deploy_artifact_deploy_artifact_source_deploy_artifact_source_type

		#Optional
		base64encoded_content = var.deploy_artifact_deploy_artifact_source_base64encoded_content
		deploy_artifact_path = var.deploy_artifact_deploy_artifact_source_deploy_artifact_path
		deploy_artifact_version = var.deploy_artifact_deploy_artifact_source_deploy_artifact_version
		image_digest = var.deploy_artifact_deploy_artifact_source_image_digest
		image_uri = var.deploy_artifact_deploy_artifact_source_image_uri
		repository_id = oci_devops_repository.test_repository.id
	}
	deploy_artifact_type = var.deploy_artifact_deploy_artifact_type
	project_id = oci_devops_project.test_project.id

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.deploy_artifact_description
	display_name = var.deploy_artifact_display_name
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `argument_substitution_mode` - (Required) (Updatable) Mode for artifact parameter substitution.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_source` - (Required) (Updatable) Specifies source of an artifact.
	* `base64encoded_content` - (Required when deploy_artifact_source_type=INLINE) (Updatable) base64 Encoded String
	* `deploy_artifact_path` - (Required when deploy_artifact_source_type=GENERIC_ARTIFACT) (Updatable) Specifies the artifact path in the repository.
	* `deploy_artifact_source_type` - (Required) (Updatable) Specifies types of artifact sources.
	* `deploy_artifact_version` - (Required when deploy_artifact_source_type=GENERIC_ARTIFACT) (Updatable) Users can set this as a placeholder value that refers to a pipeline parameter, for example, ${appVersion}.
	* `image_digest` - (Applicable when deploy_artifact_source_type=OCIR) (Updatable) Specifies image digest for the version of the image.
	* `image_uri` - (Required when deploy_artifact_source_type=OCIR) (Updatable) Specifies OCIR Image Path - optionally include tag.
	* `repository_id` - (Required when deploy_artifact_source_type=GENERIC_ARTIFACT) (Updatable) The OCID of a repository
* `deploy_artifact_type` - (Required) (Updatable) Type of the deployment artifact.
* `description` - (Optional) (Updatable) Optional description about the deployment artifact.
* `display_name` - (Optional) (Updatable) Deployment artifact display name. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `project_id` - (Required) The OCID of a project.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `argument_substitution_mode` - Mode for artifact parameter substitution.
* `compartment_id` - The OCID of a compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_source` - Specifies source of an artifact.
	* `base64encoded_content` - base64 Encoded String
	* `deploy_artifact_path` - Specifies the artifact path in the repository.
	* `deploy_artifact_source_type` - Specifies types of artifact sources.
	* `deploy_artifact_version` - Users can set this as a placeholder value that refers to a pipeline parameter, for example, ${appVersion}.
	* `image_digest` - Specifies image digest for the version of the image.
	* `image_uri` - Specifies OCIR Image Path - optionally include tag.
	* `repository_id` - The OCID of a repository
* `deploy_artifact_type` - Type of the deployment artifact.
* `description` - Optional description about the artifact to be deployed.
* `display_name` - Deployment artifact identifier, which can be renamed and is not necessarily unique. Avoid entering confidential information.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
* `id` - Unique identifier that is immutable on creation.
* `lifecycle_details` - A detailed message describing the current state. For example, can be used to provide actionable information for a resource in Failed state.
* `project_id` - The OCID of a project.
* `state` - Current state of the deployment artifact.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"orcl-cloud.free-tier-retained": "true"}`
* `time_created` - Time the deployment artifact was created. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).
* `time_updated` - Time the deployment artifact was updated. Format defined by [RFC3339](https://datatracker.ietf.org/doc/html/rfc3339).

## Import

DeployArtifacts can be imported using the `id`, e.g.

```
$ terraform import oci_devops_deploy_artifact.test_deploy_artifact "id"
```

