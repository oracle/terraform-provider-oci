---
subcategory: "Certificates Management"
layout: "oci"
page_title: "Oracle Cloud Infrastructure: oci_certificates_management_certificate_authority_version"
sidebar_current: "docs-oci-datasource-certificates_management-certificate_authority_version"
description: |-
  Provides details about a specific Certificate Authority Version in Oracle Cloud Infrastructure Certificates Management service
---

# Data Source: oci_certificates_management_certificate_authority_version
This data source provides details about a specific Certificate Authority Version resource in Oracle Cloud Infrastructure Certificates Management service.

Lists all versions for the specified certificate authority (CA).
Optionally, you can use the parameter `FilterByVersionNumberQueryParam` to limit the results to a single item that matches the specified version number.


## Example Usage

```hcl
data "oci_certificates_management_certificate_authority_version" "test_certificate_authority_version" {
	#Required
	certificate_authority_id = oci_certificates_management_certificate_authority.test_certificate_authority.id

	#Optional
	version_number = var.certificate_authority_version_version_number
}
```

## Argument Reference

The following arguments are supported:

* `certificate_authority_id` - (Required) The OCID of the certificate authority (CA).
* `version_number` - (Optional) A filter that returns only resources that match the specified version number. The default value is 0, which means that this filter is not applied. 


## Attributes Reference

The following attributes are exported:

* `items` - A list of certificate authority version summary objects.
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

