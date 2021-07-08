---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_certificate_authority"
sidebar_current: "docs-oci-resource-certificates_management-certificate_authority"
description: |-
  Provides the Certificate Authority resource in Oracle Cloud Infrastructure Certificates Management service
---

# oci_certificates_management_certificate_authority
This resource provides the Certificate Authority resource in Oracle Cloud Infrastructure Certificates Management service.

Creates a new certificate authority (CA) according to the details of the request.

## Example Usage

```hcl
resource "oci_certificates_management_certificate_authority" "test_certificate_authority" {
	#Required
	certificate_authority_config {
		#Required
		config_type = var.certificate_authority_certificate_authority_config_config_type
		subject {

			#Optional
			common_name = var.certificate_authority_certificate_authority_config_subject_common_name
			country = var.certificate_authority_certificate_authority_config_subject_country
			distinguished_name_qualifier = var.certificate_authority_certificate_authority_config_subject_distinguished_name_qualifier
			domain_component = var.certificate_authority_certificate_authority_config_subject_domain_component
			generation_qualifier = var.certificate_authority_certificate_authority_config_subject_generation_qualifier
			given_name = var.certificate_authority_certificate_authority_config_subject_given_name
			initials = var.certificate_authority_certificate_authority_config_subject_initials
			locality_name = var.certificate_authority_certificate_authority_config_subject_locality_name
			organization = var.certificate_authority_certificate_authority_config_subject_organization
			organizational_unit = var.certificate_authority_certificate_authority_config_subject_organizational_unit
			pseudonym = var.certificate_authority_certificate_authority_config_subject_pseudonym
			serial_number = var.certificate_authority_certificate_authority_config_subject_serial_number
			state_or_province_name = var.certificate_authority_certificate_authority_config_subject_state_or_province_name
			street = var.certificate_authority_certificate_authority_config_subject_street
			surname = var.certificate_authority_certificate_authority_config_subject_surname
			title = var.certificate_authority_certificate_authority_config_subject_title
			user_id = oci_identity_user.test_user.id
		}

		#Optional
		issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
		signing_algorithm = var.certificate_authority_certificate_authority_config_signing_algorithm
		validity {

			#Optional
			time_of_validity_not_after = var.certificate_authority_certificate_authority_config_validity_time_of_validity_not_after
			time_of_validity_not_before = var.certificate_authority_certificate_authority_config_validity_time_of_validity_not_before
		}
		version_name = var.certificate_authority_certificate_authority_config_version_name
	}
	compartment_id = var.compartment_id
	kms_key_id = oci_kms_key.test_key.id
	name = var.certificate_authority_name

	#Optional
	certificate_authority_rules {
		#Required
		rule_type = var.certificate_authority_certificate_authority_rules_rule_type

		#Optional
		certificate_authority_max_validity_duration = var.certificate_authority_certificate_authority_rules_certificate_authority_max_validity_duration
		leaf_certificate_max_validity_duration = var.certificate_authority_certificate_authority_rules_leaf_certificate_max_validity_duration
	}
	certificate_revocation_list_details {
		#Required
		object_storage_config {
			#Required
			object_storage_bucket_name = oci_objectstorage_bucket.test_bucket.name
			object_storage_object_name_format = var.certificate_authority_certificate_revocation_list_details_object_storage_config_object_storage_object_name_format

			#Optional
			object_storage_namespace = var.certificate_authority_certificate_revocation_list_details_object_storage_config_object_storage_namespace
		}

		#Optional
		custom_formatted_urls = var.certificate_authority_certificate_revocation_list_details_custom_formatted_urls
	}
	defined_tags = {"foo-namespace.bar-key"= "value"}
	description = var.certificate_authority_description
	freeform_tags = {"bar-key"= "value"}
}
```

## Argument Reference

The following arguments are supported:

