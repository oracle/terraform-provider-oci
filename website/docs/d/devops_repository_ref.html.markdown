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

Retrieves a repository's reference by its name with preference for branches over tags if the name is ambiguous. This can be disambiguated by using full names like "heads/<name>" or "tags/<name>".

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

* `ref_name` - (Required) A filter to return only resources that match the given reference name.
* `repository_id` - (Required) Unique repository identifier.


## Attributes Reference

The following attributes are exported:

* `commit_id` - Commit ID pointed to by the new branch.
* `full_ref_name` - Unique full reference name inside a repository.
* `object_id` - SHA-1 hash value of the object pointed to by the tag.
* `ref_name` - Unique reference name inside a repository.
* `ref_type` - The type of reference (Branch or Tag).
* `repository_id` - The OCID of the repository containing the reference.

