---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_ref"
sidebar_current: "docs-oci-datasource-devops-repository_ref"
description: |-
  Provides details about a specific Repository Ref in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_ref
This data source provides details about a specific Repository Ref resource in Oracle Cloud Infrastructure Devops service.

Gets a Repository's Ref by its name with preference for branches over tags if the name is ambiguous. Can be disambiguated by using full names like "heads/<name>" or "tags/<name>".

## Example Usage

```hcl
data "oci_devops_repository_ref" "test_repository_ref" {
	#Required
	ref_name = var.repository_ref_ref_name
	repository_id = oci_devops_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `ref_name` - (Required) A filter to return only resources that match the given Ref name.
* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `commit_id` - Commit ID pointed to by the new branch.
* `full_ref_name` - Unique full ref name inside a repository
* `object_id` - SHA-1 hash value of the object pointed to by the tag.
* `ref_name` - Unique Ref name inside a repository
* `ref_type` - The type of Ref (Branch or Tag)
* `repository_id` - The OCID of the repository containing the ref.

