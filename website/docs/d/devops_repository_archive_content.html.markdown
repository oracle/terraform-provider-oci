---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_archive_content"
sidebar_current: "docs-oci-datasource-devops-repository_archive_content"
description: |-
  Provides details about a specific Repository Archive Content in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_archive_content
This data source provides details about a specific Repository Archive Content resource in Oracle Cloud Infrastructure Devops service.

Return the archived repository information


## Example Usage

```hcl
data "oci_devops_repository_archive_content" "test_repository_archive_content" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	format = var.repository_archive_content_format
	ref_name = var.repository_archive_content_ref_name
}
```

## Argument Reference

The following arguments are supported:

* `format` - (Optional) The archive format query parm for download repo endpoint.
* `ref_name` - (Optional) A filter to return only resources that match the given Ref name.
* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:


