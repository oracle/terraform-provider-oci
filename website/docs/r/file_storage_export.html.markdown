---
layout: "oci"
page_title: "OCI: oci_file_storage_export"
sidebar_current: "docs-oci-resource-file_storage-export"
description: |-
  Creates and manages an OCI Export
---

# oci_file_storage_export
The `oci_file_storage_export` resource creates and manages an OCI Export

Creates a new export in the specified export set, path, and
file system.


## Example Usage

```hcl
resource "oci_file_storage_export" "test_export" {
	#Required
	export_set_id = "${oci_file_storage_mount_target.test_mount_target.export_set_id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	path = "${var.export_path}"
}
```

## Argument Reference

The following arguments are supported:

* `export_set_id` - (Required) The OCID of this export's export set.
* `file_system_id` - (Required) The OCID of this export's file system.
* `path` - (Required) Path used to access the associated file system.  Avoid entering confidential information.  Example: `/mediafiles` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `export_set_id` - The OCID of this export's export set.
* `file_system_id` - The OCID of this export's file system.
* `id` - The OCID of this export.
* `path` - Path used to access the associated file system.  Avoid entering confidential information.  Example: `/accounting` 
* `state` - The current state of this export.
* `time_created` - The date and time the export was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Import

Exports can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_export.test_export "id"
```
