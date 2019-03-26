---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_snapshot"
sidebar_current: "docs-oci-datasource-file_storage-snapshot"
description: |-
  Provides details about a specific Snapshot in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_snapshot
This data source provides details about a specific Snapshot resource in Oracle Cloud Infrastructure File Storage service.

Gets the specified snapshot's information.

## Example Usage

```hcl
data "oci_file_storage_snapshot" "test_snapshot" {
	#Required
	snapshot_id = "${oci_file_storage_snapshot.test_snapshot.id}"
}
```

## Argument Reference

The following arguments are supported:

* `snapshot_id` - (Required) The OCID of the snapshot.


## Attributes Reference

The following attributes are exported:

* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `file_system_id` - The OCID of the file system from which the snapshot was created. 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the snapshot.
* `name` - Name of the snapshot. This value is immutable.

	Avoid entering confidential information.

	Example: `Sunday` 
* `state` - The current state of the snapshot.
* `time_created` - The date and time the snapshot was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

