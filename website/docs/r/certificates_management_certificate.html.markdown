---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_certificate"
sidebar_current: "docs-oci-resource-certificates_management-certificate"
description: |-
  Provides the Certificate resource in Oracle Cloud Infrastructure Certificates Management service
---

# oci_certificates_management_certificate
This resource provides the Certificate resource in Oracle Cloud Infrastructure Certificates Management service.

Creates a new certificate according to the details of the request.

## Example Usage

```hcl
resource "oci_certificates_management_certificate" "test_certificate" {
	#Required
	certificate_config {
		#Required
		config_type = var.certificate_certificate_config_config_type

		#Optional
		cert_chain_pem = var.certificate_certificate_config_cert_chain_pem
		certificate_pem = var.certificate_certificate_config_certificate_pem
		certificate_profile_type = var.certificate_certificate_config_certificate_profile_type
		csr_pem = var.certificate_certificate_config_csr_pem
		issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
		key_algorithm = var.certificate_certificate_config_key_algorithm
		private_key_pem = var.certificate_certificate_config_private_key_pem
		private_key_pem_passphrase = var.certificate_certificate_config_private_key_pem_passphrase
		signature_algorithm = var.certificate_certificate_config_signature_algorithm
		subject {

			#Optional
			common_name = var.certificate_certificate_config_subject_common_name
			country = var.certificate_certificate_config_subject_country
			distinguished_name_qualifier = var.certificate_certificate_config_subject_distinguished_name_qualifier
			domain_component = var.certificate_certificate_config_subject_domain_component
			generation_qualifier = var.certificate_certificate_config_subject_generation_qualifier
			given_name = var.certificate_certificate_config_subject_given_name
			initials = var.certificate_certificate_config_subject_initials
			locality_name = var.certificate_certificate_config_subject_locality_name
			organization = var.certificate_certificate_config_subject_organization
			organizational_unit = var.certificate_certificate_config_subject_organizational_unit
			pseudonym = var.certificate_certificate_config_subject_pseudonym
			serial_number = var.certificate_certificate_config_subject_serial_number
			state_or_province_name = var.certificate_certificate_config_subject_state_or_province_name
			street = var.certificate_certificate_config_subject_street
			surname = var.certificate_certificate_config_subject_surname
			title = var.certificate_certificate_config_subject_title
			user_id = oci_identity_user.test_user.id
		}
		subject_alternative_names {

			#Optional
			type = var.certificate_certificate_config_subject_alternative_names_type
			value = var.certificate_certificate_config_subject_alternative_names_value
		}
		validity {

			#Optional
			time_of_validity_not_after = var.certificate_certificate_config_validity_time_of_validity_not_after
			time_of_validity_not_before = var.certificate_certificate_config_validity_time_of_validity_not_before
		}
		version_name = var.certificate_certificate_config_version_name
	}
	compartment_id = var.compartment_id
	name = var.certificate_name

	#Optional
	certificate_rules {
		#Required
		advance_renewal_period = var.certificate_certificate_rules_advance_renewal_period
		renewal_interval = var.certificate_certificate_rules_renewal_interval
		rule_type = var.certificate_certificate_rules_rule_type
	}
	defined_tags = {"Operations.CostCenter"= "42"}
	description = var.certificate_description
	freeform_tags = {"Department"= "Finance"}
}
```

## Argument Reference

The following arguments are supported:

