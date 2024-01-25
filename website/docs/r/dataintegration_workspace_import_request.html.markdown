---
subcategory: "Data Integration"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_dataintegration_workspace_import_request"
sidebar_current: "docs-oci-resource-dataintegration-workspace_import_request"
description: |-
  Provides the Workspace Import Request resource in Oracle Cloud Infrastructure Data Integration service
---

# oci_dataintegration_workspace_import_request
This resource provides the Workspace Import Request resource in Oracle Cloud Infrastructure Data Integration service.

Import Metadata Object

## Example Usage

```hcl
resource "oci_dataintegration_workspace_import_request" "test_workspace_import_request" {
	#Required
	bucket = var.workspace_import_request_bucket
	file_name = var.workspace_import_request_file_name
	workspace_id = oci_dataintegration_workspace.test_workspace.id

	#Optional
	are_data_asset_references_included = var.workspace_import_request_are_data_asset_references_included
	import_conflict_resolution {
		#Required
		import_conflict_resolution_type = var.workspace_import_request_import_conflict_resolution_import_conflict_resolution_type

		#Optional
		duplicate_prefix = var.workspace_import_request_import_conflict_resolution_duplicate_prefix
		duplicate_suffix = var.workspace_import_request_import_conflict_resolution_duplicate_suffix
	}
	object_key_for_import = var.workspace_import_request_object_key_for_import
	object_storage_region = var.workspace_import_request_object_storage_region
	object_storage_tenancy_id = oci_identity_tenancy.test_tenancy.id
}
```

## Argument Reference

The following arguments are supported:

* `are_data_asset_references_included` - (Optional) This field controls if the data asset references will be included during import.
* `bucket` - (Required) Name of the Object Storage bucket where the object will be imported from.
* `file_name` - (Required) Name of the zip file to be imported.
* `import_conflict_resolution` - (Optional) Import Objects Conflict resolution.
	* `duplicate_prefix` - (Optional) In case of DUPLICATE mode, prefix will be used to disambiguate the object.
	* `duplicate_suffix` - (Optional) In case of DUPLICATE mode, suffix will be used to disambiguate the object.
	* `import_conflict_resolution_type` - (Required) Import Objects Conflict resolution Type (RETAIN/DUPLICATE/REPLACE).
* `object_key_for_import` - (Optional) Key of the object inside which all the objects will be imported
* `object_storage_region` - (Optional) Region of the object storage (if using object storage of different region)
* `object_storage_tenancy_id` - (Optional) Optional parameter to point to object storage tenancy (if using Object Storage of different tenancy)
* `workspace_id` - (Required) The workspace ID.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Workspace Import Request
	* `update` - (Defaults to 20 minutes), when updating the Workspace Import Request
	* `delete` - (Defaults to 20 minutes), when destroying the Workspace Import Request


## Import

WorkspaceImportRequests can be imported using the `id`, e.g.

```
$ terraform import oci_dataintegration_workspace_import_request.test_workspace_import_request "workspaces/{workspaceId}/importRequests/{importRequestKey}" 
```

