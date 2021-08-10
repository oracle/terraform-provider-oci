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

Imports an array of dashboards and their saved searches. Here's an example of how you can use CLI to import a dashboard. For information on the details that must be passed to IMPORT, you can use the EXPORT API to obtain the Import.json file:
oci management-dashboard dashboard export --query data --export-dashboard-id "{\"dashboardIds\":[\"ocid1.managementdashboard.oc1..dashboardId1\"]}"  > Import.json.
Note that import API updates the resource if it already exist, and creates a new resource if it does not exist. To import to a different compartment, edit and change the compartmentId to the desired compartment OCID.
Here is an example of how you can use CLI to do import:

oci management-dashboard dashboard import --from-json file://Import.json 


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


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Dashboards Import
	* `update` - (Defaults to 20 minutes), when updating the Management Dashboards Import
	* `delete` - (Defaults to 20 minutes), when destroying the Management Dashboards Import


## Import

ManagementDashboardsImport can be imported using the `id`, e.g.

```
$ terraform import oci_management_dashboard_management_dashboards_import.test_management_dashboards_import "id"
```

