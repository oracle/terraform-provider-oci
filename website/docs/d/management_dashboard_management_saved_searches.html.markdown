---
subcategory: "Management Dashboard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_dashboard_management_saved_searches"
sidebar_current: "docs-oci-datasource-management_dashboard-management_saved_searches"
description: |-
  Provides the list of Management Saved Searches in Oracle Cloud Infrastructure Management Dashboard service
---

# Data Source: oci_management_dashboard_management_saved_searches
This data source provides the list of Management Saved Searches in Oracle Cloud Infrastructure Management Dashboard service.

Gets the list of saved searches in a compartment with pagination.  Returned properties are the summary.

## Example Usage

```hcl
data "oci_management_dashboard_management_saved_searches" "test_management_saved_searches" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	display_name = var.management_saved_search_display_name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) The ID of the compartment in which to list resources.
* `display_name` - (Optional) A filter to return only resources that match the entire display name given.


## Attributes Reference

The following attributes are exported:

* `management_saved_search_collection` - The list of management_saved_search_collection.

### ManagementSavedSearch Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment in which the saved search resides.
* `created_by` - The principle id of the user that created this saved search. This is automatically managed by the system. In Oracle Cloud Infrastructure the value is ignored. In EM it can skipped or otherwise it is ignored in both create and update API and system automatically sets its value.
* `data_config` - It defines how data is fetched. A functional saved search needs a valid dataConfig. See examples on how it can be constructed for various data sources.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - Description of the saved search.
* `display_name` - Display name of the saved search.
* `drilldown_config` - Drill-down configuration to define the destination of a drill-down action.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - ID of the saved search.
* `is_oob_saved_search` - Determines whether the saved search is an Out-of-the-Box (OOB) saved search. Note that OOB saved searches are only provided by Oracle and cannot be modified.
* `metadata_version` - The version of the metadata defined in the API. This is maintained and enforced by dashboard server. Currently it is 2.0.
* `nls` - JSON that contains internationalization options.
* `parameters_config` - Defines parameters for the saved search.
* `provider_id` - ID of the service (for example log-analytics) that owns the saved search. Each service has a unique ID.
* `provider_name` - Name of the service (for example, Logging Analytics) that owns the saved search.
* `provider_version` - Version of the service that owns this saved search.
* `screen_image` - Screen image of the saved search.
* `state` - Oracle Cloud Infrastructure lifecycle status. This is automatically managed by the system.
* `system_tags` - Usage of system tag keys. These predefined keys are scoped to namespaces. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `time_created` - Date and time the saved search was created.
* `time_updated` - Date and time the saved search was updated.
* `type` - Determines how the saved search is displayed in a dashboard.
* `ui_config` - It defines the visualization type of the widget saved search, the UI options of that visualization type, the binding of data to the visualization.
* `updated_by` - The principle id of the user that updated this saved search.
* `widget_template` - The UI template that the saved search uses to render itself.
* `widget_vm` - The View Model that the saved search uses to render itself.

