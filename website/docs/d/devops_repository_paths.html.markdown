---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_paths"
sidebar_current: "docs-oci-datasource-devops-repository_paths"
description: |-
  Provides the list of Repository Paths in Oracle Cloud Infrastructure Devops service
---

# Data Source: oci_devops_repository_paths
This data source provides the list of Repository Paths in Oracle Cloud Infrastructure Devops service.

Retrieves a list of files and directories in a repository.


## Example Usage

```hcl
data "oci_devops_repository_paths" "test_repository_paths" {
	#Required
	repository_id = oci_devops_repository.test_repository.id

	#Optional
	display_name = var.repository_path_display_name
	folder_path = var.repository_path_folder_path
	paths_in_subtree = var.repository_path_paths_in_subtree
	ref = var.repository_path_ref
}
```

## Argument Reference

The following arguments are supported:

* `display_name` - (Optional) A filter to return only resources that match the entire display name given.
* `folder_path` - (Optional) The fully qualified path to the folder whose contents are returned, including the folder name. For example, /examples is a fully-qualified path to a folder named examples that was created off of the root directory (/) of a repository.
* `paths_in_subtree` - (Optional) Flag to determine if files must be retrived recursively. Flag is False by default.
* `ref` - (Optional) The name of branch/tag or commit hash it points to. If names conflict, order of preference is commit > branch > tag. You can disambiguate with "heads/foobar" and "tags/foobar". If left blank repository's default branch will be used. 
* `repository_id` - (Required) Unique repository identifier.


## Attributes Reference

The following attributes are exported:

* `repository_path_collection` - The list of repository_path_collection.

### RepositoryPath Reference

The following attributes are exported:

* `items` - List of objects describing files or directories in a repository.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"foo-namespace.bar-key": "value"}`
	* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only.  See [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"bar-key": "value"}`
	* `name` - Name of file or directory.
	* `path` - Path to file or directory in a repository.
	* `sha` - SHA-1 checksum of blob or tree.
	* `size_in_bytes` - Size of file or directory.
	* `submodule_git_url` - The git URL of the submodule.
	* `type` - File or directory.

