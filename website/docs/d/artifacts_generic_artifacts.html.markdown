---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_generic_artifacts"
sidebar_current: "docs-oci-datasource-artifacts-generic_artifacts"
description: |-
  Provides the list of Generic Artifacts in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_generic_artifacts
This data source provides the list of Generic Artifacts in Oracle Cloud Infrastructure Artifacts service.

Lists artifacts in the specified repository.

## Example Usage

```hcl
data "oci_artifacts_generic_artifacts" "test_generic_artifacts" {
	#Required
	compartment_id = var.compartment_id
	repository_id = oci_artifacts_repository.test_repository.id

	#Optional
	artifact_path = var.generic_artifact_artifact_path
	display_name = var.generic_artifact_display_name
	id = var.generic_artifact_id
	sha256 = var.generic_artifact_sha256
	state = var.generic_artifact_state
	version = var.generic_artifact_version
}
```

## Argument Reference

The following arguments are supported:

* `artifact_path` - (Optional) Filter results by a prefix for the `artifactPath` and and return artifacts that begin with the specified prefix in their path. 
* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `id` - (Optional) A filter to return the resources for the specified OCID. 
* `repository_id` - (Required) A filter to return the artifacts only for the specified repository OCID. 
* `sha256` - (Optional) Filter results by a specified SHA256 digest for the artifact. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state name exactly. 
* `version` - (Optional) Filter results by a prefix for `version` and return artifacts that that begin with the specified prefix in their version. 


## Attributes Reference

The following attributes are exported:

* `generic_artifact_collection` - The list of generic_artifact_collection.

### GenericArtifact Reference

The following attributes are exported:

* `artifact_path` - A user-defined path to describe the location of an artifact. Slashes do not create a directory structure, but you can use slashes to organize the repository. An artifact path does not include an artifact version.  Example: `project01/my-web-app/artifact-abc` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The artifact name with the format of `<artifact-path>:<artifact-version>`. The artifact name is truncated to a maximum length of 255.  Example: `project01/my-web-app/artifact-abc:1.0.0` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.  Example: `ocid1.genericartifact.oc1..exampleuniqueID` 
* `repository_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.
* `sha256` - The SHA256 digest for the artifact. When you upload an artifact to the repository, a SHA256 digest is calculated and added to the artifact properties.
* `size_in_bytes` - The size of the artifact in bytes.
* `state` - The current state of the artifact.
* `time_created` - An RFC 3339 timestamp indicating when the repository was created.
* `version` - A user-defined string to describe the artifact version.  Example: `1.1.0` or `1.2-beta-2` 

