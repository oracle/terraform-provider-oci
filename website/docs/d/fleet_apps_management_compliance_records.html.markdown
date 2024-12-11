---
subcategory: "Fleet Apps Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_fleet_apps_management_compliance_records"
sidebar_current: "docs-oci-datasource-fleet_apps_management-compliance_records"
description: |-
  Provides the list of Compliance Records in Oracle Cloud Infrastructure Fleet Apps Management service
---

# Data Source: oci_fleet_apps_management_compliance_records
This data source provides the list of Compliance Records in Oracle Cloud Infrastructure Fleet Apps Management service.

Gets a list of complianceDetails.


## Example Usage

```hcl
data "oci_fleet_apps_management_compliance_records" "test_compliance_records" {

	#Optional
	compartment_id = var.compartment_id
	compliance_state = var.compliance_record_compliance_state
	entity_id = oci_fleet_apps_management_entity.test_entity.id
	product_name = var.compliance_record_product_name
	product_stack = var.compliance_record_product_stack
	resource_id = oci_cloud_guard_resource.test_resource.id
	target_name = oci_cloud_guard_target.test_target.name
}
```

## Argument Reference

The following arguments are supported:

* `compartment_id` - (Optional) The ID of the compartment in which to list resources.
* `compliance_state` - (Optional) Target Compliance State.
* `entity_id` - (Optional) Entity identifier.Ex:FleetId
* `product_name` - (Optional) Product Name.
* `product_stack` - (Optional) ProductStack name.
* `resource_id` - (Optional) Resource identifier.
* `target_name` - (Optional) Unique target name


## Attributes Reference

The following attributes are exported:

* `compliance_record_collection` - The list of compliance_record_collection.

### ComplianceRecord Reference

The following attributes are exported:

* `items` - List of compliancePolicys.
	* `compartment_id` - The OCID of the compartment.
	* `compliance_state` - Last known compliance state of target.
	* `entity_display_name` - The displayName of the entity for which the compliance is calculated.Ex.DisplayName for the Fleet
	* `entity_id` - The OCID of the entity for which the compliance is calculated.Ex.FleetId
	* `id` - The OCID of the ComplianceRecord.
	* `patch` - Details of the Patch
		* `patch_description` - Patch Description.
		* `patch_id` - patch OCID.
		* `patch_name` - patch Name.
		* `patch_type` - Type of patch.
		* `product` - Details of the Product
			* `product_name` - Product Name.
			* `product_stack` - Product Stack.
			* `product_version` - Product Version.
		* `severity` - Patch Severity.
		* `time_released` - Date on which patch was released
	* `policy` - Details of the Policy associated
		* `compliance_policy_display_name` - Compliane Policy DisplayName
		* `compliance_policy_id` - Compliance Policy Id
		* `compliance_policy_rule_display_name` - Product Name
		* `compliance_policy_rule_id` - Compliane Policy Rule Id
		* `grace_period` - Grace period in days,weeks,months or years the exemption is applicable for the rule.
		* `patch_selection` - Patch Selection Details
			* `days_since_release` - Days passed since patch release.
			* `patch_level` - Patch Name.
			* `patch_name` - Patch Name.
			* `selection_type` - Selection type for the Patch. 
	* `resource` - Details of the Resource
		* `compartment` - Compartment the resource belongs to.
		* `compartment_id` - TenancyId of the resource.
		* `resource_id` - The OCID to identify the resource.
		* `resource_name` - Name of the resource.
		* `resource_region` - Region the resource belongs to.
	* `state` - The current state of the ComplianceRecord.
	* `target` - Details of the Target
		* `target_id` - Target Identifier.
		* `target_name` - Target Name.
		* `version` - Current version.
	* `time_created` - The date and time the ComplianceRecord was created, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 
	* `time_updated` - The date and time the ComplianceRecord was updated, in the format defined by [RFC 3339](https://tools.ietf.org/html/rfc3339).  Example: `2016-08-25T21:10:29.600Z` 

