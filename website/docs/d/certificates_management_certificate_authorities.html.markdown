---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_certificate_authorities"
sidebar_current: "docs-oci-datasource-certificates_management-certificate_authorities"
description: |-
  Provides the list of Certificate Authorities in Oracle Cloud Infrastructure Certificates Management service
---

# Data Source: oci_certificates_management_certificate_authorities
This data source provides the list of Certificate Authorities in Oracle Cloud Infrastructure Certificates Management service.

Lists all certificate authorities (CAs) in the specified compartment.
Optionally, you can use the parameter `FilterByCertificateAuthorityIdQueryParam` to limit the results to a single item that matches the specified CA.


## Example Usage

```hcl
data "oci_certificates_management_certificate_authorities" "test_certificate_authorities" {

	#Optional
	certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
	compartment_id = var.compartment_id
	issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
	name = var.certificate_authority_name
	state = var.certificate_authority_state
}
```

## Argument Reference

The following arguments are supported:

* `certificate_authority_id` - (Optional) The OCID of the certificate authority (CA). If the parameter is set to null, the service lists all CAs.
* `compartment_id` - (Optional) A filter that returns only resources that match the given compartment OCID.
* `issuer_certificate_authority_id` - (Optional) The OCID of the certificate authority (CA). If the parameter is set to null, the service lists all CAs.
* `name` - (Optional) A filter that returns only resources that match the specified name.
* `state` - (Optional) A filter that returns only resources that match the given lifecycle state. The state value is case-insensitive.


## Attributes Reference

The following attributes are exported:

* `certificate_authority_collection` - The list of certificate_authority_collection.

### CertificateAuthority Reference

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

