---
subcategory: "Lustre File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_lustre_file_storage_object_storage_link"
sidebar_current: "docs-oci-resource-lustre_file_storage-object_storage_link"
description: |-
  Provides the Object Storage Link resource in Oracle Cloud Infrastructure Lustre File Storage service
---

# oci_lustre_file_storage_object_storage_link
This resource provides the Object Storage Link resource in Oracle Cloud Infrastructure Lustre File Storage service.

Creates an Object Storage link.


## Example Usage

```hcl
resource "oci_lustre_file_storage_object_storage_link" "test_object_storage_link" {
	#Required
	availability_domain = var.object_storage_link_availability_domain
	compartment_id = var.compartment_id
	file_system_path = var.object_storage_link_file_system_path
	is_overwrite = var.object_storage_link_is_overwrite
	lustre_file_system_id = oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
	object_storage_prefix = var.object_storage_link_object_storage_prefix

	#Optional
	defined_tags = {"Operations.CostCenter"= "42"}
	display_name = var.object_storage_link_display_name
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Required) The availability domain that the Lustre file system is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Required) (Updatable) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Object Storage link.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - (Optional) (Updatable) A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My Object Storage Link` 
* `file_system_path` - (Required) The path in the Lustre file system used for this Object Storage link.  Example: `myFileSystem/mount/myDirectory` 
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `is_overwrite` - (Required) (Updatable) The flag is an identifier to tell whether the job run has overwrite enabled. If `isOverwrite` is false, the file to be imported or exported will be skipped if it already exists. If `isOverwrite` is true, the file to be imported or exported will be overwritten if it already exists. 
* `lustre_file_system_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated Lustre file system. 
* `object_storage_prefix` - (Required) The Object Storage namespace and bucket name, including optional object prefix string, to use as the source for imports or destination for exports.  Example: `objectStorageNamespace:/bucketName/optionalFolder/optionalPrefix` 
* `start_export_to_object_trigger` - (Optional) (Updatable) An optional property when incremented triggers Start Export To Object. Could be set to any integer value.
* `start_import_from_object_trigger` - (Optional) (Updatable) An optional property when incremented triggers Start Import From Object. Could be set to any integer value.
* `stop_export_to_object_trigger` - (Optional) (Updatable) An optional property when incremented triggers Stop Export To Object. Could be set to any integer value.
* `stop_import_from_object_trigger` - (Optional) (Updatable) An optional property when incremented triggers Stop Import From Object. Could be set to any integer value.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `availability_domain` - The availability domain the file system is in. May be unset as a blank or NULL value.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment that contains the Lustre file system.
* `current_job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of currently running sync job. If no sync job is running, then this will be empty.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `display_name` - A user-friendly name. It does not have to be unique, and it is changeable. Avoid entering confidential information.  Example: `My Object Storage Link` 
* `file_system_path` - The path in the Lustre file system used for this Object Storage link.  Example: `myFileSystem/mount/myDirectory` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the ObjectStorageLink.
* `is_overwrite` - The flag is an identifier to tell whether the job run has overwrite enabled. If `isOverwrite` is false, the file to be imported or exported will be skipped if it already exists. If `isOverwrite` is true, the file to be imported or exported will be overwritten if it already exists. 
* `last_job_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of last succeeded sync job. If no sync job has previously run, then this will be empty.
* `lifecycle_details` - A message that describes the current state of the Object Storage link in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `lustre_file_system_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the associated Lustre file system. 
* `object_storage_prefix` - The Object Storage namespace and bucket name, including optional object prefix string, to use as the source for imports or destination for exports.  Example: `objectStorageNamespace:/bucketName/optionalFolder/optionalPrefix` 
* `state` - The current state of the Object Storage link.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - The date and time the Lustre file system was created, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2024-04-25T21:10:29.600Z` 
* `time_updated` - The date and time the Object Storage link was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2024-04-25T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Object Storage Link
	* `update` - (Defaults to 20 minutes), when updating the Object Storage Link
	* `delete` - (Defaults to 20 minutes), when destroying the Object Storage Link


## Import

ObjectStorageLinks can be imported using the `id`, e.g.

```
$ terraform import oci_lustre_file_storage_object_storage_link.test_object_storage_link "id"
```

