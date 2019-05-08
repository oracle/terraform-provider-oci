---
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_file_storage_file_system"
sidebar_current: "docs-oci-resource-file_storage-file_system"
description: |-
  Provides the File System resource in Oracle Cloud Infrastructure File Storage service
---

# oci_file_storage_file_system
This resource provides the File System resource in Oracle Cloud Infrastructure File Storage service.

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
[Overview of the IAM Service](https://docs.cloud.oracle.com/iaas/Content/Identity/Concepts/overview.htm).

For information about availability domains, see [Regions and
Availability Domains](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/regions.htm).
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
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = "${var.file_system_display_name}"
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain to create the file system in.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The OCID of the compartment to create the file system in.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the file system is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The OCID of the compartment that contains the file system.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My file system` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the file system.
* `metered_bytes` - The number of bytes consumed by the file system, including any snapshots. This number reflects the metered size of the file system and is updated asynchronously with respect to updates to the file system. 
* `state` - The current state of the file system.
* `time_created` - The date and time the file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2016-08-25T21:10:29.600Z` 

## Import

FileSystems can be imported using the `id`, e.g.

```
$ terraform import oci_file_storage_file_system.test_file_system "id"
```

