---
subcategory: "Core"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_core_instance_measured_boot_report"
sidebar_current: "docs-oci-datasource-core-instance_measured_boot_report"
description: |-
  Provides details about a specific Instance Measured Boot Report in Oracle Cloud Infrastructure Core service
---

# Data Source: oci_core_instance_measured_boot_report
This data source provides details about a specific Instance Measured Boot Report resource in Oracle Cloud Infrastructure Core service.

Gets the measured boot report for this shielded instance.

## Example Usage

```hcl
data "oci_core_instance_measured_boot_report" "test_instance_measured_boot_report" {
	#Required
	instance_id = oci_core_instance.test_instance.id
}
```

## Argument Reference

The following arguments are supported:

* `instance_id` - (Required) The [OCID](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/identifiers.htm) of the instance.


## Attributes Reference

The following attributes are exported:

* `is_policy_verification_successful` - Whether the verification succeeded, and the new values match the expected values. 
* `measurements` - A list of Trusted Platform Module (TPM) Platform Configuration Register (PCR) entries. 
	* `actual` - The list of actual PCR entries measured during boot.
		* `hash_algorithm` - The type of algorithm used to calculate the hash.
		* `pcr_index` - The index of the policy.
		* `value` - The hashed PCR value.
	* `policy` - The list of expected PCR entries to use during verification.
		* `hash_algorithm` - The type of algorithm used to calculate the hash.
		* `pcr_index` - The index of the policy.
		* `value` - The hashed PCR value.

