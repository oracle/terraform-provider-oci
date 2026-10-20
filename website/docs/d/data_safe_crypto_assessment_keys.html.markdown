---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_keys"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_keys"
description: |-
  Provides the list of Crypto Assessment Keys in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_keys
This data source provides the list of Crypto Assessment Keys in Oracle Cloud Infrastructure Data Safe service.

Gets a paginated list of cryptographic keys across targets in a compartment. Use assessmentId to narrow results to one crypto assessment.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_keys" "test_crypto_assessment_keys" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.crypto_assessment_key_access_level
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessment_type = var.crypto_assessment_key_assessment_type
	compartment_id_in_subtree = var.crypto_assessment_key_compartment_id_in_subtree
	feature = var.crypto_assessment_key_feature
	key_id = oci_kms_key.test_key.id
	key_manager_type = var.crypto_assessment_key_key_manager_type
	key_type = var.crypto_assessment_key_key_type
	target_id = oci_cloud_guard_target.test_target.id
	target_ids = var.crypto_assessment_key_target_ids
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `assessment_id` - (Optional) A filter to return only resources associated with the specified crypto assessment OCID.
* `assessment_type` - (Optional) A filter to return targets from assessments of the specified type.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `feature` - (Optional) A filter to return only records for the specified feature.
* `key_id` - (Optional) Filters key results to rows with an exact matching keyId.
* `key_manager_type` - (Optional) Filters key results to rows whose primary or secondary keystore type matches any of the specified key manager types.
* `key_type` - (Optional) Filters key results to rows with the specified key type.
* `target_id` - (Optional) A filter to return only inventory rows associated with the specified target OCID.
* `target_ids` - (Optional) A filter to return only resources associated with any of the specified target OCIDs.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_key_collection` - The list of crypto_assessment_key_collection.

### CryptoAssessmentKey Reference

The following attributes are exported:

* `items` - Cryptographic key summaries for the specified crypto assessment.
	* `age` - Age of the key in whole days, calculated from timeCreated using the current UTC date.
	* `algorithm` - Cryptographic algorithm used by the key.
	* `assessment_id` - OCID of the crypto assessment that discovered the key.
	* `feature` - Crypto feature for which the key is observed.
	* `key_cache` - Key cache setting observed for the key.
	* `key_id` - Identifier of the cryptographic key.
	* `key_type` - Type of cryptographic key observed for the assessment.
	* `keystore_type` - Primary keystore type observed for the key.
	* `secondary_keystore_type` - Secondary keystore type observed for the key, if configured.
	* `status` - Current status of the cryptographic key as observed on the target.
	* `target_id` - OCID of the target database associated with the key.
	* `time_created` - Key creation time in RFC3339 format.
	* `time_last_assessed` - The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	* `time_last_rotation` - Most recent key rotation time in RFC3339 format.
	* `wallet_location` - Wallet location observed for the key.

