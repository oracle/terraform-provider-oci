---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_export_request"
sidebar_current: "docs-oci-datasource-dataintegration-workspace_export_request"
description: |-
  Provides details about a specific Workspace Export Request in Oracle Cloud Infrastructure Data Integration service
---

# Data Source: oci_dataintegration_workspace_export_request
This data source provides details about a specific Workspace Export Request resource in Oracle Cloud Infrastructure Data Integration service.

This endpoint can be used to get the summary/details of object being exported.


## Example Usage

```hcl
data "oci_dataintegration_workspace_export_request" "test_workspace_export_request" {
	#Required
	export_request_key = var.workspace_export_request_export_request_key
	workspace_id = oci_dataintegration_workspace.test_workspace.id
}
```

## Argument Reference

The following arguments are supported:

* `export_request_key` - (Required) The key of the object export object request
* `workspace_id` - (Required) The workspace ID.


## Attributes Reference

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

