---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_repositories"
sidebar_current: "docs-oci-datasource-artifacts-container_repositories"
description: |-
  Provides the list of Container Repositories in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_container_repositories
This data source provides the list of Container Repositories in Oracle Cloud Infrastructure Artifacts service.

List container repositories in a compartment.

## Example Usage

```hcl
data "oci_artifacts_container_repositories" "test_container_repositories" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	compartment_id_in_subtree = var.container_repository_compartment_id_in_subtree
	display_name = var.container_repository_display_name
	is_public = var.container_repository_is_public
	repository_id = oci_artifacts_repository.test_repository.id
	state = var.container_repository_state
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment.
* `compartment_id_in_subtree` - (Optional) When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are inspected depending on the the setting of `accessLevel`. Default is false. Can only be set to true when calling the API on the tenancy (root compartment). 
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly. 
* `is_public` - (Optional) A filter to return resources that match the isPublic value. 
* `repository_id` - (Optional) A filter to return container images only for the specified container repository OCID. 
* `state` - (Optional) A filter to return only resources that match the given lifecycle state name exactly. 


## Attributes Reference

The following attributes are exported:

* `container_repository_collection` - The list of container_repository_collection.

### ContainerRepository Reference

The following attributes are exported:

* `billable_size_in_gbs` - Total storage size in GBs that will be charged.
* `compartment_id` - The OCID of the compartment in which the container repository exists.
* `created_by` - The id of the user or principal that created the resource.
* `display_name` - The container repository name.
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container repository.  Example: `ocid1.containerrepo.oc1..exampleuniqueID` 
* `image_count` - Total number of images.
* `is_immutable` - Whether the repository is immutable. Images cannot be overwritten in an immutable repository.
* `is_public` - Whether the repository is public. A public repository allows unauthenticated access.
* `layer_count` - Total number of layers.
* `layers_size_in_bytes` - Total storage in bytes consumed by layers.
* `readme` - Container repository readme.
	* `content` - Readme content. Avoid entering confidential information.
	* `format` - Readme format. Supported formats are text/plain and text/markdown.
* `state` - The current state of the container repository.
* `time_created` - An RFC 3339 timestamp indicating when the repository was created.
* `time_last_pushed` - An RFC 3339 timestamp indicating when an image was last pushed to the repository.

