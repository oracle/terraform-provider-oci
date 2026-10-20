---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment_certificates"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment_certificates"
description: |-
  Provides the list of Crypto Assessment Certificates in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment_certificates
This data source provides the list of Crypto Assessment Certificates in Oracle Cloud Infrastructure Data Safe service.

Lists certificates discovered across targets in a compartment, including target, wallet location, issuer, subject, validity window, expiry bucket, public key type, and status so expiring or weak certificates can be identified and prioritized.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment_certificates" "test_crypto_assessment_certificates" {
	#Required
	compartment_id = var.compartment_id

	#Optional
	access_level = var.crypto_assessment_certificate_access_level
	assessment_id = oci_database_migration_assessment.test_assessment.id
	assessment_type = var.crypto_assessment_certificate_assessment_type
	certificate_type = var.crypto_assessment_certificate_certificate_type
	compartment_id_in_subtree = var.crypto_assessment_certificate_compartment_id_in_subtree
	days_to_expiry = var.crypto_assessment_certificate_days_to_expiry
	expiry_bucket = var.crypto_assessment_certificate_expiry_bucket
	public_key_type = var.crypto_assessment_certificate_public_key_type
	signature_algorithm = var.crypto_assessment_certificate_signature_algorithm
	status = var.crypto_assessment_certificate_status
	target_id = oci_cloud_guard_target.test_target.id
	target_ids = var.crypto_assessment_certificate_target_ids
}
```

## Argument Reference

The following arguments are supported:

* `access_level` - (Optional) Valid values are RESTRICTED and ACCESSIBLE. Default is RESTRICTED. Setting this to ACCESSIBLE returns only those compartments for which the user has INSPECT permissions directly or indirectly (permissions can be on a resource in a subcompartment). When set to RESTRICTED permissions are checked and no partial results are displayed. 
* `assessment_id` - (Optional) A filter to return only resources associated with the specified crypto assessment OCID.
* `assessment_type` - (Optional) A filter to return targets from assessments of the specified type.
* `certificate_type` - (Optional) A filter to return only certificates of any of the specified types.
* `compartment_id` - (Required) A filter to return only resources that match the specified compartment OCID.
* `compartment_id_in_subtree` - (Optional) Default is false. When set to true, the hierarchy of compartments is traversed and all compartments and subcompartments in the tenancy are returned. Depends on the 'accessLevel' setting. 
* `days_to_expiry` - (Optional) A filter to return certificates whose validTill timestamp is on or before the current time plus the specified number of days. Negative values are allowed and filter certificates that expired on or before that many days ago.
* `expiry_bucket` - (Optional) A filter to return only certificates in the specified expiry bucket. Supported values are 0_15, 15_30, 30_60, 60_90, and 90_PLUS.
* `public_key_type` - (Optional) A filter to return only certificates with any of the specified public key types. Stored values are normalized forms such as RSA2048, RSA4096, or EC256.
* `signature_algorithm` - (Optional) A filter to return only certificates whose signature algorithm contains any of the specified values. For example, use SHA1 to match SHA1-based certificate signatures.
* `status` - (Optional) A filter to return only certificates with any of the specified statuses.
* `target_id` - (Optional) A filter to return only inventory rows associated with the specified target OCID.
* `target_ids` - (Optional) A filter to return only resources associated with any of the specified target OCIDs.


## Attributes Reference

The following attributes are exported:

* `crypto_assessment_certificate_collection` - The list of crypto_assessment_certificate_collection.

### CryptoAssessmentCertificate Reference

The following attributes are exported:

* `items` - Certificate summaries that match the request filters.
	* `age` - Age of the certificate in whole days, calculated from timeValidFrom using the current UTC date.
	* `assessment_id` - OCID of the crypto assessment that discovered the certificate.
	* `assessment_type` - Type of the crypto assessment that discovered the certificate.
	* `certificate_type` - Type of certificate.
	* `compartment_id` - OCID of the compartment that contains the certificate row.
	* `days_to_expiry` - Number of whole days until certificate expiration, calculated from timeValidUntil. Negative values indicate already expired certificates.
	* `expiry_bucket` - Expiry bucket populated for the certificate when the assessment runs. Supported values are 0_15, 15_30, 30_60, 60_90, and 90_PLUS.
	* `issuer` - Issuer of the certificate.
	* `public_key_type` - Public key type and size.
	* `serial_number` - Certificate serial number.
	* `signature_algorithm` - Signature algorithm used by the certificate.
	* `status` - Certificate validity status.
	* `subject` - Subject of the certificate.
	* `target_id` - OCID of the target database associated with the certificate.
	* `time_last_assessed` - The date and time the associated crypto assessment was last assessed, in RFC3339 format.
	* `time_valid_from` - Certificate validity start time in RFC3339 format.
	* `time_valid_until` - Certificate validity end time in RFC3339 format.
	* `wallet_location` - Wallet location where the certificate was discovered, if available.

