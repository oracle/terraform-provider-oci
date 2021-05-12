---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_images"
sidebar_current: "docs-oci-datasource-artifacts-container_images"
description: |-
  Provides the list of Container Images in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_container_images
This data source provides the list of Container Images in Oracle Cloud Infrastructure Artifacts service.

List container images in a compartment.

## Example Usage

```hcl
data "oci_artifacts_container_images" "test_container_images" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.container_image_compartment_id_in_subtree
	display_name = var.container_image_display_name
	image_id = oci_core_image.test_image.id
	is_versioned = var.container_image_is_versioned
	repository_id = oci_artifacts_repository.test_repository.id
	repository_name = oci_artifacts_repository.test_repository.name
	state = var.container_image_state
	version = var.container_image_version
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are inspected depending on the the setting of `accessLevel`. Default is false. Can only be set to true when calling the API on the tenancy (root compartment). 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `image_id` - (Optional) A filter to return a container image summary only for the specified container image OCID. 
* `is_versioned` - (Optional) A filter to return container images based on whether there are any associated versions. 
* `repository_id` - (Optional) A filter to return container images only for the specified container repository OCID. 
* `repository_name` - (Optional) A filter to return container images or container image signatures that match the repository name.  Example: `foo` or `foo*` 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state name exactly. 
* `version` - (Optional) A filter to return container images that match the version.  Example: `foo` or `foo*` 


## Attributes Reference

The following attributes are exported:

* `container_image_collection` - The list of container_image_collection.

### ContainerImage Reference

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

