---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_author"
sidebar_current: "docs-oci-datasource-devops-repository_author"
description: |-
  Provides details about a specific Repository Author in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_author
This data source provides details about a specific Repository Author resource in Oracle Cloud Infrastructure Devops service.

Retrieve a list of all the authors.


## Example Usage

```hcl
data "oci_devops_repository_author" "test_repository_author" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	ref_name = var.repository_author_ref_name
}
```

## Argument Reference

The following arguments are supported:

* `ref_name` - (Optional) A filter to return only resources that match the given reference name.
* `repository_id` - (Required) Unique repository identifier.


## Attributes Reference

The following attributes are exported:

* `items` - List of author objects.
	* `author_name` - Author name.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`

