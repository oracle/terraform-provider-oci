---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_refs"
sidebar_current: "docs-oci-datasource-devops-repository_refs"
description: |-
  Provides the list of Repository Refs in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_refs
This data source provides the list of Repository Refs in Oracle Cloud Infrastructure Devops service.

Returns a list of Refs.


## Example Usage

```hcl
data "oci_devops_repository_refs" "test_repository_refs" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	commit_id = oci_devops_commit.test_commit.id
	ref_name = var.repository_ref_ref_name
	ref_type = var.repository_ref_ref_type
}
```

## Argument Reference

The following arguments are supported:

* `commit_id` - (Applicable when ref_type=BRANCH) Commit id in a repository
* `ref_name` - (Optional) A filter to return only resources that match the given Ref name.
* `ref_type` - (Optional) Ref type to distinguish between branch and tag. If it is not specified, return all refs.
* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `repository_ref_collection` - The list of repository_ref_collection.

### RepositoryRef Reference

The following attributes are exported:

* `commit_id` - Commit ID pointed to by the new branch.
* `full_ref_name` - Unique full ref name inside a repository
* `object_id` - SHA-1 hash value of the object pointed to by the tag.
* `ref_name` - Unique Ref name inside a repository
* `ref_type` - The type of Ref (Branch or Tag)
* `repository_id` - The OCID of the repository containing the ref.

