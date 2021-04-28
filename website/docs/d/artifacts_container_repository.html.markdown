---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_repository"
sidebar_current: "docs-oci-datasource-artifacts-container_repository"
description: |-
  Provides details about a specific Container Repository in Oracle Cloud Infrastructure Artifacts service
---

# Data Source: oci_artifacts_container_repository
This data source provides details about a specific Container Repository resource in Oracle Cloud Infrastructure Artifacts service.

Get container repository.

## Example Usage

```hcl
data "oci_artifacts_container_repository" "test_container_repository" {
	#Required
	repository_id = oci_artifacts_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `repository_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container repository.  Example: `ocid1.containerrepo.oc1..exampleuniqueID` 


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

