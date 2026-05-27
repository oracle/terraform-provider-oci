---
subcategory: "Management Dashboard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_dashboard_management_saved_search"
sidebar_current: "docs-oci-resource-management_dashboard-management_saved_search"
description: |-
  Provides the Management Saved Search resource in Oracle Cloud Infrastructure Management Dashboard service
---

# oci_management_dashboard_management_saved_search
This resource provides the Management Saved Search resource in Oracle Cloud Infrastructure Management Dashboard service.

Creates a new saved search. 
Here's an example of how you can use CLI to create a saved search. For information on the details that must be passed to CREATE, you can use the GET API to obtain the Create.json file: 
`oci management-dashboard saved-search get --management-saved-search-id ocid1.managementsavedsearch.oc1..savedsearchId1 --query data > Create.json`. 
You can then modify the Create.json file by removing the `id` attribute and making other required changes, and use the `oci management-dashboard saved-search create` command. 


## Example Usage

```hcl
resource "oci_management_dashboard_management_saved_search" "test_management_saved_search" {
	#Required
	compartment_id = var.compartment_id
	data_config = var.management_saved_search_data_config
	description = var.management_saved_search_description
	display_name = var.management_saved_search_display_name
	is_oob_saved_search = var.management_saved_search_is_oob_saved_search
	metadata_version = var.management_saved_search_metadata_version
	nls = var.management_saved_search_nls
	provider_id = oci_management_dashboard_provider.test_provider.id
	provider_name = var.management_saved_search_provider_name
	provider_version = var.management_saved_search_provider_version
	screen_image = var.management_saved_search_screen_image
	type = var.management_saved_search_type
	ui_config = var.management_saved_search_ui_config
	widget_template = var.management_saved_search_widget_template
	widget_vm = var.management_saved_search_widget_vm
	drilldown_config = var.management_saved_search_drilldown_config
	parameters_config = var.management_saved_search_parameters_config

	#Optional
	defined_tags = {"foo-namespace.bar-key"= "value"}
	freeform_tags = {"bar-key"= "value"}
	id = var.management_saved_search_id
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Required) (Updatable) OCID of the compartment in which the saved search resides.
* `data_config` - (Required) (Updatable) It defines how data is fetched. A functional saved search needs a valid dataConfig. See examples on how it can be constructed for various data sources.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Required) (Updatable) Description of the saved search.
* `display_name` - (Required) (Updatable) Display name of the saved search.
* `drilldown_config` - (Required) (Updatable) Drill-down configuration to define the destination of a drill-down action.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - (Optional) ID of the saved search, which must only be provided for Out-of-the-Box (OOB) saved search.
* `is_oob_saved_search` - (Required) (Updatable) Determines whether the saved search is an Out-of-the-Box (OOB) saved search. Note that OOB saved searches are only provided by Oracle and cannot be modified.
* `metadata_version` - (Required) (Updatable) The version of the metadata defined in the API. This is maintained and enforced by dashboard server. Currently it is 2.0.
* `nls` - (Required) (Updatable) JSON that contains internationalization options.
* `parameters_config` - (Required) (Updatable) Defines parameters for the saved search.
* `provider_id` - (Required) (Updatable) ID of the service (for example log-analytics) that owns the saved search. Each service has a unique ID.
* `provider_name` - (Required) (Updatable) The user friendly name of the service (for example, Logging Analytics) that owns the saved search.
* `provider_version` - (Required) (Updatable) The version of the metadata of the provider. This is useful for provider to version its features and metadata. Any newly created saved search (or dashboard) should use providerVersion 3.0.0.
* `screen_image` - (Required) (Updatable) Screen image of the saved search.
* `type` - (Required) (Updatable) Determines how the saved search is displayed in a dashboard.
* `ui_config` - (Required) (Updatable) It defines the visualization type of the widget saved search, the UI options of that visualization type, the binding of data to the visualization.
* `widget_template` - (Required) (Updatable) The UI template that the saved search uses to render itself.
* `widget_vm` - (Required) (Updatable) The View Model that the saved search uses to render itself.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

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

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Saved Search
	* `update` - (Defaults to 20 minutes), when updating the Management Saved Search
	* `delete` - (Defaults to 20 minutes), when destroying the Management Saved Search


## Import

ManagementSavedSearches can be imported using the `id`, e.g.

```
$ terraform import oci_management_dashboard_management_saved_search.test_management_saved_search "id"
```

