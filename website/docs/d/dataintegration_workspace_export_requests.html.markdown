---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_export_requests"
sidebar_current: "docs-oci-datasource-dataintegration-workspace_export_requests"
description: |-
  Provides the list of Workspace Export Requests in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace_export_requests
This data source provides the list of Workspace Export Requests in Oracle Cloud Infrastructure Data Integration service.

This endpoint can be used to get the list of export object requests.


## Example Usage

```hcl
data "oci_dataintegration_workspace_export_requests" "test_workspace_export_requests" {
	#Required
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	export_status = var.workspace_export_request_export_status
	name = var.workspace_export_request_name
	projection = var.workspace_export_request_projection
	time_ended_in_millis = var.workspace_export_request_time_ended_in_millis
	time_started_in_millis = var.workspace_export_request_time_started_in_millis
}
```

## Argument Reference

The following arguments are supported:

* `export_status` - (Optional) Specifies export status to use, either -  ALL, SUCCESSFUL, IN_PROGRESS, QUEUED, FAILED .
* `name` - (Optional) Used to filter by the name of the object.
* `projection` - (Optional) This parameter allows users to specify which view of the export object response to return. SUMMARY - Summary of the export object request will be returned. This is the default option when no value is specified. DETAILS - Details of export object request will be returned. This will include details of all the objects to be exported.
* `time_ended_in_millis` - (Optional) Specifies end time of a copy object request.
* `time_started_in_millis` - (Optional) Specifies start time of a copy object request.
* `workspace_id` - (Required) The workspace ID.


## Attributes Reference

The following attributes are exported:

* `export_request_summary_collection` - The list of export_request_summary_collection.

### WorkspaceExportRequest Reference

The following attributes are exported:

* `are_references_included` - Controls if the references will be exported along with the objects
* `bucket` - The name of the Object Storage Bucket where the objects will be exported to
* `created_by` - Name of the user who initiated export request.
* `error_messages` - Contains key of the error
* `exported_items` - The array of exported object details.
	* `aggregator_key` - Aggregator key
	* `identifier` - Object identifier
	* `key` - Key of the object
	* `name` - Name of the object
	* `name_path` - Object name path
	* `object_type` - Object type
	* `object_version` - Object version
	* `time_updated_in_millis` - time at which this object was last updated.
* `file_name` - Name of the exported zip file.
* `filters` - Export multiple objects based on filters.
* `is_object_overwrite_enabled` - Flag to control whether to overwrite the object if it is already present at the provided object storage location.
* `key` - Export object request key
* `name` - Name of the export request.
* `object_keys` - The list of the objects to be exported
* `object_storage_region` - Region of the object storage (if using object storage of different region)
* `object_storage_tenancy_id` - Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
* `referenced_items` - The array of exported referenced objects.
* `status` - Export Objects request status.
* `time_ended_in_millis` - Time at which the request was completely processed.
* `time_started_in_millis` - Time at which the request started getting processed.
* `total_exported_object_count` - Number of objects that are exported.