* `certificate_authority_config` - (Required) (Updatable) The configuration details for creating a certificate authority (CA).
	* `config_type` - (Required) (Updatable) The origin of the CA.
	* `issuer_certificate_authority_id` - (Required when config_type=SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) The OCID of the private CA.
	* `signing_algorithm` - (Optional) The algorithm used to sign public key certificates that the CA issues.
	* `subject` - (Required) The subject of the certificate, which is a distinguished name that identifies the entity that owns the public key in the certificate. 
		* `common_name` - (Required when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Common name or fully-qualified domain name (RDN CN).
		* `country` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Country name (RDN C).
		* `distinguished_name_qualifier` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Distinguished name qualifier(RDN DNQ).
		* `domain_component` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Domain component (RDN DC).
		* `generation_qualifier` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Personal generational qualifier (for example, Sr., Jr. 3rd, or IV).
		* `given_name` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Personal given name (RDN G or GN).
		* `initials` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Personal initials.
		* `locality_name` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Locality (RDN L).
		* `organization` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Organization (RDN O).
		* `organizational_unit` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Organizational unit (RDN OU).
		* `pseudonym` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Subject pseudonym.
		* `serial_number` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Unique subject identifier, which is not the same as the certificate serial number (RDN SERIALNUMBER).
		* `state_or_province_name` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) State or province name (RDN ST or S).
		* `street` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Street address (RDN STREET).
		* `surname` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Personal surname (RDN SN).
		* `title` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) Title (RDN T or TITLE).
		* `user_id` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) User ID (RDN UID).
	* `validity` - (Optional) (Updatable) An object that describes a period of time during which an entity is valid. If this is not provided when you create a certificate, the validity of the issuing CA is used. 
		* `time_of_validity_not_after` - (Required when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) (Updatable) The date on which the certificate validity period ends, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
		* `time_of_validity_not_before` - (Applicable when config_type=ROOT_CA_GENERATED_INTERNALLY | SUBORDINATE_CA_ISSUED_BY_INTERNAL_CA) (Updatable) The date on which the certificate validity period begins, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `version_name` - (Optional) (Updatable) The name of the CA version. When the value is not null, a name is unique across versions of a given CA. 
* `certificate_authority_rules` - (Optional) (Updatable) A list of rules that control how the CA is used and managed.
	* `certificate_authority_max_validity_duration` - (Optional) (Updatable) A property indicating the maximum validity duration, in days, of subordinate CA's issued by this CA. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `leaf_certificate_max_validity_duration` - (Optional) (Updatable) A property indicating the maximum validity duration, in days, of leaf certificates issued by this CA. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `rule_type` - (Required) (Updatable) The type of rule, whether a renewal rule regarding when to renew the CA or an issuance expiry rule that governs how long the certificates and CAs issued by the CA are valid. (For internal use only) An internal issuance rule defines the number and type of certificates that the CA can issue. 
* `certificate_revocation_list_details` - (Optional) (Updatable) The details of the certificate revocation list (CRL).
	* `custom_formatted_urls` - (Optional) (Updatable) Optional CRL access points, expressed using a format where the version number of the issuing CA is inserted wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2. 
	* `object_storage_config` - (Required) (Updatable) The details of the Object Storage bucket configured to store the certificate revocation list (CRL).
		* `object_storage_bucket_name` - (Required) (Updatable) The name of the bucket where the CRL is stored.
		* `object_storage_namespace` - (Optional) (Updatable) The tenancy of the bucket where the CRL is stored.
		* `object_storage_object_name_format` - (Required) (Updatable) The object name in the bucket where the CRL is stored, expressed using a format where the version number of the issuing CA is inserted as part of the Object Storage object name wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2. 
* `compartment_id` - (Required) (Updatable) The compartment in which you want to create the CA.
* `defined_tags` - (Optional) (Updatable) Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - (Optional) (Updatable) A brief description of the CA.
* `freeform_tags` - (Optional) (Updatable) Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `kms_key_id` - (Required) The OCID of the Oracle Cloud Infrastructure Vault key used to encrypt the CA.
* `name` - (Required) A user-friendly name for the CA. Names are unique within a compartment. Avoid entering confidential information. Valid characters include uppercase or lowercase letters, numbers, hyphens, underscores, and periods.


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `certificate_authority_rules` - An optional list of rules that control how the CA is used and managed.
	* `certificate_authority_max_validity_duration` - A property indicating the maximum validity duration, in days, of subordinate CA's issued by this CA. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `leaf_certificate_max_validity_duration` - A property indicating the maximum validity duration, in days, of leaf certificates issued by this CA. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `rule_type` - The type of rule, whether a renewal rule regarding when to renew the CA or an issuance expiry rule that governs how long the certificates and CAs issued by the CA are valid. (For internal use only) An internal issuance rule defines the number and type of certificates that the CA can issue. 
