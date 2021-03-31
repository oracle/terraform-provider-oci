---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_image"
sidebar_current: "docs-oci-datasource-artifacts-container_image"
description: |-
  Provides details about a specific Container Image in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_container_image
This data source provides details about a specific Container Image resource in Oracle Cloud Infrastructure Artifacts service.

Get container image metadata.

## Example Usage

```hcl
data "oci_artifacts_container_image" "test_container_image" {
	#Required
	image_id = var.container_image_id
}
```

## Argument Reference

The following arguments are supported:

* `image_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image.  Example: `ocid1.containerimage.oc1..exampleuniqueID` 


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The compartment OCID to which the container image belongs. Inferred from the container repository.
* `created_by` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the user or principal that created the resource.
* `digest` - The container image digest.
* `display_name` - The repository name and the most recent version associated with the image. If there are no versions associated with the image, then last known version and digest are used instead. If the last known version is unavailable, then 'unknown' is used instead of the version.  Example: `ubuntu:latest` or `ubuntu:latest@sha256:45b23dee08af5e43a7fea6c4cf9c25ccf269ee113168c19722f87876677c5cb2` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container image.  Example: `ocid1.containerimage.oc1..exampleuniqueID` 
* `layers` - Layers of which the image is composed, ordered by the layer digest.
	* `digest` - The sha256 digest of the image layer.
	* `size_in_bytes` - The size of the layer in bytes.
	* `time_created` - An RFC 3339 timestamp indicating when the layer was created.
* `layers_size_in_bytes` - The total size of the container image layers in bytes.
* `manifest_size_in_bytes` - The size of the container image manifest in bytes.
* `pull_count` - Total number of pulls.
* `repository_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container repository.
* `repository_name` - The container repository name.
* `state` - The current state of the container image.
* `time_created` - An RFC 3339 timestamp indicating when the image was created.
* `time_last_pulled` - An RFC 3339 timestamp indicating when the image was last pulled.
* `version` - The most recent version associated with this image.
* `versions` - The versions associated with this image.
	* `created_by` - The OCID of the user or principal that pushed the version.
	* `time_created` - The creation time of the version.
	* `version` - The version name.

