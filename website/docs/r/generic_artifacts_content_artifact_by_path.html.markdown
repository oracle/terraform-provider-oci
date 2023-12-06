---
subcategory: "Generic Artifacts Content"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_generic_artifacts_content_artifact_by_path"
sidebar_current: "docs-oci-resource-generic_artifacts_content-artifact_by_path"
description: |-
  Provides the Artifact By Path resource in Oracle Cloud Infrastructure Generic Artifacts Content service. This resource supports upload/download the content of a generic artifact by specifying the repository id, artifact path, and artifact version
---

# oci_generic_artifacts_content_artifact_by_path
Provides the Artifact By Path resource in Oracle Cloud Infrastructure Generic Artifacts Content service. This resource supports upload/download the content of a generic artifact by specifying the repository id, artifact path, and artifact version

## Note
This resource is not supported to delete generic artifact.
In order to delete generic artifact, you can use `oci_artifacts_generic_artifact`

## Example Usage

```hcl
resource "oci_generic_artifacts_content_artifact_by_path" "test_artifact_by_path" {
  #Required
  artifact_path  = var.artifact_path
  repository_id    = oci_artifacts_repository.test_repository.id
  version = var.version
  source = var.source
}
```

## Argument Reference

The following arguments are supported:

* `artifact_path` - (Required) A user-defined path to describe the location of an artifact. You can use slashes to organize the repository, but slashes do not create a directory structure. An artifact path does not include an artifact version.
* `version` - (Required) A user-defined string to describe the artifact version. Example: `1.1.0` or `1.2-beta-2` 
* `repository_id` - (Required) The [OCID](/iaas/Content/General/Concepts/identifiers.htm) of the repository.
* `source` - (Optional) A path to a file on the local system to be uploaded as the artifact. Cannot be defined if `content` is defined.
* `content` - (Optional) Content to be uploaded as the artifact. Cannot be defined if `source` is defined.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

