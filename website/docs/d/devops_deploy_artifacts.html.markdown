---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_artifacts"
sidebar_current: "docs-oci-datasource-devops-deploy_artifacts"
description: |-
  Provides the list of Deploy Artifacts in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_deploy_artifacts
This data source provides the list of Deploy Artifacts in Oracle Cloud Infrastructure Devops service.

Returns a list of deployment artifacts.

## Example Usage

```hcl
data "oci_devops_deploy_artifacts" "test_deploy_artifacts" {

	#Optional
	compartment_id = var.compartment_id
	display_name = var.deploy_artifact_display_name
	id = var.deploy_artifact_id
	project_id = oci_devops_project.test_project.id
	state = var.deploy_artifact_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `id` - (Optional) Unique identifier or OCID for listing a single resource by ID.
* `project_id` - (Optional) unique project identifier
* `state` - (Optional) A filter to return only DeployArtifacts that matches the given lifecycleState.


## Attributes Reference

The following attributes are exported:

* `deploy_artifact_collection` - The list of deploy_artifact_collection.

### DeployArtifact Reference

The following attributes are exported:

* `argument_substitution_mode` - Mode for artifact parameter substitution. Options: `"NONE", "SUBSTITUTE_PLACEHOLDERS"` For Helm Deployments only "NONE" is supported.
* `compartment_id` - The OCID of a compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
* `deploy_artifact_source` - Specifies source of an artifact.
	* `base64encoded_content` - The Helm commands to be executed, base 64 encoded
	* `chart_url` - The URL of an OCIR repository.
	* `deploy_artifact_path` - Specifies the artifact path in the repository.
	* `deploy_artifact_source_type` - Specifies types of artifact sources.
	* `deploy_artifact_version` - Users can set this as a placeholder value that refers to a pipeline parameter, for example, ${appVersion}.
	* `helm_artifact_source_type` - Specifies types of artifact sources.
	* `helm_verification_key_source` - The source of the verification material.
		* `current_public_key` - Current version of Base64 encoding of the public key which is in binary GPG exported format.
		* `previous_public_key` - Previous version of Base64 encoding of the public key which is in binary GPG exported format. This would be used for key rotation scenarios.
		* `vault_secret_id` - The OCID of the Vault Secret containing the verification key versions.
		* `verification_key_source_type` - Specifies type of verification material.
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