* `certificate_revocation_list_details` - The details of the certificate revocation list (CRL).
	* `custom_formatted_urls` - Optional CRL access points, expressed using a format where the version number of the issuing CA is inserted wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2. 
	* `object_storage_config` - The details of the Object Storage bucket configured to store the certificate revocation list (CRL).
		* `object_storage_bucket_name` - The name of the bucket where the CRL is stored.
		* `object_storage_namespace` - The tenancy of the bucket where the CRL is stored.
		* `object_storage_object_name_format` - The object name in the bucket where the CRL is stored, expressed using a format where the version number of the issuing CA is inserted as part of the Object Storage object name wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2. 
* `compartment_id` - The OCID of the compartment under which the CA is created.
* `config_type` - The origin of the CA.
* `current_version` - The metadata details of the certificate authority (CA) version. This summary object does not contain the CA contents.
	* `certificate_authority_id` - The OCID of the CA.
	* `issuer_ca_version_number` - The version number of the issuing CA.
	* `revocation_status` - The current revocation status of the entity.
		* `revocation_reason` - The reason the certificate or certificate authority (CA) was revoked.
		* `time_of_revocation` - The time when the entity was revoked, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `serial_number` - A unique certificate identifier used in certificate revocation tracking, formatted as octets. Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF` 
	* `stages` - A list of rotation states for this CA version.
	* `time_created` - A optional property indicating when the CA version was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `time_of_deletion` - An optional property indicating when to delete the CA version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `validity` - An object that describes a period of time during which an entity is valid. If this is not provided when you create a certificate, the validity of the issuing CA is used. 
		* `time_of_validity_not_after` - The date on which the certificate validity period ends, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
		* `time_of_validity_not_before` - The date on which the certificate validity period begins, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `version_name` - The name of the CA version. When this value is not null, the name is unique across CA versions for a given CA. 
	* `version_number` - The version number of the CA.
* `defined_tags` - Usage of predefined tag keys. These predefined keys are scoped to namespaces. Example: `{"foo-namespace.bar-key": "value"}` 
* `description` - A brief description of the CA.
* `freeform_tags` - Simple key-value pair that is applied without any predefined name, type or scope. Exists for cross-compatibility only. Example: `{"bar-key": "value"}` 
* `id` - The OCID of the CA.
* `issuer_certificate_authority_id` - The OCID of the parent CA that issued this CA. If this is the root CA, then this value is null. 
* `kms_key_id` - The OCID of the Oracle Cloud Infrastructure Vault key used to encrypt the CA.
* `lifecycle_details` - Additional information about the current CA lifecycle state.
* `name` - A user-friendly name for the CA. Names are unique within a compartment. Avoid entering confidential information. Valid characters include uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
* `signing_algorithm` - The algorithm used to sign public key certificates that the CA issues.
* `state` - The current lifecycle state of the certificate authority.
* `subject` - The subject of the certificate, which is a distinguished name that identifies the entity that owns the public key in the certificate. 
	* `common_name` - Common name or fully-qualified domain name (RDN CN).
	* `country` - Country name (RDN C).
	* `distinguished_name_qualifier` - Distinguished name qualifier(RDN DNQ).
	* `domain_component` - Domain component (RDN DC).
	* `generation_qualifier` - Personal generational qualifier (for example, Sr., Jr. 3rd, or IV).
	* `given_name` - Personal given name (RDN G or GN).
	* `initials` - Personal initials.
	* `locality_name` - Locality (RDN L).
	* `organization` - Organization (RDN O).
	* `organizational_unit` - Organizational unit (RDN OU).
	* `pseudonym` - Subject pseudonym.
	* `serial_number` - Unique subject identifier, which is not the same as the certificate serial number (RDN SERIALNUMBER).
	* `state_or_province_name` - State or province name (RDN ST or S).
	* `street` - Street address (RDN STREET).
	* `surname` - Personal surname (RDN SN).
	* `title` - Title (RDN T or TITLE).
	* `user_id` - User ID (RDN UID).
* `time_created` - A property indicating when the CA was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the CA version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Certificate Authority
	* `update` - (Defaults to 20 minutes), when updating the Certificate Authority
	* `delete` - (Defaults to 20 minutes), when destroying the Certificate Authority


## Import

CertificateAuthorities can be imported using the `id`, e.g.

```
$ terraform import oci_certificates_management_certificate_authority.test_certificate_authority "id"
```

