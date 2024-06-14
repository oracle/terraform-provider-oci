---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_masking_policy_health_report_management"
sidebar_current: "docs-oci-resource-data_safe-masking_policy_health_report_management"
description: |-
  Provides Masking Policy Health Report Management resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_masking_policy_health_report_management
This resource provides Pre-masking Report Management resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified pre-masking health report.

## Example Usage

```hcl
resource "oci_data_safe_masking_policy_health_report_management" "test_pre_masking_report_management" { 
  #Required
  target_id = oci_data_safe_target_database.test_target_database.id
  masking_policy_id = oci_data_safe_masking_policy.test_masking_policy.id
}
```

## Argument Reference

The following arguments are supported:

* `target_id` - (Required) The OCID of the target database masked.
* `masking_policy_id` - (Required) The OCID of the masking policy.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the masking report.
* `id` - The OCID of the pre-masking report.
* `masking_policy_id` - The OCID of the masking policy used.
* `display_name`: The display name of the pre-masking report,
* `time_created`: The date and time the pre-masking report was created, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339),
* `time_updated`: The date and time the pre-masking report was updated, in the format defined by [RFC3339](https://tools.ietf.org/html/rfc3339),
* `description`: Description for the pre-masking report,
* `state` - The current state of the pre-masking report.
* `target_id` - The OCID of the target database used for pre-masking check.


## Import

Import is not supported for this resource.