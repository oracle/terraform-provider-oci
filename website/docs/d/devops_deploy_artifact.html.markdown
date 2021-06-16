---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_deploy_artifact"
sidebar_current: "docs-oci-datasource-devops-deploy_artifact"
description: |-
  Provides details about a specific Deploy Artifact in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_deploy_artifact
This data source provides details about a specific Deploy Artifact resource in Oracle Cloud Infrastructure Devops service.

Retrieves a deployment artifact by identifier.

## Example Usage

```hcl
data "oci_devops_deploy_artifact" "test_deploy_artifact" {
	#Required
	deploy_artifact_id = oci_devops_deploy_artifact.test_deploy_artifact.id
}
```

## Argument Reference

The following arguments are supported:

* `deploy_artifact_id` - (Required) Unique artifact identifier.


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

