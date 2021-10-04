---
subcategory: "Devops"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_devops_repository_mirror"
sidebar_current: "docs-oci-resource-devops-repository_mirror"
description: |-
  Provides the Repository Mirror resource in Oracle Cloud Infrastructure Devops service
---

# oci_devops_repository_mirror
This resource provides the Repository Mirror resource in Oracle Cloud Infrastructure Devops service.

Synchronize a mirrored repository to the latest version from external providers


## Example Usage

```hcl
resource "oci_devops_repository_mirror" "test_repository_mirror" {
	#Required
	repository_id = oci_devops_repository.test_repository.id
}
```

## Argument Reference

The following arguments are supported:

* `repository_id` - (Required) unique Repository identifier.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Repository Mirror
	* `update` - (Defaults to 20 minutes), when updating the Repository Mirror
	* `delete` - (Defaults to 20 minutes), when destroying the Repository Mirror


## Import

RepositoryMirror can be imported using the `id`, e.g.

```
$ terraform import oci_devops_repository_mirror.test_repository_mirror "id"
```

