---
subcategory: "Management Dashboard"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_management_dashboard_management_dashboards_import"
sidebar_current: "docs-oci-resource-management_dashboard-management_dashboards_import"
description: |-
  Provides the Management Dashboards Import resource in Oracle Cloud Infrastructure Management Dashboard service
---

# oci_management_dashboard_management_dashboards_import
This resource provides the Management Dashboards Import resource in Oracle Cloud Infrastructure Management Dashboard service.

Import an array of dashboards and their saved searches.


## Example Usage

```hcl
resource "oci_management_dashboard_management_dashboards_import" "test_management_dashboards_import" {
	#Optional
	import_details = var.sample_import_details
	import_details_file = var.sample_import_details_file_path
}
```

## Argument Reference

The following arguments are supported:

* `import_details` - (Optional) Array of Dashboards to import. The `import_details` is mandatory if `import_details_path` is not passed. Value should be stringified JSON of [ManagementDashboardImportDetails](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/managementdashboard/20200901/ManagementDashboardImportDetails/)
* `import_details_path` - (Optional) Array of Dashboards to import. The `import_details_path` is mandatory if `import_details` is not passed. Value should be the path to the JSON file containing [ManagementDashboardImportDetails](https://docs.cloud.oracle.com/en-us/iaas/api/#/en/managementdashboard/20200901/ManagementDashboardImportDetails/)

** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Import

ManagementDashboardsImport can be imported using the `id`, e.g.

```
$ terraform import oci_management_dashboard_management_dashboards_import.test_management_dashboards_import "id"
```

