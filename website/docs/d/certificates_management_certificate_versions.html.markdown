---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_certificate_versions"
sidebar_current: "docs-oci-datasource-certificates_management-certificate_versions"
description: |-
  Provides the list of Certificate Versions in Oracle Cloud Infrastructure Certificates Management service
---

# Data Source: oci_certificates_management_certificate_versions
This data source provides the list of Certificate Versions in Oracle Cloud Infrastructure Certificates Management service.

Lists all certificate versions for the specified certificate.
Optionally, you can use the parameter `FilterByVersionNumberQueryParam` to limit the result set to a single item that matches the specified version number.


## Example Usage

```hcl
data "oci_certificates_management_certificate_versions" "test_certificate_versions" {
	#Required
	certificate_id = oci_certificates_management_certificate.test_certificate.id

	#Optional
	version_number = var.certificate_version_version_number
}
```

## Argument Reference

The following arguments are supported:

* `certificate_id` - (Required) The OCID of the certificate.
* `version_number` - (Optional) A filter that returns only resources that match the specified version number. The default value is 0, which means that this filter is not applied. 


## Attributes Reference

The following attributes are exported:

* `certificate_version_collection` - The list of certificate_version_collection.

### CertificateVersion Reference

The following attributes are exported:

* `items` - A list of certificate version summary objects.
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

