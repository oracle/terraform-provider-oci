---
layout: "oci"
page_title: "OCI: oci_file_storage_exports"
sidebar_current: "docs-oci-datasource-file_storage-exports"
description: |-
Provides a list of Exports
---
# Data Source: oci_file_storage_exports
The Exports data source allows access to the list of OCI exports

Lists export resources by compartment, file system, or export
set. You must specify an export set ID, a file system ID, and
/ or a compartment ID.


## Example Usage

```hcl
data "oci_file_storage_exports" "test_exports" {

	#Optional
	compartment_id = "${var.compartment_id}"
	export_set_id = "${oci_file_storage_export_set.test_export_set.id}"
	file_system_id = "${oci_file_storage_file_system.test_file_system.id}"
	id = "${var.export_id}"
	state = "${var.export_state}"
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The OCID of the compartment.
* `export_set_id` - (Optional) The OCID of the export set.
* `file_system_id` - (Optional) The OCID of the file system.
* `id` - (Optional) Filter results by OCID. Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `exports` - The list of exports.

### Export Reference

The following attributes are exported:

* `export_set_id` - The OCID of this export's export set.
* `file_system_id` - The OCID of this export's file system.
* `id` - The OCID of this export.
* `path` - Path used to access the associated file system.  Avoid entering confidential information.  Example: `/accounting` 
* `state` - The current state of this export.
* `time_created` - The date and time the export was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