* `certificate_config` - (Required) (Updatable) The details of the contents of the certificate and certificate metadata.
	* `cert_chain_pem` - (Required when config_type=IMPORTED) (Updatable) The certificate chain (in PEM format) for the imported certificate.
	* `certificate_pem` - (Required when config_type=IMPORTED) (Updatable) The certificate (in PEM format) for the imported certificate.
	* `certificate_profile_type` - (Required when config_type=ISSUED_BY_INTERNAL_CA) The name of the profile used to create the certificate, which depends on the type of certificate you need.
	* `config_type` - (Required) (Updatable) The origin of the certificate.
	* `csr_pem` - (Required when config_type=MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA) (Updatable) The certificate signing request (in PEM format).
	* `issuer_certificate_authority_id` - (Required when config_type=ISSUED_BY_INTERNAL_CA | MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA) The OCID of the private CA.
	* `key_algorithm` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) The algorithm to use to create key pairs.
	* `private_key_pem` - (Required when config_type=IMPORTED) (Updatable) The private key (in PEM format) for the imported certificate.
	* `private_key_pem_passphrase` - (Applicable when config_type=IMPORTED) (Updatable) An optional passphrase for the private key.
	* `signature_algorithm` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) The algorithm to use to sign the public key certificate.
	* `subject` - (Required when config_type=ISSUED_BY_INTERNAL_CA) The subject of the certificate, which is a distinguished name that identifies the entity that owns the public key in the certificate. 
		* `common_name` - (Required when config_type=ISSUED_BY_INTERNAL_CA) Common name or fully-qualified domain name (RDN CN).
		* `country` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Country name (RDN C).
		* `distinguished_name_qualifier` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Distinguished name qualifier(RDN DNQ).
		* `domain_component` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Domain component (RDN DC).
		* `generation_qualifier` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Personal generational qualifier (for example, Sr., Jr. 3rd, or IV).
		* `given_name` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Personal given name (RDN G or GN).
		* `initials` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Personal initials.
		* `locality_name` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Locality (RDN L).
		* `organization` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Organization (RDN O).
		* `organizational_unit` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Organizational unit (RDN OU).
		* `pseudonym` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Subject pseudonym.
		* `serial_number` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Unique subject identifier, which is not the same as the certificate serial number (RDN SERIALNUMBER).
		* `state_or_province_name` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) State or province name (RDN ST or S).
		* `street` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Street address (RDN STREET).
		* `surname` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Personal surname (RDN SN).
		* `title` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) Title (RDN T or TITLE).
		* `user_id` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) User ID (RDN UID).
	* `subject_alternative_names` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA) A list of subject alternative names.
		* `type` - (Required when config_type=ISSUED_BY_INTERNAL_CA) The subject alternative name type. Currently only DNS domain or host names and IP addresses are supported.
		* `value` - (Required when config_type=ISSUED_BY_INTERNAL_CA) The subject alternative name.
	* `validity` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA | MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA) (Updatable) An object that describes a period of time during which an entity is valid. If this is not provided when you create a certificate, the validity of the issuing CA is used. 
		* `time_of_validity_not_after` - (Required when config_type=ISSUED_BY_INTERNAL_CA | MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA) (Updatable) The date on which the certificate validity period ends, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
		* `time_of_validity_not_before` - (Applicable when config_type=ISSUED_BY_INTERNAL_CA | MANAGED_EXTERNALLY_ISSUED_BY_INTERNAL_CA) (Updatable) The date on which the certificate validity period begins, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `version_name` - (Optional) (Updatable) A name for the certificate. When the value is not null, a name is unique across versions of a given certificate.
* `certificate_rules` - (Optional) (Updatable) An optional list of rules that control how the certificate is used and managed.
	* `advance_renewal_period` - (Required) (Updatable) A property specifying the period of time, in days, before the certificate's targeted renewal that the process should occur. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `renewal_interval` - (Required) (Updatable) A property specifying how often, in days, a certificate should be renewed. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `rule_type` - (Required) (Updatable) The type of rule.
* `compartment_id` - (Required) (Updatable) The OCID of the compartment where you want to create the certificate.
* `defined_tags` - (Optional) (Updatable) Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - (Optional) (Updatable) A brief description of the certificate. Avoid entering confidential information.
* `freeform_tags` - (Optional) (Updatable) Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `name` - (Required) A user-friendly name for the certificate. Names are unique within a compartment. Avoid entering confidential information. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods. 


** IMPORTANT **
Any change to a property that does not support update will force the destruction and recreation of the resource with the new property values

## Attributes Reference

The following attributes are exported:

