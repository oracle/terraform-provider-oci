---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_repository"
sidebar_current: "docs-oci-resource-artifacts-repository"
description: |-
  Provides the Repository resource in Oracle Cloud Infrastructure Artifacts service
---

# oci_artifacts_repository
This resource provides the Repository resource in Oracle Cloud Infrastructure Artifacts service.

Creates a new repository for storing artifacts.

## Example Usage

```hcl
resource "oci_artifacts_repository" "test_repository" {
	#Required
	compartment_id = var.compartment_id
	is_immutable = var.repository_is_immutable
	repository_type = var.repository_repository_type

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.repository_description
	display_name = var.repository_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository's compartment. 
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A short description of the repository. It can be updated later.
* `display_name` - (Optional) (Updatable) A user-friendly display name for the repository. If not present, will be auto-generated. It can be modified later. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_immutable` - (Required) Whether to make the repository immutable. The artifacts of an immutable repository cannot be overwritten.
* `repository_type` - (Required) (Updatable) The repository's supported artifact type.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the repository's compartment.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The repository description.
* `display_name` - The repository name.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the repository.  Example: `ocid1.artifactrepository.oc1..exampleuniqueID` 
* `is_immutable` - Whether the repository is immutable. The artifacts of an immutable repository cannot be overwritten.
* `repository_type` - The repository's supported artifact type.
* `state` - The current state of the repository.
* `time_created` - An RFC 3339 timestamp indicating when the repository was created.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Repository
	* `update` - (Defaults to 20 minutes), when updating the Repository
	* `delete` - (Defaults to 20 minutes), when destroying the Repository


## Import

Repositories can be imported using the `id`, e.g.

```
$ terraform import oci_artifacts_repository.test_repository "id"
```

