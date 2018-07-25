---
layout: "oci"
page_title: "OCI: oci_file_storage_snapshot"
sidebar_current: "docs-oci-resource-file_storage-snapshot"
description: |-
  Creates and manages an OCI Snapshot
---

# oci_file_storage_snapshot
The `oci_file_storage_snapshot` resource creates and manages an OCI Snapshot

Creates a new snapshot of the specified file system. You
can access the snapshot at `.snapshot/<name>`.


## Example Usage

```hcl
resource "oci_file_storage_snapshot" "test_snapshot" {
	#Required
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	name = "${var.snapshot_name}"
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Required) The OCID of this export's file system.
* `name` - (Required) Name of the snapshot. This value is immutable. It must also be unique with respect to all other non-DELETED snapshots on the associated file system.  Avoid entering confidential information.  Example: `Sunday` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `file_system_id` - The OCID of the file system from which the snapshot was created. 
* `id` - The OCID of the snapshot.
* `name` - Name of the snapshot. This value is immutable.  Avoid entering confidential information.  Example: `Sunday` 
* `state` - The current state of the snapshot.
* `time_created` - The date and time the snapshot was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Import

Snapshots can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_snapshot.test_snapshot "id"
```
