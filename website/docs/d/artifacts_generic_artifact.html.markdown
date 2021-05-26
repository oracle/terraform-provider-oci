---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_generic_artifact"
sidebar_current: "docs-oci-datasource-artifacts-generic_artifact"
description: |-
  Provides details about a specific Generic Artifact in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_generic_artifact
This data source provides details about a specific Generic Artifact resource in Oracle Cloud Infrastructure Artifacts service.

Gets information about an artifact with a specified [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm).

## Example Usage

```hcl
data "oci_artifacts_generic_artifact" "test_generic_artifact" {
	#Required
	artifact_id = oci_artifacts_artifact.test_artifact.id
}
```

## Argument Reference

The following arguments are supported:

* `artifact_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the artifact.  Example: `ocid1.genericartifact.oc1..exampleuniqueID` 


## Attributes Reference

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

