---
layout: "oci"
page_title: "OCI: oci_file_storage_snapshots"
sidebar_current: "docs-oci-datasource-file_storage-snapshots"
description: |-
  Provides a list of Snapshots
---

# Data Source: oci_file_storage_snapshots
The Snapshots data source allows access to the list of OCI snapshots

Lists snapshots of the specified file system.


## Example Usage

```hcl
data "oci_file_storage_snapshots" "test_snapshots" {
	#Required
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"

	#Optional
	id = "${var.snapshot_id}"
	state = "${var.snapshot_state}"
}
```

## Argument Reference

The following arguments are supported:

* `file_system_id` - (Required) The OCID of the file system.
* `id` - (Optional) Filter results by OCID. Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `snapshots` - The list of snapshots.

### Snapshot Reference

The following attributes are exported:

* `file_system_id` - The OCID of the file system from which the snapshot was created. 
* `id` - The OCID of the snapshot.
* `name` - Name of the snapshot. This value is immutable.  Avoid entering confidential information.  Example: `Sunday` 
* `state` - The current state of the snapshot.
* `time_created` - The date and time the snapshot was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

