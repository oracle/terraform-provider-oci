---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_object_content"
sidebar_current: "docs-oci-datasource-devops-repository_object_content"
description: |-
  Provides details about a specific Repository Object Content in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_object_content
This data source provides details about a specific Repository Object Content resource in Oracle Cloud Infrastructure Devops service.

Retrieve contents of a specified object.


## Example Usage

```hcl
data "oci_devops_repository_object_content" "test_repository_object_content" {
	#Required
	repository_id = oci_devops_repository.test_repository.id
	sha = var.repository_object_content_sha

	#Optional
	file_path = var.repository_object_content_file_path
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Optional) A filter to return only commits that affect any of the specified paths.
* `repository_id` - (Required) Unique repository identifier.
* `sha` - (Required) The SHA of a blob or tree.


## Attributes Reference

The following attributes are exported:


