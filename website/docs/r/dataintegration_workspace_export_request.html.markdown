---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_export_request"
sidebar_current: "docs-oci-resource-dataintegration-workspace_export_request"
description: |-
  Provides the Workspace Export Request resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace_export_request
This resource provides the Workspace Export Request resource in Oracle Cloud Infrastructure Data Integration service.

Export Metadata Object

## Example Usage

```hcl
resource "oci_dataintegration_workspace_export_request" "test_workspace_export_request" {
	#Required
	bucket = var.workspace_export_request_bucket
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	are_references_included = var.workspace_export_request_are_references_included
	file_name = var.workspace_export_request_file_name
	filters = var.workspace_export_request_filters
	is_object_overwrite_enabled = var.workspace_export_request_is_object_overwrite_enabled
	object_keys = var.workspace_export_request_object_keys
	object_storage_region = var.workspace_export_request_object_storage_region
	object_storage_tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `are_references_included` - (Optional) This field controls if the references will be exported along with the objects
* `bucket` - (Required) Name of the Object Storage bucket where the object will be exported.
* `file_name` - (Optional) Name of the exported zip file.
* `filters` - (Optional) Filters for exported objects
* `is_object_overwrite_enabled` - (Optional) Flag to control whether to overwrite the object if it is already present at the provided object storage location.
* `object_keys` - (Optional) Field is used to specify which object keys to export
* `object_storage_region` - (Optional) Region of the object storage (if using object storage of different region)
* `object_storage_tenancy_id` - (Optional) Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
* `workspace_id` - (Required) The workspace ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Workspace Export Request
	* `update` - (Defaults to 20 minutes), when updating the Workspace Export Request
	* `delete` - (Defaults to 20 minutes), when destroying the Workspace Export Request


## Import

WorkspaceExportRequests can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace_export_request.test_workspace_export_request "workspaces/{workspaceId}/exportRequests/{exportRequestKey}" 
```

