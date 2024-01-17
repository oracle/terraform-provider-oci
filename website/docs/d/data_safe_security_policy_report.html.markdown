---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_security_policy_report"
sidebar_current: "docs-oci-datasource-data_safe-security_policy_report"
description: |-
  Provides details about a specific Security Policy Report in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_security_policy_report
This data source provides details about a specific Security Policy Report resource in Oracle Cloud Infrastructure Data Safe service.

Gets a security policy report by the specified OCID of the security policy report resource.

## Example Usage

```hcl
data "oci_data_safe_security_policy_report" "test_security_policy_report" {
	#Required
	security_policy_report_id = oci_data_safe_security_policy_report.test_security_policy_report.id
}
```

## Argument Reference

The following arguments are supported:

* `security_policy_report_id` - (Required) The OCID of the security policy report resource.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the security policy report.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the security policy report.
* `display_name` - The display name of the security policy report.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the security policy report.
* `lifecycle_details` - Details about the current state of the security policy report.
* `state` - The current state of the security policy report.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_id` - The OCID of the of the  target database.
* `time_created` - The date and time the security policy report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).
* `time_updated` - The date and time the security policy report was last updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339).

