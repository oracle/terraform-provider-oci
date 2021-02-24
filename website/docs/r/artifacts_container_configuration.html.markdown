---
subcategory: "Artifacts"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_artifacts_container_configuration"
sidebar_current: "docs-oci-resource-artifacts-container_configuration"
description: |-
  Provides the Container Configuration resource in Oracle Cloud Infrastructure Artifacts service
---

# oci_artifacts_container_configuration
This resource provides the Container Configuration resource in Oracle Cloud Infrastructure Artifacts service.



## Example Usage

```hcl
resource "oci_artifacts_container_configuration" "test_container_configuration" {
}
```

## Argument Reference

The following arguments are supported:



** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `is_repository_created_on_first_push` - Whether to create a new container repository when a container is pushed to a new repository path. Repositories created in this way belong to the root compartment. 
* `namespace` - The tenancy namespace used in the container repository path.

## Import

ContainerConfiguration can be imported using the `id`, e.g.

```
$ terraform import oci_artifacts_container_configuration.test_container_configuration "container/configuration/compartmentId/{compartmentId}" 
```

