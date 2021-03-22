---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_repository"
sidebar_current: "docs-oci-resource-artifacts-container_repository"
description: |-
  Provides the Container Repository resource in Oracle Cloud Infrastructure Artifacts service
---

# oci_artifacts_container_repository
This resource provides the Container Repository resource in Oracle Cloud Infrastructure Artifacts service.

Create a new empty container repository. Avoid entering confidential information.

## Example Usage

```hcl
resource "oci_artifacts_container_repository" "test_container_repository" {
	#Required
	compartment_id = var.compartment_id
	display_name = var.container_repository_display_name

	#Optional
	is_immutable = var.container_repository_is_immutable
	is_public = var.container_repository_is_public
	readme {
		#Required
		content = var.container_repository_readme_content
		format = var.container_repository_readme_format
	}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to create the resource. 
* `display_name` - (Required) The container repository name.
* `is_immutable` - (Optional) (Updatable) Whether the repository is immutable. Images cannot be overwritten in an immutable repository.
* `is_public` - (Optional) (Updatable) Whether the repository is public. A public repository allows unauthenticated access.
* `readme` - (Optional) (Updatable) Container repository readme.
	* `content` - (Required) (Updatable) Readme content. Avoid entering confidential information.
	* `format` - (Required) (Updatable) Readme format. Supported formats are text/plain and text/markdown.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Container Repository
	* `update` - (Defaults to 20 minutes), when updating the Container Repository
	* `delete` - (Defaults to 20 minutes), when destroying the Container Repository


## Import

ContainerRepositories can be imported using the `id`, e.g.

```
$ terraform import oci_artifacts_container_repository.test_container_repository "container/repositories/{repositoryId}" 
```

