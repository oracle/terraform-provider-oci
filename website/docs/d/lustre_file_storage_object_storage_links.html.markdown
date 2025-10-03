---
subcategory: "Lustre File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_lustre_file_storage_object_storage_links"
sidebar_current: "docs-oci-datasource-lustre_file_storage-object_storage_links"
description: |-
  Provides the list of Object Storage Links in Oracle Cloud Infrastructure Lustre File Storage service
---

# Data Source: oci_lustre_file_storage_object_storage_links
This data source provides the list of Object Storage Links in Oracle Cloud Infrastructure Lustre File Storage service.

Gets a list of Object Storage links.


## Example Usage

```hcl
data "oci_lustre_file_storage_object_storage_links" "test_object_storage_links" {

	#Optional
	availability_domain = var.object_storage_link_availability_domain
	compartment_id = var.compartment_id
	display_name = var.object_storage_link_display_name
	id = var.object_storage_link_id
	lustre_file_system_id = oci_lustre_file_storage_lustre_file_system.test_lustre_file_system.id
	state = var.object_storage_link_state
}
```

## Argument Reference

The following arguments are supported:

* `availability_domain` - (Optional) The name of the availability domain.  Example: `Uocm:PHX-AD-1` 
* `compartment_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the given display name exactly.
* `id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Object Storage link.
* `lustre_file_system_id` - (Optional) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Lustre file system.
* `state` - (Optional) A filter to return only resources that match the given lifecycle state. The state value is case-insensitive. 


## Attributes Reference

The following attributes are exported:

* `object_storage_link_collection` - The list of object_storage_link_collection.

### ObjectStorageLink Reference

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

