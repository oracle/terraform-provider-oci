---
subcategory: "Object Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_objectstorage_namespace_metadata"
sidebar_current: "docs-oci-datasource-objectstorage-namespace_metadata"
description: |-
  Provides details about namespace metadata in Oracle Cloud Infrastructure Object Storage service
---

# Data Source: oci_objectstorage_namespace_metadata

This data source provides the default compartment assignments for the Amazon S3 Compatibility API and Swift API in an Object Storage namespace.

## Example Usage

```hcl
data "oci_objectstorage_namespace" "test_namespace" {
  compartment_id = var.compartment_id
}

data "oci_objectstorage_namespace_metadata" "test_namespace_metadata" {
  namespace = data.oci_objectstorage_namespace.test_namespace.namespace
}

output "default_s3_compartment_id" {
  value = data.oci_objectstorage_namespace_metadata.test_namespace_metadata.default_s3compartment_id
}

output "default_swift_compartment_id" {
  value = data.oci_objectstorage_namespace_metadata.test_namespace_metadata.default_swift_compartment_id
}
```

## Argument Reference

The following argument is supported:

* `namespace` - (Required) The Object Storage namespace used for the request.

## Attributes Reference

The following attributes are exported:

* `default_s3compartment_id` - The OCID of the default compartment for buckets created through the Amazon S3 Compatibility API.
* `default_swift_compartment_id` - The OCID of the default compartment for buckets created through the Swift API.
* `namespace` - The Object Storage namespace to which the metadata belongs.
