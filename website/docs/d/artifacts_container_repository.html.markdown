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
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - The container repository name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the container repository.  Example: `ocid1.containerrepo.oc1..exampleuniqueID` 
* `image_count` - Total number of images.
* `is_immutable` - Whether the repository is immutable. Images cannot be overwritten in an immutable repository.
* `is_public` - Whether the repository is public. A public repository allows unauthenticated access.
* `is_top_level` - Whether the repository is top-level. A top-level repository reserves its name in namespace as a prefix and allows images to be hosted under child paths of that prefix.
* `layer_count` - Total number of layers.
* `layers_size_in_bytes` - Total storage in bytes consumed by layers.
* `multi_region_peers` - List of peer regions that define the multi-region group. Peer regions must be specified using canonical region names (for example, us-phoenix-1). The peer set must match a supported multi-region endpoint group and MUST include the local OCIR region where the repository is created. 
* `namespace` - The tenancy namespace used in the container repository path.
* `readme` - Container repository readme.
	* `content` - Readme content. Avoid entering confidential information.
	* `format` - Readme format. Supported formats are text/plain and text/markdown.
* `repository_type` - The type of container repository. This determines the repository behavior and required settings.
* `state` - The current state of the container repository.
* `system_tags` - The system tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - An RFC 3339 timestamp indicating when the repository was created.
* `time_last_pushed` - An RFC 3339 timestamp indicating when an image was last pushed to the repository.
* `upstream_secret_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Oracle Cloud Infrastructure Vault secret that contains the upstream registry password as base64 encoded string for a pull-through cache repository. 
* `upstream_url` - The upstream registry URL for a pull-through cache repository.
* `upstream_username` - The upstream registry username for a pull-through cache repository.

