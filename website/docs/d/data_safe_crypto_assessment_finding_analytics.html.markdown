---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_finding_analytics"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_finding_analytics"
description: |-
  Provides the list of Crypto Assessment Finding Analytics in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_finding_analytics
This data source provides the list of Crypto Assessment Finding Analytics in Oracle Cloud Infrastructure Data Safe service.

Lists findings in a compartment with the number of affected targets.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_finding_analytics" "test_crypto_assessment_finding_analytics" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.crypto_assessment_finding_analytic_access_level
	category = var.crypto_assessment_finding_analytic_category
	compartment_id_in_subtree = var.crypto_assessment_finding_analytic_compartment_id_in_subtree
	finding_key = var.crypto_assessment_finding_analytic_finding_key
	is_quantum_readiness_check = var.crypto_assessment_finding_analytic_is_quantum_readiness_check
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `category` - (Optional) A filter to return only findings in the specified category key.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `finding_key` - (Optional) A filter to return only findings with any of the specified finding keys.
* `is_quantum_readiness_check` - (Optional) A filter to return only findings that are or are not part of quantum-readiness checks.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_finding_analytics_collection` - The list of crypto_assessment_finding_analytics_collection.

### CryptoAssessmentFindingAnalytic Reference

The following attributes are exported:

* `items` - Array of crypto finding analytics summaries.
	* `category` - Category key to which the finding belongs.
	* `finding_key` - Unique key of the finding.
	* `priority` - Numeric priority of the finding. 1 is CRITICAL, 2 is HIGH, 3 is MEDIUM, and 4 is LOW.
	* `severity` - Text severity derived from priority using the static mapping 1=CRITICAL, 2=HIGH, 3=MEDIUM, 4=LOW.
	* `short_remediation` - Short remediation for the finding.
	* `short_summary` - Short summary of the finding.
	* `target_count` - Number of targets impacted by this finding in the queried scope.
	* `title` - Display title of the finding.

