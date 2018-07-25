---
layout: "oci"
page_title: "OCI: oci_objectstorage_namespace"
sidebar_current: "docs-oci-datasource-object_storage-namespace"
description: |-
  Provides details about a specific Namespace
---

# Data Source: oci_objectstorage_namespace
The Namespace data source provides details about a specific Namespace

Namespaces are unique. Namespaces are either the tenancy name or a random string automatically generated during
account creation. You cannot edit a namespace.


## Example Usage

```hcl
data "oci_objectstorage_namespace" "test_namespace" {
}
```

## Attributes Reference

The following attributes are exported:

* `namespace` - (Computed) The name of the user's namespace.

