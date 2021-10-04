---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_authors"
sidebar_current: "docs-oci-datasource-devops-repository_authors"
description: |-
  Provides the list of Repository Authors in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_authors
This data source provides the list of Repository Authors in Oracle Cloud Infrastructure Devops service.

Get a list of all the authors


## Example Usage

```hcl
data "oci_devops_repository_authors" "test_repository_authors" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	ref_name = var.repository_author_ref_name
}
```

## Argument Reference

The following arguments are supported:

* `ref_name` - (Optional) A filter to return only resources that match the given Ref name.
* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `repository_author_collection` - The list of repository_author_collection.

### RepositoryAuthor Reference

The following attributes are exported:

* `items` - List of author objects.
	* `author_name` - Author name
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`

