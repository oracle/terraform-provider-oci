---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_fleet_compliance_report"
sidebar_current: "docs-oci-datasource-fleet_apps_management-fleet_compliance_report"
description: |-
  Provides details about a specific Fleet Compliance Report in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_fleet_compliance_report
This data source provides details about a specific Fleet Compliance Report resource in Oracle Cloud Infrastructure Fleet Apps Management service.

Retrieve compiane report for a Fleet

## Example Usage

```hcl
data "oci_fleet_apps_management_fleet_compliance_report" "test_fleet_compliance_report" {
	#Required
	compliance_report_id = oci_data_safe_report.test_report.id
	fleet_id = oci_fleet_apps_management_fleet.test_fleet.id
}
```

## Argument Reference

The following arguments are supported:

* `compliance_report_id` - (Required) compliance report identifier
* `fleet_id` - (Required) unique Fleet identifier


## Attributes Reference

The following attributes are exported:

* `compliance_state` - Last known compliance state of fleet.
* `fleet_id` - The fleet OCID.
* `id` - Compliance Report Identifier
* `resources` - Resources assocaited with the Fleet.
	* `compartment` - Compartment the resource belongs to.
	* `compliance_state` - Last known compliance state of fleet.
	* `products` - Products assocaited with the Fleet.Only products belonging to managed targets will be shown.
		* `product_name` - Product Name
		* `targets` - Managed Targets associated with the Product.
			* `compliance_state` - Last known compliance state of target.
			* `installed_patches` - Installed Patches for the Target.
				* `patch_description` - The OCID of the work request to start the analysis.
				* `patch_name` - The OCID to identify this analysis results.
				* `patch_type` - Type of patch.
				* `time_applied` - Time the patch was applied
				* `time_released` - Date on which patch was released.
			* `recommended_patches` - Recommended Patches for the Target.
				* `patch_description` - The OCID of the work request to start the analysis.
				* `patch_name` - The OCID to identify this analysis results.
				* `patch_type` - Type of patch.
				* `time_applied` - Time the patch was applied
				* `time_released` - Date on which patch was released.
			* `target_id` - Target Identifier.
			* `target_name` - Target Name.
			* `version` - Current version.
	* `resource_id` - The OCID to identify the resource.
	* `resource_name` - Display name of the resource.
	* `resource_region` - Region the resource belongs to.
	* `resource_type` - Type of the resource.
	* `tenancy_id` - TenancyId of the resource.
	* `tenancy_name` - Tenancy the resource belongs to.

