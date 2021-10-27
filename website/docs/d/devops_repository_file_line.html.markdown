---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_file_line"
sidebar_current: "docs-oci-datasource-devops-repository_file_line"
description: |-
  Provides details about a specific Repository File Line in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_file_line
This data source provides details about a specific Repository File Line resource in Oracle Cloud Infrastructure Devops service.

Get lines of a specified file. Supports starting line number and limit.


## Example Usage

```hcl
data "oci_devops_repository_file_line" "test_repository_file_line" {
	#Required
	file_path = var.repository_file_line_file_path
	repository_id = oci_devops_repository.test_repository.id
	revision = var.repository_file_line_revision

	#Optional
	start_line_number = var.repository_file_line_start_line_number
}
```

## Argument Reference

The following arguments are supported:

* `file_path` - (Required) Path to a file within a repository.
* `repository_id` - (Required) unique Repository identifier.
* `revision` - (Required) Retrive file lines from specific revision.
* `start_line_number` - (Optional) Line number from where to start returning file lines. 1 indexed.


## Attributes Reference

The following attributes are exported:

* `lines` - The list of lines in the file
	* `line_content` - The content of the line
	* `line_number` - The line number

