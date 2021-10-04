---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_object"
sidebar_current: "docs-oci-datasource-devops-repository_object"
description: |-
  Provides details about a specific Repository Object in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_object
This data source provides details about a specific Repository Object resource in Oracle Cloud Infrastructure Devops service.

Get blob of specific branch name/commit id and file path


## Example Usage

```hcl
data "oci_devops_repository_object" "test_repository_object" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	file_path = var.repository_object_file_path
	ref_name = var.repository_object_ref_name
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Optional) A filter to return only commits that affect any of the specified paths.
* `ref_name` - (Optional) A filter to return only resources that match the given Ref name.
* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `is_binary` - flag to determine is the object contains binary file content or not.
* `sha` - SHA-1 hash of git object
* `size_in_bytes` - Size in Bytes
* `type` - The type of git object.

