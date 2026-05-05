---
subcategory: "Management Dashboard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_dashboard_management_saved_search"
sidebar_current: "docs-oci-datasource-management_dashboard-management_saved_search"
description: |-
  Provides details about a specific Management Saved Search in Oracle Cloud Infrastructure Management Dashboard service
---

# Data Source: oci_management_dashboard_management_saved_search
This data source provides details about a specific Management Saved Search resource in Oracle Cloud Infrastructure Management Dashboard service.

Gets a saved search by ID.

## Example Usage

```hcl
data "oci_management_dashboard_management_saved_search" "test_management_saved_search" {
	#Required
	management_saved_search_id = oci_management_dashboard_management_saved_search.test_management_saved_search.id
}
```

## Argument Reference

The following arguments are supported:

* `management_saved_search_id` - (Required) A unique saved search identifier.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - OCID of the compartment in which the saved search resides.
* `created_by` - The principal id of the user that created this saved search. This is automatically managed by the system. In Oracle Cloud Infrastructure the value is ignored. In EM it can skipped or otherwise it is ignored in both create and update API and system automatically sets its value.
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

