---
subcategory: "Data Safe"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_data_safe_crypto_assessment"
sidebar_current: "docs-oci-datasource-data_safe-crypto_assessment"
description: |-
  Provides details about a specific Crypto Assessment in Oracle Cloud Infrastructure Data Safe service
---

# Data Source: oci_data_safe_crypto_assessment
This data source provides details about a specific Crypto Assessment resource in Oracle Cloud Infrastructure Data Safe service.

Gets the details of the specified crypto assessment.

## Example Usage

```hcl
data "oci_data_safe_crypto_assessment" "test_crypto_assessment" {
	#Required
	crypto_assessment_id = oci_data_safe_crypto_assessment.test_crypto_assessment.id
}
```

## Argument Reference

The following arguments are supported:

* `crypto_assessment_id` - (Required) The OCID of the crypto assessment.


## Attributes Reference

The following attributes are exported:

* `compartment_id` - The OCID of the compartment that contains the crypto assessment.
* `crypto_posture` - Cryptographic posture details captured by the assessment.
	* `backup_status` - Backup encryption status observed for the assessment.
	* `encrypted_backup_pieces_count` - Number of encrypted backup pieces.
	* `fips_mode_configured` - Common FIPS mode configured for the assessment when the target uses common FIPS configuration.
	* `fips_status` - Overall FIPS status for the assessment when the target uses common FIPS configuration.
	* `network_encryption` - Network encryption details.
	* `nne` - Native network encryption posture details.
		* `are_weak_options_allowed` - Indicates if weak NNE options are allowed.
		* `encryption_configured` - Configured network encryption algorithm.
		* `fips_mode_configured` - FIPS mode configured for NNE when the target uses legacy per-feature FIPS configuration.
		* `integrity` - NNE integrity algorithm(s).
		* `key_exchange` - Observed NNE key exchange setting.
		* `quantum_readiness` - Quantum-readiness classification for NNE posture.
		* `server_encryption` - Observed server-side encryption requirement.
		* `server_integrity` - NNE server integrity algorithm(s).
		* `status` - NNE enablement status.
	* `tde` - Transparent data encryption posture details.
		* `db_credentials_encryption_observed` - Observed DB credentials encryption algorithm.
		* `encrypted_tablespaces_count` - Number of encrypted tablespaces detected.
		* `encryption_configured` - Configured TDE encryption algorithm.
		* `fips_mode_configured` - FIPS mode configured for TDE when the target uses legacy per-feature FIPS configuration.
		* `integrity_configured` - Configured TDE integrity-related setting.
		* `key_cache_status` - The observed TDE key cache status.
		* `key_store_type` - The observed TDE key store type.
		* `master_key_encryption_algorithm` - The observed encryption algorithm used by the master key.
		* `master_key_id` - The observed TDE master key identifier.
		* `quantum_readiness` - Quantum-readiness classification for TDE posture.
		* `redo_encryption_observed` - Observed redo log encryption algorithm.
		* `status` - TDE enablement status.
		* `time_master_key_last_rotation` - The last observed rotation time for the TDE master key, in RFC3339 format.
		* `unencrypted_tablespaces_count` - Number of unencrypted tablespaces detected.
		* `wallet_location` - The observed wallet location for TDE keys.
	* `tls` - TLS posture details.
		* `are_weak_cipher_suites_allowed` - Indicates if weak TLS cipher suites are allowed.
		* `cipher_suites_configured` - TLS cipher suites configured on target.
		* `fips_mode_configured` - FIPS mode configured for TLS when the target uses legacy per-feature FIPS configuration.
		* `is_mtls_configured` - Whether TLS client authentication is configured.
		* `quantum_readiness` - Quantum-readiness classification for TLS posture.
		* `revocation_mode` - Certificate revocation checking mode.
		* `status` - TLS enablement status.
		* `versions` - TLS versions configured on target.
		* `wallet_location` - TLS wallet location observed on target.
	* `unencrypted_backup_pieces_count` - Number of unencrypted backup pieces.
* `crypto_provider` - Cryptographic provider and version information observed on the target.
* `database_architecture` - The architecture of the assessed target database.
* `database_name` - The name of the assessed target database.
* `database_version` - The version of the assessed target database.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm) Example: `{"Operations.CostCenter": "42"}` 
* `description` - The description of the crypto assessment.
* `display_name` - The display name of the crypto assessment.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm)  Example: `{"Department": "Finance"}` 
* `id` - The OCID of the crypto assessment.
* `is_assessment_scheduled` - Indicates whether scheduled execution is active for this crypto assessment. When false, the schedule value is retained but scheduled execution is paused.
* `issue_count` - Number of crypto issues detected in this assessment.
* `lifecycle_details` - Details about the current lifecycle state of the crypto assessment.
* `posture_category` - Overall posture category for the crypto assessment.
* `schedule` - The schedule used to run the crypto assessment periodically. The schedule uses the format: <version-string>;<version-specific-schedule>

	For v1, the version-specific schedule format is: <ss> <mm> <hh> <day-of-week> <day-of-month>

	Specify either day-of-week for weekly schedules or day-of-month for monthly schedules. Do not specify both. For monthly schedules, day-of-month must be between 1 and 28. If the service generates a default monthly schedule for an assessment created on day 29, 30, or 31 of a month, it uses day 28. 
* `state` - The current lifecycle state of the crypto assessment.
* `system_tags` - System tags for this resource. Each key is predefined and scoped to a namespace. For more information, see Resource Tags. Example: `{"orcl-cloud.free-tier-retained": "true"}` 
* `target_database_group_id` - The OCID of the target database group. This is returned when `targetType` is `TARGET_DATABASE_GROUP`.
* `target_id` - The OCID of the target database.
* `target_type` - The target type of the crypto assessment.
* `targets_with_issues_count` - Number of assessed targets with one or more crypto issues. For a target database assessment, this value is 1 when the target has issues and 0 otherwise. For a target database group assessment, this value is the number of targets in the group that have issues. 
* `time_created` - The date and time the crypto assessment was created, in RFC3339 format.
* `time_last_assessed` - The date and time the crypto posture was last assessed, in RFC3339 format.
* `time_updated` - The date and time the crypto assessment was last updated, in RFC3339 format.
* `triggered_by` - The actor that created the assessment.
* `type` - The type of this crypto assessment.

