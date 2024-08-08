---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_calculate_audit_volume_available"
sidebar_current: "docs-oci-resource-data_safe-calculate_audit_volume_available"
description: |-
  Provides the Calculate Audit Volume Available resource in Oracle Cloud Infrastructure Data Safe service
---

# oci_data_safe_calculate_audit_volume_available
This resource provides the Calculate Audit Volume Available resource in Oracle Cloud Infrastructure Data Safe service.

Calculates the volume of audit events available on the target database to be collected. Measurable up to the defined retention period of the audit target resource.

## Example Usage

```hcl
resource "oci_data_safe_calculate_audit_volume_available" "test_calculate_audit_volume_available" {
	#Required
	audit_profile_id = oci_data_safe_audit_profile.test_audit_profile.id

	#Optional
	audit_collection_start_time = var.calculate_audit_volume_available_audit_collection_start_time
	database_unique_name = var.calculate_audit_volume_available_database_unique_name
	trail_locations = var.calculate_audit_volume_available_trail_locations
}
```

## Argument Reference

The following arguments are supported:

* `audit_collection_start_time` - (Optional) The date from which the audit trail must start collecting data in UTC, in the format defined by RFC3339. If not specified, this will default to the date based on the retention period.
* `audit_profile_id` - (Required) The OCID of the audit.
* `database_unique_name` - (Optional) Unique name of the database associated to the peer target database.
* `trail_locations` - (Optional) The trail locations for which the audit data volume has to be calculated.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `available_audit_volumes` - List of available audit volumes.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/oracle/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Calculate Audit Volume Available
	* `update` - (Defaults to 20 minutes), when updating the Calculate Audit Volume Available
	* `delete` - (Defaults to 20 minutes), when destroying the Calculate Audit Volume Available


## Import

CalculateAuditVolumeAvailable can be imported using the `id`, e.g.

```
$ terraform import oci_data_safe_calculate_audit_volume_available.test_calculate_audit_volume_available "id"
```

