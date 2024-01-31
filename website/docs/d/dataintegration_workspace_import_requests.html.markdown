---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_import_requests"
sidebar_current: "docs-oci-datasource-dataintegration-workspace_import_requests"
description: |-
  Provides the list of Workspace Import Requests in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace_import_requests
This data source provides the list of Workspace Import Requests in Oracle Cloud Infrastructure Data Integration service.

This endpoint can be used to get the list of import object requests.


## Example Usage

```hcl
data "oci_dataintegration_workspace_import_requests" "test_workspace_import_requests" {
	#Required
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	import_status = var.workspace_import_request_import_status
	name = var.workspace_import_request_name
	projection = var.workspace_import_request_projection
	time_ended_in_millis = var.workspace_import_request_time_ended_in_millis
	time_started_in_millis = var.workspace_import_request_time_started_in_millis
}
```

## Argument Reference

The following arguments are supported:

* `import_status` - (Optional) Specifies import status to use, either -  ALL, SUCCESSFUL, IN_PROGRESS, QUEUED, FAILED .
* `name` - (Optional) Used to filter by the name of the object.
* `projection` - (Optional) This parameter allows users to specify which view of the import object response to return. SUMMARY - Summary of the import object request will be returned. This is the default option when no value is specified. DETAILS - Details of import object request will be returned. This will include details of all the objects to be exported.
* `time_ended_in_millis` - (Optional) Specifies end time of a copy object request.
* `time_started_in_millis` - (Optional) Specifies start time of a copy object request.
* `workspace_id` - (Required) The workspace ID.


## Attributes Reference

The following attributes are exported:

* `import_request_summary_collection` - The list of import_request_summary_collection.

### WorkspaceImportRequest Reference

The following attributes are exported:

* `are_data_asset_references_included` - This field controls if the data asset references will be included during import.
* `bucket` - The name of the Object Storage Bucket where the objects will be imported from
* `created_by` - Name of the user who initiated import request.
* `error_messages` - Contains key of the error
* `file_name` - Name of the zip file from which objects will be imported.
* `import_conflict_resolution` - Import Objects Conflict resolution.
	* `duplicate_prefix` - In case of DUPLICATE mode, prefix will be used to disambiguate the object.
	* `duplicate_suffix` - In case of DUPLICATE mode, suffix will be used to disambiguate the object.
	* `import_conflict_resolution_type` - Import Objects Conflict resolution Type (RETAIN/DUPLICATE/REPLACE).
* `imported_objects` - The array of imported object details.
	* `aggregator_key` - Aggregator key
	* `identifier` - Object identifier
	* `name` - Name of the object
	* `name_path` - Object name path
	* `new_key` - New key of the object
	* `object_type` - Object type
	* `object_version` - Object version
	* `old_key` - Old key of the object
	* `resolution_action` - Object resolution action
	* `time_updated_in_millis` - time at which this object was last updated.
* `key` - Import object request key
* `name` - Name of the import request.
* `object_key_for_import` - Key of the object inside which all the objects will be imported
* `object_storage_region` - Region of the object storage (if using object storage of different region)
* `object_storage_tenancy_id` - Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
* `status` - Import Objects request status.
* `time_ended_in_millis` - Time at which the request was completely processed.
* `time_started_in_millis` - Time at which the request started getting processed.
* `total_imported_object_count` - Number of objects that are imported.

