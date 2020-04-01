---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_namespace"
sidebar_current: "docs-oci-datasource-objectstorage-namespace"
description: |-
  Provides details about a specific Namespace in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_namespace
This data source provides details about a specific Namespace resource in Oracle Cloud Infrastructure Object Storage service.

Each Oracle Cloud Infrastructure tenant is assigned one unique and uneditable Object Storage namespace. The namespace
is a system-generated string assigned during account creation. For some older tenancies, the namespace string may be
the tenancy name in all lower-case letters. You cannot edit a namespace.

GetNamespace returns the name of the Object Storage namespace for the user making the request.
If an optional compartmentId query parameter is provided, GetNamespace returns the namespace name of the corresponding
tenancy, provided the user has access to it.


## Example Usage

```hcl
data "oci_objectstorage_namespace" "test_namespace" {

	#Optional
	compartment_id = "${var.compartment_id}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) This is an optional field representing either the tenancy [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) or the compartment [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) within the tenancy whose Object Storage namespace is to be retrieved. 


## Attributes Reference

The following attributes are exported:

* `namespace` - (Computed) The name of the user's namespace.

