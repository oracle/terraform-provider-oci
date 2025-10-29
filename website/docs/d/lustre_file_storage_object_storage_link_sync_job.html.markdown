---
subcategory: "Lustre File Storage"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_lustre_file_storage_object_storage_link_sync_job"
sidebar_current: "docs-oci-datasource-lustre_file_storage-object_storage_link_sync_job"
description: |-
  Provides details about a specific Object Storage Link Sync Job in Oracle Cloud Infrastructure Lustre File Storage service
---

# Data Source: oci_lustre_file_storage_object_storage_link_sync_job
This data source provides details about a specific Object Storage Link Sync Job resource in Oracle Cloud Infrastructure Lustre File Storage service.

Gets details of a sync job associated with an Object Storage link when `objectStorageLink` and a unique ID are provided.


## Example Usage

```hcl
data "oci_lustre_file_storage_object_storage_link_sync_job" "test_object_storage_link_sync_job" {
	#Required
	object_storage_link_id = oci_lustre_file_storage_object_storage_link.test_object_storage_link.id
	sync_job_id = oci_database_migration_job.test_job.id
}
```

## Argument Reference

The following arguments are supported:

* `object_storage_link_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Object Storage link.
* `sync_job_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the sync job.


## Attributes Reference

The following attributes are exported:

* `bytes_transferred` - Bytes transferred during the sync. This value changes while the sync is still in progress. 
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Operations.CostCenter": "42"}` 
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).  Example: `{"Department": "Finance"}` 
* `id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the sync job.
* `is_overwrite` - The flag is an identifier to tell whether this specific job run has overwrite enabled. If `isOverwrite` is false, the file to be imported or exported will be skipped if it already exists. If `isOverwrite` is true, the file to be imported or exported will be overwritten if it already exists. 
* `job_type` - The type of the sync job. 
* `lifecycle_details` - A message that describes the current state of the sync job in more detail. For example, can be used to provide actionable information for a resource in the Failed state. 
* `lustre_file_system_path` - The path in the Lustre file system used for this Object Storage link.  Example: `myFileSystem/mount/myDirectory` 
* `object_storage_path` - The Object Storage namespace and bucket name, including optional object prefix string, to use as the source for imports or destination for exports.  Example: `objectStorageNamespace:/bucketName/optionalFolder/optionalPrefix` 
* `objects_transferred` - Count of total files that transferred successfully. 
* `parent_id` - The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the Object Storage link.
* `skipped_error_count` - Count of files or objects that failed to export or import due to errors. 
* `state` - The current state of the sync job. 
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace.  Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_finished` - The date and time the job finished, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2020-07-25T21:10:29.600Z` 
* `time_started` - The date and time the job was started, expressed in [RFC 3339](https://tools.ietf.org/rfc/rfc3339) timestamp format.  Example: `2020-07-25T21:10:29.600Z` 
* `total_objects_scanned` - Total object count for scanned files for import or export as part of this sync job. 

