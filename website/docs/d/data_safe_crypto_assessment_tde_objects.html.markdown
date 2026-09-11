---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_tde_objects"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_tde_objects"
description: |-
  Provides the list of Crypto Assessment Tde Objects in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_tde_objects
This data source provides the list of Crypto Assessment Tde Objects in Oracle Cloud Infrastructure Data Safe service.

Lists TDE object encryption summaries across targets in a compartment. Use assessmentId to narrow results to one crypto assessment, and objectType to return either tablespace-level or column-level TDE observations.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_tde_objects" "test_crypto_assessment_tde_objects" {
	#Required
	compartment_id = var.compartment_id
	object_type = var.crypto_assessment_tde_object_object_type

	#Optional
	access_level = var.crypto_assessment_tde_object_access_level
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessment_type = var.crypto_assessment_tde_object_assessment_type
	compartment_id_in_subtree = var.crypto_assessment_tde_object_compartment_id_in_subtree
	encryption_observed = var.crypto_assessment_tde_object_encryption_observed
	encryption_status = var.crypto_assessment_tde_object_encryption_status
	quantum_readiness = var.crypto_assessment_tde_object_quantum_readiness
	target_id = oci_cloud_guard_target.test_target.id
	target_ids = var.crypto_assessment_tde_object_target_ids
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `assessment_id` - (Optional) A filter to return only resources associated with the specified crypto assessment OCID.
* `assessment_type` - (Optional) A filter to return targets from assessments of the specified type.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `encryption_observed` - (Optional) Filters TDE object summary rows by any of the specified observed encryption algorithms.
* `encryption_status` - (Optional) Filters TDE object summary rows by derived encryption status. NOT_SUPPORTED maps to rows where encryptionObserved is NOT_SUPPORTED, UNENCRYPTED maps to rows where encryptionObserved is null or NONE, and ENCRYPTED maps to rows where the observed encryption algorithm is any other value.
* `object_type` - (Required) A required filter to return only TDE objects of the specified type.
* `quantum_readiness` - (Optional) Filters TDE object summary rows by quantum-readiness category.
* `target_id` - (Optional) A filter to return only inventory rows associated with the specified target OCID.
* `target_ids` - (Optional) A filter to return only resources associated with any of the specified target OCIDs.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_tde_object_collection` - The list of crypto_assessment_tde_object_collection.

### CryptoAssessmentTdeObject Reference

The following attributes are exported:

* `items` - TDE object encryption summary items.
	* `assessment_id` - OCID of the crypto assessment that discovered the TDE object.
	* `column_name` - Name of the encrypted column. This field is returned when objectType is COLUMN.
	* `encryption_observed` - Encryption algorithm observed for the TDE object.
	* `mode_observed` - Encryption mode observed for the tablespace. This field is returned when objectType is TABLESPACE.
	* `quantum_readiness` - Quantum-readiness classification for the observed TDE object encryption.
	* `schema_name` - Name of the schema containing the encrypted column. This field is returned when objectType is COLUMN.
	* `size_in_gbs` - Tablespace size in gigabytes. This field is returned when objectType is TABLESPACE.
	* `table_name` - Name of the table containing the encrypted column. This field is returned when objectType is COLUMN.
	* `tablespace_name` - Name of the tablespace. This field is returned when objectType is TABLESPACE.
	* `target_id` - OCID of the target database associated with the TDE object.
	* `time_last_assessed` - The date and time the associated crypto assessment was last assessed, in RFC3339 format.