* `certificate_profile_type` - The name of the profile used to create the certificate, which depends on the type of certificate you need.
* `certificate_revocation_list_details` - The details of the certificate revocation list (CRL).
	* `custom_formatted_urls` - Optional CRL access points, expressed using a format where the version number of the issuing CA is inserted wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2. 
	* `object_storage_config` - The details of the Object Storage bucket configured to store the certificate revocation list (CRL).
		* `object_storage_bucket_name` - The name of the bucket where the CRL is stored.
		* `object_storage_namespace` - The tenancy of the bucket where the CRL is stored.
		* `object_storage_object_name_format` - The object name in the bucket where the CRL is stored, expressed using a format where the version number of the issuing CA is inserted as part of the Object Storage object name wherever you include a pair of curly braces. This versioning scheme helps avoid collisions when new CA versions are created. For example, myCrlFileIssuedFromCAVersion{}.crl becomes myCrlFileIssuedFromCAVersion2.crl for CA version 2. 
* `certificate_rules` - A list of rules that control how the certificate is used and managed.
	* `advance_renewal_period` - A property specifying the period of time, in days, before the certificate's targeted renewal that the process should occur. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `renewal_interval` - A property specifying how often, in days, a certificate should be renewed. Expressed in [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601#Time_intervals) format. 
	* `rule_type` - The type of rule.
* `compartment_id` - The OCID of the compartment where you want to create the certificate.
* `config_type` - The origin of the certificate.
* `current_version` - The details of the certificate version. This object does not contain the certificate contents.
	* `certificate_id` - The OCID of the certificate.
	* `issuer_ca_version_number` - The version number of the issuing certificate authority (CA).
	* `revocation_status` - The current revocation status of the entity.
		* `revocation_reason` - The reason the certificate or certificate authority (CA) was revoked.
		* `time_of_revocation` - The time when the entity was revoked, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `serial_number` - A unique certificate identifier used in certificate revocation tracking, formatted as octets. Example: `03 AC FC FA CC B3 CB 02 B8 F8 DE F5 85 E7 7B FF` 
	* `stages` - A list of rotation states for this certificate version.
	* `subject_alternative_names` - A list of subject alternative names.
		* `type` - The subject alternative name type. Currently only DNS domain or host names and IP addresses are supported.
		* `value` - The subject alternative name.
	* `time_created` - A optional property indicating the time when the certificate version was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `time_of_deletion` - An optional property indicating when to delete the certificate version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `validity` - An object that describes a period of time during which an entity is valid. If this is not provided when you create a certificate, the validity of the issuing CA is used. 
		* `time_of_validity_not_after` - The date on which the certificate validity period ends, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
		* `time_of_validity_not_before` - The date on which the certificate validity period begins, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
	* `version_name` - The name of the certificate version. When the value is not null, a name is unique across versions of a given certificate. 
	* `version_number` - The version number of the certificate.
* `defined_tags` - Defined tags for this resource. Each key is predefined and scoped to a namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Operations.CostCenter": "42"}` 
* `description` - A brief description of the certificate. Avoid entering confidential information.
* `freeform_tags` - Free-form tags for this resource. Each tag is a simple key-value pair with no predefined name, type, or namespace. For more information, see [Resource Tags](https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm). Example: `{"Department": "Finance"}` 
* `id` - The OCID of the certificate.
* `issuer_certificate_authority_id` - The OCID of the certificate authority (CA) that issued the certificate.
* `key_algorithm` - The algorithm used to create key pairs.
* `lifecycle_details` - Additional information about the current lifecycle state of the certificate.
* `name` - A user-friendly name for the certificate. Names are unique within a compartment. Avoid entering confidential information. Valid characters are uppercase or lowercase letters, numbers, hyphens, underscores, and periods.
* `signature_algorithm` - The algorithm used to sign the public key certificate.
* `state` - The current lifecycle state of the certificate.
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
* `time_created` - A property indicating when the certificate was created, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 
* `time_of_deletion` - An optional property indicating when to delete the certificate version, expressed in [RFC 3339](https://tools.ietf.org/html/rfc3339) timestamp format. Example: `2019-04-03T21:10:29.600Z` 

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://registry.terraform.io/providers/hashicorp/oci/latest/docs/guides/changing_timeouts) for certain operations:
	* `create` - (Defaults to 20 minutes), when creating the Certificate
	* `update` - (Defaults to 20 minutes), when updating the Certificate
	* `delete` - (Defaults to 20 minutes), when destroying the Certificate


## Import

Certificates can be imported using the `id`, e.g.

```
$ terraform import oci_certificates_management_certificate.test_certificate "id"
```

