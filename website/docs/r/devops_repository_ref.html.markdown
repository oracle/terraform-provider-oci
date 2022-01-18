---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_ref"
sidebar_current: "docs-oci-resource-devops-repository_ref"
description: |-
  Provides the Repository Ref resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_repository_ref
This resource provides the Repository Ref resource in Oracle Cloud Infrastructure Devops service.

Creates a new reference or updates an existing one.


## Example Usage

```hcl
resource "oci_devops_repository_ref" "test_repository_ref" {
	#Required
	ref_name = var.repository_ref_ref_name
	ref_type = var.repository_ref_ref_type
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	commit_id = oci_devops_commit.test_commit.id
	object_id = oci_objectstorage_object.test_object.id
}
```

## Argument Reference

The following arguments are supported:

* `commit_id` - (Required when ref_type=BRANCH) (Updatable) Commit ID pointed to by the new branch.
* `object_id` - (Required when ref_type=TAG) (Updatable) SHA-1 hash value of the object pointed to by the tag.
* `ref_name` - (Required) A filter to return only resources that match the given reference name.
* `ref_type` - (Required) (Updatable) The type of reference (Branch or Tag).
* `repository_id` - (Required) Unique repository identifier.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `commit_id` - Commit ID pointed to by the new branch.
* `full_ref_name` - Unique full reference name inside a repository.
* `object_id` - SHA-1 hash value of the object pointed to by the tag.
* `ref_name` - Unique reference name inside a repository.
* `ref_type` - The type of reference (Branch or Tag).
* `repository_id` - The OCID of the repository containing the reference.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Repository Ref
	* `update` - (Defaults to 20 minutes), when updating the Repository Ref
	* `delete` - (Defaults to 20 minutes), when destroying the Repository Ref


## Import

RepositoryRefs can be imported using the `id`, e.g.

```
$ terraform import oci_devops_repository_ref.test_repository_ref "repositories/{repositoryId}/refs/{refName}" 
```

