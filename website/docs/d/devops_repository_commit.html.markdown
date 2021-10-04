---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_commit"
sidebar_current: "docs-oci-datasource-devops-repository_commit"
description: |-
  Provides details about a specific Repository Commit in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_commit
This data source provides details about a specific Repository Commit resource in Oracle Cloud Infrastructure Devops service.

Gets a Repository's Commit by commitId

## Example Usage

```hcl
data "oci_devops_repository_commit" "test_repository_commit" {
	#Required
	commit_id = oci_devops_commit.test_commit.id
	repository_id = oci_devops_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `commit_id` - (Required) A filter to return only resources that match the given commit Id.
* `repository_id` - (Required) unique Repository identifier.


## Attributes Reference

The following attributes are exported:

* `author_email` - The email of the author of the repository.
* `author_name` - The name of the author of the repository.
* `commit_id` - Commit hash pointed to by Ref name
* `commit_message` - The commit message.
* `committer_email` - The email of who create the commit.
* `committer_name` - The name of who create the commit.
* `parent_commit_ids` - An array of parent commit ids of created commit.
* `time_created` - The time at which commit was created.
* `tree_id` - Tree information for the specified commit

