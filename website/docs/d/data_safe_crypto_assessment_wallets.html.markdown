---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_wallets"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_wallets"
description: |-
  Provides the list of Crypto Assessment Wallets in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_wallets
This data source provides the list of Crypto Assessment Wallets in Oracle Cloud Infrastructure Data Safe service.

Gets wallet details across targets in a compartment. Use assessmentId to narrow results to one crypto assessment.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_wallets" "test_crypto_assessment_wallets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.crypto_assessment_wallet_access_level
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessment_type = var.crypto_assessment_wallet_assessment_type
	compartment_id_in_subtree = var.crypto_assessment_wallet_compartment_id_in_subtree
	feature = var.crypto_assessment_wallet_feature
	target_id = oci_cloud_guard_target.test_target.id
	target_ids = var.crypto_assessment_wallet_target_ids
	wallet_encryption_algorithm = var.crypto_assessment_wallet_wallet_encryption_algorithm
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `assessment_id` - (Optional) A filter to return only resources associated with the specified crypto assessment OCID.
* `assessment_type` - (Optional) A filter to return targets from assessments of the specified type.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `feature` - (Optional) A filter to return only wallets for the specified feature.
* `target_id` - (Optional) A filter to return only inventory rows associated with the specified target OCID.
* `target_ids` - (Optional) A filter to return only resources associated with any of the specified target OCIDs.
* `wallet_encryption_algorithm` - (Optional) A filter to return only wallets whose encryption algorithm exactly matches any of the specified values, case-insensitively.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_wallet_collection` - The list of crypto_assessment_wallet_collection.

### CryptoAssessmentWallet Reference

The following attributes are exported:

* `items` - Wallet details for the specified crypto assessment.
	* `assessment_id` - OCID of the crypto assessment that discovered the wallet.
	* `auto_login` - Whether wallet auto-login is enabled.
	* `feature` - Crypto feature for which wallet details are reported.
	* `target_id` - OCID of the target database associated with the wallet.
	* `time_created` - Wallet creation time in RFC3339 format.
	* `time_last_assessed` - The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	* `wallet_encryption_algorithm` - Wallet encryption algorithm.
	* `wallet_location` - Wallet path for the feature.

