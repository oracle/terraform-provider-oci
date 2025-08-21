---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_audit_profile_target_overrides"
sidebar_current: "docs-oci-datasource-data_safe-audit_profile_target_overrides"
description: |-
  Provides the list of Audit Profile Target Overrides in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_audit_profile_target_overrides
This data source provides the list of Audit Profile Target Overrides in Oracle Cloud Infrastructure Data Safe service.

Gets a list of all targets whose audit settings override the target group setting.

## Example Usage

```hcl
data "oci_data_safe_audit_profile_target_overrides" "test_audit_profile_target_overrides" {
	#Required
	audit_profile_id = oci_data_safe_audit_profile.test_audit_profile.id

	#Optional
	display_name = var.audit_profile_target_override_display_name
}
```

## Argument Reference

The following arguments are supported:

* `audit_profile_id` - (Required) The OCID of the audit.
* `display_name` - (Optional) A filter to return only resources that match the specified display name. 


## Attributes Reference

The following attributes are exported:

* `target_override_collection` - The list of target_override_collection.

### AuditProfileTargetOverride Reference

The following attributes are exported:

* `items` - Array of target database override summary.
	* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
	* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
	* `is_paid_usage_enabled` - Indicates if you want to continue collecting audit records beyond the free limit of one million audit records per month per target database, potentially incurring additional charges. The default value is inherited from the global settings.  You can change at the global level or at the target level. 
	* `offline_months` - Number of months the audit records will be stored offline in the offline archive. Minimum: 0; Maximum: 72 months. If you have a requirement to store the audit data even longer (up to 168 months) in the offline archive, please contact the Oracle Support. 
	* `offline_months_source` - The name or the OCID of the resource from which the offline month retention setting is sourced. For example a target database group OCID or global.
	* `online_months` - Number of months the audit records will be stored online in the audit repository for immediate reporting and analysis. Minimum: 1; Maximum: 12 months 
	* `online_months_source` - The name or the OCID of the resource from which the online month retention setting is sourced. For example a target database group OCID or global.
	* `paid_usage_source` - The name or the OCID of the resource from which the paid usage setting is sourced. For example a target database group OCID or global.
	* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
	* `target_database_id` - The OCID of the target database that overrides from the audit profile of the target database group.
* `targets_conforming_count` - Number of target databases within the target database group that conform with the audit profile of the target database group.
* `targets_count` - Number of target databases within the target database group.
* `targets_overriding_count` - Number of target databases within the target database group that override the audit profile of the target database group.
* `targets_overriding_offline_months_count` - Number of target databases within the group that override the offline retention setting of the audit profile for the target database group.
* `targets_overriding_online_months_count` - Number of target databases within the group that override the online retention setting of the audit profile for the target database group.
* `targets_overriding_paid_usage_count` - Number of target databases within the group that override the paid usage setting of the audit profile for the target database group.

