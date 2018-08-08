---
layout: "oci"
page_title: "OCI: oci_file_storage_file_system"
sidebar_current: "docs-oci-resource-file_storage-file_system"
description: |-
  Creates and manages an OCI FileSystem
---

# oci_file_storage_file_system
The `oci_file_storage_file_system` resource creates and manages an OCI FileSystem

Creates a new file system in the specified compartment and
availability domain. Instances can mount file systems in
another availability domain, but doing so might increase
latency when compared to mounting instances in the same
availability domain.

After you create a file system, you can associate it with a mount
target. Instances can then mount the file system by connecting to the
mount target's IP address. You can associate a file system with
more than one mount target at a time.

For information about access control and compartments, see
[Overview of the IAM Service](https://docs.us-phoenix-1.oraclecloud.com/Content/Identity/Concepts/overview.htm).

For information about availability domains, see [Regions and
Availability Domains](/Content/General/Concepts/regions.htm).
To get a list of availability domains, use the
`ListAvailabilityDomains` operation in the Identity and Access
Management Service API.

All Oracle Cloud Infrastructure resources, including
file systems, get an Oracle-assigned, unique ID called an Oracle
Cloud Identifier (OCID).  When you create a resource, you can
find its OCID in the response. You can also retrieve a
resource's OCID by using a List API operation on that resource
type or by viewing the resource in the Console.


## Example Usage

```hcl
resource "oci_file_storage_file_system" "test_file_system" {
	#Required
	availability_domain = "${var.file_system_availability_domain}"
	compartment_id = "${var.compartment_id}"

	#Optional
	display_name = "${var.file_system_display_name}"
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) The OCID of the compartment to create the file system in.
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the file system is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the file system.
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system` 
* `id` - The OCID of the file system.
* `metered_bytes` - The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system. 
* `state` - The current state of the file system.
* `time_created` - The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Import

FileSystems can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_file_system.test_file_system "id"
```
