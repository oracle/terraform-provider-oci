---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_calculate_audit_volume_collected"
sidebar_current: "docs-oci-resource-data_safe-calculate_audit_volume_collected"
description: |-
  Provides the Calculate Audit Volume Collected resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_calculate_audit_volume_collected
This resource provides the Calculate Audit Volume Collected resource in Oracle Cloud Infrastructure Data Safe service.

Calculates the volume of audit events collected by data safe.

## Example Usage

```hcl
resource "oci_data_safe_calculate_audit_volume_collected" "test_calculate_audit_volume_collected" {
	#Required
	audit_profile_id = oci_data_safe_audit_profile.test_audit_profile.id
	time_from_month = var.calculate_audit_volume_collected_time_from_month

	#Optional
	time_to_month = var.calculate_audit_volume_collected_time_to_month
}
```

## Argument Reference

The following arguments are supported:

* `audit_profile_id` - (Required) The OCID of the audit.
* `time_from_month` - (Required) The date from which the audit volume collected by data safe has to be calculated, in the format defined by RFC3339.
* `time_to_month` - (Optional) The date from which the audit volume collected by data safe has to be calculated, in the format defined by RFC3339. If not specified, this will default to the current date.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `collected_audit_volumes` - List of collected audit volumes.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Calculate Audit Volume Collected
	* `update` - (Defaults to 20 minutes), when updating the Calculate Audit Volume Collected
	* `delete` - (Defaults to 20 minutes), when destroying the Calculate Audit Volume Collected


## Import

CalculateAuditVolumeCollected can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_calculate_audit_volume_collected.test_calculate_audit_volume_collected "id"
```

