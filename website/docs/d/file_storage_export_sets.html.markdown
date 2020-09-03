---
subcategory: "File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_export_sets"
sidebar_current: "docs-oci-datasource-file_storage-export_sets"
description: |-
  Provides the list of Export Sets in Oracle Cloud Infrastructure File Storage service
---

# Data Source: oci_file_storage_export_sets
This data source provides the list of Export Sets in Oracle Cloud Infrastructure File Storage service.

Lists the export set resources in the specified compartment.


## Example Usage

```hcl
data "oci_file_storage_export_sets" "test_export_sets" {
	#Required
	availability_domain = var.export_set_availability_domain
	compartment_id = var.compartment_id

	#Optional
	display_name = var.export_set_display_name
	id = var.export_set_id
	state = var.export_set_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment.
* `display_name` - (Optional) A user-friendly name. It does not have to be unique, and it is changeable.  Example: `My resource` 
* `id` - (Optional) Filter results by OCID. Must be an OCID of the correct type for the resouce type. 
* `state` - (Optional) Filter results by the specified lifecycle state. Must be a valid state for the resource type. 


## Attributes Reference

The following attributes are exported:

* `export_sets` - The list of export_sets.

### ExportSet Reference

The following attributes are exported:

* `availability_domain` - The availability domain the export set is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the export set.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My export set` 
* `id` - The OCID of the export set.
* `max_fs_stat_bytes` - Controls the maximum `tbytes`, `fbytes`, and `abytes`, values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tbytes` value reported by `FSSTAT` will be `maxFsStatBytes`. The value of `fbytes` and `abytes` will be `maxFsStatBytes` minus the metered size of the file system. If the metered size is larger than `maxFsStatBytes`, then `fbytes` and `abytes` will both be '0'. 
* `max_fs_stat_files` - Controls the maximum `tfiles`, `ffiles`, and `afiles` values reported by `NFS FSSTAT` calls through any associated mount targets. This is an advanced feature. For most applications, use the default value. The `tfiles` value reported by `FSSTAT` will be `maxFsStatFiles`. The value of `ffiles` and `afiles` will be `maxFsStatFiles` minus the metered size of the file system. If the metered size is larger than `maxFsStatFiles`, then `ffiles` and `afiles` will both be '0'. 
* `state` - The current state of the export set.
* `time_created` - The date and time the export set was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 
* `vcn_id` - The OCID of the virtual cloud network (VCN) the export set is in.

