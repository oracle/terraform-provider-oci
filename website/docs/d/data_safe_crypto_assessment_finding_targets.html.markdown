---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_finding_targets"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_finding_targets"
description: |-
  Provides the list of Crypto Assessment Finding Targets in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_finding_targets
This data source provides the list of Crypto Assessment Finding Targets in Oracle Cloud Infrastructure Data Safe service.

For a selected finding, lists targets where it occurs in assessments.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_finding_targets" "test_crypto_assessment_finding_targets" {
	#Required
	compartment_id = var.compartment_id
	finding_key = var.crypto_assessment_finding_target_finding_key

	#Optional
	access_level = var.crypto_assessment_finding_target_access_level
	assessment_type = var.crypto_assessment_finding_target_assessment_type
	compartment_id_in_subtree = var.crypto_assessment_finding_target_compartment_id_in_subtree
	is_quantum_readiness_check = var.crypto_assessment_finding_target_is_quantum_readiness_check
	status = var.crypto_assessment_finding_target_status
	target_id = oci_cloud_guard_target.test_target.id
	target_ids = var.crypto_assessment_finding_target_target_ids
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `assessment_type` - (Optional) A filter to return targets from assessments of the specified type.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `finding_key` - (Required) The finding keys for which target occurrences are listed.
* `is_quantum_readiness_check` - (Optional) A filter to return only findings that are or are not part of quantum-readiness checks.
* `status` - (Optional) A filter to return only finding target rows with the specified status.
* `target_id` - (Optional) Filters results to targets with an exact matching target OCID.
* `target_ids` - (Optional) A filter to return only resources associated with any of the specified target OCIDs.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_finding_target_collection` - The list of crypto_assessment_finding_target_collection.

### CryptoAssessmentFindingTarget Reference

The following attributes are exported:

* `items` - Array of target-level finding occurrences.
	* `assessment_id` - The crypto assessment OCID associated with this finding occurrence.
	* `database_version` - Database version of the affected target.
	* `finding_key` - Unique key of the finding affecting this target.
	* `is_quantum_readiness_check` - Indicates whether this finding is part of quantum-readiness checks.
	* `observed_value` - The observed value for the selected finding on this target.
	* `priority` - Numeric priority of the finding. 1 is CRITICAL, 2 is HIGH, 3 is MEDIUM, and 4 is LOW.
	* `severity` - Text severity derived from priority using the static mapping 1=CRITICAL, 2=HIGH, 3=MEDIUM, 4=LOW.
	* `target_id` - The OCID of the affected target.

