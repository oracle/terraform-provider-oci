---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_certificates"
sidebar_current: "docs-oci-datasource-certificates_management-certificates"
description: |-
  Provides the list of Certificates in Oracle Cloud Infrastructure Certificates Management service
---

# Data Source: oci_certificates_management_certificates
This data source provides the list of Certificates in Oracle Cloud Infrastructure Certificates Management service.

Lists all certificates that match the query parameters.
Optionally, you can use the parameter `FilterByCertificateIdQueryParam` to limit the result set to a single item that matches the specified certificate.


## Example Usage

```hcl
data "oci_certificates_management_certificates" "test_certificates" {

	#Optional
	certificate_id = oci_certificates_management_certificate.test_certificate.id
	compartment_id = var.compartment_id
	issuer_certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id
	name = var.certificate_name
	state = var.certificate_state
}
```

## Argument Reference

The following arguments are supported:

* `certificate_id` - (Optional) The OCID of the certificate. If the parameter is set to null, the service lists all certificates.
* `compartment_id` - (Optional) A filter that returns only resources that match the given compartment OCID.
* `issuer_certificate_authority_id` - (Optional) The OCID of the certificate authority (CA). If the parameter is set to null, the service lists all CAs.
* `name` - (Optional) A filter that returns only resources that match the specified name.
* `state` - (Optional) A filter that returns only resources that match the given lifecycle state. The state value is case-insensitive.


## Attributes Reference

The following attributes are exported:

* `certificate_collection` - The list of certificate_collection.

### Certificate Reference

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

