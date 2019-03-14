---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_namespace"
sidebar_current: "docs-oci-datasource-object_storage-namespace"
description: |-
  Provides details about a specific Namespace in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_namespace
This data source provides details about a specific Namespace resource in Oracle Cloud Infrastructure Object Storage service.

Each Oracle Cloud Infrastructure tenant is assigned one unique and uneditable Object Storage namespace. The namespace
is a system-generated string assigned during account creation. For some older tenancies, the namespace string may be
the tenancy name in all lower-case letters. You cannot edit a namespace.

GetNamespace returns the name of the Object Storage namespace for the user making the request.


## Example Usage

```hcl
data "oci_objectstorage_namespace" "test_namespace" {
}
```

## Attributes Reference

The following attributes are exported:

* `namespace` - (Computed) The name of the user's namespace.

