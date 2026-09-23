---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_backup_sets"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_backup_sets"
description: |-
  Provides the list of Crypto Assessment Backup Sets in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_backup_sets
This data source provides the list of Crypto Assessment Backup Sets in Oracle Cloud Infrastructure Data Safe service.

Gets backup set summaries across targets in a compartment. Use assessmentId to narrow results to one crypto assessment.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_backup_sets" "test_crypto_assessment_backup_sets" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.crypto_assessment_backup_set_access_level
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessment_type = var.crypto_assessment_backup_set_assessment_type
	backup_set_key = var.crypto_assessment_backup_set_backup_set_key
	compartment_id_in_subtree = var.crypto_assessment_backup_set_compartment_id_in_subtree
	is_encrypted = var.crypto_assessment_backup_set_is_encrypted
	target_id = oci_cloud_guard_target.test_target.id
	target_ids = var.crypto_assessment_backup_set_target_ids
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `assessment_id` - (Optional) A filter to return only resources associated with the specified crypto assessment OCID.
* `assessment_type` - (Optional) A filter to return targets from assessments of the specified type.
* `backup_set_key` - (Optional) Filters backup set summary rows to an exact matching backupSetKey.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `is_encrypted` - (Optional) Filters backup set summary rows by whether the backup set is encrypted.
* `target_id` - (Optional) A filter to return only inventory rows associated with the specified target OCID.
* `target_ids` - (Optional) A filter to return only resources associated with any of the specified target OCIDs.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_backup_set_collection` - The list of crypto_assessment_backup_set_collection.

### CryptoAssessmentBackupSet Reference

The following attributes are exported:

* `items` - Backup set summary items for the specified crypto assessment.
	* `algorithm_observed` - Encryption algorithm observed for the backup set when encryption is enabled.
	* `assessment_id` - OCID of the crypto assessment that discovered the backup set.
	* `backup_pieces` - Number of backup pieces in the set.
	* `backup_set_key` - Backup set key.
	* `backup_type` - Backup type observed for the set.
	* `cipher_mode_observed` - Cipher mode observed for the backup set when encryption is enabled.
	* `is_compressed` - Indicates whether the backup set is compressed.
	* `is_encrypted` - Indicates whether the backup set is encrypted.
	* `set_stamp` - Backup set stamp.
	* `size_in_gbs` - Backup set size in gigabytes.
	* `status` - Current status of the backup set.
	* `target_id` - OCID of the target database associated with the backup set.
	* `time_created` - Backup set creation time in RFC3339 format.
	* `time_last_assessed` - The date and time the associated crypto assessment was last assessed, in RFC3339 format.

