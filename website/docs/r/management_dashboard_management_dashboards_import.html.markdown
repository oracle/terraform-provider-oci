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
Api doc link for the resource: https://docs.oracle.com/iaas/api/#/en/managementdashboard/latest/ManagementDashboardsImport

Example terraform configs related to the resource : https://github.com/oracle/terraform-provider-oci/tree/master/examples/management_dashboard

Imports an array of dashboards and their saved searches. 
Here's an example of how you can use CLI to import a dashboard. For information on the details that must be passed to IMPORT, you can use the EXPORT API to obtain the Import.json file: 
`oci management-dashboard dashboard export --query data --export-dashboard-id "{\"dashboardIds\":[\"ocid1.managementdashboard.oc1..dashboardId1\"]}"  > Import.json`. 
Note that import API updates the resource if it already exists, and creates a new resource if it does not exist. To import to a different compartment, edit and change the compartmentId to the desired compartment OCID. 
Here's an example of how you can use CLI to import:
`oci management-dashboard dashboard import --from-json file://Import.json`


## Example Usage

```hcl
resource "oci_management_dashboard_management_dashboards_import" "test_management_dashboards_import" {
	#Optional
	override_dashboard_compartment_ocid = var.management_dashboards_import_override_dashboard_compartment_ocid
	override_same_name = var.management_dashboards_import_override_same_name
	override_saved_search_compartment_ocid = var.management_dashboards_import_override_saved_search_compartment_ocid
}
```

## Argument Reference

The following arguments are supported:


* `override_dashboard_compartment_ocid` - (Optional) If this attribute is set, the dashboard resources are created or updated in the compartment specified by OCID. If this attribute is not set, the compartment specified in the JSON metadata is used. 
* `override_same_name` - (Optional) By default, if a resource with the same OCID exists in the target compartment, it is updated during the import process, otherwise, a new resource is created. However, if this attribute is set to true, then during the import process if a resource with the same displayName exists in the compartment, then it is updated even if the OCIDs are different. This is useful when importing the same resource multiple times. If the compartment and displayName remain the same, the resource is only updated and multiple copies of a resource are not created. 
* `override_saved_search_compartment_ocid` - (Optional) If this attribute is set, the saved search resources are created or updated in the compartment specified by OCID. If this attribute is not set, the compartment specified in the JSON metadata is used. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:


## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Management Dashboards Import
	* `update` - (Defaults to 20 minutes), when updating the Management Dashboards Import
	* `delete` - (Defaults to 20 minutes), when destroying the Management Dashboards Import


## Import

ManagementDashboardsImport can be imported using the `id`, e.g.

```
$ terraform import oci_management_dashboard_management_dashboards_import.test_management_dashboards_import "id"
```

