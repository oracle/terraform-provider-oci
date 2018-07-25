---
layout: "oci"
page_title: "OCI: oci_file_storage_file_systems"
sidebar_current: "docs-oci-datasource-file_storage-file_systems"
description: |-
  Provides a list of FileSystems
---

# Data Source: oci_file_storage_file_systems
The FileSystems data source allows access to the list of OCI file_systems

Lists the file system resources in the specified compartment.


## Example Usage

```hcl
data "oci_file_storage_file_systems" "test_file_systems" {
	#Required
	availability_domain = "${var.file_system_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.file_system_display_name}"
	id = "${var.file_system_id}"
	state = "${var.file_system_state}"
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

* `file_systems` - The list of file_systems.

### FileSystem Reference

The following attributes are exported:

* `availability_domain` - The availability domain the file system is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the file system.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system` 
* `id` - The OCID of the file system.
* `metered_bytes` - The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system.  
* `state` - The current state of the file system.
* `time_created` - The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

