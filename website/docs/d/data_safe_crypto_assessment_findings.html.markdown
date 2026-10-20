---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_findings"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_findings"
description: |-
  Provides the list of Crypto Assessment Findings in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_findings
This data source provides the list of Crypto Assessment Findings in Oracle Cloud Infrastructure Data Safe service.

Lists crypto deviation findings for the specified crypto assessment.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_findings" "test_crypto_assessment_findings" {
	#Required
	crypto_assessment_id = oci_data_safe_crypto_assessment.test_crypto_assessment.id

	#Optional
	category = var.crypto_assessment_finding_category
	finding_key = var.crypto_assessment_finding_finding_key
	is_quantum_readiness_check = var.crypto_assessment_finding_is_quantum_readiness_check
	status = var.crypto_assessment_finding_status
	title = var.crypto_assessment_finding_title
}
```

## Argument Reference

The following arguments are supported:

* `category` - (Optional) A filter to return only findings in the specified category key.
* `crypto_assessment_id` - (Required) The OCID of the crypto assessment.
* `finding_key` - (Optional) A filter to return only findings with the specified finding key.
* `is_quantum_readiness_check` - (Optional) A filter to return only findings that are or are not part of quantum-readiness checks.
* `status` - (Optional) A filter to return only findings with the specified status.
* `title` - (Optional) A filter to return only findings with the specified title.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_finding_collection` - The list of crypto_assessment_finding_collection.

### CryptoAssessmentFinding Reference

The following attributes are exported:

* `assessment_type` - Assessment type recorded in the findings table.
* `database_version` - Database version for the target database associated with the specified crypto assessment.
* `items` - Array of crypto deviation findings.
	* `category` - Category value recorded for the finding.
	* `compliance` - Compliance mapping recorded for the finding.
	* `expected_value` - Expected value recorded for the finding.
	* `finding_key` - Unique key identifier for the finding.
	* `is_quantum_readiness_check` - Indicates whether this finding is part of quantum-readiness checks.
	* `observed_value` - Observed values recorded for the finding.
	* `priority` - Numeric priority of the finding. 1 is CRITICAL, 2 is HIGH, 3 is MEDIUM, and 4 is LOW.
	* `recommended_value` - Recommended value recorded for the finding.
	* `remediation` - Remediation text recorded for the finding.
	* `severity` - Text severity derived from priority using the static mapping 1=CRITICAL, 2=HIGH, 3=MEDIUM, 4=LOW.
	* `short_remediation` - Short remediation for the finding.
	* `short_summary` - Short summary of the finding.
	* `status` - Status recorded for the finding.
	* `summary` - Summary recorded for the finding.
	* `title` - Human-readable title for the finding.
	* `url` - URL recorded for the finding.
* `summary` - Aggregate finding counts for the specified crypto assessment.
	* `backup_status` - Aggregate finding counts for one crypto finding category.
		* `findings` - Findings with FAIL or EVALUATE status.
		* `pass_checks` - Findings with PASS status.
		* `total_checks` - Total findings across all statuses in this category.
	* `critical` - FAIL or EVALUATE findings with priority 1.
	* `data_encryption_status` - Aggregate finding counts for one crypto finding category.
		* `findings` - Findings with FAIL or EVALUATE status.
		* `pass_checks` - Findings with PASS status.
		* `total_checks` - Total findings across all statuses in this category.
	* `high` - FAIL or EVALUATE findings with priority 2.
	* `low` - FAIL or EVALUATE findings with priority 4.
	* `med` - FAIL or EVALUATE findings with priority 3.
	* `network_encryption_status` - Aggregate finding counts for one crypto finding category.
		* `findings` - Findings with FAIL or EVALUATE status.
		* `pass_checks` - Findings with PASS status.
		* `total_checks` - Total findings across all statuses in this category.
	* `status_counts` - Counts keyed by finding status. All supported statuses are included with a zero count when absent.
	* `total_checks` - Total findings across all statuses.
	* `total_findings` - Findings with FAIL or EVALUATE status.
	* `wallet_status` - Aggregate finding counts for one crypto finding category.
		* `findings` - Findings with FAIL or EVALUATE status.
		* `pass_checks` - Findings with PASS status.
		* `total_checks` - Total findings across all statuses in this category.
* `target_id` - The OCID of the target database for the specified crypto assessment.

