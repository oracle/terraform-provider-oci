---
subcategory: "Management Dashboard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_dashboard_management_dashboards_export"
sidebar_current: "docs-oci-datasource-management_dashboard-management_dashboards_export"
description: |-
  Provides details about a specific Management Dashboards Export in Oracle Cloud Infrastructure Management Dashboard service
---

# Data Source: oci_management_dashboard_management_dashboards_export
This data source provides details about a specific Management Dashboards Export resource in Oracle Cloud Infrastructure Management Dashboard service.

Exports an array of dashboards and their saved searches. Export is designed to work with importDashboard. Here's an example of how you can use CLI to export a dashboard. $oci management-dashboard dashboard export --query data --export-dashboard-id "{\"dashboardIds\":[\"ocid1.managementdashboard.oc1..dashboardId1\"]}"  > dashboards.json

## Example Usage

```hcl
data "oci_management_dashboard_management_dashboards_export" "test_management_dashboards_export" {
	#Required
	export_dashboard_id = oci_management_dashboard_export_dashboard.test_export_dashboard.id
}
```

## Argument Reference

The following arguments are supported:

* `export_dashboard_id` - (Required) List of dashboardIds in plain text. The syntax is '{"dashboardIds":["dashboardId1", "dashboardId2", ...]}'. Escaping is needed when using in Oracle Cloud Infrastructure CLI. For example, "{\"dashboardIds\":[\"ocid1.managementdashboard.oc1..dashboardId1\"]}" .


## Attributes Reference

The following attributes are exported:

* `export_details` - String containing Array of Dashboards exported, check [ManagementDashboardExportDetails](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/managementdashboard/20200901/datatypes/ManagementDashboardExportDetails) for exact contents in the string value. The value of `export_details` can be used to pass as `import_details` (CompartmentIds may have to be changed) in `oci_management_dashboard_management_dashboards_import` resource.
